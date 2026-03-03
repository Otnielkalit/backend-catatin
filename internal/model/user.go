package model

type LoginRequest struct {
	Username    string `json:"username" form:"username" validate:"required"`
	PhoneNumber string `json:"phone_number" form:"phone_number" validate:"required"`
	Pin         string `json:"pin" form:"pin" validate:"required,len=6"`
}
