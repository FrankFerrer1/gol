package main

import (
    "github.com/gofiber/fiber/v2"
)

func main() {
    inject := initializeAll()

    app := fiber.New()

    app.Get("/", func(c *fiber.Ctx) error {
        inject.Database.Stats()
        return c.SendString("Hello World!")
    })


    app.Listen(":9000")
}
