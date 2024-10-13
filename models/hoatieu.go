package models

import (
	"time"

	"gorm.io/gorm"
)

type Hoatieu struct {
	gorm.Model
	Status bool   `gorm:"column:status;" json:"status"`
	Rank   string `gorm:"column:rank" json:"rank"`
	Image  string `gorm:"column:image" json:"image"`
	Name   string `gorm:"column:name" json:"name"`
}

type ObjectHoaTieu struct {
	ID     uint   `gorm:"column:id" json:"id"`
	Status bool   `gorm:"column:status;" json:"status"`
	Rank   string `gorm:"column:rank" json:"rank"`
	Image  string `gorm:"column:image" json:"image"`
	Name   string `gorm:"column:name" json:"name"`
}

func GetAllNavigator_Model(limit int, page int, showHidden bool) (*[]ObjectHoaTieu, int64, error) {
	var results []ObjectHoaTieu
	totalRecords := int64(0)
	var err error

	// Tính offset dựa trên limit và page
	offset := (page - 1) * limit

	if showHidden {

		// Truy vấn để đếm tổng số bản ghi
		err = db.Table("hoatieus").Where("deleted_at IS NULL").Count(&totalRecords).Error
		if err != nil {
			return nil, 0, err
		}

		// Truy vấn dữ liệu không có điều kiện status
		err = db.Table("hoatieus").
			Where("deleted_at IS NULL").
			Order("created_at DESC").
			Limit(limit).
			Offset(offset).
			Find(&results).Error
	} else {

		// Truy vấn để đếm tổng số bản ghi
		err = db.Table("hoatieus").Where("status = ? AND deleted_at IS NULL", 1).Count(&totalRecords).Error
		if err != nil {
			return nil, 0, err
		}
		// Truy vấn dữ liệu dựa trên limit và điều kiện status = 1
		err = db.Table("hoatieus").
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
func CreateNavigator_Model(data map[string]interface{}) error {
	// Create a News object using the provided data
	item := Hoatieu{
		Name:   data["name"].(string),
		Image:  data["image"].(string),
		Status: data["status"].(bool),
		Rank:   data["rank"].(string),
	}

	// Insert into the database
	result := db.Create(&item)

	if err := result.Error; err != nil {
		return err
	}
	return nil
}

func UpdateNavigator_Model(id string, data map[string]interface{}) error {
	item := Hoatieu{
		Name:   data["name"].(string),
		Image:  data["image"].(string),
		Status: data["status"].(bool),
		Rank:   data["rank"].(string),
	}

	if err := db.Model(&item).Where("id = ?", id).Updates(map[string]interface{}{
		"name":       item.Name,
		"image":      item.Image,
		"status":     item.Status,
		"rank":       item.Rank,
		"updated_at": time.Now(),
	}).Error; err != nil {
		return err
	}
	return nil
}

func DeleteNavigator_Model(id string) error {
	// Tìm kiếm bản ghi dựa trên ID
	var dichvu Hoatieu
	if err := db.Select("deleted_at").Where("id = ?", id).First(&dichvu).Error; err != nil {
		// Nếu không tìm thấy bản ghi, trả về lỗi
		return err
	}

	// Tiến hành xóa bản ghi
	if err := db.Delete(&Hoatieu{}, id).Error; err != nil {
		return err
	}

	return nil
}
