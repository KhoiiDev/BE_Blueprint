package models

import (
	"time"

	"gorm.io/gorm"
)

type Dichvu struct {
	gorm.Model
	Title   string `gorm:"column:title" json:"title"`
	Image   string `gorm:"column:image" json:"image"`
	Pdfdata string `gorm:"column:pdfdata" json:"pdfdata"`
	Status  bool   `gorm:"column:status" json:"status"`
}

type ObjectDichvu struct {
	ID      uint   `gorm:"column:id" json:"id"`
	Title   string `gorm:"column:title" json:"title"`
	Image   string `gorm:"column:image" json:"image"`
	Pdfdata string `gorm:"column:pdfdata" json:"pdfdata"`
	Status  bool   `gorm:"column:status" json:"status"`
}

func GetDichvu_Model(limit int, page int, showHidden bool) (*[]ObjectDichvu, int64, error) {
	var results []ObjectDichvu
	totalRecords := int64(0)
	var err error

	// Tính offset dựa trên limit và page
	offset := (page - 1) * limit

	// Truy vấn để đếm tổng số bản ghi
	err = db.Table("dichvus").Where("deleted_at IS NULL").Count(&totalRecords).Error
	if err != nil {
		return nil, 0, err
	}

	if showHidden {

		// Truy vấn dữ liệu không có điều kiện status
		err = db.Table("dichvus").
			Where("deleted_at IS NULL").
			Order("created_at DESC").
			Limit(limit).
			Offset(offset).
			Find(&results).Error
	} else {
		// Truy vấn dữ liệu dựa trên limit và điều kiện status = 1
		err = db.Table("dichvus").
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
func CreateDichvu_Model(data map[string]interface{}) error {
	// Create a News object using the provided data
	item := Dichvu{
		Title:   data["title"].(string),
		Image:   data["image"].(string),
		Status:  data["status"].(bool),
		Pdfdata: data["pdfdata"].(string),
	}

	// Insert into the database
	result := db.Create(&item)

	if err := result.Error; err != nil {
		return err
	}
	return nil
}

func UpdateDichvu_Model(id string, data map[string]interface{}) error {
	item := Dichvu{
		Title:   data["title"].(string),
		Image:   data["image"].(string),
		Status:  data["status"].(bool),
		Pdfdata: data["pdfdata"].(string),
	}

	if err := db.Model(&item).Where("id = ?", id).Updates(map[string]interface{}{
		"title":      item.Title,
		"image":      item.Image,
		"status":     item.Status,
		"pdfdata":    item.Pdfdata,
		"updated_at": time.Now(),
	}).Error; err != nil {
		return err
	}
	return nil
}

func DeleteDichvu_Model(id string) error {
	// Tìm kiếm bản ghi dựa trên ID
	var dichvu Dichvu
	if err := db.Select("deleted_at").Where("id = ?", id).First(&dichvu).Error; err != nil {
		// Nếu không tìm thấy bản ghi, trả về lỗi
		return err
	}

	// Tiến hành xóa bản ghi
	if err := db.Delete(&Dichvu{}, id).Error; err != nil {
		return err
	}

	return nil
}
