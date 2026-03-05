package expense

import (
	"be-catatin/internal/entity"
	budgetRepo "be-catatin/internal/repository/budget"
	expenseRepo "be-catatin/internal/repository/expense"
	"be-catatin/pkg/cloudinary"
	"fmt"
	"mime/multipart"
	"time"
)

type Usecase interface {
	Create(userID uint, categoryID uint, title string, amount float64, transactionDate string, fileHeader *multipart.FileHeader) (*entity.Expense, error)
	FindAll(userID uint) ([]*entity.Expense, error)
	FindByID(id uint, userID uint) (*entity.Expense, error)
	Delete(id uint, userID uint) error
}

type usecase struct {
	repo       expenseRepo.Repository
	budgetRepo budgetRepo.Repository
	cloudinary *cloudinary.CloudinaryService
}

func NewUsecase(repo expenseRepo.Repository, budgetRepo budgetRepo.Repository, cloudinary *cloudinary.CloudinaryService) Usecase {
	return &usecase{repo, budgetRepo, cloudinary}
}

func (u *usecase) Create(userID uint, categoryID uint, title string, amount float64, transactionDate string, fileHeader *multipart.FileHeader) (*entity.Expense, error) {
	// Parse Transaction Date
	layout := "2006-01-02"
	t, err := time.Parse(layout, transactionDate)
	if err != nil {
		return nil, err
	}

	expense := &entity.Expense{
		UserID:          userID,
		CategoryID:      categoryID,
		Title:           title,
		Amount:          amount,
		TransactionDate: t,
	}

	// Optional Image Upload
	if fileHeader != nil {
		file, err := fileHeader.Open()
		if err != nil {
			return nil, err
		}
		defer file.Close()

		// Generate a unique filename identifier for Cloudinary
		filename := fmt.Sprintf("expense_%d_%s", userID, time.Now().Format("20060102150405"))

		url, _, err := u.cloudinary.UploadImage(file, filename)
		if err != nil {
			return nil, err
		}
		expense.ImgPath = url
	}

	err = u.repo.Create(expense)
	if err != nil {
		return nil, err
	}

	// Automaticaly deduct budget for that month if it exists
	budgets, err := u.budgetRepo.FindAll(userID, int(t.Month()), t.Year())
	if err == nil && len(budgets) > 0 {
		budget := budgets[0]
		budget.Amount -= amount
		_ = u.budgetRepo.Update(budget)
	}

	return expense, nil
}

func (u *usecase) FindAll(userID uint) ([]*entity.Expense, error) {
	return u.repo.FindAll(userID)
}

func (u *usecase) FindByID(id uint, userID uint) (*entity.Expense, error) {
	return u.repo.FindByID(id, userID)
}

func (u *usecase) Delete(id uint, userID uint) error {
	expense, err := u.repo.FindByID(id, userID)
	if err != nil {
		return err
	}

	// Only attempt Cloudinary delete if ImgPath exists
	if expense.ImgPath != "" {
		publicID := u.cloudinary.GetPublicIDFromURL(expense.ImgPath)
		if publicID != "" {
			_ = u.cloudinary.DeleteImage(publicID) // Fire-and-forget deletion
		}
	}

	return u.repo.Delete(id, userID)
}
