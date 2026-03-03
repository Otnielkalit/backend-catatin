package analytics

import (
	"gorm.io/gorm"
)

type CategoryTotal struct {
	CategoryID   uint
	CategoryName string
	TotalAmount  float64
}

type Repository interface {
	GetExpenseAnalytics(userID uint, month int, year int) ([]CategoryTotal, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db}
}

func (r *repository) GetExpenseAnalytics(userID uint, month int, year int) ([]CategoryTotal, error) {
	var results []CategoryTotal

	// Using GORM to execute grouped query joining expenses and categories
	err := r.db.Table("expenses").
		Select("categories.id as category_id, categories.name as category_name, SUM(expenses.amount) as total_amount").
		Joins("JOIN categories ON expenses.category_id = categories.id").
		Where("expenses.user_id = ? AND EXTRACT(MONTH FROM expenses.transaction_date) = ? AND EXTRACT(YEAR FROM expenses.transaction_date) = ?", userID, month, year).
		Group("categories.id, categories.name").
		Order("total_amount DESC").
		Scan(&results).Error

	if err != nil {
		return nil, err
	}

	return results, nil
}
