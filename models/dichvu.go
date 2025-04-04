package models

import (
	"time"

	"gorm.io/gorm"
)

type Dichvu struct {
	gorm.Model
	Title    string `gorm:"column:title" json:"title"`
	SubTitle string `gorm:"column:subtitle" json:"subtitle"`
	Content  string `gorm:"column:content" json:"content"`
	Image    string `gorm:"column:image" json:"image"`
	Pdfurl   string `gorm:"column:pdfurl" json:"pdfurl"`
	Postdate string `gorm:"column:postdate" json:"postdate"`
	Status   bool   `gorm:"column:status" json:"status"`
}

type ObjectDichvu struct {
	ID       uint   `gorm:"column:id" json:"id"`
	Title    string `gorm:"column:title" json:"title"`
	Image    string `gorm:"column:image" json:"image"`
	SubTitle string `gorm:"column:subtitle" json:"subtitle"`
	Content  string `gorm:"column:content" json:"content"`
	Postdate string `gorm:"column:postdate" json:"postdate"`

	Pdfurl string `gorm:"column:pdfurl" json:"pdfurl"`
	Status bool   `gorm:"column:status" json:"status"`
}

func GetDichvu_Model(limit int, page int, name string, showHidden bool) (*[]ObjectDichvu, int64, error) {
	var results []ObjectDichvu
	totalRecords := int64(0)
	var err error

	// Tính offset dựa trên limit và page
	offset := (page - 1) * limit

	// Truy vấn để đếm tổng số bản ghi

	if showHidden {
		err = db.Table("dichvus").Where("deleted_at IS NULL").Count(&totalRecords).Error
		if err != nil {
			return nil, 0, err
		}
		// Truy vấn dữ liệu không có điều kiện status
		err = db.Table("dichvus").
			Where("deleted_at IS NULL").
			Order("created_at DESC").
			Where("title LIKE ?", "%"+name+"%").
			Limit(limit).
			Offset(offset).
			Find(&results).Error
	} else {
		err = db.Table("servicelists").Where("status = ? AND deleted_at IS NULL", 1).Count(&totalRecords).Error
		if err != nil {
			return nil, 0, err
		}
		// Truy vấn dữ liệu dựa trên limit và điều kiện status = 1
		err = db.Table("dichvus").
			Where("status = ? AND deleted_at IS NULL", 1).
			Order("created_at DESC").
			Where("title LIKE ?", "%"+name+"%").
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
		Title:    data["title"].(string),
		SubTitle: data["subtitle"].(string),
		Content:  data["content"].(string),
		Image:    data["image"].(string),
		Postdate: data["postdate"].(string),
		Status:   data["status"].(bool),
		Pdfurl:   data["pdfurl"].(string),
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
		Title:    data["title"].(string),
		SubTitle: data["subtitle"].(string),
		Content:  data["content"].(string),
		Image:    data["image"].(string),
		Postdate: data["postdate"].(string),

		Status: data["status"].(bool),
		Pdfurl: data["pdfurl"].(string),
	}

	if err := db.Model(&item).Where("id = ?", id).Updates(map[string]interface{}{
		"title":      item.Title,
		"subtitle":   item.SubTitle,
		"content":    item.Content,
		"postdate":   item.Postdate,
		"image":      item.Image,
		"status":     item.Status,
		"pdfurl":     item.Pdfurl,
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
