package v1

import (
	news_service "be-hoatieu/services/news"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

type News struct {
	Title    string    `gorm:"column:title" json:"title"`
	Url      string    `gorm:"column:url" json:"url"`
	Image    string    `gorm:"column:image" json:"image"`
	Status   bool      `gorm:"column:status" json:"status"`
	Content  string    `gorm:"column:content" json:"content"`
	PostDate time.Time `gorm:"column:postdate" json:"postdate"`
}

func GetNews_Component(c *fiber.Ctx) error {
	item := news_service.News{}

	limit := c.Query("limit")

	limitStr, err := strconv.Atoi(limit)

	data, err := item.GetNews_Service(limitStr)

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
