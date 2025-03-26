package v1

import (
	dichvu_service "be-hoatieu/services/dichvu"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type Dichvu struct {
	Title    string `gorm:"column:title" json:"title"`
	SubTitle string `gorm:"column:subtitle" json:"subtitle"`
	Content  string `gorm:"column:content" json:"content"`
	Image    string `gorm:"column:image" json:"image"`
	Pdfurl   string `gorm:"column:pdfurl" json:"pdfurl"`

	Status   bool   `gorm:"column:status" json:"status"`
	Postdate string `gorm:"column:postdate" json:"postdate"`
}

func GetDichvu_Component(c *fiber.Ctx) error {
	item := dichvu_service.Dichvu{}

	limit := c.Query("limit")
	page := c.Query("page")
	showHiddenItem := c.Query("showHiddenItem")
	name := c.Query("name")

	limitStr, err := strconv.Atoi(limit)
	PageStr, err := strconv.Atoi(page)
	showHidden, err := strconv.ParseBool(showHiddenItem)

	data, totalRecords, err := item.GetDichvu_Service(limitStr, PageStr, name, showHidden)

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

func CreateDichvu_Component(c *fiber.Ctx) error {
	form := &Dichvu{}

	// Check, if received JSON data is valid.
	if err := c.BodyParser(form); err != nil {
		// Return status 400 and error message.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	DichvuService := dichvu_service.Dichvu{
		Title:    form.Title,
		SubTitle: form.SubTitle,
		Content:  form.Content,
		Postdate: form.Postdate,
		Image:    form.Image,
		Status:   form.Status,
		Pdfurl:   form.Pdfurl,
	}
	if err := DichvuService.CreateDichvu_Service(); err != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"error": true,
			"msg":   "Thêm mới không thành công",
		})
	}

	data := make(map[string]string)
	data["Title"] = form.Title
	data["SubTitle"] = form.SubTitle
	data["Content"] = form.Content
	data["Postdate"] = form.Postdate
	data["Image"] = form.Image
	data["Pdfurl"] = form.Pdfurl
	data["Status"] = strconv.FormatBool(form.Status)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data":    data,
	})
}

func UpdateDichvu_Component(c *fiber.Ctx) error {
	form := &Dichvu{}
	if err := c.BodyParser(form); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	NewsService := dichvu_service.Dichvu{
		Title:    form.Title,
		SubTitle: form.SubTitle,
		Content:  form.Content,
		Image:    form.Image,
		Postdate: form.Postdate,
		Status:   form.Status,
		Pdfurl:   form.Pdfurl,
	}

	err := NewsService.UpdateDichvu_Service(c.Params("id"))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   "Register false",
		})
	}

	return c.Status(fiber.StatusOK).JSON(NewsService)
}

func DeleteDichvu_Component(c *fiber.Ctx) error {
	item := dichvu_service.Dichvu{}

	data, err := item.DeleteDichvu_Service(c.Params("id"))
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
