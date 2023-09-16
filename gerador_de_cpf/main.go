package main

import (
	"github.com/gofiber/fiber/v2"
	"strconv"
)

func main() { 
	app := fiber.New()

	app.Get("/:ammount?", func(c *fiber.Ctx) error {
		if c.Params("ammount") == "" {
			return c.SendString("empty ammount")
		}
		return c.SendString(getCpf(c.Params("ammount")))
	})

    	app.Listen(":3000")
}

func getCpf(ammount string) string {

	response := ""

	ammountInt, err := strconv.Atoi(ammount)
	if err != nil {
		return "invalid ammount, use a integer"
	}

	for i := 1; i <= ammountInt; i++ {
		response += createCpf()
	}

	return response
}

func createCpf() string {

	return ""
}