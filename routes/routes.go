package routes

import (
	"context"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/nahumsa/streaming-pipeline-clickhouse/event"
	"github.com/nahumsa/streaming-pipeline-clickhouse/repositories"
)

func SetupRoutes(app *fiber.App, eventRepo repositories.EventRepository) {
	app.Post("/api/v1/sendEvent", func(c *fiber.Ctx) error {
		var req event.RequestEvent

		if err := c.BodyParser(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "cannot parse JSON",
			})
		}

		ctx := context.Background()
		if err := eventRepo.InsertEvent(ctx, req.Event); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": fmt.Sprintf("failed to insert event %s", err),
			})
		}

		return c.Status(fiber.StatusAccepted).JSON(req)
	})
}
