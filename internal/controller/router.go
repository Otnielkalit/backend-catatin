package controller

import (
	budgetController "be-catatin/internal/controller/budget"
	categoryController "be-catatin/internal/controller/category"
	expenseController "be-catatin/internal/controller/expense"
	userController "be-catatin/internal/controller/user"

	budgetRepo "be-catatin/internal/repository/budget"
	categoryRepo "be-catatin/internal/repository/category"
	expenseRepo "be-catatin/internal/repository/expense"
	userRepo "be-catatin/internal/repository/user"

	budgetUsecase "be-catatin/internal/usecase/budget"
	categoryUsecase "be-catatin/internal/usecase/category"
	expenseUsecase "be-catatin/internal/usecase/expense"
	userUsecase "be-catatin/internal/usecase/user"

	"be-catatin/pkg/cloudinary"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SetupRoutes(app *fiber.App, db *gorm.DB) {
	// Initialize Cloudinary
	cloudinarySvc := cloudinary.NewCloudinaryService()

	api := app.Group("/api/v1")

	// Dependencies Injection untuk User
	userRepository := userRepo.NewRepository(db)
	userUcase := userUsecase.NewUsecase(userRepository)
	userCtrl := userController.NewController(userUcase)

	// Dependencies Injection untuk Category
	categoryRepository := categoryRepo.NewRepository(db)
	categoryUcase := categoryUsecase.NewUsecase(categoryRepository)
	categoryCtrl := categoryController.NewController(categoryUcase)

	// Dependencies Injection untuk Budget
	budgetRepository := budgetRepo.NewRepository(db)
	budgetUcase := budgetUsecase.NewUsecase(budgetRepository)
	budgetCtrl := budgetController.NewController(budgetUcase)

	// Dependencies Injection untuk Expense
	expenseRepository := expenseRepo.NewRepository(db)
	expenseUcase := expenseUsecase.NewUsecase(expenseRepository, cloudinarySvc)
	expenseCtrl := expenseController.NewController(expenseUcase)

	// Routes
	userGroup := api.Group("/users")
	userGroup.Post("/login", userCtrl.Login)

	categoryGroup := api.Group("/categories")
	categoryGroup.Post("/", categoryCtrl.Create)
	categoryGroup.Get("/", categoryCtrl.FindAll)
	categoryGroup.Get("/:id", categoryCtrl.FindByID)
	categoryGroup.Delete("/:id", categoryCtrl.Delete)

	budgetGroup := api.Group("/budgets")
	budgetGroup.Post("/", budgetCtrl.Create)
	budgetGroup.Get("/", budgetCtrl.FindAll)
	budgetGroup.Get("/:id", budgetCtrl.FindByID)
	budgetGroup.Delete("/:id", budgetCtrl.Delete)

	expenseGroup := api.Group("/expenses")
	expenseGroup.Post("/", expenseCtrl.Create)
	expenseGroup.Get("/", expenseCtrl.FindAll)
	expenseGroup.Get("/:id", expenseCtrl.FindByID)
	expenseGroup.Delete("/:id", expenseCtrl.Delete)
}
