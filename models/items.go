package models

import (
	"time"

	"gorm.io/gorm"
)

type Item struct {
	gorm.Model
	Title    string `gorm:"column:title" json:"title"`
	Image    string `gorm:"column:image" json:"image"`
	Status   bool   `gorm:"column:status" json:"status"`
	Content  string `gorm:"column:content" json:"content"`
	Postdate string `gorm:"column:postdate" json:"postdate"`
	Pdfurl   string `gorm:"column:pdfurl" json:"pdfurl"`
	ItemType string `gorm:"column:itemtype" json:"itemtype"`
}

type ObjectItems struct {
	ID       uint   `gorm:"column:id" json:"id"`
	Title    string `gorm:"column:title" json:"title"`
	Image    string `gorm:"column:image" json:"image"`
	Status   bool   `gorm:"column:status" json:"status"`
	Content  string `gorm:"column:content" json:"content"`
	Pdfurl   string `gorm:"column:pdfurl" json:"pdfurl"`
	Postdate string `gorm:"column:postdate" json:"postdate"`
	ItemType string `gorm:"column:itemtype" json:"itemtype"`
}

func GetItems_Model(limit int, page int, showHidden bool, item_type string) (*[]ObjectItems, int64, error) {
	var results []ObjectItems
	totalRecords := int64(0)
	var err error

	// Tính offset dựa trên limit và page
	offset := (page - 1) * limit

	if showHidden {
		// Truy vấn để đếm tổng số bản ghi
		err = db.Table("items").Where("deleted_at IS NULL").Count(&totalRecords).Error
		if err != nil {
			return nil, 0, err
		}

		// Truy vấn dữ liệu không có điều kiện status
		err = db.Table("items").
			Where("itemtype = ? AND deleted_at IS NULL", item_type).
			Order("created_at DESC").
			Limit(limit).
			Offset(offset).
			Find(&results).Error
	} else {
		// Truy vấn dữ liệu dựa trên limit và điều kiện status = 1
		err = db.Table("items").
			Where("itemtype = ? AND deleted_at IS NULL", item_type).
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

// Corrected Createitems_Model function
func CreateItems_Model(data map[string]interface{}) error {
	// Create a items object using the provided data
	item := Item{
		Title:    data["title"].(string),
		Image:    data["image"].(string),
		Pdfurl:   data["pdfurl"].(string),
		Status:   data["status"].(bool),
		Content:  data["content"].(string),
		Postdate: data["postdate"].(string),
		ItemType: data["itemtype"].(string),
	}

	// Insert into the database
	result := db.Create(&item)

	if err := result.Error; err != nil {
		return err
	}
	return nil
}

func UpdateItems_Model(id string, data map[string]interface{}) error {
	item := Item{
		Title:    data["title"].(string),
		Image:    data["image"].(string),
		Pdfurl:   data["pdfurl"].(string),
		Status:   data["status"].(bool),
		Content:  data["content"].(string),
		Postdate: data["postdate"].(string),
	}

	if err := db.Model(&item).Where("id = ?", id).Updates(map[string]interface{}{
		"title":      item.Title,
		"image":      item.Image,
		"pdfurl":     item.Pdfurl,
		"status":     item.Status,
		"postdate":   item.Postdate,
		"content":    item.Content,
		"updated_at": time.Now(),
	}).Error; err != nil {
		return err
	}
	return nil
}

func DeleteItems_Model(id string) error {
	// Tìm kiếm bản ghi dựa trên ID
	var item Item
	if err := db.Select("deleted_at").Where("id = ?", id).First(&item).Error; err != nil {
		// Nếu không tìm thấy bản ghi, trả về lỗi
		return err
	}

	// Tiến hành xóa bản ghi
	if err := db.Delete(&Item{}, id).Error; err != nil {
		return err
	}

	return nil
}
