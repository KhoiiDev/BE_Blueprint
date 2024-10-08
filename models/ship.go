package models

import (
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

func GetShip_Model(limit int, page int) (*[]ObjectShip, int64, error) {
	var results []ObjectShip
	var totalRecords int64

	// Tính offset dựa trên limit và page
	offset := (page - 1) * limit

	// Truy vấn để đếm tổng số bản ghi có status = 1
	err := db.Table("ships").Where("status = ?", 1).Count(&totalRecords).Error
	if err != nil {
		return nil, 0, err
	}

	// Truy vấn dữ liệu dựa trên limit, offset và điều kiện status = 1
	err = db.Table("ships").
		Where("status = ?", 1).
		Order("created_at DESC").
		Offset(offset).
		Limit(limit).
		Find(&results).Error

	if err != nil {
		return nil, 0, err
	}

	return &results, totalRecords, nil
}
