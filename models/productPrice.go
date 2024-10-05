package models

import (
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

func GetProductPrice_Model(limit int, PageStr int) (*[]ObjectProductPrice, int64, error) {
	var results []ObjectProductPrice
	var totalRecords int64

	//  Truy vấn để đếm tổng số bản ghi có status = 1
	err := db.Table("product_prices").Where("status = ?", 1).Count(&totalRecords).Error
	if err != nil {
		return nil, 0, err
	}

	// Truy vấn dữ liệu dựa trên limit và điều kiện status = 1
	err = db.Table("product_prices").
		Where("status = ?", 1).
		Order("created_at DESC").
		Limit(limit).
		Find(&results).Error

	if err != nil {
		return nil, 0, err
	}

	return &results, totalRecords, nil
}
