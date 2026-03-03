package entity

type User struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	Username    string `gorm:"type:varchar(100);not null;unique" json:"username"`
	PhoneNumber string `gorm:"type:varchar(20);not null;unique" json:"phone_number"`
	Pin         string `gorm:"type:varchar(255);not null" json:"-"`
}
