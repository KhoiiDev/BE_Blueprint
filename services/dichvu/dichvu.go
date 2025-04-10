package services

import (
	"be-hoatieu/models"

	"gorm.io/gorm"
)

type Dichvu struct {
	gorm.Model
	Title      string `gorm:"column:title" json:"title"`
	TitleEN    string `gorm:"column:title_en" json:"title_en"`
	SubTitle   string `gorm:"column:subtitle" json:"subtitle"`
	SubTitleEN string `gorm:"column:subtitle_en" json:"subtitle_en"`
	Content    string `gorm:"column:content" json:"content"`
	ContentEN  string `gorm:"column:content_en" json:"content_en"`
	Image      string `gorm:"column:image" json:"image"`
	Pdfurl     string `gorm:"column:pdfurl" json:"pdfurl"`
	Postdate   string `gorm:"column:postdate" json:"postdate"`
	Status     bool   `gorm:"column:status" json:"status"`
}

type ObjectDichvu struct {
	ID         uint   `gorm:"column:id" json:"id"`
	Title      string `gorm:"column:title" json:"title"`
	TitleEN    string `gorm:"column:title_en" json:"title_en"`
	SubTitle   string `gorm:"column:subtitle" json:"subtitle"`
	SubTitleEN string `gorm:"column:subtitle_en" json:"subtitle_en"`
	Content    string `gorm:"column:content" json:"content"`
	ContentEN  string `gorm:"column:content_en" json:"content_en"`
	Image      string `gorm:"column:image" json:"image"`
	Pdfurl     string `gorm:"column:pdfurl" json:"pdfurl"`
	Postdate   string `gorm:"column:postdate" json:"postdate"`
	Status     bool   `gorm:"column:status" json:"status"`
}

// Lấy danh sách dịch vụ có phân trang + tìm kiếm + lọc ẩn/hiện
func (a *Dichvu) GetDichvu_Service(limit int, page int, name string, showHidden bool) (*[]models.ObjectDichvu, int64, error) {
	item, totalRecords, err := models.GetDichvu_Model(limit, page, name, showHidden)
	if err != nil {
		return nil, 0, err
	}
	return item, totalRecords, nil
}

// Tạo mới dịch vụ
func (n *Dichvu) CreateDichvu_Service() error {
	item := map[string]interface{}{
		"title":       n.Title,
		"title_en":    n.TitleEN,
		"subtitle":    n.SubTitle,
		"subtitle_en": n.SubTitleEN,
		"content":     n.Content,
		"content_en":  n.ContentEN,
		"image":       n.Image,
		"postdate":    n.Postdate,
		"status":      n.Status,
		"pdfurl":      n.Pdfurl,
	}
	if err := models.CreateDichvu_Model(item); err != nil {
		return err
	}
	return nil
}

// Cập nhật dịch vụ
func (a *Dichvu) UpdateDichvu_Service(id string) error {
	item := map[string]interface{}{
		"title":       a.Title,
		"title_en":    a.TitleEN,
		"subtitle":    a.SubTitle,
		"subtitle_en": a.SubTitleEN,
		"content":     a.Content,
		"content_en":  a.ContentEN,
		"image":       a.Image,
		"postdate":    a.Postdate,
		"status":      a.Status,
		"pdfurl":      a.Pdfurl,
	}
	if err := models.UpdateDichvu_Model(id, item); err != nil {
		return err
	}
	return nil
}

// Xóa dịch vụ
func (a *Dichvu) DeleteDichvu_Service(id string) (bool, error) {
	if err := models.DeleteDichvu_Model(id); err != nil {
		return false, err
	}
	return true, nil
}
