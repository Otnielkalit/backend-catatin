package expense

import (
	"strconv"

	"be-catatin/internal/model"
	expenseUsecase "be-catatin/internal/usecase/expense"
	"be-catatin/pkg/response"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validate = validator.New()

type Controller struct {
	usecase expenseUsecase.Usecase
}

func NewController(usecase expenseUsecase.Usecase) *Controller {
	return &Controller{usecase}
}

func (c *Controller) Create(ctx *fiber.Ctx) error {
	// Setup req model for validation mapping
	var req model.CreateExpenseRequest

	// Manual Form Extraction
	userIDStr := ctx.FormValue("user_id")
	categoryIDStr := ctx.FormValue("category_id")
	req.Title = ctx.FormValue("title")
	amountStr := ctx.FormValue("amount")
	req.TransactionDate = ctx.FormValue("transaction_date")

	userID, _ := strconv.ParseUint(userIDStr, 10, 32)
	categoryID, _ := strconv.ParseUint(categoryIDStr, 10, 32)
	amount, _ := strconv.ParseFloat(amountStr, 64)

	req.UserID = uint(userID)
	req.CategoryID = uint(categoryID)
	req.Amount = amount

	if err := validate.Struct(&req); err != nil {
		return response.Error(ctx, fiber.StatusBadRequest, err.Error())
	}

	// Handle File Option: Try "image_path" first, then "image", then "img_path"
	file, err := ctx.FormFile("image_path")
	if err != nil {
		file, err = ctx.FormFile("image")
		if err != nil {
			file, err = ctx.FormFile("img_path")
			if err != nil {
				// Provide a clear error if none match.
				return response.Error(ctx, fiber.StatusBadRequest, "Gambar wajib diunggah (gunakan key 'image', 'image_path', atau 'img_path')")
			}
		}
	}

	expense, err := c.usecase.Create(req.UserID, req.CategoryID, req.Title, req.Amount, req.TransactionDate, file)
	if err != nil {
		return response.Error(ctx, fiber.StatusInternalServerError, err.Error())
	}

	return response.Success(ctx, fiber.StatusCreated, expense)
}

func (c *Controller) FindAll(ctx *fiber.Ctx) error {
	var req model.GetExpenseRequest
	if err := ctx.BodyParser(&req); err != nil {
		return response.Error(ctx, fiber.StatusBadRequest, "invalid request body")
	}

	if err := validate.Struct(&req); err != nil {
		return response.Error(ctx, fiber.StatusBadRequest, err.Error())
	}

	expenses, err := c.usecase.FindAll(req.UserID)
	if err != nil {
		return response.Error(ctx, fiber.StatusInternalServerError, err.Error())
	}

	return response.Success(ctx, fiber.StatusOK, expenses)
}

func (c *Controller) FindByID(ctx *fiber.Ctx) error {
	idStr := ctx.Params("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		return response.Error(ctx, fiber.StatusBadRequest, "invalid expense id")
	}

	var req model.GetExpenseRequest
	if err := ctx.BodyParser(&req); err != nil {
		return response.Error(ctx, fiber.StatusBadRequest, "invalid request body")
	}

	if err := validate.Struct(&req); err != nil {
		return response.Error(ctx, fiber.StatusBadRequest, err.Error())
	}

	expense, err := c.usecase.FindByID(uint(id), req.UserID)
	if err != nil {
		return response.Error(ctx, fiber.StatusNotFound, "expense not found")
	}

	return response.Success(ctx, fiber.StatusOK, expense)
}

func (c *Controller) Delete(ctx *fiber.Ctx) error {
	idStr := ctx.Params("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		return response.Error(ctx, fiber.StatusBadRequest, "invalid expense id")
	}

	var req model.GetExpenseRequest
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
