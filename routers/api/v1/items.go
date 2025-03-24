package v1

import (
	items_service "be-hoatieu/services/items"
	"strconv"

	// "time"

	"github.com/gofiber/fiber/v2"
)

type Items struct {
	Title    string `gorm:"column:title" json:"title"`
	SubTitle string `gorm:"column:subtitle" json:"subtitle"`
	Image    string `gorm:"column:image" json:"image"`
	Pdfurl   string `gorm:"column:pdfurl" json:"pdfurl"`
	Status   bool   `gorm:"column:status" json:"status"`
	Content  string `gorm:"column:content" json:"content"`
	Videourl string `gorm:"column:videourl" json:"videourl"`

	PostDate string `gorm:"column:postdate" json:"postdate"`
	ItemType string `gorm:"column:itemtype" json:"itemtype"`
}

func GetItems_Component(c *fiber.Ctx) error {
	item := items_service.Items{}

	limit := c.Query("limit")
	page := c.Query("page")
	showHiddenItem := c.Query("showHiddenItem")
	item_type := c.Query("itemType")

	limitStr, err := strconv.Atoi(limit)
	PageStr, err := strconv.Atoi(page)
	showHidden, err := strconv.ParseBool(showHiddenItem)

	data, totalRecords, err := item.GetItems_Service(limitStr, PageStr, showHidden, item_type)

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

func CreateItems_Component(c *fiber.Ctx) error {
	form := &Items{}

	// Check, if received JSON data is valid.
	if err := c.BodyParser(form); err != nil {
		// Return status 400 and error message.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	ItemsService := items_service.Items{
		Title:    form.Title,
		SubTitle: form.SubTitle,
		Image:    form.Image,
		Pdfurl:   form.Pdfurl,
		Status:   form.Status,
		Content:  form.Content,
		Videourl: form.Videourl,
		Postdate: form.PostDate,
		ItemType: form.ItemType,
	}
	if err := ItemsService.CreateItems_Service(); err != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"error": true,
			"msg":   "Thêm mới không thành công",
		})
	}

	data := make(map[string]string)
	data["Title"] = form.Title
	data["SubTitle"] = form.SubTitle
	data["Image"] = form.Image
	data["Pdfurl"] = form.Pdfurl
	data["Content"] = form.Content
	data["Videourl"] = form.Videourl
	data["PostDate"] = form.PostDate
	data["ItemType"] = form.ItemType
	data["Status"] = strconv.FormatBool(form.Status)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data":    data,
	})
}

func UpdateItems_Component(c *fiber.Ctx) error {
	form := &Items{}
	if err := c.BodyParser(form); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	ItemsService := items_service.Items{
		Title:    form.Title,
		SubTitle: form.SubTitle,
		Image:    form.Image,
		Pdfurl:   form.Pdfurl,
		Status:   form.Status,
		Content:  form.Content,
		Videourl: form.Videourl,

		Postdate: form.PostDate,
	}

	err := ItemsService.UpdateItems_Service(c.Params("id"))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   "Register false",
		})
	}

	return c.Status(fiber.StatusOK).JSON(ItemsService)
}

func DeleteItems_Component(c *fiber.Ctx) error {
	item := items_service.Items{}

	data, err := item.DeleteItems_Service(c.Params("id"))
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
