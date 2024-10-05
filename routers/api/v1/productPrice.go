package v1

import (
	productPrice_service "be-hoatieu/services/productPrice"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type ProductPrice struct {
	Title   string `gorm:"column:title" json:"title"`
	Url     string `gorm:"column:url" json:"url"`
	Image   string `gorm:"column:image" json:"image"`
	Status  bool   `gorm:"column:status" json:"status"`
	Content string `gorm:"column:content" json:"content"`
}

func GetProductPrice_Component(c *fiber.Ctx) error {
	item := productPrice_service.ProductPrice{}

	limit := c.Query("limit")
	page := c.Query("page")

	limitStr, err := strconv.Atoi(limit)
	PageStr, err := strconv.Atoi(page)

	data, totalRecords, err := item.GetProductPrice_Service(limitStr, PageStr)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success":      true,
		"data":         data,
		"totalRecords": totalRecords,
	})

}
