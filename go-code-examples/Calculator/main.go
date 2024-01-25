package main

import (
	function "calculator/functions"
	"strconv"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Post("/submit", func(c *fiber.Ctx) error {
		num1, err := strconv.ParseFloat(c.FormValue("num1"), 64)
		if err != nil {
			return err
		}
		num2, err := strconv.ParseFloat(c.FormValue("num2"), 64)
		if err != nil {
			return err
		}
		operator := c.FormValue("operator", "+")
		result := function.Calculator(num1, num2, operator)
		return c.SendString(strconv.FormatFloat(result, 'f', -1, 64))
	})

	app.Listen(":3003")
}
