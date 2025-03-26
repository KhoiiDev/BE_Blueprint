package models

import (
	"time"

	"gorm.io/gorm"
)

type Servicelist struct {
	gorm.Model
	Title    string `gorm:"column:title" json:"title"`
	SubTitle string `gorm:"column:subtitle" json:"subtitle"`
	Pdfurl   string `gorm:"column:pdfurl" json:"pdfurl"`
	Postdate string `gorm:"column:postdate" json:"postdate"`

	Image   string `gorm:"column:image" json:"image"`
	Status  bool   `gorm:"column:status" json:"status"`
	Content string `gorm:"column:content" json:"content"`
}

type ObjectServicelist struct {
	ID       uint   `gorm:"column:id" json:"id"`
	Title    string `gorm:"column:title" json:"title"`
	SubTitle string `gorm:"column:subtitle" json:"subtitle"`
	Pdfurl   string `gorm:"column:pdfurl" json:"pdfurl"`
	Postdate string `gorm:"column:postdate" json:"postdate"`

	Image   string `gorm:"column:image" json:"image"`
	Status  bool   `gorm:"column:status" json:"status"`
	Content string `gorm:"column:content" json:"content"`
}

func GetServiceList_Model(limit int, page int, name string, showHidden bool) (*[]ObjectServicelist, int64, error) {
	var results []ObjectServicelist
	totalRecords := int64(0)
	var err error

	// Tính offset dựa trên limit và page
	offset := (page - 1) * limit

	if showHidden {
		// Truy vấn để đếm tổng số bản ghi
		err = db.Table("servicelists").Where("deleted_at IS NULL").Count(&totalRecords).Error
		if err != nil {
			return nil, 0, err
		}

		// Truy vấn dữ liệu không có điều kiện status
		err = db.Table("servicelists").
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
		err = db.Table("servicelists").
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
func CreateServiceList_Model(data map[string]interface{}) error {
	// Create a News object using the provided data
	item := Servicelist{
		Title:    data["title"].(string),
		SubTitle: data["subtitle"].(string),
		Postdate: data["postdate"].(string),
		Image:    data["image"].(string),
		Pdfurl:   data["pdfurl"].(string),
		Status:   data["status"].(bool),
		Content:  data["content"].(string),
	}

	// Insert into the database
	result := db.Create(&item)

	if err := result.Error; err != nil {
		return err
	}
	return nil
}

func UpdateServiceList_Model(id string, data map[string]interface{}) error {
	item := Servicelist{
		Title:    data["title"].(string),
		SubTitle: data["subtitle"].(string),
		Postdate: data["postdate"].(string),
		Image:    data["image"].(string),
		Pdfurl:   data["pdfurl"].(string),
		Status:   data["status"].(bool),
		Content:  data["content"].(string),
	}

	if err := db.Model(&item).Where("id = ?", id).Updates(map[string]interface{}{
		"title":      item.Title,
		"image":      item.Image,
		"subtitle":   item.SubTitle,
		"pdfurl":     item.Pdfurl,
		"postdate":   item.Postdate,
		"status":     item.Status,
		"content":    item.Content,
		"updated_at": time.Now(),
	}).Error; err != nil {
		return err
	}
	return nil
}

func DeleteServiceList_Model(id string) error {
	// Tìm kiếm bản ghi dựa trên ID
	var dichvu Servicelist
	if err := db.Select("deleted_at").Where("id = ?", id).First(&dichvu).Error; err != nil {
		// Nếu không tìm thấy bản ghi, trả về lỗi
		return err
	}

	// Tiến hành xóa bản ghi
	if err := db.Delete(&Servicelist{}, id).Error; err != nil {
		return err
	}

	return nil
}
