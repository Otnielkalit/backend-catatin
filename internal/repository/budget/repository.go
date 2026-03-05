package budget

import (
	"be-catatin/internal/entity"

	"gorm.io/gorm"
)

type Repository interface {
	Create(budget *entity.Budget) error
	FindAll(userID uint, month, year int) ([]*entity.Budget, error)
	FindByID(id uint, userID uint) (*entity.Budget, error)
	Delete(id uint, userID uint) error
	Update(budget *entity.Budget) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db}
}

func (r *repository) Create(budget *entity.Budget) error {
	return r.db.Create(budget).Error
}

func (r *repository) FindAll(userID uint, month, year int) ([]*entity.Budget, error) {
	var budgets []*entity.Budget
	query := r.db.Where("user_id = ?", userID)
	if month > 0 {
		query = query.Where("month = ?", month)
	}
	if year > 0 {
		query = query.Where("year = ?", year)
	}
	err := query.Find(&budgets).Error
	return budgets, err
}

func (r *repository) FindByID(id uint, userID uint) (*entity.Budget, error) {
	var budget entity.Budget
	err := r.db.Where("id = ? AND user_id = ?", id, userID).First(&budget).Error
	return &budget, err
}

func (r *repository) Delete(id uint, userID uint) error {
	return r.db.Where("id = ? AND user_id = ?", id, userID).Delete(&entity.Budget{}).Error
}

func (r *repository) Update(budget *entity.Budget) error {
	return r.db.Save(budget).Error
}
