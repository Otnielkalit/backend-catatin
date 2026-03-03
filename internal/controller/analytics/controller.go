package analytics

import (
	"be-catatin/internal/model"
	analyticsUsecase "be-catatin/internal/usecase/analytics"
	"be-catatin/pkg/response"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validate = validator.New()

type Controller struct {
	usecase analyticsUsecase.Usecase
}

func NewController(usecase analyticsUsecase.Usecase) *Controller {
	return &Controller{usecase}
}

func (c *Controller) GetExpenseAnalytics(ctx *fiber.Ctx) error {
	var req model.ExpenseAnalyticsRequest

	if err := ctx.BodyParser(&req); err != nil {
		return response.Error(ctx, fiber.StatusBadRequest, "invalid request body")
	}

	// Read month and year from Query Params
	monthStr := ctx.Query("month")
	yearStr := ctx.Query("year")

	if monthStr != "" {
		if month, err := strconv.Atoi(monthStr); err == nil {
			req.Month = month
		}
	}

	if yearStr != "" {
		if year, err := strconv.Atoi(yearStr); err == nil {
			req.Year = year
		}
	}

	if err := validate.Struct(&req); err != nil {
		return response.Error(ctx, fiber.StatusBadRequest, err.Error())
	}

	result, err := c.usecase.GetExpenseAnalytics(req.UserID, req.Month, req.Year)
	if err != nil {
		return response.Error(ctx, fiber.StatusInternalServerError, err.Error())
	}

	return response.Success(ctx, fiber.StatusOK, result)
}
