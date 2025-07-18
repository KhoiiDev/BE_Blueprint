package models

import (
	"time"

	"gorm.io/gorm"
)

type Ship struct {
	gorm.Model
	Name   string `gorm:"column:name" json:"name"`
	Image  string `gorm:"column:image" json:"image"`
	Status bool   `gorm:"column:status" json:"status"`
}

type ObjectShip struct {
	ID     uint   `gorm:"column:id" json:"id"`
	Name   string `gorm:"column:name" json:"name"`
	Image  string `gorm:"column:image" json:"image"`
	Status bool   `gorm:"column:status" json:"status"`
}

func GetShip_Model(limit int, page int, name string, showHidden bool) (*[]ObjectShip, int64, error) {
	var results []ObjectShip
	totalRecords := int64(0)
	var err error

	// Tính offset dựa trên limit và page
	offset := (page - 1) * limit

	if showHidden {

		// Truy vấn để đếm tổng số bản ghi
		err = db.Table("ships").Where("deleted_at IS NULL").Count(&totalRecords).Error
		if err != nil {
			return nil, 0, err
		}

		// Truy vấn dữ liệu không có điều kiện status
		err = db.Table("ships").
			Where("deleted_at IS NULL").
			Where("name LIKE ?", "%"+name+"%").
			Order("created_at").
			Limit(limit).
			Offset(offset).
			Find(&results).Error
	} else {

		// Truy vấn để đếm tổng số bản ghi
		err = db.Table("ships").Where("status = ? AND deleted_at IS NULL", 1).Count(&totalRecords).Error
		if err != nil {
			return nil, 0, err
		}
		// Truy vấn dữ liệu dựa trên limit và điều kiện status = 1
		err = db.Table("ships").
			Where("status = ? AND deleted_at IS NULL", 1).
			Order("created_at").
			Where("name LIKE ?", "%"+name+"%").
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
func CreateShip_Model(data map[string]interface{}) error {
	// Create a News object using the provided data
	item := Ship{
		Name:   data["name"].(string),
		Image:  data["image"].(string),
		Status: data["status"].(bool),
	}

	// Insert into the database
	result := db.Create(&item)

	if err := result.Error; err != nil {
		return err
	}
	return nil
}

func UpdateShip_Model(id string, data map[string]interface{}) error {
	item := Ship{
		Name:   data["name"].(string),
		Image:  data["image"].(string),
		Status: data["status"].(bool),
	}

	if err := db.Model(&item).Where("id = ?", id).Updates(map[string]interface{}{
		"name":       item.Name,
		"image":      item.Image,
		"status":     item.Status,
		"updated_at": time.Now(),
	}).Error; err != nil {
		return err
	}
	return nil
}

func DeleteShip_Model(id string) error {
	// Tìm kiếm bản ghi dựa trên ID
	var dichvu Ship
	if err := db.Select("deleted_at").Where("id = ?", id).First(&dichvu).Error; err != nil {
		// Nếu không tìm thấy bản ghi, trả về lỗi
		return err
	}

	// Tiến hành xóa bản ghi
	if err := db.Delete(&Ship{}, id).Error; err != nil {
		return err
	}

	return nil
}
