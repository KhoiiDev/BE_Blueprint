package news_service

import (
	"be-hoatieu/models"

	"gorm.io/gorm"
)

type News struct {
	gorm.Model
	Title    string `gorm:"column:title" json:"title"`
	SubTitle string `gorm:"column:subtitle" json:"subtitle"`
	Image    string `gorm:"column:image" json:"image"`
	Status   bool   `gorm:"column:status" json:"status"`
	Content  string `gorm:"column:content" json:"content"`
	Postdate string `gorm:"column:postdate" json:"postdate"`
}

type ObjectNews struct {
	ID       uint   `gorm:"column:ID" json:"ID"`
	Title    string `gorm:"column:title" json:"title"`
	SubTitle string `gorm:"column:subtitle" json:"subtitle"`
	Image    string `gorm:"column:image" json:"image"`
	Status   bool   `gorm:"column:status" json:"status"`
	Content  string `gorm:"column:content" json:"content"`
	Postdate string `gorm:"column:postdate" json:"postdate"`
}

func (a *News) GetNews_Service(limit int, page int, showHidden bool) (*[]models.ObjectNews, int64, error) {
	item, totalRecords, err := models.GetNews_Model(limit, page, showHidden)
	if err != nil {
		return nil, totalRecords, err
	}
	return item, totalRecords, nil
}

func (n *News) CreateNews_Service() error {
	item := map[string]interface{}{
		"title":    n.Title,
		"subtitle": n.SubTitle,
		"image":    n.Image,
		"status":   n.Status,
		"content":  n.Content,
		"postdate": n.Postdate,
	}
	if err := models.CreateNews_Model(item); err != nil { // Change function name to match your model's function
		return err
	}
	return nil
}

func (a *News) UpdateNews_Service(id string) error {
	item := map[string]interface{}{
		"title":    a.Title,
		"subtitle": a.SubTitle,
		"image":    a.Image,
		"status":   a.Status,
		"content":  a.Content,
		"postdate": a.Postdate,
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
