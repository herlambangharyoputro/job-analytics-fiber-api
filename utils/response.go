package utils

import "github.com/gofiber/fiber/v2"

type Response struct {
    Success bool        `json:"success"`
    Message string      `json:"message"`
    Data    interface{} `json:"data,omitempty"`
    Error   string      `json:"error,omitempty"`
}

func SuccessResponse(c *fiber.Ctx, message string, data interface{}) error {
    return c.Status(fiber.StatusOK).JSON(Response{
        Success: true,
        Message: message,
        Data:    data,
    })
}

func ErrorResponse(c *fiber.Ctx, statusCode int, message string, err error) error {
    errorMsg := ""
    if err != nil {
        errorMsg = err.Error()
    }

    return c.Status(statusCode).JSON(Response{
        Success: false,
        Message: message,
        Error:   errorMsg,
    })
}

func CreatedResponse(c *fiber.Ctx, message string, data interface{}) error {
    return c.Status(fiber.StatusCreated).JSON(Response{
        Success: true,
        Message: message,
        Data:    data,
    })
}
