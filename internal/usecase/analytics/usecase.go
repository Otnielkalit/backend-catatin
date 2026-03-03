package analytics

import (
	"be-catatin/internal/model"
	analyticsRepo "be-catatin/internal/repository/analytics"
	"math"
)

type Usecase interface {
	GetExpenseAnalytics(userID uint, month int, year int) ([]model.ExpenseAnalyticsResponse, error)
}

type usecase struct {
	repo analyticsRepo.Repository
}

func NewUsecase(repo analyticsRepo.Repository) Usecase {
	return &usecase{repo}
}

// Pre-defined palette for chart colors (Hex strings)
var colorPalette = []string{
	"#FF5733", "#33FF57", "#3357FF", "#FF33A8", "#A833FF",
	"#33FFF0", "#F0FF33", "#FF8C33", "#8CFF33", "#338CFF",
}

func (u *usecase) GetExpenseAnalytics(userID uint, month int, year int) ([]model.ExpenseAnalyticsResponse, error) {
	categoryTotals, err := u.repo.GetExpenseAnalytics(userID, month, year)
	if err != nil {
		return nil, err
	}

	// 1. Calculate overall total for the month
	var grandTotal float64 = 0
	for _, ct := range categoryTotals {
		grandTotal += ct.TotalAmount
	}

	// 2. Build response with percentages and distinct colors
	var responses []model.ExpenseAnalyticsResponse
	for i, ct := range categoryTotals {
		// Calculate Percentage out of 100, rounded to 2 decimals
		var percentage float64 = 0
		if grandTotal > 0 {
			percentage = math.Round((ct.TotalAmount/grandTotal)*100*100) / 100
		}

		// Pick a distinct color from palette (loop back if categories > palette length)
		color := colorPalette[i%len(colorPalette)]

		responses = append(responses, model.ExpenseAnalyticsResponse{
			CategoryID:   ct.CategoryID,
			CategoryName: ct.CategoryName,
			TotalAmount:  ct.TotalAmount,
			Percentage:   percentage,
			Color:        color,
		})
	}

	// Fallback empty array instead of nil if no data
	if len(responses) == 0 {
		return []model.ExpenseAnalyticsResponse{}, nil
	}

	return responses, nil
}
