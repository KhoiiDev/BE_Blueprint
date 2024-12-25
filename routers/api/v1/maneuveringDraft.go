package v1

import (
	maneuveringDraft_service "be-hoatieu/services/maneuveringDraft"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type ManeuveringDraft struct {
	PostDate string `gorm:"column:postdate" json:"postdate"`
	Pdfurl   string `gorm:"column:pdfurl" json:"pdfurl"`
	Status   bool   `gorm:"column:status" json:"status"`
}

func GetManeuveringDraft_Component(c *fiber.Ctx) error {
	item := maneuveringDraft_service.ManeuveringDraft{}

	limit := c.Query("limit")
	page := c.Query("page")
	date := c.Query("date")
	showHiddenItem := c.Query("showHiddenItem")

	limitStr, err := strconv.Atoi(limit)
	PageStr, err := strconv.Atoi(page)
	showHidden, err := strconv.ParseBool(showHiddenItem)

	data, totalRecords, err := item.GetManeuveringDraft_Service(limitStr, PageStr, showHidden, date)

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

func CreateManeuveringDraft_Component(c *fiber.Ctx) error {
	form := &ManeuveringDraft{}
	// Check, if received JSON data is valid.
	if err := c.BodyParser(form); err != nil {
		// Return status 400 and error message.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	ManeuveringDraftService := maneuveringDraft_service.ManeuveringDraft{
		Pdfurl:  form.Pdfurl,
		Status:   form.Status,
		PostDate: form.PostDate,
	}
	if err := ManeuveringDraftService.CreateManeuveringDraft_Service(); err != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"error": true,
			"msg":   "Thêm mới không thành công",
		})
	}

	data := make(map[string]string)
	data["Pdfurl"] = form.Pdfurl
	data["PostDate"] = form.PostDate
	data["Status"] = strconv.FormatBool(form.Status)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data":    data,
	})
}

func UpdateManeuveringDraft_Component(c *fiber.Ctx) error {
	form := &ManeuveringDraft{}
	if err := c.BodyParser(form); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	ManeuveringDraftService := maneuveringDraft_service.ManeuveringDraft{
		Pdfurl:  form.Pdfurl,
		Status:   form.Status,
		PostDate: form.PostDate,
	}

	err := ManeuveringDraftService.UpdateManeuveringDraft_Service(c.Params("id"))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   "Register false",
		})
	}

	return c.Status(fiber.StatusOK).JSON(ManeuveringDraftService)
}

func DeleteManeuveringDraft_Component(c *fiber.Ctx) error {
	item := maneuveringDraft_service.ManeuveringDraft{}

	data, err := item.DeleteManeuveringDraft_Service(c.Params("id"))
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
