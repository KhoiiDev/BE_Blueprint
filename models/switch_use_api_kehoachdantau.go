package models

import (
	"log"
	"time"

	"gorm.io/gorm"
)

// Switch model
type Switch struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Flag      bool           `gorm:"column:flag;default:false" json:"flag"`
	CreatedAt time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
}

// GetSwitch_Model function
func GetSwitch_Model(id uint) (*Switch, error) {
	var switchItem Switch
	err := db.Where("id = ?", id).
		First(&switchItem).Error

	if err != nil {
		return nil, err
	}
	return &switchItem, nil
}

// CreateSwitch_Model function
func CreateSwitch_Model(flag bool) error {
	switchItem := Switch{
		Flag: flag,
	}

	result := db.Create(&switchItem)
	if err := result.Error; err != nil {
		return err
	}
	return nil
}

// UpdateSwitch_Model function
func UpdateSwitch_Model(id uint, flag bool) error {
	switchItem := Switch{
		Flag: flag,
	}

	if err := db.Model(&switchItem).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"flag":       flag,
			"updated_at": time.Now(),
		}).Error; err != nil {
		return err
	}
	return nil
}

// DeleteSwitch_Model function
func DeleteSwitch_Model(id uint) error {
	var switchItem Switch
	if err := db.Select("deleted_at").
		Where("id = ?", id).
		First(&switchItem).Error; err != nil {
		return err
	}

	if err := db.Delete(&Switch{}, id).Error; err != nil {
		return err
	}
	return nil
}

// Hàm khởi tạo dòng Switch mặc định
func initializeDefaultSwitch() {
	var count int64
	if err := db.Model(&Switch{}).Count(&count).Error; err != nil {
		log.Println("Error checking Switch table:", err)
		return
	}

	if count == 0 {
		defaultSwitch := Switch{
			Flag: false, // Giá trị mặc định
		}
		if err := db.Create(&defaultSwitch).Error; err != nil {
			log.Println("Error creating default Switch:", err)
			return
		}
		log.Println("Created default Switch with ID:", defaultSwitch.ID)
	} else {
		log.Println("Switch table already has data, skipping default creation")
	}
}
