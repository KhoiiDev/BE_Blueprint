package TideCalendar_service

import (
	"be-hoatieu/models"
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

func (a *TideCalendar) GetTideCalendar_Service() (*[]models.ObjectTideCalendar, error) {
	item, err := models.GetTideCalendar_Model()
	if err != nil {
		return nil, err
	}
	return item, nil
}
