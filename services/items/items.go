package items_service

import (
	"be-hoatieu/models"

	"gorm.io/gorm"
)

type Items struct {
	gorm.Model
	Title      string `gorm:"column:title" json:"title"`
	TitleEN    string `gorm:"column:title_en" json:"title_en"`
	SubTitle   string `gorm:"column:subtitle" json:"subtitle"`
	SubTitleEN string `gorm:"column:subtitle_en" json:"subtitle_en"`
	Image      string `gorm:"column:image" json:"image"`
	Status     bool   `gorm:"column:status" json:"status"`
	Content    string `gorm:"column:content" json:"content"`
	ContentEN  string `gorm:"column:content_en" json:"content_en"`
	Pdfurl     string `gorm:"column:pdfurl" json:"pdfurl"`
	Videourl   string `gorm:"column:videourl" json:"videourl"`
	Postdate   string `gorm:"column:postdate" json:"postdate"`
	ItemType   string `gorm:"column:itemtype" json:"itemtype"`
}

type ObjectItems struct {
	ID         uint   `gorm:"column:id" json:"id"`
	Title      string `gorm:"column:title" json:"title"`
	TitleEN    string `gorm:"column:title_en" json:"title_en"`
	SubTitle   string `gorm:"column:subtitle" json:"subtitle"`
	SubTitleEN string `gorm:"column:subtitle_en" json:"subtitle_en"`
	Image      string `gorm:"column:image" json:"image"`
	Status     bool   `gorm:"column:status" json:"status"`
	Content    string `gorm:"column:content" json:"content"`
	ContentEN  string `gorm:"column:content_en" json:"content_en"`
	Pdfurl     string `gorm:"column:pdfurl" json:"pdfurl"`
	Videourl   string `gorm:"column:videourl" json:"videourl"`
	Postdate   string `gorm:"column:postdate" json:"postdate"`
	ItemType   string `gorm:"column:itemtype" json:"itemtype"`
}

func (a *Items) GetItems_Service(limit int, page int, showHidden bool, name string, item_type string) (*[]models.ObjectItems, int64, error) {
	item, totalRecords, err := models.GetItems_Model(limit, page, showHidden, name, item_type)
	if err != nil {
		return nil, totalRecords, err
	}
	return item, totalRecords, nil
}

func (n *Items) CreateItems_Service() error {
	item := map[string]interface{}{
		"title":       n.Title,
		"title_en":    n.TitleEN,
		"subtitle":    n.SubTitle,
		"subtitle_en": n.SubTitleEN,
		"image":       n.Image,
		"pdfurl":      n.Pdfurl,
		"status":      n.Status,
		"videourl":    n.Videourl,
		"content":     n.Content,
		"content_en":  n.ContentEN,
		"postdate":    n.Postdate,
		"itemtype":    n.ItemType,
	}
	if err := models.CreateItems_Model(item); err != nil {
		return err
	}
	return nil
}

func (a *Items) UpdateItems_Service(id string) error {
	item := map[string]interface{}{
		"title":       a.Title,
		"title_en":    a.TitleEN,
		"subtitle":    a.SubTitle,
		"subtitle_en": a.SubTitleEN,
		"image":       a.Image,
		"pdfurl":      a.Pdfurl,
		"videourl":    a.Videourl,
		"status":      a.Status,
		"content":     a.Content,
		"content_en":  a.ContentEN,
		"postdate":    a.Postdate,
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
