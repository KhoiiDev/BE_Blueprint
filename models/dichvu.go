package models

import (
	"time"

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

func GetDichvu_Model(limit int, page int, name string, showHidden bool) (*[]ObjectDichvu, int64, error) {
	var results []ObjectDichvu
	totalRecords := int64(0)
	offset := (page - 1) * limit
	var err error

	if showHidden {
		err = db.Table("dichvus").Where("deleted_at IS NULL").Count(&totalRecords).Error
		if err != nil {
			return nil, 0, err
		}
		err = db.Table("dichvus").
			Where("deleted_at IS NULL").
			Where("title LIKE ?", "%"+name+"%").
			Order("created_at DESC").
			Limit(limit).
			Offset(offset).
			Find(&results).Error
	} else {
		err = db.Table("dichvus").
			Where("status = ? AND deleted_at IS NULL", true).
			Count(&totalRecords).Error
		if err != nil {
			return nil, 0, err
		}
		err = db.Table("dichvus").
			Where("status = ? AND deleted_at IS NULL", true).
			Where("title LIKE ?", "%"+name+"%").
			Order("created_at DESC").
			Limit(limit).
			Offset(offset).
			Find(&results).Error
	}

	if err != nil {
		return nil, 0, err
	}

	return &results, totalRecords, nil
}

func CreateDichvu_Model(data map[string]interface{}) error {
	item := Dichvu{
		Title:      data["title"].(string),
		TitleEN:    data["title_en"].(string),
		SubTitle:   data["subtitle"].(string),
		SubTitleEN: data["subtitle_en"].(string),
		Content:    data["content"].(string),
		ContentEN:  data["content_en"].(string),
		Image:      data["image"].(string),
		Pdfurl:     data["pdfurl"].(string),
		Postdate:   data["postdate"].(string),
		Status:     data["status"].(bool),
	}

	return db.Create(&item).Error
}

func UpdateDichvu_Model(id string, data map[string]interface{}) error {
	item := Dichvu{
		Title:      data["title"].(string),
		TitleEN:    data["title_en"].(string),
		SubTitle:   data["subtitle"].(string),
		SubTitleEN: data["subtitle_en"].(string),
		Content:    data["content"].(string),
		ContentEN:  data["content_en"].(string),
		Image:      data["image"].(string),
		Pdfurl:     data["pdfurl"].(string),
		Postdate:   data["postdate"].(string),
		Status:     data["status"].(bool),
	}

	if err := db.Model(&item).Where("id = ?", id).Updates(map[string]interface{}{
		"title":       item.Title,
		"title_en":    item.TitleEN,
		"subtitle":    item.SubTitle,
		"subtitle_en": item.SubTitleEN,
		"content":     item.Content,
		"content_en":  item.ContentEN,
		"image":       item.Image,
		"pdfurl":      item.Pdfurl,
		"postdate":    item.Postdate,
		"status":      item.Status,
		"updated_at":  time.Now(),
	}).Error; err != nil {
		return err
	}

	return nil
}

func CheckDeletedAtDichvu_Model(id string) (bool, error) {
	var item Dichvu
	if err := db.Select("deleted_at").Where("id = ?", id).First(&item).Error; err != nil {
		return false, err
	}
	return item.DeletedAt.Valid, nil
}

func DeleteDichvu_Model(id string) error {
	isDeleted, err := CheckDeletedAtDichvu_Model(id)
	if err != nil {
		return err
	}
	if !isDeleted {
		if err := db.Delete(&Dichvu{}, id).Error; err != nil {
			return err
		}
	}
	return nil
}
