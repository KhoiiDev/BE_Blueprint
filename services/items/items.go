package items_service

import (
	"be-hoatieu/models"

	"gorm.io/gorm"
)

type Items struct {
	gorm.Model
	Title    string `gorm:"column:title" json:"title"`
	SubTitle string `gorm:"column:subtitle" json:"subtitle"`
	Image    string `gorm:"column:image" json:"image"`
	Status   bool   `gorm:"column:status" json:"status"`
	Pdfurl   string `gorm:"column:pdfurl" json:"pdfurl"`
	Content  string `gorm:"column:content" json:"content"`
	Videourl string `gorm:"column:videourl" json:"videourl"`

	Postdate string `gorm:"column:postdate" json:"postdate"`
	ItemType string `gorm:"column:itemtype" json:"itemtype"`
}

type ObjectItems struct {
	ID       uint   `gorm:"column:ID" json:"ID"`
	Title    string `gorm:"column:title" json:"title"`
	SubTitle string `gorm:"column:subtitle" json:"subtitle"`
	Image    string `gorm:"column:image" json:"image"`
	Status   bool   `gorm:"column:status" json:"status"`
	Pdfurl   string `gorm:"column:pdfurl" json:"pdfurl"`
	Videourl string `gorm:"column:videourl" json:"videourl"`
	Content  string `gorm:"column:content" json:"content"`
	Postdate string `gorm:"column:postdate" json:"postdate"`
	ItemType string `gorm:"column:itemtype" json:"itemtype"`
}

func (a *Items) GetItems_Service(limit int, page int, showHidden bool, item_type string) (*[]models.ObjectItems, int64, error) {
	item, totalRecords, err := models.GetItems_Model(limit, page, showHidden, item_type)
	if err != nil {
		return nil, totalRecords, err
	}
	return item, totalRecords, nil
}

func (n *Items) CreateItems_Service() error {
	item := map[string]interface{}{
		"title":    n.Title,
		"subtitle": n.SubTitle,

		"image":    n.Image,
		"pdfurl":   n.Pdfurl,
		"status":   n.Status,
		"videourl": n.Videourl,

		"content":  n.Content,
		"postdate": n.Postdate,
		"itemtype": n.ItemType,
	}
	if err := models.CreateItems_Model(item); err != nil { // Change function name to match your model's function
		return err
	}
	return nil
}

func (a *Items) UpdateItems_Service(id string) error {
	item := map[string]interface{}{
		"title":    a.Title,
		"subtitle": a.SubTitle,
		"image":    a.Image,
		"pdfurl":   a.Pdfurl,
		"videourl": a.Videourl,
		"status":   a.Status,
		"content":  a.Content,
		"postdate": a.Postdate,
	}
	if err := models.UpdateItems_Model(id, item); err != nil {
		return err
	}
	return nil
}

func (a *Items) DeleteItems_Service(id string) (bool, error) {
	if err := models.DeleteItems_Model(id); err != nil {
		return false, err
	}
	return true, nil
}
