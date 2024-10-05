package v1

import (
	carousel_service "be-hoatieu/services/carousel"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type Carousel struct {
	Image  string `gorm:"column:image" json:"image"`
	Status bool   `gorm:"column:status" json:"status"`
}

func GetCarousel_Component(c *fiber.Ctx) error {
	item := carousel_service.Carousel{}
	limit := c.Query("limit")

	limitStr, err := strconv.Atoi(limit)

	data, err := item.GetCarousel_Service(limitStr)

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
