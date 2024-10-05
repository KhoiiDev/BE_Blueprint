package models

import (
	"time"

	"gorm.io/gorm"
)

type News struct {
	gorm.Model
	Title    string    `gorm:"column:title" json:"title"`
	Url      string    `gorm:"column:url" json:"url"`
	Image    string    `gorm:"column:image" json:"image"`
	Status   bool      `gorm:"column:status" json:"status"`
	Content  string    `gorm:"column:content" json:"content"`
	PostDate time.Time `gorm:"column:postdate" json:"postdate"`
}

type ObjectNews struct {
	ID       uint      `gorm:"column:ID" json:"ID"`
	Title    string    `gorm:"column:title" json:"title"`
	Url      string    `gorm:"column:url" json:"url"`
	Image    string    `gorm:"column:image" json:"image"`
	Status   bool      `gorm:"column:status" json:"status"`
	Content  string    `gorm:"column:content" json:"content"`
	PostDate time.Time `gorm:"column:postdate" json:"postdate"`
}

func GetNews_Model(limit int) (*[]ObjectNews, error) {
	var results []ObjectNews

	// Truy vấn dữ liệu dựa trên limit và điều kiện status = 1
	err := db.Table("news").
		Where("status = ?", 1).
		Order("created_at DESC").
		Limit(limit).
		Find(&results).Error

	if err != nil {
		return nil, err
	}

	return &results, nil
}
