package main

import (
	"ecommerce-backend-go/db"
	"ecommerce-backend-go/pkg/sessionx"
	"ecommerce-backend-go/routes"
	"html/template"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env")
	}

	// Initialize MySQL DB
	db.Init()

	// Initialize global session store
	sessionx.InitSessionStore()

	// Setup HTML template engine and auto-reload
	engine := html.New("../templates", ".html")
	engine.Reload(true) // Enables auto reload
	engine.AddFunc("safe", func(s string) template.HTML { return template.HTML(s) })

	// Parse all templates including layout/header/footer/menu
	engine.Layout("base") // Optional if using base layout template

	// Create new Fiber app with the engine
	app := fiber.New(fiber.Config{
		Views: engine,
	})
	app.Static("/static", "../public")

	// Register all routes
	routes.Setup(app)

	// Start the app
	log.Fatal(app.Listen(":3000"))
}
