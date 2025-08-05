package models

import (
	"time"

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
	Pdfurl     string `gorm:"column:pdfurl" json:"pdfurl"`
}

type ObjectNews struct {
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

func GetNews_Model(limit int, page int, name string, showHidden bool) (*[]ObjectNews, int64, error) {
	var results []ObjectNews
	totalRecords := int64(0)
	var err error

	offset := (page - 1) * limit

	if showHidden {
		err = db.Table("news").Where("deleted_at IS NULL").Count(&totalRecords).Error
		if err != nil {
			return nil, 0, err
		}

		err = db.Table("news").
			Where("deleted_at IS NULL").
			Where("title LIKE ?", "%"+name+"%").
			Order("created_at DESC").
			Limit(limit).
			Offset(offset).
			Find(&results).Error
	} else {
		err = db.Table("news").
			Where("status = ? AND deleted_at IS NULL", 1).
			Count(&totalRecords).Error
		if err != nil {
			return nil, 0, err
		}

		err = db.Table("news").
			Where("status = ? AND deleted_at IS NULL", 1).
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

func CreateNews_Model(data map[string]interface{}) error {
	item := News{
		Title:      data["title"].(string),
		TitleEN:    data["title_en"].(string),
		SubTitle:   data["subtitle"].(string),
		SubTitleEN: data["subtitle_en"].(string),
		Image:      data["image"].(string),
		Status:     data["status"].(bool),
		Content:    data["content"].(string),
		ContentEN:  data["content_en"].(string),
		Postdate:   data["postdate"].(string),
		Pdfurl:     data["pdfurl"].(string),
	}

	if err := db.Create(&item).Error; err != nil {
		return err
	}
	return nil
}

func UpdateNews_Model(id string, data map[string]interface{}) error {
	item := News{
		Title:      data["title"].(string),
		TitleEN:    data["title_en"].(string),
		SubTitle:   data["subtitle"].(string),
		SubTitleEN: data["subtitle_en"].(string),
		Image:      data["image"].(string),
		Status:     data["status"].(bool),
		Content:    data["content"].(string),
		ContentEN:  data["content_en"].(string),
		Postdate:   data["postdate"].(string),
		Pdfurl:     data["pdfurl"].(string),
	}

	if err := db.Model(&item).Where("id = ?", id).Updates(map[string]interface{}{
		"title":       item.Title,
		"title_en":    item.TitleEN,
		"subtitle":    item.SubTitle,
		"subtitle_en": item.SubTitleEN,
		"image":       item.Image,
		"status":      item.Status,
		"content":     item.Content,
		"content_en":  item.ContentEN,
		"postdate":    item.Postdate,
		"pdfurl":      item.Pdfurl,
		"updated_at":  time.Now(),
	}).Error; err != nil {
		return err
	}
	return nil
}

func CheckDeletedAt_Model(id string) (bool, error) {
	var news News
	if err := db.Select("deleted_at").Where("id = ?", id).First(&news).Error; err != nil {
		return false, err
	}
	return news.DeletedAt.Valid, nil
}

func DeleteNews_Model(id string) error {
	isDeleted, err := CheckDeletedAt_Model(id)
	if err != nil {
		return err
	}
	if !isDeleted {
		if err := db.Delete(&News{}, id).Error; err != nil {
			return err
		}
	}
	return nil
}
