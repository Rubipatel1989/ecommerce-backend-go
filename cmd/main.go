package main

import (
	"ecommerce-backend-go/db"
	"ecommerce-backend-go/routes"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("../.env")

	if err != nil {
		log.Fatal("Error loading .env")
	}

	db.Init()

	engine := html.New("../templates", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	routes.Setup(app)

	log.Fatal(app.Listen(":3000"))
}
