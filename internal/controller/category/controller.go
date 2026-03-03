package category

import (
	"strconv"

	"be-catatin/internal/model"
	categoryUsecase "be-catatin/internal/usecase/category"
	"be-catatin/pkg/response"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validate = validator.New()

type Controller struct {
	usecase categoryUsecase.Usecase
}

func NewController(usecase categoryUsecase.Usecase) *Controller {
	return &Controller{usecase}
}

func (c *Controller) Create(ctx *fiber.Ctx) error {
	var req model.CreateCategoryRequest
	if err := ctx.BodyParser(&req); err != nil {
		return response.Error(ctx, fiber.StatusBadRequest, "invalid request body")
	}

	if err := validate.Struct(&req); err != nil {
		return response.Error(ctx, fiber.StatusBadRequest, err.Error())
	}

	category, err := c.usecase.Create(req.UserID, req.Name)
	if err != nil {
		return response.Error(ctx, fiber.StatusInternalServerError, err.Error())
	}

	return response.Success(ctx, fiber.StatusCreated, category)
}

func (c *Controller) FindAll(ctx *fiber.Ctx) error {
	var req struct {
		UserID uint `json:"user_id" form:"user_id"`
	}

	if err := ctx.BodyParser(&req); err != nil {
		return response.Error(ctx, fiber.StatusBadRequest, "invalid request body")
	}

	if req.UserID == 0 {
		return response.Error(ctx, fiber.StatusBadRequest, "user_id is required")
	}

	categories, err := c.usecase.FindAll(req.UserID)
	if err != nil {
		return response.Error(ctx, fiber.StatusInternalServerError, err.Error())
	}

	return response.Success(ctx, fiber.StatusOK, categories)
}

func (c *Controller) FindByID(ctx *fiber.Ctx) error {
	idStr := ctx.Params("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		return response.Error(ctx, fiber.StatusBadRequest, "invalid category id")
	}

	var req struct {
		UserID uint `json:"user_id" form:"user_id"`
	}

	if err := ctx.BodyParser(&req); err != nil {
		return response.Error(ctx, fiber.StatusBadRequest, "invalid request body")
	}

	if req.UserID == 0 {
		return response.Error(ctx, fiber.StatusBadRequest, "user_id is required")
	}

	category, err := c.usecase.FindByID(uint(id), req.UserID)
	if err != nil {
		return response.Error(ctx, fiber.StatusNotFound, "category not found")
	}

	return response.Success(ctx, fiber.StatusOK, category)
}

func (c *Controller) Delete(ctx *fiber.Ctx) error {
	idStr := ctx.Params("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		return response.Error(ctx, fiber.StatusBadRequest, "invalid category id")
	}

	var req struct {
		UserID uint `json:"user_id" form:"user_id"`
	}

	if err := ctx.BodyParser(&req); err != nil {
		return response.Error(ctx, fiber.StatusBadRequest, "invalid request body")
	}

	if req.UserID == 0 {
		return response.Error(ctx, fiber.StatusBadRequest, "user_id is required")
	}

	err = c.usecase.Delete(uint(id), req.UserID)
	if err != nil {
		return response.Error(ctx, fiber.StatusInternalServerError, err.Error())
	}

	return response.Success(ctx, fiber.StatusOK, nil)
}
