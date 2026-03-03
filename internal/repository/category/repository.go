package category

import (
	"be-catatin/internal/entity"

	"gorm.io/gorm"
)

type Repository interface {
	Create(category *entity.Category) error
	FindAll(userID uint) ([]*entity.Category, error)
	FindByID(id uint, userID uint) (*entity.Category, error)
	Delete(id uint, userID uint) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db}
}

func (r *repository) Create(category *entity.Category) error {
	return r.db.Create(category).Error
}

func (r *repository) FindAll(userID uint) ([]*entity.Category, error) {
	var categories []*entity.Category
	err := r.db.Where("user_id = ?", userID).Find(&categories).Error
	return categories, err
}

func (r *repository) FindByID(id uint, userID uint) (*entity.Category, error) {
	var category entity.Category
	err := r.db.Where("id = ? AND user_id = ?", id, userID).First(&category).Error
	return &category, err
}

func (r *repository) Delete(id uint, userID uint) error {
	return r.db.Where("id = ? AND user_id = ?", id, userID).Delete(&entity.Category{}).Error
}
