package v1

import (
	introduction_service "be-hoatieu/services/introduction"

	"github.com/gofiber/fiber/v2"
)

type Introduction struct {
	Content string `gorm:"column:content" json:"content"`
	Status  bool   `gorm:"column:status" json:"status"`
}

func GetIntroduction_Component(c *fiber.Ctx) error {
	item := introduction_service.Introduction{}

	data, err := item.GetIntroduction_Service()

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data":    data,
	})
}
