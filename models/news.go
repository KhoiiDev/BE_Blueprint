package models

import (
	"time"

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
	ID       uint   `gorm:"column:id" json:"id"`
	Title    string `gorm:"column:title" json:"title"`
	SubTitle string `gorm:"column:subtitle" json:"subtitle"`
	Image    string `gorm:"column:image" json:"image"`
	Status   bool   `gorm:"column:status" json:"status"`
	Content  string `gorm:"column:content" json:"content"`
	Postdate string `gorm:"column:postdate" json:"postdate"`
}

func GetNews_Model(limit int, page int, showHidden bool) (*[]ObjectNews, int64, error) {
	var results []ObjectNews
	totalRecords := int64(0)
	var err error

	// Tính offset dựa trên limit và page
	offset := (page - 1) * limit

	if showHidden {
		// Truy vấn để đếm tổng số bản ghi
		err = db.Table("news").Where("deleted_at IS NULL").Count(&totalRecords).Error
		if err != nil {
			return nil, 0, err
		}

		// Truy vấn dữ liệu không có điều kiện status
		err = db.Table("news").
			Where("deleted_at IS NULL").
			Order("created_at DESC").
			Limit(limit).
			Offset(offset).
			Find(&results).Error
	} else {

		// Truy vấn để đếm tổng số bản ghi
		err = db.Table("news").Where("status = ? AND deleted_at IS NULL", 1).Count(&totalRecords).Error
		if err != nil {
			return nil, 0, err
		}
		// Truy vấn dữ liệu dựa trên limit và điều kiện status = 1
		err = db.Table("news").
			Where("status = ? AND deleted_at IS NULL", 1).
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

// Corrected CreateNews_Model function
func CreateNews_Model(data map[string]interface{}) error {
	// Create a News object using the provided data
	item := News{
		Title:    data["title"].(string),
		SubTitle: data["subtitle"].(string),
		Image:    data["image"].(string),
		Status:   data["status"].(bool),
		Content:  data["content"].(string),
		Postdate: data["postdate"].(string), // Ensure the type matches
	}

	// Insert into the database
	result := db.Create(&item)

	if err := result.Error; err != nil {
		return err
	}
	return nil
}

func UpdateNews_Model(id string, data map[string]interface{}) error {
	item := News{
		Title:    data["title"].(string),
		SubTitle: data["subtitle"].(string),
		Image:    data["image"].(string),
		Status:   data["status"].(bool),
		Content:  data["content"].(string),
		Postdate: data["postdate"].(string), // Ensure the type matches
	}

	if err := db.Model(&item).Where("id = ?", id).Updates(map[string]interface{}{
		"title":      item.Title,
		"subtitle":   item.SubTitle,
		"image":      item.Image,
		"status":     item.Status,
		"postdate":   item.Postdate,
		"content":    item.Content,
		"updated_at": time.Now(),
	}).Error; err != nil {
		return err
	}
	return nil
}

func CheckDeletedAt_Model(id string) (bool, error) {
	var new News
	if err := db.Select("deleted_at").Where("id = ?", id).First(&new).Error; err != nil {
		return false, err
	}
	return new.DeletedAt.Valid, nil
}
func DeleteNews_Model(id string) error {
	// Check if DeletedAt is not NULL
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
