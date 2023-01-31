package main

import (
    "log"

    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()
    app.Use(cors.New())
    database.ConnectDB()
	defer database.DB.Close()
	
    api := app.Group("/api")
	url.Register(api, database.DB)

    log.Fatal(app.Listen(":4000"))
}