package v1

import (
	maneuveringDraft_service "be-hoatieu/services/maneuveringDraft"

	"github.com/gofiber/fiber/v2"
)

type ManeuveringDraft struct {
	Image  string `gorm:"column:image" json:"image"`
	Status bool   `gorm:"column:status" json:"status"`
}

func GetManeuveringDraft_Component(c *fiber.Ctx) error {
	item := maneuveringDraft_service.ManeuveringDraft{}

	data, err := item.GetManeuveringDraft_Service()

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
