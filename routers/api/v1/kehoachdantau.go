package v1

import (
	kehoachdantau_service "be-hoatieu/services/kehoachdantau"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type Kehoachdantau struct {
	Name     string `gorm:"column:name" json:"name"`       // Tên tàu
	Country  string `gorm:"column:country" json:"country"` // Quốc gia
	Agency   string `gorm:"column:agency" json:"agency"`   // Đại lý
	Dwt      string `gorm:"column:dwt" json:"dwt"`         // Trọng tải toàn phần
	Grt      string `gorm:"column:grt" json:"grt"`         // Dung tích toàn phần
	Loa      string `gorm:"column:loa" json:"loa"`         // Chiều dài tổng
	Draft    string `gorm:"column:draft" json:"draft"`     // Mớn nước
	Fromkh   string `gorm:"column:fromkh" json:"fromkh"`   // Từ cảng
	Tokh     string `gorm:"column:tokh" json:"tokh"`       // Đến cảng
	Pob      string `gorm:"column:pob" json:"pob"`         // Số người trên tàu
	NameHT   string `gorm:"column:nameHT" json:"nameHT"`   // Tên hệ thống
	RangeHT  string `gorm:"column:rangeHT" json:"rangeHT"` // Tầm hoạt động hệ thống
	ItemType string `gorm:"column:itemtype" json:"itemtype"`
	PostDate string `gorm:"column:postdate" json:"postdate"`
}

func GetKehoachdantau_Component(c *fiber.Ctx) error {
	kehoach := kehoachdantau_service.Kehoachdantau{}

	limit := c.Query("limit")
	page := c.Query("page")
	showHiddenItem := c.Query("showHiddenItem")
	name := c.Query("name")
	// item_type := c.Query("itemType")
	item_type := c.Query("dendoi")
	ngay := c.Query("ngay")

	limitStr, err := strconv.Atoi(limit)
	PageStr, err := strconv.Atoi(page)
	showHidden, err := strconv.ParseBool(showHiddenItem)

	data, totalRecords, err := kehoach.GetKehoachdantau_Service(limitStr, PageStr, showHidden, name, item_type, ngay)

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

func CreateKehoachdantau_Component(c *fiber.Ctx) error {
	form := &Kehoachdantau{}

	if err := c.BodyParser(form); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	KehoachdantauService := kehoachdantau_service.Kehoachdantau{
		Name:     form.Name,
		Country:  form.Country,
		Agency:   form.Agency,
		Dwt:      form.Dwt,
		Grt:      form.Grt,
		Loa:      form.Loa,
		Draft:    form.Draft,
		Fromkh:   form.Fromkh,
		Tokh:     form.Tokh,
		Pob:      form.Pob,
		NameHT:   form.NameHT,
		RangeHT:  form.RangeHT,
		ItemType: form.ItemType,
		PostDate: form.PostDate,
	}

	if err := KehoachdantauService.CreateKehoachdantau_Service(); err != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"error": true,
			"msg":   "Thêm mới không thành công",
		})
	}

	data := make(map[string]string)
	data["Name"] = form.Name
	data["Country"] = form.Country
	data["Agency"] = form.Agency
	data["Dwt"] = form.Dwt
	data["Grt"] = form.Grt
	data["Loa"] = form.Loa
	data["Draft"] = form.Draft
	data["Fromkh"] = form.Fromkh
	data["Tokh"] = form.Tokh
	data["Pob"] = form.Pob
	data["NameHT"] = form.NameHT
	data["RangeHT"] = form.RangeHT
	data["itemtype"] = form.ItemType
	data["postdate"] = form.PostDate

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data":    data,
	})
}

func UpdateKehoachdantau_Component(c *fiber.Ctx) error {
	form := &Kehoachdantau{}
	if err := c.BodyParser(form); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	KehoachdantauService := kehoachdantau_service.Kehoachdantau{
		Name:     form.Name,
		Country:  form.Country,
		Agency:   form.Agency,
		Dwt:      form.Dwt,
		Grt:      form.Grt,
		Loa:      form.Loa,
		Draft:    form.Draft,
		Fromkh:   form.Fromkh,
		Tokh:     form.Tokh,
		Pob:      form.Pob,
		NameHT:   form.NameHT,
		RangeHT:  form.RangeHT,
		ItemType: form.ItemType,
		PostDate: form.PostDate,
	}

	err := KehoachdantauService.UpdateKehoachdantau_Service(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   "Register false",
		})
	}

	return c.Status(fiber.StatusOK).JSON(KehoachdantauService)
}

func DeleteKehoachdantau_Component(c *fiber.Ctx) error {
	kehoach := kehoachdantau_service.Kehoachdantau{}

	data, err := kehoach.DeleteKehoachdantau_Service(c.Params("id"))
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
