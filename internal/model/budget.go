package model

type CreateBudgetRequest struct {
	UserID uint    `json:"user_id" form:"user_id" validate:"required"`
	Amount float64 `json:"amount" form:"amount" validate:"required,gt=0"`
	Month  int     `json:"month" form:"month" validate:"required,min=1,max=12"`
	Year   int     `json:"year" form:"year" validate:"required,gt=0"`
}

type GetBudgetRequest struct {
	UserID uint `json:"user_id" form:"user_id" query:"user_id" validate:"required"`
	Month  int  `json:"month" form:"month" query:"month"`
	Year   int  `json:"year" form:"year" query:"year"`
}
