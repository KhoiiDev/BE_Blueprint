package v1

import (
	TideCalendar_service "be-hoatieu/services/tideCalendar"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type TideCalendar struct {
	Pdfurl   string `gorm:"column:pdfurl" json:"pdfurl"`
	PostDate string `gorm:"column:postdate" json:"postdate"`
	Title    string `gorm:"column:title" json:"title"`

	Status bool `gorm:"column:status" json:"status"`
}

func GetTideCalendar_Component(c *fiber.Ctx) error {
	item := TideCalendar_service.TideCalendar{}

	limit := c.Query("limit")
	page := c.Query("page")
	showHiddenItem := c.Query("showHiddenItem")
	date := c.Query("date")
	name := c.Query("name")

	limitStr, err := strconv.Atoi(limit)
	PageStr, err := strconv.Atoi(page)
	showHidden, err := strconv.ParseBool(showHiddenItem)

	data, totalRecords, err := item.GetTideCalendar_Service(limitStr, PageStr, showHidden, name, date)

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

func CreateTideCalendar_Component(c *fiber.Ctx) error {
	form := &TideCalendar{}
	// Check, if received JSON data is valid.
	if err := c.BodyParser(form); err != nil {
		// Return status 400 and error message.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	TideCalendarService := TideCalendar_service.TideCalendar{
		Pdfurl:   form.Pdfurl,
		Status:   form.Status,
		PostDate: form.PostDate,
	}
	if err := TideCalendarService.CreateTideCalendar_Service(); err != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"error": true,
			"msg":   "Thêm mới không thành công",
		})
	}

	data := make(map[string]string)
	data["Pdfurl"] = form.Pdfurl
	data["PostDate"] = form.PostDate
	data["Title"] = form.Title
	data["Status"] = strconv.FormatBool(form.Status)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data":    data,
	})
}

func UpdateTideCalendar_Component(c *fiber.Ctx) error {
	form := &TideCalendar{}
	if err := c.BodyParser(form); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	TideCalendarService := TideCalendar_service.TideCalendar{
		Pdfurl:   form.Pdfurl,
		Status:   form.Status,
		Title:    form.Title,
		PostDate: form.PostDate,
	}

	err := TideCalendarService.UpdateTideCalendar_Service(c.Params("id"))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   "Register false",
		})
	}

	return c.Status(fiber.StatusOK).JSON(TideCalendarService)
}

func DeleteTideCalendar_Component(c *fiber.Ctx) error {
	item := TideCalendar_service.TideCalendar{}

	data, err := item.DeleteTideCalendar_Service(c.Params("id"))
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
