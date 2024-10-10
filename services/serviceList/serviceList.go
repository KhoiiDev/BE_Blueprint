package serviceList_service

import (
	"be-hoatieu/models"

	"gorm.io/gorm"
)

type Servicelist struct {
	gorm.Model
	Title   string `gorm:"column:title" json:"title"`
	Image   string `gorm:"column:image" json:"image"`
	Status  bool   `gorm:"column:status" json:"status"`
	Content string `gorm:"column:content" json:"content"`
}

type ObjectServicelist struct {
	ID      uint   `gorm:"column:id" json:"id"`
	Title   string `gorm:"column:title" json:"title"`
	Image   string `gorm:"column:image" json:"image"`
	Status  bool   `gorm:"column:status" json:"status"`
	Content string `gorm:"column:content" json:"content"`
}

func (a *Servicelist) GetServiceList_Service(limit int, PageStr int, showHidden bool) (*[]models.ObjectServicelist, int64, error) {
	item, totalRecords, err := models.GetServiceList_Model(limit, PageStr, showHidden)
	if err != nil {
		return nil, 0, err
	}
	return item, totalRecords, nil
}

func (n *Servicelist) CreateServiceList_Service() error {
	item := map[string]interface{}{
		"title":   n.Title,
		"image":   n.Image,
		"status":  n.Status,
		"content": n.Content,
	}
	if err := models.CreateServiceList_Model(item); err != nil { // Change function name to match your model's function
		return err
	}
	return nil
}

func (a *Servicelist) UpdateServiceList_Service(id string) error {
	item := map[string]interface{}{
		"title":   a.Title,
		"image":   a.Image,
		"status":  a.Status,
		"content": a.Content,
	}
	if err := models.UpdateServiceList_Model(id, item); err != nil {
		return err
	}
	return nil
}

func (a *Servicelist) DeleteServiceList_Service(id string) (bool, error) {
	if err := models.DeleteServiceList_Model(id); err != nil {
		return false, err
	}
	return true, nil
}
