package carousel_service

import (
	"be-hoatieu/models"

	"gorm.io/gorm"
)

type Carousel struct {
	gorm.Model
	Image  string `gorm:"column:image" json:"image"`
	Status bool   `gorm:"column:status" json:"status"`
}

type ObjectCarousel struct {
	ID     uint   `gorm:"column:ID" json:"ID"`
	Image  string `gorm:"column:image" json:"image"`
	Status bool   `gorm:"column:status" json:"status"`
}

func (a *Carousel) GetCarousel_Service(limit int) (*[]models.ObjectCarousel, error) {
	item, err := models.GetCarousel_Model(limit)
	if err != nil {
		return nil, err
	}
	return item, nil
}
