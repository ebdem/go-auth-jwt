package main

import (
	"github.com/ebdem/golang/go-auth/database"
	"github.com/ebdem/golang/go-auth/routes"
	"github.com/gofiber/fiber"
	"github.com/gofiber/fiber/middleware/cors"
)

func main() {
	// Initialize database
	database.Connect()

	// Initialize fiber
	app := fiber.New()

	// Fix CORS
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	// Initialize routes
	routes.Setup(app)

	err := app.Listen(":8000")
	if err != nil {
		return
	}
}
