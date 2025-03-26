package dichvu_service

import (
	"be-hoatieu/models"

	"gorm.io/gorm"
)

type Dichvu struct {
	gorm.Model
	Title    string `gorm:"column:title" json:"title"`
	SubTitle string `gorm:"column:subtitle" json:"subtitle"`
	Content  string `gorm:"column:content" json:"content"`
	Postdate string `gorm:"column:postdate" json:"postdate"`

	Image  string `gorm:"column:image" json:"image"`
	Pdfurl string `gorm:"column:pdfurl" json:"pdfurl"`
	Status bool   `gorm:"column:status" json:"status"`
}

type ObjectDichvu struct {
	ID       uint   `gorm:"column:id" json:"id"`
	Title    string `gorm:"column:title" json:"title"`
	SubTitle string `gorm:"column:subtitle" json:"subtitle"`
	Content  string `gorm:"column:content" json:"content"`
	Image    string `gorm:"column:image" json:"image"`
	Pdfurl   string `gorm:"column:pdfurl" json:"pdfurl"`
	Postdate string `gorm:"column:postdate" json:"postdate"`

	Status bool `gorm:"column:status" json:"status"`
}

func (a *Dichvu) GetDichvu_Service(limit int, page int, name string, showHidden bool) (*[]models.ObjectDichvu, int64, error) {
	item, totalRecords, err := models.GetDichvu_Model(limit, page, name, showHidden)
	if err != nil {
		return nil, 0, err
	}
	return item, totalRecords, nil
}

func (n *Dichvu) CreateDichvu_Service() error {
	item := map[string]interface{}{
		"title":    n.Title,
		"subtitle": n.SubTitle,
		"content":  n.Content,
		"image":    n.Image,
		"postdate": n.Postdate,

		"status": n.Status,
		"pdfurl": n.Pdfurl,
	}
	if err := models.CreateDichvu_Model(item); err != nil { // Change function name to match your model's function
		return err
	}
	return nil
}

func (a *Dichvu) UpdateDichvu_Service(id string) error {
	item := map[string]interface{}{
		"title":    a.Title,
		"subtitle": a.SubTitle,
		"content":  a.Content,
		"image":    a.Image,
		"postdate": a.Postdate,
		"status":   a.Status,
		"pdfurl":   a.Pdfurl,
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
