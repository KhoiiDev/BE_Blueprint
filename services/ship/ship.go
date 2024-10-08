package ship_service

import (
	"be-hoatieu/models"

	"gorm.io/gorm"
)

type Ship struct {
	gorm.Model
	Title   string `gorm:"column:title" json:"title"`
	Url     string `gorm:"column:url" json:"url"`
	Image   string `gorm:"column:image" json:"image"`
	Status  bool   `gorm:"column:status" json:"status"`
	Content string `gorm:"column:content" json:"content"`
}

type ObjectShip struct {
	ID      uint   `gorm:"column:ID" json:"ID"`
	Title   string `gorm:"column:title" json:"title"`
	Url     string `gorm:"column:url" json:"url"`
	Image   string `gorm:"column:image" json:"image"`
	Status  bool   `gorm:"column:status" json:"status"`
	Content string `gorm:"column:content" json:"content"`
}

func (a *Ship) GetShip_Service(limit int, page int) (*[]models.ObjectShip, int64, error) {
	item, totalRecords, err := models.GetShip_Model(limit, page)
	if err != nil {
		return nil, 0, err
	}
	return item, totalRecords, nil
}
