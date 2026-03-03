package category

import (
	"be-catatin/internal/entity"
	categoryRepo "be-catatin/internal/repository/category"
)

type Usecase interface {
	Create(userID uint, name string) (*entity.Category, error)
	FindAll(userID uint) ([]*entity.Category, error)
	FindByID(id uint, userID uint) (*entity.Category, error)
	Delete(id uint, userID uint) error
}

type usecase struct {
	repo categoryRepo.Repository
}

func NewUsecase(repo categoryRepo.Repository) Usecase {
	return &usecase{repo}
}

func (u *usecase) Create(userID uint, name string) (*entity.Category, error) {
	category := &entity.Category{
		UserID: userID,
		Name:   name,
	}

	err := u.repo.Create(category)
	if err != nil {
		return nil, err
	}

	return category, nil
}

func (u *usecase) FindAll(userID uint) ([]*entity.Category, error) {
	return u.repo.FindAll(userID)
}

func (u *usecase) FindByID(id uint, userID uint) (*entity.Category, error) {
	return u.repo.FindByID(id, userID)
}

func (u *usecase) Delete(id uint, userID uint) error {
	// Optional: verify if it exists first
	_, err := u.repo.FindByID(id, userID)
	if err != nil {
		return err // e.g. record not found
	}

	return u.repo.Delete(id, userID)
}
