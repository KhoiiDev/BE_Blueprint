package v1

import (
	product_service "be-hoatieu/services/product"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type Product struct {
	Title   string `gorm:"column:title" json:"title"`
	Image   string `gorm:"column:image" json:"image"`
	Url     string `gorm:"column:url" json:"url"`
	Content string `gorm:"column:content" json:"content"`
	Status  bool   `gorm:"column:status" json:"status"`
}

func GetServiceProduct_Component(c *fiber.Ctx) error {
	item := product_service.Product{}

	limit := c.Query("limit")

	limitStr, err := strconv.Atoi(limit)

	data, err := item.GetProduct_Service(limitStr)

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
