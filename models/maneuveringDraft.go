package models

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type ManeuveringDraft struct {
	gorm.Model
	PostDate string `gorm:"column:postdate" json:"postdate"`
	Pdfurl   string `gorm:"column:pdfurl" json:"pdfurl"`
	Status   bool   `gorm:"column:status" json:"status"`
}

type ObjectManeuveringDraft struct {
	ID       uint   `gorm:"column:id" json:"id"`
	PostDate string `gorm:"column:postdate" json:"postdate"`
	Pdfurl   string `gorm:"column:pdfurl" json:"pdfurl"`
	Status   bool   `gorm:"column:status" json:"status"`
}

func GetManeuveringDraft_Model(limit int, page int, showHidden bool, name string, date string) (*[]ObjectManeuveringDraft, int64, error) {
	var results []ObjectManeuveringDraft
	totalRecords := int64(0)
	var err error

	// Tính offset dựa trên limit và page
	offset := (page - 1) * limit

	if showHidden {

		// Truy vấn để đếm tổng số bản ghi
		err = db.Table("maneuvering_drafts").Where("deleted_at IS NULL").Count(&totalRecords).Error
		if err != nil {
			return nil, 0, err
		}

		// Truy vấn dữ liệu không có điều kiện status
		err = db.Table("maneuvering_drafts").
			Where("deleted_at IS NULL").
			Order("created_at DESC").
			Limit(limit).
			Offset(offset).
			Find(&results).Error
	} else {
		if date != "" {
			// Xử lý lỗi ngay khi parse date
			parsedDate, parseErr := time.Parse("02/01/2006", date)
			if parseErr != nil {
				return nil, 0, fmt.Errorf("Định dạng ngày không hợp lệ: %v", parseErr)
			}
			formattedDate := parsedDate.Format("2006-01-02")

			// Truy vấn dữ liệu dựa trên limit, status = 1 và PostDate
			err = db.Table("maneuvering_drafts").
				Where("status = ? AND deleted_at IS NULL AND postdate = ?", 1, formattedDate).
				Order("created_at DESC").
				Limit(1).
				Find(&results).Error
		} else {
			// Truy vấn dữ liệu dựa trên limit, status = 1 và PostDate
			err = db.Table("maneuvering_drafts").
				Where("status = ? AND deleted_at IS NULL", 1).
				Order("created_at DESC").
				Limit(1).
				Find(&results).Error
		}

	}

	if err != nil {
		return nil, 0, err
	}

	return &results, totalRecords, nil
}

// Corrected CreateNews_Model function
func CreateManeuveringDraft_Model(data map[string]interface{}) error {
	// Create a News object using the provided data
	item := ManeuveringDraft{
		Pdfurl:   data["pdfurl"].(string),
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

func UpdateManeuveringDraft_Model(id string, data map[string]interface{}) error {
	item := ManeuveringDraft{
		Pdfurl:   data["pdfurl"].(string),
		PostDate: data["postdate"].(string),
		Status:   data["status"].(bool),
	}

	if err := db.Model(&item).Where("id = ?", id).Updates(map[string]interface{}{
		"pdfurl":     item.Pdfurl,
		"postdate":   item.PostDate,
		"status":     item.Status,
		"updated_at": time.Now(),
	}).Error; err != nil {
		return err
	}
	return nil
}

func DeleteManeuveringDraft_Model(id string) error {
	// Tìm kiếm bản ghi dựa trên ID
	var item ManeuveringDraft
	if err := db.Select("deleted_at").Where("id = ?", id).First(&item).Error; err != nil {
		// Nếu không tìm thấy bản ghi, trả về lỗi
		return err
	}

	// Tiến hành xóa bản ghi
	if err := db.Delete(&ManeuveringDraft{}, id).Error; err != nil {
		return err
	}

	return nil
}
