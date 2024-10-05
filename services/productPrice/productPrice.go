package productPrice_service

import (
	"be-hoatieu/models"

	"gorm.io/gorm"
)

type ProductPrice struct {
	gorm.Model
	Title   string `gorm:"column:title" json:"title"`
	Url     string `gorm:"column:url" json:"url"`
	Image   string `gorm:"column:image" json:"image"`
	Status  bool   `gorm:"column:status" json:"status"`
	Content string `gorm:"column:content" json:"content"`
}

type ObjectProductPrice struct {
	ID      uint   `gorm:"column:ID" json:"ID"`
	Title   string `gorm:"column:title" json:"title"`
	Url     string `gorm:"column:url" json:"url"`
	Image   string `gorm:"column:image" json:"image"`
	Status  bool   `gorm:"column:status" json:"status"`
	Content string `gorm:"column:content" json:"content"`
}

func (a *ProductPrice) GetProductPrice_Service(limit int, PageStr int) (*[]models.ObjectProductPrice, int64, error) {
	item, totalRecords, err := models.GetProductPrice_Model(limit, PageStr)
	if err != nil {
		return nil, 0, err
	}
	return item, totalRecords, nil
}
