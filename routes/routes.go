package routes

import (
	"context"
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/nahumsa/streaming-pipeline-clickhouse/event"
	"github.com/nahumsa/streaming-pipeline-clickhouse/repositories"
)

type (
	ErrorResponse struct {
		Error       bool
		FailedField string
		Tag         string
		Value       interface{}
	}

	GlobalErrorHandlerResp struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
	}
)

var validate *validator.Validate

func SetupRoutes(app *fiber.App, eventRepo repositories.EventRepository) {
	app.Post("/api/v1/sendEvent", func(c *fiber.Ctx) error {
		var req event.RequestEvent

		if err := c.BodyParser(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error":  "cannot parse JSON",
				"detail": fmt.Sprintf("%s", err),
			})
		}

		validate = validator.New(validator.WithRequiredStructEnabled())
		err := validate.Struct(req)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error":  "invalid JSON",
				"detail": fmt.Sprintf("%s", err),
			})
		}

		ctx := context.Background()
		if err := eventRepo.InsertEvent(ctx, req.Event); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": fmt.Sprintf("failed to insert event %s", err),
			})
		}

		return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": "success"})
	})
}
