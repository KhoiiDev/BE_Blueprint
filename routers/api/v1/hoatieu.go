package v1

import (
	hoatieu_service "be-hoatieu/services/hoatieu"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type Hoatieu struct {
	Code              string `form:"code" `
	Status            bool   `form:"status" `
	UserId            uint   `form:"userid" `
	HangId            uint   `form:"hangid" `
	BPCTac            uint   `form:"bpctac" `
	NgaySinh          string `form:"ngaysinh" `
	Phone             string `form:"phone" `
	Phone2            string `form:"phone2" `
	Sex               string `form:"sex" `
	Tuoi              string `form:"tuoi" `
	ChungChiHoaTieuID uint   `form:"chungchihoatieuid"`
}

func GetAllNavigator_Component(c *fiber.Ctx) error {
	item := hoatieu_service.Hoatieu{}
	limit := c.Query("limit")
	page := c.Query("page")

	limitStr, err := strconv.Atoi(limit)
	PageStr, err := strconv.Atoi(page)
	data, totalRecords, err := item.GetAllNavigator_Service(limitStr, PageStr)

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

// func CreateHoatieu_Component(c *fiber.Ctx) error {
// 	form := &Hoatieu{}

// 	// Check, if received JSON data is valid.
// 	if err := c.BodyParser(form); err != nil {
// 		// Return status 400 and error message.
// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
// 			"error": true,
// 			"msg":   err.Error(),
// 		})
// 	}

// 	hangService := hoatieu_service.Hoatieu{
// 		Code:              form.Code,
// 		UserId:            form.UserId,
// 		HangId:            form.HangId,
// 		Status:            form.Status,
// 		BPCTac:            form.BPCTac,
// 		NgaySinh:          form.NgaySinh,
// 		Phone2:            form.Phone2,
// 		Sex:               form.Sex,
// 		Tuoi:              form.Tuoi,
// 		ChungChiHoaTieuID: form.ChungChiHoaTieuID,
// 	}
// 	if err := hangService.CreateHoaTieu_Service(); err != nil {
// 		return c.Status(fiber.StatusOK).JSON(fiber.Map{
// 			"error": true,
// 			"msg":   "Thêm mới không thành công",
// 		})
// 	}
// 	data := hangService

// 	return c.Status(fiber.StatusOK).JSON(fiber.Map{
// 		"success": true,
// 		"data":    data,
// 	})
// }

// func UpdateHoaTieu_Component(c *fiber.Ctx) error {
// 	form := &Hoatieu{}
// 	if err := c.BodyParser(form); err != nil {
// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
// 			"error": true,
// 			"msg":   err.Error(),
// 		})
// 	}
// 	hoatieuService := hoatieu_service.Hoatieu{
// 		Code:              form.Code,
// 		UserId:            form.UserId,
// 		HangId:            form.HangId,
// 		Status:            form.Status,
// 		BPCTac:            form.BPCTac,
// 		NgaySinh:          form.NgaySinh,
// 		Phone:             form.Phone,
// 		Phone2:            form.Phone2,
// 		Sex:               form.Sex,
// 		Tuoi:              form.Tuoi,
// 		ChungChiHoaTieuID: form.ChungChiHoaTieuID,
// 	}

// 	err := hoatieuService.UpdateHoaTieu_Service(c.Params("id"))

// 	if err != nil {
// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
// 			"error": true,
// 			"msg":   "Register false",
// 		})
// 	}

//		return c.Status(fiber.StatusOK).JSON(hoatieuService)
//	}

// func GetAllHoaTieuUserID_Component(c *fiber.Ctx) error {
// 	item := hoatieu_service.Hoatieu{}

// 	data, err := item.GetAllHoaTieuUserID_Service(c.Params("id"))

// 	if err != nil {
// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
// 			"error": true,
// 			"msg":   err.Error(),
// 		})
// 	}

// 	return c.Status(fiber.StatusOK).JSON(fiber.Map{
// 		"success": true,
// 		"data":    data,
// 	})
// }
// func GetByHoaTieuUserID_Component(c *fiber.Ctx) error {
// 	item := hoatieu_service.Hoatieu{}

// 	data, err := item.GetByHoaTieuUserID_Service(c.Params("id"))

// 	if err != nil {
// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
// 			"error": true,
// 			"msg":   err.Error(),
// 		})
// 	}

// 	return c.Status(fiber.StatusOK).JSON(fiber.Map{
// 		"success": true,
// 		"data":    data,
// 	})
// }
// func SearchHoaTieu_Component(c *fiber.Ctx) error {
// 	name := c.Query("name")
// 	cangs, err := hoatieu_service.SearchHoaTieu_Service(name)
// 	if err != nil {
// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
// 			"error":   true,
// 			"message": "Search failed: " + err.Error(),
// 		})
// 	}

// 	return c.Status(fiber.StatusOK).JSON(fiber.Map{
// 		"success": true,
// 		"data":    cangs,
// 	})
// }
// func SearchUpDateHoaTieu_Component(c *fiber.Ctx) error {
// 	name := c.Query("name")
// 	id := c.Query("id")
// 	ngay := c.Query("ngay")
// 	cangs, err := hoatieu_service.SearchUpDateHoaTieu_Service(id, ngay, name)
// 	if err != nil {
// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
// 			"error":   true,
// 			"message": "Search failed: " + err.Error(),
// 		})
// 	}

// 	return c.Status(fiber.StatusOK).JSON(fiber.Map{
// 		"success": true,
// 		"data":    cangs,
// 	})
// }

// func SearchHT2UpDateHoaTieu_Component(c *fiber.Ctx) error {
// 	name := c.Query("name")
// 	id := c.Query("id")
// 	ngay := c.Query("ngay")
// 	cangs, err := hoatieu_service.SearchHT2UpDateHoaTieu_Service(id, ngay, name)
// 	if err != nil {
// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
// 			"error":   true,
// 			"message": "Search failed: " + err.Error(),
// 		})
// 	}

// 	return c.Status(fiber.StatusOK).JSON(fiber.Map{
// 		"success": true,
// 		"data":    cangs,
// 	})
// }
// func SearchHTBangKeSanLuongUpDateHoaTieu_Component(c *fiber.Ctx) error {
// 	name := c.Query("name")
// 	ngay := c.Query("ngay")
// 	cangs, err := hoatieu_service.SearchHTBangKeSanLuongUpDateHoaTieu_Service(ngay, name)
// 	if err != nil {
// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
// 			"error":   true,
// 			"message": "Search failed: " + err.Error(),
// 		})
// 	}

// 	return c.Status(fiber.StatusOK).JSON(fiber.Map{
// 		"success": true,
// 		"data":    cangs,
// 	})
// }

// func SearchHoaTieuRole_Component(c *fiber.Ctx) error {
// 	name := c.Query("name")
// 	role := c.Query("role")
// 	cangs, err := hoatieu_service.SearchHoaTieuRole_Service(name, role)
// 	if err != nil {
// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
// 			"error":   true,
// 			"message": "Search failed: " + err.Error(),
// 		})
// 	}

// 	return c.Status(fiber.StatusOK).JSON(fiber.Map{
// 		"success": true,
// 		"data":    cangs,
// 	})
// }
// func GetMultipleHoaTieuTrue_Component(c *fiber.Ctx) error {
// 	item := hoatieu_service.Hoatieu{}
// 	name := c.Query("name")
// 	data, err := item.GetMultipleHoaTieuTrue_Service(name)

// 	if err != nil {
// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
// 			"error": true,
// 			"msg":   err.Error(),
// 		})
// 	}

// 	return c.Status(fiber.StatusOK).JSON(fiber.Map{
// 		"success": true,
// 		"data":    data,
// 	})
// }
