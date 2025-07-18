package models

import (
	"time"

	"gorm.io/gorm"
)

type Header struct {
	gorm.Model
	CompanyName string `gorm:"column:company_name" json:"company_name"`
	Address     string `gorm:"column:address" json:"address"`
	Fax         string `gorm:"column:fax" json:"fax"`
	Email       string `gorm:"column:email" json:"email"`
	NumberPhone string `gorm:"column:number_phone" json:"number_phone"`
	BranchName  string `gorm:"column:branch_name" json:"branch_name"`
}

type ObjectHeader struct {
	ID          uint   `gorm:"column:id" json:"id"`
	CompanyName string `gorm:"column:company_name" json:"company_name"`
	Address     string `gorm:"column:address" json:"address"`
	Fax         string `gorm:"column:fax" json:"fax"`
	Email       string `gorm:"column:email" json:"email"`
	NumberPhone string `gorm:"column:number_phone" json:"number_phone"`
	BranchName  string `gorm:"column:branch_name" json:"branch_name"`
}

func GetHeader_Model(limit int, page int, showHidden bool, name string) (*[]ObjectHeader, int64, error) {
	var results []ObjectHeader
	var totalRecords int64
	var err error

	// Tính offset dựa trên limit và page
	offset := (page - 1) * limit

	// Tạo truy vấn cơ bản
	query := db.Table("headers").Where("deleted_at IS NULL")

	if name != "undefined" && name != "" {
		query = query.Where("company_name LIKE ?", "%"+name+"%")
	}

	// Tính tổng số bản ghi
	err = query.Count(&totalRecords).Error
	if err != nil {
		return nil, 0, err
	}

	// Lấy dữ liệu với limit và offset
	err = query.Order("created_at").
		Limit(limit).
		Offset(offset).
		Find(&results).Error

	if err != nil {
		return nil, 0, err
	}

	return &results, totalRecords, nil
}

func CreateHeader_Model(data map[string]interface{}) error {
	// Create a header object using the provided data
	header := Header{
		CompanyName: data["company_name"].(string),
		Address:     data["address"].(string),
		Fax:         data["fax"].(string),
		Email:       data["email"].(string),
		NumberPhone: data["number_phone"].(string),
		BranchName:  data["branch_name"].(string),
	}

	// Insert into the database
	result := db.Create(&header)

	if err := result.Error; err != nil {
		return err
	}
	return nil
}

func UpdateHeader_Model(id string, data map[string]interface{}) error {
	header := Header{
		CompanyName: data["company_name"].(string),
		Address:     data["address"].(string),
		Fax:         data["fax"].(string),
		Email:       data["email"].(string),
		NumberPhone: data["number_phone"].(string),
		BranchName:  data["branch_name"].(string),
	}

	if err := db.Model(&header).Where("id = ?", id).Updates(map[string]interface{}{
		"company_name": header.CompanyName,
		"address":      header.Address,
		"fax":          header.Fax,
		"email":        header.Email,
		"number_phone": header.NumberPhone,
		"branch_name":  header.BranchName,
		"updated_at":   time.Now(),
	}).Error; err != nil {
		return err
	}
	return nil
}

func DeleteHeader_Model(id string) error {
	var header Header
	if err := db.Select("deleted_at").Where("id = ?", id).First(&header).Error; err != nil {
		// Nếu không tìm thấy bản ghi, trả về lỗi
		return err
	}

	// Tiến hành xóa bản ghi
	if err := db.Delete(&Header{}, id).Error; err != nil {
		return err
	}

	return nil
}
