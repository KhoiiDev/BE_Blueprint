package hoatieu_service

import (
	"be-hoatieu/models"

	"gorm.io/gorm"
)

type Hoatieu struct {
	gorm.Model
	Status bool   `gorm:"column:status;" json:"status"`
	Rank   string `gorm:"column:rank" json:"rank"`
	Image  string `gorm:"column:image" json:"image"`
	Name   string `gorm:"column:name" json:"name"`
	Sort   uint   `gorm:"column:sort" json:"sort"`
}

type ObjectHoaTieu struct {
	ID     uint   `gorm:"column:id" json:"id"`
	Status bool   `gorm:"column:status;" json:"status"`
	Rank   string `gorm:"column:rank" json:"rank"`
	Image  string `gorm:"column:image" json:"image"`
	Name   string `gorm:"column:name" json:"name"`
	Sort   uint   `gorm:"column:sort" json:"sort"`
}

func (a *Hoatieu) GetAllNavigator_Service(limit int, page int, name string, showHidden bool) (*[]models.ObjectHoaTieu, int64, error) {
	item, totalRecords, err := models.GetAllNavigator_Model(limit, page, name, showHidden)
	if err != nil {
		return nil, 0, err
	}
	return item, totalRecords, nil
}

func (n *Hoatieu) CreateNavigator_Service() error {
	item := map[string]interface{}{
		"name":   n.Name,
		"image":  n.Image,
		"status": n.Status,
		"rank":   n.Rank,
		"sort":   n.Sort,
	}
	if err := models.CreateNavigator_Model(item); err != nil { // Change function name to match your model's function
		return err
	}
	return nil
}

func (a *Hoatieu) UpdateNavigator_Service(id string) error {
	item := map[string]interface{}{
		"name":   a.Name,
		"image":  a.Image,
		"status": a.Status,
		"rank":   a.Rank,
		"sort":   a.Sort,
	}
	if err := models.UpdateNavigator_Model(id, item); err != nil {
		return err
	}
	return nil
}

func (a *Hoatieu) DeleteNavigator_Service(id string) (bool, error) {
	if err := models.DeleteNavigator_Model(id); err != nil {
		return false, err
	}
	return true, nil
}
