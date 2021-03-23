package Router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kriangkrai/GoFiberSQLLogin/Controller"
)

func Gets(c *fiber.Ctx) error {
	data, err := Controller.GetsData()
	if err != nil {
		return nil
	}
	return c.JSON(data)
}
