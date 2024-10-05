package models

import (
	"gorm.io/gorm"
)

type Carousel struct {
	gorm.Model
	Image  string `gorm:"column:image" json:"image"`
	Status bool   `gorm:"column:status" json:"status"`
}

type ObjectCarousel struct {
	ID     uint   `gorm:"column:id" json:"id"`
	Image  string `gorm:"column:image" json:"image"`
	Status bool   `gorm:"column:status" json:"status"`
}

func GetCarousel_Model(limit int) (*[]ObjectCarousel, error) {
	var results []ObjectCarousel

	// Truy vấn dữ liệu dựa trên limit và điều kiện status = 1
	err := db.Table("carousels").
		Where("status = ?", 1).
		Order("created_at DESC").
		Limit(limit).
		Find(&results).Error

	if err != nil {
		return nil, err
	}

	return &results, nil
}
