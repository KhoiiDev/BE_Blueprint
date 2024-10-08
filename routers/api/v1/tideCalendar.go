package v1

import (
	TideCalendar_service "be-hoatieu/services/tideCalendar"
	"time"

	"github.com/gofiber/fiber/v2"
)

type TideCalendar struct {
	Url      string    `gorm:"column:url" json:"url"`
	PostDate time.Time `gorm:"column:created_at" json:"createdAt"`
	Status   bool      `gorm:"column:status" json:"status"`
}

func GetTideCalendar_Component(c *fiber.Ctx) error {
	item := TideCalendar_service.TideCalendar{}

	data, err := item.GetTideCalendar_Service()

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
