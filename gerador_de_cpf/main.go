package main

import (
"github.com/gofiber/fiber/v2"
"fmt"
)
func main() {
    app := fiber.New()

    app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("Hello, World!")
    })
	fmt.Println("yeah")
    app.Listen(":3000")
}