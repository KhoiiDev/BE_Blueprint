package models

import (
	"gorm.io/gorm"
)

type Introduction struct {
	gorm.Model
	Content string `gorm:"column:content" json:"content"`
	Status  bool   `gorm:"column:status" json:"status"`
}

type ObjectIntroduction struct {
	ID      uint   `gorm:"column:ID" json:"ID"`
	Content string `gorm:"column:content" json:"content"`
	Status  bool   `gorm:"column:status" json:"status"`
}

func GetIntroduction_Model() (*[]ObjectIntroduction, error) {
	var results []ObjectIntroduction

	// Truy vấn dữ liệu dựa trên limit và điều kiện status = 1
	err := db.Table("introductions").
		Where("status = ?", 1).
		Order("created_at DESC").
		Limit(1).
		Find(&results).Error

	if err != nil {
		return nil, err
	}

	return &results, nil
}
