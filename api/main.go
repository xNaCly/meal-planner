package main

import (
	"errors"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/xnacly/meal-planner/api/models"
	"github.com/xnacly/meal-planner/api/routes"
)

func main() {
	databaseUrl := os.Getenv("MEAL_DATABASE")
	if len(databaseUrl) == 0 {
		log.Fatalln("Missing MEAL_DATABASE env variable, required for database access")
	}
	app := fiber.New(fiber.Config{
		AppName:      "meal-planner",
		ServerHeader: "meal-planner",
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError

			var e *fiber.Error
			if errors.As(err, &e) {
				code = e.Code
			}

			return c.Status(code).JSON(models.ApiResponse{
				Success: false,
				Code:    code,
				Message: err.Error(),
				Data:    nil,
			})
		},
	})
	app.Use(cors.New(cors.Config{
		AllowMethods:     "",
		Next:             nil,
		AllowOrigins:     "*",
		AllowHeaders:     "",
		AllowCredentials: false,
		ExposeHeaders:    "",
		MaxAge:           0,
	}))

	app.Use(logger.New(logger.Config{
		TimeFormat: "2006-01-02 15:04:05",
	}))

	routes.RegisterRoutes(app, routes.Routes...)

	app.Use(func(c *fiber.Ctx) error {
		return fiber.NewError(404, "Requested ressource not found")
	})

	log.Fatal(app.Listen(":8080"))
}
