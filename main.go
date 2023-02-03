package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func GetEnvConfig(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	return os.Getenv(key)
}

func main() {
	r := fiber.New()
	r.Get("/getdata/:type?", getTotalPower)
	r.Get("/setdata", setupRandomData)
	r.Listen(":8080")
}

func setupRandomData(c *fiber.Ctx) error {
	err := SetRandomData()
	if err != nil {
		return c.JSON(fiber.Map{
			"msg": err.Error(),
		})
	} else {
		return c.JSON(fiber.Map{
			"msg": "Setup Data Success",
		})
	}
}

func getTotalPower(c *fiber.Ctx) error {
	if c.Params("type") != "" {
		powertype := c.Params("type")
		pw, err := GetSumPower(powertype)
		if err != nil {
			return c.JSON(fiber.Map{
				"error": "No Such Power Type",
			})
		} else {
			return c.JSON(fiber.Map{
				powertype: pw,
			})
		}
	} else {
		return c.JSON(fiber.Map{
			"error": "Power Type Not Specified",
		})
	}
}
