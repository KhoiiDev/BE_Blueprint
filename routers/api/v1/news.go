package v1

import (
	news_service "be-hoatieu/services/news"
	"strconv"

	// "time"

	"github.com/gofiber/fiber/v2"
)

type News struct {
	Title    string `gorm:"column:title" json:"title"`
	SubTitle string `gorm:"column:subtitle" json:"subtitle"`
	Image    string `gorm:"column:image" json:"image"`
	Status   bool   `gorm:"column:status" json:"status"`
	Content  string `gorm:"column:content" json:"content"`
	PostDate string `gorm:"column:postdate" json:"postdate"`
}

func GetNews_Component(c *fiber.Ctx) error {
	item := news_service.News{}

	limit := c.Query("limit")
	page := c.Query("page")
	showHiddenItem := c.Query("showHiddenItem")

	limitStr, err := strconv.Atoi(limit)
	PageStr, err := strconv.Atoi(page)
	showHidden, err := strconv.ParseBool(showHiddenItem)

	data, totalRecords, err := item.GetNews_Service(limitStr, PageStr, showHidden)

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

	// Check, if received JSON data is valid.
	if err := c.BodyParser(form); err != nil {
		// Return status 400 and error message.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	NewsService := news_service.News{
		Title:    form.Title,
		SubTitle: form.SubTitle,
		Image:    form.Image,
		Status:   form.Status,
		Content:  form.Content,
		Postdate: form.PostDate,
	}
	if err := NewsService.CreateNews_Service(); err != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"error": true,
			"msg":   "Thêm mới không thành công",
		})
	}

	data := make(map[string]string)
	data["Title"] = form.Title
	data["SubTitle"] = form.SubTitle
	data["Image"] = form.Image
	data["Content"] = form.Content
	data["PostDate"] = form.PostDate
	data["Status"] = strconv.FormatBool(form.Status)

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
		Title:    form.Title,
		SubTitle: form.SubTitle,
		Image:    form.Image,
		Status:   form.Status,
		Content:  form.Content,
		Postdate: form.PostDate,
	}

	err := NewsService.UpdateNews_Service(c.Params("id"))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   "Register false",
		})
	}

	return c.Status(fiber.StatusOK).JSON(NewsService)
}

func DeleteNews_Component(c *fiber.Ctx) error {
	item := news_service.News{}

	data, err := item.DeleteNews_Service(c.Params("id"))
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
