package v1

import (
	switch_service "be-hoatieu/services/switch"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type Switch struct {
	Flag bool `gorm:"column:flag;default:false" json:"flag"`
}

func GetSwitch_Component(c *fiber.Ctx) error {
	switchItem := switch_service.Switch{}
	idStr := c.Params("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   "Invalid ID format",
		})
	}

	data, err := switchItem.GetSwitch_Service(uint(id))
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

func CreateSwitch_Component(c *fiber.Ctx) error {
	form := &Switch{}

	// Check if received JSON data is valid
	if err := c.BodyParser(form); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	SwitchService := switch_service.Switch{
		Flag: form.Flag,
	}
	if err := SwitchService.CreateSwitch_Service(); err != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"error": true,
			"msg":   "Thêm mới không thành công",
		})
	}

	data := make(map[string]string)
	data["Flag"] = strconv.FormatBool(form.Flag)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data":    data,
	})
}

func UpdateSwitch_Component(c *fiber.Ctx) error {
	form := &Switch{}
	if err := c.BodyParser(form); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	SwitchService := switch_service.Switch{
		Flag: form.Flag,
	}

	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   "Invalid ID format",
		})
	}

	err = SwitchService.UpdateSwitch_Service(uint(id))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   "Update failed",
		})
	}

	return c.Status(fiber.StatusOK).JSON(SwitchService)
}

func DeleteSwitch_Component(c *fiber.Ctx) error {
	switchItem := switch_service.Switch{}

	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   "Invalid ID format",
		})
	}

	data, err := switchItem.DeleteSwitch_Service(uint(id))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "Delete failed: " + err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data":    data,
	})
}
