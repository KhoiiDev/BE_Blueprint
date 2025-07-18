package models

import (
	"time"

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
	Postdate   string `gorm:"column:postdate" json:"postdate"`
	Pdfurl     string `gorm:"column:pdfurl" json:"pdfurl"`
	Videourl   string `gorm:"column:videourl" json:"videourl"`
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

func GetItems_Model(limit int, page int, showHidden bool, name string, item_type string) (*[]ObjectItems, int64, error) {
	var results []ObjectItems
	var totalRecords int64
	var err error

	offset := (page - 1) * limit

	query := db.Table("items").Where("deleted_at IS NULL")

	if item_type != "" {
		query = query.Where("itemtype = ?", item_type)
	}

	if name != "undefined" && name != "" {
		query = query.Where("title LIKE ?", "%"+name+"%")
	}

	if !showHidden {
		query = query.Where("status = ?", 1)
	}

	err = query.Count(&totalRecords).Error
	if err != nil {
		return nil, 0, err
	}

	err = query.Order("created_at").
		Limit(limit).
		Offset(offset).
		Find(&results).Error

	if err != nil {
		return nil, 0, err
	}

	return &results, totalRecords, nil
}

func CreateItems_Model(data map[string]interface{}) error {
	item := Items{
		Title:      data["title"].(string),
		TitleEN:    data["title_en"].(string),
		SubTitle:   data["subtitle"].(string),
		SubTitleEN: data["subtitle_en"].(string),
		Image:      data["image"].(string),
		Pdfurl:     data["pdfurl"].(string),
		Status:     data["status"].(bool),
		Content:    data["content"].(string),
		ContentEN:  data["content_en"].(string),
		Videourl:   data["videourl"].(string),
		Postdate:   data["postdate"].(string),
		ItemType:   data["itemtype"].(string),
	}

	result := db.Create(&item)
	if err := result.Error; err != nil {
		return err
	}
	return nil
}

func UpdateItems_Model(id string, data map[string]interface{}) error {
	item := Items{
		Title:      data["title"].(string),
		TitleEN:    data["title_en"].(string),
		SubTitle:   data["subtitle"].(string),
		SubTitleEN: data["subtitle_en"].(string),
		Image:      data["image"].(string),
		Pdfurl:     data["pdfurl"].(string),
		Videourl:   data["videourl"].(string),
		Status:     data["status"].(bool),
		Content:    data["content"].(string),
		ContentEN:  data["content_en"].(string),
		Postdate:   data["postdate"].(string),
	}

	if err := db.Model(&item).Where("id = ?", id).Updates(map[string]interface{}{
		"title":       item.Title,
		"title_en":    item.TitleEN,
		"subtitle":    item.SubTitle,
		"subtitle_en": item.SubTitleEN,
		"image":       item.Image,
		"pdfurl":      item.Pdfurl,
		"videourl":    item.Videourl,
		"status":      item.Status,
		"content":     item.Content,
		"content_en":  item.ContentEN,
		"postdate":    item.Postdate,
		"updated_at":  time.Now(),
	}).Error; err != nil {
		return err
	}
	return nil
}

func DeleteItems_Model(id string) error {
	var item Items
	if err := db.Select("deleted_at").Where("id = ?", id).First(&item).Error; err != nil {
		return err
	}

	if err := db.Delete(&Items{}, id).Error; err != nil {
		return err
	}

	return nil
}

func CheckDeletedAtItems_Model(id string) (bool, error) {
	var item Items
	if err := db.Select("deleted_at").Where("id = ?", id).First(&item).Error; err != nil {
		return false, err
	}
	return item.DeletedAt.Valid, nil
}

func DeleteItems_Model_WithCheck(id string) error {
	isDeleted, err := CheckDeletedAtItems_Model(id)
	if err != nil {
		return err
	}
	if !isDeleted {
		if err := db.Delete(&Items{}, id).Error; err != nil {
			return err
		}
	}
	return nil
}
