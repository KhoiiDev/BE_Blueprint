package v1

import (
	serviceList_service "be-hoatieu/services/serviceList"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// Struct để nhận form data
type ServiceList struct {
	Title      string `gorm:"column:title" json:"title"`
	TitleEN    string `gorm:"column:title_en" json:"title_en"`
	SubTitle   string `gorm:"column:subtitle" json:"subtitle"`
	SubTitleEN string `gorm:"column:subtitle_en" json:"subtitle_en"`
	Image      string `gorm:"column:image" json:"image"`
	Postdate   string `gorm:"column:postdate" json:"postdate"`
	Status     bool   `gorm:"column:status" json:"status"`
	Pdfurl     string `gorm:"column:pdfurl" json:"pdfurl"`
	Content    string `gorm:"column:content" json:"content"`
	ContentEN  string `gorm:"column:content_en" json:"content_en"`
}

func GetServiceList_Component(c *fiber.Ctx) error {
	item := serviceList_service.Servicelist{}

	limit := c.Query("limit")
	page := c.Query("page")
	showHiddenItem := c.Query("showHiddenItem")
	name := c.Query("name")

	limitStr, err := strconv.Atoi(limit)
	PageStr, err := strconv.Atoi(page)
	showHidden, err := strconv.ParseBool(showHiddenItem)

	data, totalRecords, err := item.GetServiceList_Service(limitStr, PageStr, name, showHidden)

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

	if err := c.BodyParser(form); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	ListService := serviceList_service.Servicelist{
		Title:      form.Title,
		TitleEN:    form.TitleEN,
		SubTitle:   form.SubTitle,
		SubTitleEN: form.SubTitleEN,
		Pdfurl:     form.Pdfurl,
		Postdate:   form.Postdate,
		Image:      form.Image,
		Status:     form.Status,
		Content:    form.Content,
		ContentEN:  form.ContentEN,
	}

	if err := ListService.CreateServiceList_Service(); err != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"error": true,
			"msg":   "Thêm mới không thành công",
		})
	}

	data := make(map[string]string)
	data["Title"] = form.Title
	data["TitleEN"] = form.TitleEN
	data["SubTitle"] = form.SubTitle
	data["SubTitleEN"] = form.SubTitleEN
	data["Pdfurl"] = form.Pdfurl
	data["Image"] = form.Image
	data["Postdate"] = form.Postdate
	data["Content"] = form.Content
	data["ContentEN"] = form.ContentEN
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
		Title:      form.Title,
		TitleEN:    form.TitleEN,
		SubTitle:   form.SubTitle,
		SubTitleEN: form.SubTitleEN,
		Pdfurl:     form.Pdfurl,
		Postdate:   form.Postdate,
		Image:      form.Image,
		Status:     form.Status,
		Content:    form.Content,
		ContentEN:  form.ContentEN,
	}

	err := ListService.UpdateServiceList_Service(c.Params("id"))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   "Cập nhật không thành công",
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
			"message": "Xoá không thành công: " + err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data":    data,
	})
}
