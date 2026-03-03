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

	// Read from JSON Body or Query Params
	// For analytics, GET request typically uses Query string or Body depending on the architecture.
	// Since standard practice here used BodyParser for UserID, we'll continue using it for ease.
	if err := ctx.BodyParser(&req); err != nil {
		// Try Query if Body isn't provided (flexible)
		month, _ := strconv.Atoi(ctx.Query("month"))
		year, _ := strconv.Atoi(ctx.Query("year"))
		userID, _ := strconv.ParseUint(ctx.Query("user_id"), 10, 32)
		if month > 0 && year > 0 && userID > 0 {
			req.Month = month
			req.Year = year
			req.UserID = uint(userID)
		} else {
			return response.Error(ctx, fiber.StatusBadRequest, "invalid request format")
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
