package v1

import (
	header_service "be-hoatieu/services/header" // Update service import
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Header struct {
	gorm.Model
	CompanyName string `gorm:"column:company_name" json:"company_name"`
	Address     string `gorm:"column:address" json:"address"`
	Fax         string `gorm:"column:fax" json:"fax"`
	Email       string `gorm:"column:email" json:"email"`
	NumberPhone string `gorm:"column:number_phone" json:"number_phone"`
	BranchName  string `gorm:"column:branch_name" json:"branch_name"`
}

func GetHeader_Component(c *fiber.Ctx) error {
	header := header_service.Header{}

	limit := c.Query("limit")
	page := c.Query("page")
	name := c.Query("name")

	limitStr, err := strconv.Atoi(limit)
	PageStr, err := strconv.Atoi(page)

	data, totalRecords, err := header.GetHeader_Service(limitStr, PageStr, name)

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

func CreateHeader_Component(c *fiber.Ctx) error {
	form := &Header{}

	if err := c.BodyParser(form); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	HeaderService := header_service.Header{
		CompanyName: form.CompanyName,
		Address:     form.Address,
		Fax:         form.Fax,
		Email:       form.Email,
		NumberPhone: form.NumberPhone,
		BranchName:  form.BranchName,
	}
	if err := HeaderService.CreateHeader_Service(); err != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"error": true,
			"msg":   "Thêm mới không thành công",
		})
	}

	data := make(map[string]string)
	data["CompanyName"] = form.CompanyName
	data["Address"] = form.Address
	data["Fax"] = form.Fax
	data["Email"] = form.Email
	data["NumberPhone"] = form.NumberPhone
	data["BranchName"] = form.BranchName

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data":    data,
	})
}

func UpdateHeader_Component(c *fiber.Ctx) error {
	form := &Header{}
	if err := c.BodyParser(form); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	HeaderService := header_service.Header{
		CompanyName: form.CompanyName,
		Address:     form.Address,
		Fax:         form.Fax,
		Email:       form.Email,
		NumberPhone: form.NumberPhone,
		BranchName:  form.BranchName,
	}

	err := HeaderService.UpdateHeader_Service(c.Params("id"))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   "Update false",
		})
	}

	return c.Status(fiber.StatusOK).JSON(HeaderService)
}

func DeleteHeader_Component(c *fiber.Ctx) error {
	header := header_service.Header{}

	data, err := header.DeleteHeader_Service(c.Params("id"))
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
