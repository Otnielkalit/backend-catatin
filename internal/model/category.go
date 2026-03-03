package model

type CreateCategoryRequest struct {
	UserID uint   `json:"user_id" form:"user_id" validate:"required"`
	Name   string `json:"name" form:"name" validate:"required"`
}
