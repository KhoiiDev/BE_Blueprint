package models

import "gorm.io/gorm"

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

func GetProduct_Model(limit int) (*[]ObjectProduct, error) {
	var results []ObjectProduct

	// Truy vấn dữ liệu dựa trên limit và điều kiện status = 1
	err := db.Table("products").
		Where("status = ?", 1).
		Order("created_at DESC").
		Limit(limit).
		Find(&results).Error

	if err != nil {
		return nil, err
	}

	return &results, nil
}
