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
	ID     uint   `gorm:"column:id" json:"id"`
	Image  string `gorm:"column:image" json:"image"`
	Status bool   `gorm:"column:status" json:"status"`
}

func (a *Carousel) GetCarousel_Service(limit int, page int, showHidden bool) (*[]models.ObjectCarousel, int64, error) {
	item, totalRecords, err := models.GetCarousel_Model(limit, page, showHidden)
	if err != nil {
		return nil, 0, err
	}
	return item, totalRecords, nil
}

func (n *Carousel) CreateCarousel_Service() error {
	item := map[string]interface{}{
		"image":  n.Image,
		"status": n.Status,
	}
	if err := models.CreateCarousel_Model(item); err != nil { // Change function name to match your model's function
		return err
	}
	return nil
}

func (a *Carousel) UpdateCarousel_Service(id string) error {
	item := map[string]interface{}{
		"image":  a.Image,
		"status": a.Status,
	}
	if err := models.UpdateCarousel_Model(id, item); err != nil {
		return err
	}
	return nil
}

func (a *Carousel) DeleteCarousel_Service(id string) (bool, error) {
	if err := models.DeleteCarousel_Model(id); err != nil {
		return false, err
	}
	return true, nil
}
