package user

import (
	"be-catatin/internal/model"
	userUsecase "be-catatin/internal/usecase/user"
	"be-catatin/pkg/response"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validate = validator.New()

type Controller struct {
	usecase userUsecase.Usecase
}

func NewController(usecase userUsecase.Usecase) *Controller {
	return &Controller{usecase}
}

func (c *Controller) Login(ctx *fiber.Ctx) error {
	var req model.LoginRequest
	if err := ctx.BodyParser(&req); err != nil {
		return response.Error(ctx, fiber.StatusBadRequest, "invalid request body")
	}

	if err := validate.Struct(&req); err != nil {
		return response.Error(ctx, fiber.StatusBadRequest, err.Error())
	}

	user, err := c.usecase.LoginOrCreate(req.PhoneNumber, req.Username, req.Pin)
	if err != nil {
		return response.Error(ctx, fiber.StatusInternalServerError, err.Error())
	}

	return response.Success(ctx, fiber.StatusOK, user)
}
