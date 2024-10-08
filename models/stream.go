package models

import (
	"time"

	"gorm.io/gorm"
)

type Stream struct {
	gorm.Model
	Url      string    `gorm:"column:url" json:"url"`
	Title    string    `gorm:"column:title" json:"title"`
	PostDate time.Time `gorm:"column:created_at" json:"createdAt"`
	Status   bool      `gorm:"column:status" json:"status"`
}

type ObjectStream struct {
	ID       uint      `gorm:"column:ID" json:"ID"`
	Title    string    `gorm:"column:title" json:"title"`
	Url      string    `gorm:"column:url" json:"url"`
	PostDate time.Time `gorm:"column:created_at" json:"createdAt"`
	Status   bool      `gorm:"column:status" json:"status"`
}

func GetStream_Model() (*[]ObjectStream, error) {
	var results []ObjectStream
	// Truy vấn dữ liệu dựa trên limit, offset và điều kiện status = 1
	err := db.Table("streams").
		Where("status = ?", 1).
		Order("created_at DESC").
		Find(&results).Error

	if err != nil {
		return nil, err
	}
	return &results, nil
}
