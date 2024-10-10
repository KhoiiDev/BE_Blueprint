package v1

import (
	ship_service "be-hoatieu/services/ship"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type Ship struct {
	Name   string `gorm:"column:name" json:"name"`
	Image  string `gorm:"column:image" json:"image"`
	Status bool   `gorm:"column:status" json:"status"`
}

func GetShip_Component(c *fiber.Ctx) error {
	item := ship_service.Ship{}
	limit := c.Query("limit")
	page := c.Query("page")
	showHiddenItem := c.Query("showHiddenItem")

	limitStr, err := strconv.Atoi(limit)
	PageStr, err := strconv.Atoi(page)
	showHidden, err := strconv.ParseBool(showHiddenItem)

	data, totalRecords, err := item.GetShip_Service(limitStr, PageStr, showHidden)

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

func CreateShip_Component(c *fiber.Ctx) error {
	form := &Ship{}

	// Check, if received JSON data is valid.
	if err := c.BodyParser(form); err != nil {
		// Return status 400 and error message.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	ShipService := ship_service.Ship{
		Name:   form.Name,
		Image:  form.Image,
		Status: form.Status,
	}
	if err := ShipService.CreateShip_Service(); err != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"error": true,
			"msg":   "Thêm mới không thành công",
		})
	}

	data := make(map[string]string)
	data["Name"] = form.Name
	data["Image"] = form.Image
	data["Status"] = strconv.FormatBool(form.Status)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data":    data,
	})
}

func UpdateShip_Component(c *fiber.Ctx) error {
	form := &Hoatieu{}
	if err := c.BodyParser(form); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	ShipService := ship_service.Ship{
		Name:   form.Name,
		Image:  form.Image,
		Status: form.Status,
	}

	err := ShipService.UpdateShip_Service(c.Params("id"))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   "Register false",
		})
	}

	return c.Status(fiber.StatusOK).JSON(ShipService)
}

func DeleteShip_Component(c *fiber.Ctx) error {
	item := ship_service.Ship{}

	data, err := item.DeleteShip_Service(c.Params("id"))
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
