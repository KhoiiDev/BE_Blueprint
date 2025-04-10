package models

import (
	"time"

	"gorm.io/gorm"
)

type Servicelist struct {
	gorm.Model
	Title      string `gorm:"column:title" json:"title"`
	TitleEN    string `gorm:"column:title_en" json:"title_en"`
	SubTitle   string `gorm:"column:subtitle" json:"subtitle"`
	SubTitleEN string `gorm:"column:subtitle_en" json:"subtitle_en"`
	Pdfurl     string `gorm:"column:pdfurl" json:"pdfurl"`
	Postdate   string `gorm:"column:postdate" json:"postdate"`
	Image      string `gorm:"column:image" json:"image"`
	Status     bool   `gorm:"column:status" json:"status"`
	Content    string `gorm:"column:content" json:"content"`
	ContentEN  string `gorm:"column:content_en" json:"content_en"`
}

type ObjectServicelist struct {
	ID         uint   `gorm:"column:id" json:"id"`
	Title      string `gorm:"column:title" json:"title"`
	TitleEN    string `gorm:"column:title_en" json:"title_en"`
	SubTitle   string `gorm:"column:subtitle" json:"subtitle"`
	SubTitleEN string `gorm:"column:subtitle_en" json:"subtitle_en"`
	Pdfurl     string `gorm:"column:pdfurl" json:"pdfurl"`
	Postdate   string `gorm:"column:postdate" json:"postdate"`
	Image      string `gorm:"column:image" json:"image"`
	Status     bool   `gorm:"column:status" json:"status"`
	Content    string `gorm:"column:content" json:"content"`
	ContentEN  string `gorm:"column:content_en" json:"content_en"`
}

func GetServiceList_Model(limit int, page int, name string, showHidden bool) (*[]ObjectServicelist, int64, error) {
	var results []ObjectServicelist
	totalRecords := int64(0)
	var err error

	offset := (page - 1) * limit

	if showHidden {
		err = db.Table("servicelists").Where("deleted_at IS NULL").Count(&totalRecords).Error
		if err != nil {
			return nil, 0, err
		}

		err = db.Table("servicelists").
			Where("deleted_at IS NULL").
			Order("created_at DESC").
			Where("title LIKE ?", "%"+name+"%").
			Limit(limit).
			Offset(offset).
			Find(&results).Error
	} else {
		err = db.Table("servicelists").Where("status = ? AND deleted_at IS NULL", 1).Count(&totalRecords).Error
		if err != nil {
			return nil, 0, err
		}
		err = db.Table("servicelists").
			Where("status = ? AND deleted_at IS NULL", 1).
			Order("created_at DESC").
			Where("title LIKE ?", "%"+name+"%").
			Limit(limit).
			Offset(offset).
			Find(&results).Error
	}

	if err != nil {
		return nil, 0, err
	}

	return &results, totalRecords, nil
}

func CreateServiceList_Model(data map[string]interface{}) error {
	item := Servicelist{
		Title:      data["title"].(string),
		TitleEN:    data["title_en"].(string),
		SubTitle:   data["subtitle"].(string),
		SubTitleEN: data["subtitle_en"].(string),
		Postdate:   data["postdate"].(string),
		Image:      data["image"].(string),
		Pdfurl:     data["pdfurl"].(string),
		Status:     data["status"].(bool),
		Content:    data["content"].(string),
		ContentEN:  data["content_en"].(string),
	}

	result := db.Create(&item)
	if err := result.Error; err != nil {
		return err
	}
	return nil
}

func UpdateServiceList_Model(id string, data map[string]interface{}) error {
	item := Servicelist{
		Title:      data["title"].(string),
		TitleEN:    data["title_en"].(string),
		SubTitle:   data["subtitle"].(string),
		SubTitleEN: data["subtitle_en"].(string),
		Postdate:   data["postdate"].(string),
		Image:      data["image"].(string),
		Pdfurl:     data["pdfurl"].(string),
		Status:     data["status"].(bool),
		Content:    data["content"].(string),
		ContentEN:  data["content_en"].(string),
	}

	if err := db.Model(&item).Where("id = ?", id).Updates(map[string]interface{}{
		"title":       item.Title,
		"title_en":    item.TitleEN,
		"subtitle":    item.SubTitle,
		"subtitle_en": item.SubTitleEN,
		"pdfurl":      item.Pdfurl,
		"postdate":    item.Postdate,
		"image":       item.Image,
		"status":      item.Status,
		"content":     item.Content,
		"content_en":  item.ContentEN,
		"updated_at":  time.Now(),
	}).Error; err != nil {
		return err
	}
	return nil
}

func DeleteServiceList_Model(id string) error {
	var dichvu Servicelist
	if err := db.Select("deleted_at").Where("id = ?", id).First(&dichvu).Error; err != nil {
		return err
	}

	if err := db.Delete(&Servicelist{}, id).Error; err != nil {
		return err
	}

	return nil
}
