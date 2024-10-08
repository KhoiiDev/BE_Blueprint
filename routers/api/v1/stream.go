package v1

import (
	stream_service "be-hoatieu/services/stream"
	"time"

	"github.com/gofiber/fiber/v2"
)

type Stream struct {
	Url      string    `gorm:"column:url" json:"url"`
	Title    string    `gorm:"column:title" json:"title"`
	PostDate time.Time `gorm:"column:created_at" json:"createdAt"`
	Status   bool      `gorm:"column:status" json:"status"`
}

func GetStream_Component(c *fiber.Ctx) error {
	item := stream_service.Stream{}

	data, err := item.GetStream_Service()

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
