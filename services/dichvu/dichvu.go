package dichvu_service

import (
	"be-hoatieu/models"

	"gorm.io/gorm"
)

type Dichvu struct {
	gorm.Model
	Title   string `gorm:"column:title" json:"title"`
	Image   string `gorm:"column:image" json:"image"`
	Pdfdata string `gorm:"column:pdfdata" json:"pdfdata"`
	Status  bool   `gorm:"column:status" json:"status"`
}

type ObjectDichvu struct {
	ID      uint   `gorm:"column:id" json:"id"`
	Title   string `gorm:"column:title" json:"title"`
	Image   string `gorm:"column:image" json:"image"`
	Pdfdata string `gorm:"column:pdfdata" json:"pdfdata"`
	Status  bool   `gorm:"column:status" json:"status"`
}

func (a *Dichvu) GetDichvu_Service(limit int, page int, showHidden bool) (*[]models.ObjectDichvu, int64, error) {
	item, totalRecords, err := models.GetDichvu_Model(limit, page, showHidden)
	if err != nil {
		return nil, 0, err
	}
	return item, totalRecords, nil
}

func (n *Dichvu) CreateDichvu_Service() error {
	item := map[string]interface{}{
		"title":   n.Title,
		"image":   n.Image,
		"status":  n.Status,
		"pdfdata": n.Pdfdata,
	}
	if err := models.CreateDichvu_Model(item); err != nil { // Change function name to match your model's function
		return err
	}
	return nil
}

func (a *Dichvu) UpdateDichvu_Service(id string) error {
	item := map[string]interface{}{
		"title":   a.Title,
		"image":   a.Image,
		"status":  a.Status,
		"pdfdata": a.Pdfdata,
	}
	if err := models.UpdateDichvu_Model(id, item); err != nil {
		return err
	}
	return nil
}

func (a *Dichvu) DeleteDichvu_Service(id string) (bool, error) {
	if err := models.DeleteDichvu_Model(id); err != nil {
		return false, err
	}
	return true, nil
}
