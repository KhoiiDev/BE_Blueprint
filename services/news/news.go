package news_service

import (
	"be-hoatieu/models"

	"gorm.io/gorm"
)

type News struct {
	gorm.Model
	Title      string `gorm:"column:title" json:"title"`
	TitleEN    string `gorm:"column:title_en" json:"title_en"`
	SubTitle   string `gorm:"column:subtitle" json:"subtitle"`
	SubTitleEN string `gorm:"column:subtitle_en" json:"subtitle_en"`
	Image      string `gorm:"column:image" json:"image"`
	Status     bool   `gorm:"column:status" json:"status"`
	Content    string `gorm:"column:content" json:"content"`
	ContentEN  string `gorm:"column:content_en" json:"content_en"`
	Postdate   string `gorm:"column:postdate" json:"postdate"`
}

type ObjectNews struct {
	ID         uint   `gorm:"column:ID" json:"ID"`
	Title      string `gorm:"column:title" json:"title"`
	TitleEN    string `gorm:"column:title_en" json:"title_en"`
	SubTitle   string `gorm:"column:subtitle" json:"subtitle"`
	SubTitleEN string `gorm:"column:subtitle_en" json:"subtitle_en"`
	Image      string `gorm:"column:image" json:"image"`
	Status     bool   `gorm:"column:status" json:"status"`
	Content    string `gorm:"column:content" json:"content"`
	ContentEN  string `gorm:"column:content_en" json:"content_en"`
	Postdate   string `gorm:"column:postdate" json:"postdate"`
}

func (a *News) GetNews_Service(limit int, page int, name string, showHidden bool) (*[]models.ObjectNews, int64, error) {
	item, totalRecords, err := models.GetNews_Model(limit, page, name, showHidden)
	if err != nil {
		return nil, totalRecords, err
	}
	return item, totalRecords, nil
}

func (n *News) CreateNews_Service() error {
	item := map[string]interface{}{
		"title":       n.Title,
		"title_en":    n.TitleEN,
		"subtitle":    n.SubTitle,
		"subtitle_en": n.SubTitleEN,
		"image":       n.Image,
		"status":      n.Status,
		"content":     n.Content,
		"content_en":  n.ContentEN,
		"postdate":    n.Postdate,
	}
	if err := models.CreateNews_Model(item); err != nil {
		return err
	}
	return nil
}

func (a *News) UpdateNews_Service(id string) error {
	item := map[string]interface{}{
		"title":       a.Title,
		"title_en":    a.TitleEN,
		"subtitle":    a.SubTitle,
		"subtitle_en": a.SubTitleEN,
		"image":       a.Image,
		"status":      a.Status,
		"content":     a.Content,
		"content_en":  a.ContentEN,
		"postdate":    a.Postdate,
	}
	if err := models.UpdateNews_Model(id, item); err != nil {
		return err
	}
	return nil
}

func (a *News) DeleteNews_Service(id string) (bool, error) {
	if err := models.DeleteNews_Model(id); err != nil {
		return false, err
	}
	return true, nil
}
