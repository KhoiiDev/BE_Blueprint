package serviceList_service

import (
	"be-hoatieu/models"

	"gorm.io/gorm"
)

// Struct lưu trữ dữ liệu nhận vào từ client hoặc form
type Servicelist struct {
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
	Pdfurl     string `gorm:"column:pdfurl" json:"pdfurl"`
}

// DTO struct trả về cho client (không cần gorm.Model)
type ObjectServicelist struct {
	ID         uint   `gorm:"column:id" json:"id"`
	Title      string `gorm:"column:title" json:"title"`
	TitleEN    string `gorm:"column:title_en" json:"title_en"`
	SubTitle   string `gorm:"column:subtitle" json:"subtitle"`
	SubTitleEN string `gorm:"column:subtitle_en" json:"subtitle_en"`
	Image      string `gorm:"column:image" json:"image"`
	Status     bool   `gorm:"column:status" json:"status"`
	Content    string `gorm:"column:content" json:"content"`
	ContentEN  string `gorm:"column:content_en" json:"content_en"`
	Postdate   string `gorm:"column:postdate" json:"postdate"`
	Pdfurl     string `gorm:"column:pdfurl" json:"pdfurl"`
}

// Lấy danh sách ServiceList có phân trang, tìm kiếm
func (a *Servicelist) GetServiceList_Service(limit int, page int, name string, showHidden bool) (*[]models.ObjectServicelist, int64, error) {
	item, totalRecords, err := models.GetServiceList_Model(limit, page, name, showHidden)
	if err != nil {
		return nil, 0, err
	}
	return item, totalRecords, nil
}

// Tạo mới một ServiceList
func (s *Servicelist) CreateServiceList_Service() error {
	item := map[string]interface{}{
		"title":       s.Title,
		"title_en":    s.TitleEN,
		"subtitle":    s.SubTitle,
		"subtitle_en": s.SubTitleEN,
		"postdate":    s.Postdate,
		"pdfurl":      s.Pdfurl,
		"image":       s.Image,
		"status":      s.Status,
		"content":     s.Content,
		"content_en":  s.ContentEN,
	}
	if err := models.CreateServiceList_Model(item); err != nil {
		return err
	}
	return nil
}

// Cập nhật ServiceList theo ID
func (s *Servicelist) UpdateServiceList_Service(id string) error {
	item := map[string]interface{}{
		"title":       s.Title,
		"title_en":    s.TitleEN,
		"subtitle":    s.SubTitle,
		"subtitle_en": s.SubTitleEN,
		"postdate":    s.Postdate,
		"pdfurl":      s.Pdfurl,
		"image":       s.Image,
		"status":      s.Status,
		"content":     s.Content,
		"content_en":  s.ContentEN,
	}
	if err := models.UpdateServiceList_Model(id, item); err != nil {
		return err
	}
	return nil
}

// Xoá ServiceList theo ID
func (s *Servicelist) DeleteServiceList_Service(id string) (bool, error) {
	if err := models.DeleteServiceList_Model(id); err != nil {
		return false, err
	}
	return true, nil
}
