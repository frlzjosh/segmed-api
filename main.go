package main

import (
	"log"

	"github.com/gofiber/fiber"
	"github.com/gofiber/fiber/middleware/cors"
)

func main() {
	server := fiber.New()
	server.Use(cors.New())
	err := server.Listen(":3000")
	if err != nil {
		log.Println(err)
	}
}
