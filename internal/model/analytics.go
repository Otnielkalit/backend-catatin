package model

type ExpenseAnalyticsRequest struct {
	UserID uint `json:"user_id" form:"user_id" validate:"required"`
	Month  int  `json:"month" form:"month" validate:"required,min=1,max=12"`
	Year   int  `json:"year" form:"year" validate:"required,min=2000"`
}

type ExpenseAnalyticsResponse struct {
	CategoryID   uint    `json:"category_id"`
	CategoryName string  `json:"category_name"`
	TotalAmount  float64 `json:"total_amount"`
	Percentage   float64 `json:"percentage"`
	Color        string  `json:"color"`
}
