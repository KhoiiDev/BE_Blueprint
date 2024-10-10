package ship_service

import (
	"be-hoatieu/models"

	"gorm.io/gorm"
)

type Ship struct {
	gorm.Model
	Name   string `gorm:"column:name" json:"name"`
	Image  string `gorm:"column:image" json:"image"`
	Status bool   `gorm:"column:status" json:"status"`
}

type ObjectShip struct {
	ID     uint   `gorm:"column:id" json:"id"`
	Name   string `gorm:"column:name" json:"name"`
	Image  string `gorm:"column:image" json:"image"`
	Status bool   `gorm:"column:status" json:"status"`
}

func (a *Ship) GetShip_Service(limit int, page int, showHidden bool) (*[]models.ObjectShip, int64, error) {
	item, totalRecords, err := models.GetShip_Model(limit, page, showHidden)
	if err != nil {
		return nil, 0, err
	}
	return item, totalRecords, nil
}

func (n *Ship) CreateShip_Service() error {
	item := map[string]interface{}{
		"name":   n.Name,
		"image":  n.Image,
		"status": n.Status,
	}
	if err := models.CreateShip_Model(item); err != nil { // Change function name to match your model's function
		return err
	}
	return nil
}

func (a *Ship) UpdateShip_Service(id string) error {
	item := map[string]interface{}{
		"name":   a.Name,
		"image":  a.Image,
		"status": a.Status,
	}
	if err := models.UpdateShip_Model(id, item); err != nil {
		return err
	}
	return nil
}

func (a *Ship) DeleteShip_Service(id string) (bool, error) {
	if err := models.DeleteShip_Model(id); err != nil {
		return false, err
	}
	return true, nil
}
