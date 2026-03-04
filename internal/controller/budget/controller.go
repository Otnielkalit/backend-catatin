package budget

import (
	"strconv"

	"be-catatin/internal/model"
	budgetUsecase "be-catatin/internal/usecase/budget"
	"be-catatin/pkg/response"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validate = validator.New()

type Controller struct {
	usecase budgetUsecase.Usecase
}

func NewController(usecase budgetUsecase.Usecase) *Controller {
	return &Controller{usecase}
}

func (c *Controller) Create(ctx *fiber.Ctx) error {
	var req model.CreateBudgetRequest
	if err := ctx.BodyParser(&req); err != nil {
		return response.Error(ctx, fiber.StatusBadRequest, "invalid request body")
	}

	if err := validate.Struct(&req); err != nil {
		return response.Error(ctx, fiber.StatusBadRequest, err.Error())
	}

	budget, err := c.usecase.Create(req.UserID, req.Amount, req.Month, req.Year)
	if err != nil {
		return response.Error(ctx, fiber.StatusInternalServerError, err.Error())
	}

	return response.Success(ctx, fiber.StatusCreated, budget)
}

func (c *Controller) FindAll(ctx *fiber.Ctx) error {
	var req model.GetBudgetRequest
	if err := ctx.BodyParser(&req); err != nil {
		// Ignore body parser error as this is a GET request and parameters might be in query
	}

	if err := ctx.QueryParser(&req); err != nil {
		return response.Error(ctx, fiber.StatusBadRequest, "invalid query parameters")
	}

	if err := validate.Struct(&req); err != nil {
		return response.Error(ctx, fiber.StatusBadRequest, err.Error())
	}

	budgets, err := c.usecase.FindAll(req.UserID, req.Month, req.Year)
	if err != nil {
		return response.Error(ctx, fiber.StatusInternalServerError, err.Error())
	}

	return response.Success(ctx, fiber.StatusOK, budgets)
}

func (c *Controller) FindByID(ctx *fiber.Ctx) error {
	idStr := ctx.Params("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		return response.Error(ctx, fiber.StatusBadRequest, "invalid budget id")
	}

	var req model.GetBudgetRequest
	if err := ctx.BodyParser(&req); err != nil {
		return response.Error(ctx, fiber.StatusBadRequest, "invalid request body")
	}

	if err := validate.Struct(&req); err != nil {
		return response.Error(ctx, fiber.StatusBadRequest, err.Error())
	}

	budget, err := c.usecase.FindByID(uint(id), req.UserID)
	if err != nil {
		return response.Error(ctx, fiber.StatusNotFound, "budget not found")
	}

	return response.Success(ctx, fiber.StatusOK, budget)
}

func (c *Controller) Delete(ctx *fiber.Ctx) error {
	idStr := ctx.Params("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		return response.Error(ctx, fiber.StatusBadRequest, "invalid budget id")
	}

	var req model.GetBudgetRequest
	if err := ctx.BodyParser(&req); err != nil {
		return response.Error(ctx, fiber.StatusBadRequest, "invalid request body")
	}

	if err := validate.Struct(&req); err != nil {
		return response.Error(ctx, fiber.StatusBadRequest, err.Error())
	}

	err = c.usecase.Delete(uint(id), req.UserID)
	if err != nil {
		return response.Error(ctx, fiber.StatusInternalServerError, err.Error())
	}

	return response.Success(ctx, fiber.StatusOK, nil)
}
