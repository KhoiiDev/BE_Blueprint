package kehoachdantau_service

import (
	"be-hoatieu/models"

	"gorm.io/gorm"
)

type Kehoachdantau struct {
	gorm.Model
	Name    string `gorm:"column:name" json:"name"`       // Tên tàu
	Country string `gorm:"column:country" json:"country"` // Quốc gia
	Agency  string `gorm:"column:agency" json:"agency"`   // Đại lý
	Dwt     string `gorm:"column:dwt" json:"dwt"`         // Trọng tải toàn phần
	Grt     string `gorm:"column:grt" json:"grt"`         // Dung tích toàn phần
	Loa     string `gorm:"column:loa" json:"loa"`         // Chiều dài tổng
	Draft   string `gorm:"column:draft" json:"draft"`     // Mớn nước
	Fromkh  string `gorm:"column:fromkh" json:"fromkh"`   // Từ cảng
	Tokh    string `gorm:"column:tokh" json:"tokh"`       // Đến cảng
	Pob     string `gorm:"column:pob" json:"pob"`         // Số người trên tàu
	NameHT  string `gorm:"column:nameHT" json:"nameHT"`   // Tên hệ thống
	RangeHT string `gorm:"column:rangeHT" json:"rangeHT"` // Tầm hoạt động hệ thống

	ItemType string `gorm:"column:itemtype" json:"itemtype"`
	PostDate string `gorm:"column:postdate" json:"postdate"`
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

func (k *Kehoachdantau) GetKehoachdantau_Service(limit int, page int, showHidden bool, name string, item_type string, ngay string) (*[]models.ObjectKehoachdantau, int64, error) {
	kehoach, totalRecords, err := models.GetKehoachdantau_Model(limit, page, showHidden, name, item_type, ngay)
	if err != nil {
		return nil, totalRecords, err
	}
	return kehoach, totalRecords, nil
}

func (k *Kehoachdantau) CreateKehoachdantau_Service() error {
	kehoach := map[string]interface{}{
		"name":     k.Name,
		"country":  k.Country,
		"agency":   k.Agency,
		"dwt":      k.Dwt,
		"grt":      k.Grt,
		"loa":      k.Loa,
		"draft":    k.Draft,
		"fromkh":   k.Fromkh,
		"tokh":     k.Tokh,
		"pob":      k.Pob,
		"nameHT":   k.NameHT,
		"rangeHT":  k.RangeHT,
		"itemtype": k.ItemType,
		"postdate": k.PostDate,
	}
	if err := models.CreateKehoachdantau_Model(kehoach); err != nil {
		return err
	}
	return nil
}

func (k *Kehoachdantau) UpdateKehoachdantau_Service(id string) error {
	kehoach := map[string]interface{}{
		"name":     k.Name,
		"country":  k.Country,
		"agency":   k.Agency,
		"dwt":      k.Dwt,
		"grt":      k.Grt,
		"loa":      k.Loa,
		"draft":    k.Draft,
		"fromkh":   k.Fromkh,
		"tokh":     k.Tokh,
		"pob":      k.Pob,
		"nameHT":   k.NameHT,
		"rangeHT":  k.RangeHT,
		"itemtype": k.ItemType,
		"postdate": k.PostDate,
	}
	if err := models.UpdateKehoachdantau_Model(id, kehoach); err != nil {
		return err
	}
	return nil
}

func (k *Kehoachdantau) DeleteKehoachdantau_Service(id string) (bool, error) {
	if err := models.DeleteKehoachdantau_Model(id); err != nil {
		return false, err
	}
	return true, nil
}
