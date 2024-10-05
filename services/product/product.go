package product_service

import (
	"be-hoatieu/models"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Title   string `gorm:"column:title" json:"title"`
	Image   string `gorm:"column:image" json:"image"`
	Url     string `gorm:"column:url" json:"url"`
	Content string `gorm:"column:content" json:"content"`
	Status  bool   `gorm:"column:status" json:"status"`
}

type ObjectProduct struct {
	ID      uint   `gorm:"column:ID" json:"ID"`
	Title   string `gorm:"column:title" json:"title"`
	Image   string `gorm:"column:image" json:"image"`
	Url     string `gorm:"column:url" json:"url"`
	Content string `gorm:"column:content" json:"content"`
	Status  bool   `gorm:"column:status" json:"status"`
}

func (a *Product) GetProduct_Service(limit int) (*[]models.ObjectProduct, error) {
	item, err := models.GetProduct_Model(limit)
	if err != nil {
		return nil, err
	}
	return item, nil
}
