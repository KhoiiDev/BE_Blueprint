package models

import (
	"time"

	"gorm.io/gorm"
)

type Introduction struct {
	gorm.Model
	Content string `gorm:"column:content" json:"content"`
	Status  bool   `gorm:"column:status" json:"status"`
}

type ObjectIntroduction struct {
	ID      uint   `gorm:"column:id" json:"id"`
	Content string `gorm:"column:content" json:"content"`
	Status  bool   `gorm:"column:status" json:"status"`
}

func GetIntroduction_Model(limit int, page int, showHidden bool) (*[]ObjectIntroduction, int64, error) {
	var results []ObjectIntroduction
	totalRecords := int64(0)
	var err error

	// Tính offset dựa trên limit và page
	offset := (page - 1) * limit

	if showHidden {
		// Truy vấn để đếm tổng số bản ghi
		err = db.Table("introductions").Where("deleted_at IS NULL").Count(&totalRecords).Error
		if err != nil {
			return nil, 0, err
		}

		// Truy vấn dữ liệu không có điều kiện status
		err = db.Table("introductions").
			Where("deleted_at IS NULL").
			Order("created_at DESC").
			Limit(limit).
			Offset(offset).
			Find(&results).Error
	} else {
		// Truy vấn dữ liệu dựa trên limit và điều kiện status = 1
		err = db.Table("introductions").
			Where("status = ? AND deleted_at IS NULL", 1).
			Order("created_at DESC").
			Limit(1).
			Find(&results).Error
	}

	if err != nil {
		return nil, 0, err
	}

	return &results, totalRecords, nil
}

// Corrected CreateNews_Model function
func CreateIntroduction_Model(data map[string]interface{}) error {
	// Create a News object using the provided data
	item := Introduction{
		Content: data["content"].(string),
		Status:  data["status"].(bool),
	}

	// Insert into the database
	result := db.Create(&item)

	if err := result.Error; err != nil {
		return err
	}
	return nil
}

func UpdateIntroduction_Model(id string, data map[string]interface{}) error {
	item := Introduction{

		Content: data["content"].(string),
		Status:  data["status"].(bool),
	}

	if err := db.Model(&item).Where("id = ?", id).Updates(map[string]interface{}{
		"content":    item.Content,
		"status":     item.Status,
		"updated_at": time.Now(),
	}).Error; err != nil {
		return err
	}
	return nil
}

func DeleteIntroduction_Model(id string) error {
	// Tìm kiếm bản ghi dựa trên ID
	var intro Introduction
	if err := db.Select("deleted_at").Where("id = ?", id).First(&intro).Error; err != nil {
		// Nếu không tìm thấy bản ghi, trả về lỗi
		return err
	}

	// Tiến hành xóa bản ghi
	if err := db.Delete(&Introduction{}, id).Error; err != nil {
		return err
	}

	return nil
}
