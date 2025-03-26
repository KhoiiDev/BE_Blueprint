package models

import (
	"time"

	"gorm.io/gorm"
)

type Items struct {
	gorm.Model
	Title    string `gorm:"column:title" json:"title"`
	SubTitle string `gorm:"column:subtitle" json:"subtitle"`
	Image    string `gorm:"column:image" json:"image"`
	Status   bool   `gorm:"column:status" json:"status"`
	Content  string `gorm:"column:content" json:"content"`
	Postdate string `gorm:"column:postdate" json:"postdate"`
	Pdfurl   string `gorm:"column:pdfurl" json:"pdfurl"`
	Videourl string `gorm:"column:videourl" json:"videourl"`
	ItemType string `gorm:"column:itemtype" json:"itemtype"`
}

type ObjectItems struct {
	ID       uint   `gorm:"column:id" json:"id"`
	Title    string `gorm:"column:title" json:"title"`
	SubTitle string `gorm:"column:subtitle" json:"subtitle"`
	Image    string `gorm:"column:image" json:"image"`
	Status   bool   `gorm:"column:status" json:"status"`
	Content  string `gorm:"column:content" json:"content"`
	Pdfurl   string `gorm:"column:pdfurl" json:"pdfurl"`
	Videourl string `gorm:"column:videourl" json:"videourl"`

	Postdate string `gorm:"column:postdate" json:"postdate"`
	ItemType string `gorm:"column:itemtype" json:"itemtype"`
}

func GetItems_Model(limit int, page int, showHidden bool, name string, item_type string) (*[]ObjectItems, int64, error) {
	var results []ObjectItems
	var totalRecords int64
	var err error

	// Tính offset dựa trên limit và page
	offset := (page - 1) * limit

	// Tạo truy vấn cơ bản
	query := db.Table("items").Where("deleted_at IS NULL")

	// Thêm điều kiện cho item_type
	if item_type != "" {
		query = query.Where("itemtype = ?", item_type)
	}

	if name != "undefined" && name != "" {
		query = query.Where("title LIKE ?", "%"+name+"%")
	}

	if !showHidden {
		// Thêm điều kiện status nếu không hiển thị item ẩn
		query = query.Where("status = ?", 1)
	}

	// Tính tổng số bản ghi
	err = query.Count(&totalRecords).Error
	if err != nil {
		return nil, 0, err
	}

	// Lấy dữ liệu với limit và offset
	err = query.Order("created_at DESC").
		Limit(limit).
		Offset(offset).
		Find(&results).Error

	if err != nil {
		return nil, 0, err
	}

	return &results, totalRecords, nil
}

// Corrected Createitems_Model function
func CreateItems_Model(data map[string]interface{}) error {
	// Create a items object using the provided data
	item := Items{
		Title:    data["title"].(string),
		SubTitle: data["subtitle"].(string),
		Image:    data["image"].(string),
		Pdfurl:   data["pdfurl"].(string),
		Status:   data["status"].(bool),
		Content:  data["content"].(string),
		Videourl: data["videourl"].(string),
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
	item := Items{
		Title:    data["title"].(string),
		SubTitle: data["subtitle"].(string),
		Image:    data["image"].(string),
		Pdfurl:   data["pdfurl"].(string),
		Videourl: data["videourl"].(string),

		Status:   data["status"].(bool),
		Content:  data["content"].(string),
		Postdate: data["postdate"].(string),
	}

	if err := db.Model(&item).Where("id = ?", id).Updates(map[string]interface{}{
		"title":      item.Title,
		"subtitle":   item.SubTitle,
		"pdfurl":     item.Pdfurl,
		"image":      item.Image,
		"videourl":   item.Videourl,
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
	var item Items
	if err := db.Select("deleted_at").Where("id = ?", id).First(&item).Error; err != nil {
		// Nếu không tìm thấy bản ghi, trả về lỗi
		return err
	}

	// Tiến hành xóa bản ghi
	if err := db.Delete(&Items{}, id).Error; err != nil {
		return err
	}

	return nil
}
