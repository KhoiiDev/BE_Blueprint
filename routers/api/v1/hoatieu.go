package v1

import (
	hoatieu_service "be-hoatieu/services/HoaTieu"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type Hoatieu struct {
	Status bool   `gorm:"column:status;" json:"status"`
	Rank   string `gorm:"column:rank" json:"rank"`
	Image  string `gorm:"column:image" json:"image"`
	Name   string `gorm:"column:name" json:"name"`
}

func GetAllNavigator_Component(c *fiber.Ctx) error {
	item := hoatieu_service.Hoatieu{}
	limit := c.Query("limit")
	page := c.Query("page")
	showHiddenItem := c.Query("showHiddenItem")

	limitStr, err := strconv.Atoi(limit)
	PageStr, err := strconv.Atoi(page)
	showHidden, err := strconv.ParseBool(showHiddenItem)

	data, totalRecords, err := item.GetAllNavigator_Service(limitStr, PageStr, showHidden)

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

func CreateNavigator_Component(c *fiber.Ctx) error {
	form := &Hoatieu{}

	// Check, if received JSON data is valid.
	if err := c.BodyParser(form); err != nil {
		// Return status 400 and error message.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	HoaTieuService := hoatieu_service.Hoatieu{
		Name:   form.Name,
		Image:  form.Image,
		Status: form.Status,
		Rank:   form.Rank,
	}
	if err := HoaTieuService.CreateNavigator_Service(); err != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"error": true,
			"msg":   "Thêm mới không thành công",
		})
	}

	data := make(map[string]string)
	data["Name"] = form.Name
	data["Image"] = form.Image
	data["Rank"] = form.Rank
	data["Status"] = strconv.FormatBool(form.Status)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data":    data,
	})
}

func UpdateNavigator_Component(c *fiber.Ctx) error {
	form := &Hoatieu{}
	if err := c.BodyParser(form); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	HoaTieuService := hoatieu_service.Hoatieu{
		Name:   form.Name,
		Image:  form.Image,
		Status: form.Status,
		Rank:   form.Rank,
	}

	err := HoaTieuService.UpdateNavigator_Service(c.Params("id"))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   "Register false",
		})
	}

	return c.Status(fiber.StatusOK).JSON(HoaTieuService)
}

func DeleteNavigator_Component(c *fiber.Ctx) error {
	item := hoatieu_service.Hoatieu{}

	data, err := item.DeleteNavigator_Service(c.Params("id"))
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
