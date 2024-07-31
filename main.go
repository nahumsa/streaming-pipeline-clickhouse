package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/nahumsa/streaming-pipeline-clickhouse/config"
	"github.com/nahumsa/streaming-pipeline-clickhouse/repositories"
	"github.com/nahumsa/streaming-pipeline-clickhouse/routes"
)

func main() {
	app := fiber.New()
	app.Use(logger.New())

	env := config.New()

	conn, err := repositories.SetupClickhouse(env)
	if err != nil {
		log.Fatalf("failed to connect to ClickHouse: %v", err)
	}

	repo := repositories.NewClickhouseRepository(conn)

	routes.SetupRoutes(app, repo)

	log.Fatal(app.Listen(env.ServerHost))
}
