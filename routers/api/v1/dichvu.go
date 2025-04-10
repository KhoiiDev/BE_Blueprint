package v1

import (
	dichvu_service "be-hoatieu/services/dichvu"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type Dichvu struct {
	Title      string `gorm:"column:title" json:"title"`
	TitleEN    string `gorm:"column:title_en" json:"title_en"`
	SubTitle   string `gorm:"column:subtitle" json:"subtitle"`
	SubTitleEN string `gorm:"column:subtitle_en" json:"subtitle_en"`
	Content    string `gorm:"column:content" json:"content"`
	ContentEN  string `gorm:"column:content_en" json:"content_en"`
	Image      string `gorm:"column:image" json:"image"`
	Pdfurl     string `gorm:"column:pdfurl" json:"pdfurl"`
	Status     bool   `gorm:"column:status" json:"status"`
	Postdate   string `gorm:"column:postdate" json:"postdate"`
}

func GetDichvu_Component(c *fiber.Ctx) error {
	item := dichvu_service.Dichvu{}

	limit := c.Query("limit")
	page := c.Query("page")
	showHiddenItem := c.Query("showHiddenItem")
	name := c.Query("name")

	limitStr, _ := strconv.Atoi(limit)
	pageStr, _ := strconv.Atoi(page)
	showHidden, _ := strconv.ParseBool(showHiddenItem)

	data, totalRecords, err := item.GetDichvu_Service(limitStr, pageStr, name, showHidden)
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

	if err := c.BodyParser(form); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	DichvuService := dichvu_service.Dichvu{
		Title:      form.Title,
		TitleEN:    form.TitleEN,
		SubTitle:   form.SubTitle,
		SubTitleEN: form.SubTitleEN,
		Content:    form.Content,
		ContentEN:  form.ContentEN,
		Image:      form.Image,
		Pdfurl:     form.Pdfurl,
		Status:     form.Status,
		Postdate:   form.Postdate,
	}
	if err := DichvuService.CreateDichvu_Service(); err != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"error": true,
			"msg":   "Thêm mới không thành công",
		})
	}

	data := map[string]string{
		"Title":      form.Title,
		"TitleEN":    form.TitleEN,
		"SubTitle":   form.SubTitle,
		"SubTitleEN": form.SubTitleEN,
		"Content":    form.Content,
		"ContentEN":  form.ContentEN,
		"Postdate":   form.Postdate,
		"Image":      form.Image,
		"Pdfurl":     form.Pdfurl,
		"Status":     strconv.FormatBool(form.Status),
	}

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

	DichvuService := dichvu_service.Dichvu{
		Title:      form.Title,
		TitleEN:    form.TitleEN,
		SubTitle:   form.SubTitle,
		SubTitleEN: form.SubTitleEN,
		Content:    form.Content,
		ContentEN:  form.ContentEN,
		Image:      form.Image,
		Pdfurl:     form.Pdfurl,
		Status:     form.Status,
		Postdate:   form.Postdate,
	}

	err := DichvuService.UpdateDichvu_Service(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   "Cập nhật không thành công",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data":    DichvuService,
	})
}

func DeleteDichvu_Component(c *fiber.Ctx) error {
	item := dichvu_service.Dichvu{}

	data, err := item.DeleteDichvu_Service(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "Xoá thất bại: " + err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data":    data,
	})
}
