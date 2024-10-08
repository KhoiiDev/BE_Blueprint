package models

import (
	"time"

	"gorm.io/gorm"
)

type TideCalendar struct {
	gorm.Model
	Url      string    `gorm:"column:url" json:"url"`
	PostDate time.Time `gorm:"column:created_at" json:"createdAt"`
	Status   bool      `gorm:"column:status" json:"status"`
}

type ObjectTideCalendar struct {
	ID       uint      `gorm:"column:ID" json:"ID"`
	Url      string    `gorm:"column:url" json:"url"`
	PostDate time.Time `gorm:"column:created_at" json:"createdAt"`
	Status   bool      `gorm:"column:status" json:"status"`
}

func GetTideCalendar_Model() (*[]ObjectTideCalendar, error) {
	var results []ObjectTideCalendar
	// Truy vấn dữ liệu dựa trên limit, offset và điều kiện status = 1
	err := db.Table("tide_calendars").
		Where("status = ?", 1).
		Order("created_at DESC").
		Limit(1).
		Find(&results).Error

	if err != nil {
		return nil, err
	}

	return &results, nil
}
