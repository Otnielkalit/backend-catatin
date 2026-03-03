package entity

type Budget struct {
	ID     uint    `gorm:"primaryKey" json:"id"`
	UserID uint    `gorm:"not null" json:"user_id"`
	Amount float64 `gorm:"type:decimal(15,2);not null" json:"amount"`
	Month  int     `gorm:"not null" json:"month"`
	Year   int     `gorm:"not null" json:"year"`
	User   *User   `gorm:"foreignKey:UserID" json:"user,omitempty"`
}
