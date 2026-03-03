package expense

import (
	"be-catatin/internal/entity"

	"gorm.io/gorm"
)

type Repository interface {
	Create(expense *entity.Expense) error
	FindAll(userID uint) ([]*entity.Expense, error)
	FindByID(id uint, userID uint) (*entity.Expense, error)
	Delete(id uint, userID uint) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db}
}

func (r *repository) Create(expense *entity.Expense) error {
	return r.db.Create(expense).Error
}

func (r *repository) FindAll(userID uint) ([]*entity.Expense, error) {
	var expenses []*entity.Expense
	err := r.db.Where("user_id = ?", userID).Preload("Category").Find(&expenses).Error
	return expenses, err
}

func (r *repository) FindByID(id uint, userID uint) (*entity.Expense, error) {
	var expense entity.Expense
	err := r.db.Where("id = ? AND user_id = ?", id, userID).Preload("Category").First(&expense).Error
	return &expense, err
}

func (r *repository) Delete(id uint, userID uint) error {
	return r.db.Where("id = ? AND user_id = ?", id, userID).Delete(&entity.Expense{}).Error
}
