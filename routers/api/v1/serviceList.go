package v1

import (
	serviceList_service "be-hoatieu/services/serviceList"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type ServiceList struct {
	Title    string `gorm:"column:title" json:"title"`
	SubTitle string `gorm:"column:subtitle" json:"subtitle"`
	Image    string `gorm:"column:image" json:"image"`
	Postdate string `gorm:"column:postdate" json:"postdate"`
	Status   bool   `gorm:"column:status" json:"status"`

	Pdfurl string `gorm:"column:pdfurl" json:"pdfurl"`

	Content string `gorm:"column:content" json:"content"`
}

func GetServiceList_Component(c *fiber.Ctx) error {
	item := serviceList_service.Servicelist{}

	limit := c.Query("limit")
	page := c.Query("page")
	showHiddenItem := c.Query("showHiddenItem")

	limitStr, err := strconv.Atoi(limit)
	PageStr, err := strconv.Atoi(page)
	showHidden, err := strconv.ParseBool(showHiddenItem)

	data, totalRecords, err := item.GetServiceList_Service(limitStr, PageStr, showHidden)

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

func CreateServiceList_Component(c *fiber.Ctx) error {
	form := &ServiceList{}

	// Check, if received JSON data is valid.
	if err := c.BodyParser(form); err != nil {
		// Return status 400 and error message.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	ListService := serviceList_service.Servicelist{
		Title:    form.Title,
		SubTitle: form.SubTitle,
		Pdfurl:   form.Pdfurl,
		Postdate: form.Postdate,
		Image:    form.Image,
		Status:   form.Status,
		Content:  form.Content,
	}
	if err := ListService.CreateServiceList_Service(); err != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"error": true,
			"msg":   "Thêm mới không thành công",
		})
	}

	data := make(map[string]string)
	data["Title"] = form.Title
	data["SubTitle"] = form.SubTitle
	data["Pdfurl"] = form.Pdfurl
	data["Image"] = form.Image
	data["Postdate"] = form.Postdate
	data["Content"] = form.Content
	data["Status"] = strconv.FormatBool(form.Status)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data":    data,
	})
}

func UpdateServiceList_Component(c *fiber.Ctx) error {
	form := &ServiceList{}
	if err := c.BodyParser(form); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	ListService := serviceList_service.Servicelist{
		Title:    form.Title,
		Image:    form.Image,
		SubTitle: form.SubTitle,
		Postdate: form.Postdate,
		Pdfurl:   form.Pdfurl,
		Status:   form.Status,
		Content:  form.Content,
	}

	err := ListService.UpdateServiceList_Service(c.Params("id"))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   "Register false",
		})
	}

	return c.Status(fiber.StatusOK).JSON(ListService)
}

func DeleteServiceList_Component(c *fiber.Ctx) error {
	item := serviceList_service.Servicelist{}

	data, err := item.DeleteServiceList_Service(c.Params("id"))
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
