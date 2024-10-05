package product_service

import (
	"be-hoatieu/models"
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
	PostDate time.Time `gorm:"column:postdate" json:"postdate"`
}

func (a *News) GetNews_Service(limit int) (*[]models.ObjectNews, error) {
	item, err := models.GetNews_Model(limit)
	if err != nil {
		return nil, err
	}
	return item, nil
}
