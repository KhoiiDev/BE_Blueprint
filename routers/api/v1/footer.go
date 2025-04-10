package v1

import (
	footer_service "be-hoatieu/services/footer" // Update service import
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Footer struct {
	gorm.Model
	CompanyName string `gorm:"column:company_name" json:"company_name"`
	Address     string `gorm:"column:address" json:"address"`
	Fax         string `gorm:"column:fax" json:"fax"`
	Mst         string `gorm:"column:mst" json:"mst"`
	Email       string `gorm:"column:email" json:"email"`
	NumberPhone string `gorm:"column:number_phone" json:"number_phone"`
	BranchName  string `gorm:"column:branch_name" json:"branch_name"`
	Linkfb      string `gorm:"column:linkfb" json:"linkfb"`
}

func GetFooter_Component(c *fiber.Ctx) error {
	footer := footer_service.Footer{}

	limit := c.Query("limit")
	page := c.Query("page")
	// showHiddenItem := c.Query("showHiddenItem")
	name := c.Query("name")
	// Removed item_type and name as they might not be relevant for Footer

	limitStr, err := strconv.Atoi(limit)
	PageStr, err := strconv.Atoi(page)
	// showHidden, err := strconv.ParseBool(showHiddenItem)

	data, totalRecords, err := footer.GetFooter_Service(limitStr, PageStr, name)

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

func CreateFooter_Component(c *fiber.Ctx) error {
	form := &Footer{}

	if err := c.BodyParser(form); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	FooterService := footer_service.Footer{
		CompanyName: form.CompanyName,
		Address:     form.Address,
		Fax:         form.Fax,
		Mst:         form.Mst,
		Email:       form.Email,
		NumberPhone: form.NumberPhone,
		Linkfb:      form.Linkfb,
	}
	if err := FooterService.CreateFooter_Service(); err != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"error": true,
			"msg":   "Thêm mới không thành công",
		})
	}

	data := make(map[string]string)
	data["CompanyName"] = form.CompanyName
	data["Address"] = form.Address
	data["Fax"] = form.Fax
	data["Mst"] = form.Mst
	data["Email"] = form.Email
	data["NumberPhone"] = form.NumberPhone
	data["BranchName"] = form.BranchName
	data["Linkfb"] = form.Linkfb

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data":    data,
	})
}

func UpdateFooter_Component(c *fiber.Ctx) error {
	form := &Footer{}
	if err := c.BodyParser(form); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	FooterService := footer_service.Footer{
		CompanyName: form.CompanyName,
		Address:     form.Address,
		Fax:         form.Fax,
		Mst:         form.Mst,
		Email:       form.Email,
		NumberPhone: form.NumberPhone,
		BranchName:  form.BranchName,
		Linkfb:      form.Linkfb,
	}

	err := FooterService.UpdateFooter_Service(c.Params("id"))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   "Update false",
		})
	}

	return c.Status(fiber.StatusOK).JSON(FooterService)
}

func DeleteFooter_Component(c *fiber.Ctx) error {
	footer := footer_service.Footer{}

	data, err := footer.DeleteFooter_Service(c.Params("id"))
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
