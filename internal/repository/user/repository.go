package user

import (
	"be-catatin/internal/entity"

	"gorm.io/gorm"
)

type Repository interface {
	FindByPhoneOrUsername(phone string, username string) (*entity.User, error)
	Create(user *entity.User) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db}
}

func (r *repository) FindByPhoneOrUsername(phone string, username string) (*entity.User, error) {
	var user entity.User
	err := r.db.Where("phone_number = ? OR username = ?", phone, username).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil // Return nil, nil when not found to easily handle creation
		}
		return nil, err
	}
	return &user, nil
}

func (r *repository) Create(user *entity.User) error {
	return r.db.Create(user).Error
}
