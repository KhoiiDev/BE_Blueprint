package TideCalendar_service

import (
	"be-hoatieu/models"
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

func (a *Stream) GetStream_Service() (*[]models.ObjectStream, error) {
	item, err := models.GetStream_Model()
	if err != nil {
		return nil, err
	}
	return item, nil
}
