package entity

type Category struct {
	ID     uint   `gorm:"primaryKey" json:"id"`
	UserID uint   `gorm:"not null" json:"user_id"`
	Name   string `gorm:"type:varchar(100);not null" json:"name"`
	User   *User  `gorm:"foreignKey:UserID" json:"user,omitempty"`
}
