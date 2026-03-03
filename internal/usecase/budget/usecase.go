package budget

import (
	"be-catatin/internal/entity"
	budgetRepo "be-catatin/internal/repository/budget"
)

type Usecase interface {
	Create(userID uint, amount float64, month int, year int) (*entity.Budget, error)
	FindAll(userID uint) ([]*entity.Budget, error)
	FindByID(id uint, userID uint) (*entity.Budget, error)
	Delete(id uint, userID uint) error
}

type usecase struct {
	repo budgetRepo.Repository
}

func NewUsecase(repo budgetRepo.Repository) Usecase {
	return &usecase{repo}
}

func (u *usecase) Create(userID uint, amount float64, month int, year int) (*entity.Budget, error) {
	budget := &entity.Budget{
		UserID: userID,
		Amount: amount,
		Month:  month,
		Year:   year,
	}

	err := u.repo.Create(budget)
	if err != nil {
		return nil, err
	}

	return budget, nil
}

func (u *usecase) FindAll(userID uint) ([]*entity.Budget, error) {
	return u.repo.FindAll(userID)
}

func (u *usecase) FindByID(id uint, userID uint) (*entity.Budget, error) {
	return u.repo.FindByID(id, userID)
}

func (u *usecase) Delete(id uint, userID uint) error {
	_, err := u.repo.FindByID(id, userID)
	if err != nil {
		return err
	}

	return u.repo.Delete(id, userID)
}
