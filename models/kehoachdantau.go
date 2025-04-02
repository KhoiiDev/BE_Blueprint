package models

import (
	"time"

	"gorm.io/gorm"
)

type Kehoachdantau struct {
	gorm.Model
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
	Postdate string `gorm:"column:postdate" json:"postdate"`
}

type ObjectKehoachdantau struct {
	ID       uint   `gorm:"column:id" json:"id"`
	Name     string `gorm:"column:name" json:"name"`
	Country  string `gorm:"column:country" json:"country"`
	Agency   string `gorm:"column:agency" json:"agency"`
	Dwt      string `gorm:"column:dwt" json:"dwt"`
	Grt      string `gorm:"column:grt" json:"grt"`
	Loa      string `gorm:"column:loa" json:"loa"`
	Draft    string `gorm:"column:draft" json:"draft"`
	Fromkh   string `gorm:"column:fromkh" json:"fromkh"`
	Tokh     string `gorm:"column:tokh" json:"tokh"`
	Pob      string `gorm:"column:pob" json:"pob"`
	NameHT   string `gorm:"column:nameHT" json:"nameHT"`
	RangeHT  string `gorm:"column:rangeHT" json:"rangeHT"`
	ItemType string `gorm:"column:itemtype" json:"itemtype"`
	Postdate string `gorm:"column:postdate" json:"postdate"`
}

// func GetKehoachdantau_Model(limit int, page int, showHidden bool, name string, item_type string, ngay string) (*[]ObjectKehoachdantau, int64, error) {
// 	var results []ObjectKehoachdantau
// 	var totalRecords int64
// 	var err error

// 	offset := (page - 1) * limit
// 	query := db.Table("kehoachdantaus").Where("deleted_at IS NULL")

// 	// Lọc theo tên tàu nếu có
// 	if name != "undefined" && name != "" {
// 		query = query.Where("name LIKE ?", "%"+name+"%")
// 	}

// 	err = query.Count(&totalRecords).Error
// 	if err != nil {
// 		return nil, 0, err
// 	}

// 	err = query.Order("created_at DESC").
// 		Limit(limit).
// 		Offset(offset).
// 		Find(&results).Error

// 	if err != nil {
// 		return nil, 0, err
// 	}

// 	return &results, totalRecords, nil
// }

func GetKehoachdantau_Model(limit int, page int, showHidden bool, name string, item_type string, ngay string) (*[]ObjectKehoachdantau, int64, error) {
	var results []ObjectKehoachdantau
	var totalRecords int64
	var err error

	offset := (page - 1) * limit
	query := db.Debug().Table("kehoachdantaus").Where("deleted_at IS NULL")

	// Lọc theo tên tàu nếu có
	if name != "undefined" && name != "" {
		query = query.Where("name LIKE ?", "%"+name+"%")
	}

	// Lọc theo loại tàu nếu có
	if item_type != "undefined" && item_type != "" {
		query = query.Where("itemtype = ?", item_type)
	}

	// Lọc theo ngày nếu có
	if ngay != "undefined" && ngay != "" {
		query = query.Where("postdate = ?", ngay)
	}

	// // Lọc theo trạng thái ẩn/hiện
	// if !showHidden {
	// 	query = query.Where("is_hidden = ?", false)
	// }

	err = query.Count(&totalRecords).Error
	if err != nil {
		return nil, 0, err
	}

	err = query.Order("created_at DESC").
		Limit(limit).
		Offset(offset).
		Find(&results).Error

	if err != nil {
		return nil, 0, err
	}

	return &results, totalRecords, nil
}

func CreateKehoachdantau_Model(data map[string]interface{}) error {
	kehoach := Kehoachdantau{
		Name:     data["name"].(string),
		Country:  data["country"].(string),
		Agency:   data["agency"].(string),
		Dwt:      data["dwt"].(string),
		Grt:      data["grt"].(string),
		Loa:      data["loa"].(string),
		Draft:    data["draft"].(string),
		Fromkh:   data["fromkh"].(string),
		Tokh:     data["tokh"].(string),
		Pob:      data["pob"].(string),
		NameHT:   data["nameHT"].(string),
		RangeHT:  data["rangeHT"].(string),
		ItemType: data["itemtype"].(string),
		Postdate: data["postdate"].(string),
	}

	result := db.Create(&kehoach)
	if err := result.Error; err != nil {
		return err
	}
	return nil
}

func UpdateKehoachdantau_Model(id string, data map[string]interface{}) error {
	kehoach := Kehoachdantau{
		Name:     data["name"].(string),
		Country:  data["country"].(string),
		Agency:   data["agency"].(string),
		Dwt:      data["dwt"].(string),
		Grt:      data["grt"].(string),
		Loa:      data["loa"].(string),
		Draft:    data["draft"].(string),
		Fromkh:   data["fromkh"].(string),
		Tokh:     data["tokh"].(string),
		Pob:      data["pob"].(string),
		NameHT:   data["nameHT"].(string),
		RangeHT:  data["rangeHT"].(string),
		ItemType: data["itemtype"].(string),
		Postdate: data["postdate"].(string),
	}

	if err := db.Model(&kehoach).Where("id = ?", id).Updates(map[string]interface{}{
		"name":       kehoach.Name,
		"country":    kehoach.Country,
		"agency":     kehoach.Agency,
		"dwt":        kehoach.Dwt,
		"grt":        kehoach.Grt,
		"loa":        kehoach.Loa,
		"draft":      kehoach.Draft,
		"fromkh":     kehoach.Fromkh,
		"tokh":       kehoach.Tokh,
		"pob":        kehoach.Pob,
		"nameHT":     kehoach.NameHT,
		"rangeHT":    kehoach.RangeHT,
		"itemtype":   kehoach.ItemType,
		"postdate":   kehoach.Postdate,
		"updated_at": time.Now(),
	}).Error; err != nil {
		return err
	}
	return nil
}

func DeleteKehoachdantau_Model(id string) error {
	var kehoach Kehoachdantau
	if err := db.Select("deleted_at").Where("id = ?", id).First(&kehoach).Error; err != nil {
		return err
	}

	if err := db.Delete(&Kehoachdantau{}, id).Error; err != nil {
		return err
	}
	return nil
}
