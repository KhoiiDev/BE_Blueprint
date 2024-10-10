package models

import (
	"time"

	"gorm.io/gorm"
)

type TideCalendar struct {
	gorm.Model
	Pdfuri   string `gorm:"column:pdfuri" json:"pdfuri"`
	PostDate string `gorm:"column:postdate" json:"postdate"`
	Status   bool   `gorm:"column:status" json:"status"`
}

type ObjectTideCalendar struct {
	ID       uint   `gorm:"column:id" json:"id"`
	Pdfuri   string `gorm:"column:pdfuri" json:"pdfuri"`
	PostDate string `gorm:"column:postdate" json:"postdate"`
	Status   bool   `gorm:"column:status" json:"status"`
}

func GetTideCalendar_Model(limit int, page int, showHidden bool) (*[]ObjectTideCalendar, int64, error) {
	var results []ObjectTideCalendar
	totalRecords := int64(0)
	var err error

	// Tính offset dựa trên limit và page
	offset := (page - 1) * limit

	if showHidden {

		// Truy vấn để đếm tổng số bản ghi
		err = db.Table("tide_calendars").Where("deleted_at IS NULL").Count(&totalRecords).Error
		if err != nil {
			return nil, 0, err
		}

		// Truy vấn dữ liệu không có điều kiện status
		err = db.Table("tide_calendars").
			Where("deleted_at IS NULL").
			Order("created_at DESC").
			Limit(limit).
			Offset(offset).
			Find(&results).Error
	} else {

		// Truy vấn để đếm tổng số bản ghi
		err = db.Table("tide_calendars").Where("status = ? AND deleted_at IS NULL", 1).Count(&totalRecords).Error
		if err != nil {
			return nil, 0, err
		}
		// Truy vấn dữ liệu dựa trên limit và điều kiện status = 1
		err = db.Table("tide_calendars").
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
func CreateTideCalendar_Model(data map[string]interface{}) error {
	// Create a News object using the provided data
	item := TideCalendar{
		Pdfuri:   data["pdfuri"].(string),
		PostDate: data["postdate"].(string),
		Status:   data["status"].(bool),
	}

	// Insert into the database
	result := db.Create(&item)

	if err := result.Error; err != nil {
		return err
	}
	return nil
}

func UpdateTideCalendar_Model(id string, data map[string]interface{}) error {
	item := TideCalendar{
		Pdfuri:   data["pdfuri"].(string),
		PostDate: data["postdate"].(string),
		Status:   data["status"].(bool),
	}

	if err := db.Model(&item).Where("id = ?", id).Updates(map[string]interface{}{
		"pdfuri":     item.Pdfuri,
		"postdate":   item.PostDate,
		"status":     item.Status,
		"updated_at": time.Now(),
	}).Error; err != nil {
		return err
	}
	return nil
}

func DeleteTideCalendar_Model(id string) error {
	// Tìm kiếm bản ghi dựa trên ID
	var item TideCalendar
	if err := db.Select("deleted_at").Where("id = ?", id).First(&item).Error; err != nil {
		// Nếu không tìm thấy bản ghi, trả về lỗi
		return err
	}

	// Tiến hành xóa bản ghi
	if err := db.Delete(&TideCalendar{}, id).Error; err != nil {
		return err
	}

	return nil
}
