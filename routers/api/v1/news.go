package v1

import (
	news_service "be-hoatieu/services/news"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type News struct {
	Title      string `gorm:"column:title" json:"title"`
	TitleEN    string `gorm:"column:title_en" json:"title_en"`
	SubTitle   string `gorm:"column:subtitle" json:"subtitle"`
	SubTitleEN string `gorm:"column:subtitle_en" json:"subtitle_en"`
	Image      string `gorm:"column:image" json:"image"`
	Status     bool   `gorm:"column:status" json:"status"`
	Content    string `gorm:"column:content" json:"content"`
	ContentEN  string `gorm:"column:content_en" json:"content_en"`
	PostDate   string `gorm:"column:postdate" json:"postdate"`
	Pdfurl     string `gorm:"column:pdfurl" json:"pdfurl"`
}

func GetNews_Component(c *fiber.Ctx) error {
	item := news_service.News{}

	limit := c.Query("limit")
	page := c.Query("page")
	name := c.Query("name")
	showHiddenItem := c.Query("showHiddenItem")

	limitStr, _ := strconv.Atoi(limit)
	pageStr, _ := strconv.Atoi(page)
	showHidden, _ := strconv.ParseBool(showHiddenItem)

	data, totalRecords, err := item.GetNews_Service(limitStr, pageStr, name, showHidden)
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

func CreateNews_Component(c *fiber.Ctx) error {
	form := &News{}
	if err := c.BodyParser(form); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	NewsService := news_service.News{
		Title:      form.Title,
		TitleEN:    form.TitleEN,
		SubTitle:   form.SubTitle,
		SubTitleEN: form.SubTitleEN,
		Image:      form.Image,
		Status:     form.Status,
		Content:    form.Content,
		ContentEN:  form.ContentEN,
		Postdate:   form.PostDate,
		Pdfurl:     form.Pdfurl,
	}
	if err := NewsService.CreateNews_Service(); err != nil {
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
		"Image":      form.Image,
		"Content":    form.Content,
		"ContentEN":  form.ContentEN,
		"PostDate":   form.PostDate,
		"Pdfurl":     form.Pdfurl,
		"Status":     strconv.FormatBool(form.Status),
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data":    data,
	})
}

func UpdateNews_Component(c *fiber.Ctx) error {
	form := &News{}
	if err := c.BodyParser(form); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	NewsService := news_service.News{
		Title:      form.Title,
		TitleEN:    form.TitleEN,
		SubTitle:   form.SubTitle,
		SubTitleEN: form.SubTitleEN,
		Image:      form.Image,
		Status:     form.Status,
		Content:    form.Content,
		ContentEN:  form.ContentEN,
		Postdate:   form.PostDate,
		Pdfurl:     form.Pdfurl,
	}

	err := NewsService.UpdateNews_Service(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   "Cập nhật không thành công",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data":    NewsService,
	})
}

func DeleteNews_Component(c *fiber.Ctx) error {
	item := news_service.News{}

	data, err := item.DeleteNews_Service(c.Params("id"))
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
