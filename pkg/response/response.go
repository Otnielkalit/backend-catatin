package response

import "github.com/gofiber/fiber/v2"

type BaseResponse struct {
	Message string      `json:"message"`
	Status  int         `json:"status"`
	Data    interface{} `json:"data"`
}

func Success(c *fiber.Ctx, status int, data interface{}) error {
	return c.Status(status).JSON(BaseResponse{
		Message: "success",
		Status:  status,
		Data:    data,
	})
}

func Error(c *fiber.Ctx, status int, message string) error {
	return c.Status(status).JSON(BaseResponse{
		Message: message,
		Status:  status,
		Data:    map[string]interface{}{},
	})
}
