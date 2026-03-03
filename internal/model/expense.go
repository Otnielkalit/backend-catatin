package model

type CreateExpenseRequest struct {
	UserID          uint    `json:"user_id" form:"user_id" validate:"required"`
	CategoryID      uint    `json:"category_id" form:"category_id" validate:"required"`
	Title           string  `json:"title" form:"title" validate:"required"`
	Amount          float64 `json:"amount" form:"amount" validate:"required,gt=0"`
	TransactionDate string  `json:"transaction_date" form:"transaction_date" validate:"required"`
	// Image processing will happen through parsing the form file directly.
}

type GetExpenseRequest struct {
	UserID uint `json:"user_id" form:"user_id" validate:"required"`
}
