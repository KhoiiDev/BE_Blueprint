package v1

import (
	carousel_service "be-hoatieu/services/carousel"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type Carousel struct {
	Image  string `gorm:"column:image" json:"image"`
	Status bool   `gorm:"column:status" json:"status"`
}

func GetCarousel_Component(c *fiber.Ctx) error {
	item := carousel_service.Carousel{}

	limit := c.Query("limit")
	page := c.Query("page")
	showHiddenItem := c.Query("showHiddenItem")

	limitStr, err := strconv.Atoi(limit)
	pageStr, err := strconv.Atoi(page)
	showHiddenStr, err := strconv.ParseBool(showHiddenItem)

	data, totalRecords, err := item.GetCarousel_Service(limitStr, pageStr, showHiddenStr)

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

func CreateCarousel_Component(c *fiber.Ctx) error {
	form := &Carousel{}
	// Check, if received JSON data is valid.
	if err := c.BodyParser(form); err != nil {
		// Return status 400 and error message.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	CarouselService := carousel_service.Carousel{
		Image:  form.Image,
		Status: form.Status,
	}
	if err := CarouselService.CreateCarousel_Service(); err != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"error": true,
			"msg":   "Thêm mới không thành công",
		})
	}

	data := make(map[string]string)
	data["Image"] = form.Image
	data["Status"] = strconv.FormatBool(form.Status)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data":    data,
	})
}

func UpdateCarousel_Component(c *fiber.Ctx) error {
	form := &Carousel{}
	if err := c.BodyParser(form); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	CarouselService := carousel_service.Carousel{
		Image:  form.Image,
		Status: form.Status,
	}

	err := CarouselService.UpdateCarousel_Service(c.Params("id"))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   "Register false",
		})
	}

	return c.Status(fiber.StatusOK).JSON(CarouselService)
}

func DeleteCarousel_Component(c *fiber.Ctx) error {
	item := carousel_service.Carousel{}

	data, err := item.DeleteCarousel_Service(c.Params("id"))
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
