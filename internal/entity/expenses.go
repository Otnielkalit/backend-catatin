package entity

import "time"

type Expense struct {
	ID              uint      `gorm:"primaryKey" json:"id"`
	CategoryID      uint      `gorm:"not null" json:"category_id"`
	UserID          uint      `gorm:"not null" json:"user_id"`
	Title           string    `gorm:"type:varchar(255);not null" json:"title"`
	Amount          float64   `gorm:"type:decimal(15,2);not null" json:"amount"`
	TransactionDate time.Time `gorm:"type:date;not null" json:"transaction_date"`
	ImgPath         string    `gorm:"type:text;column:img_path" json:"img_path"`
	Category        *Category `gorm:"foreignKey:CategoryID" json:"category,omitempty"`
	User            *User     `gorm:"foreignKey:UserID" json:"user,omitempty"`
}
