package models

import (
	"time"

	"gorm.io/gorm"
)

type Footer struct {
	gorm.Model
	CompanyName string `gorm:"column:company_name" json:"company_name"`
	Address     string `gorm:"column:address" json:"address"`
	Fax         string `gorm:"column:fax" json:"fax"`
	Mst         string `gorm:"column:mst" json:"mst"`
	Email       string `gorm:"column:email" json:"email"`
	NumberPhone string `gorm:"column:number_phone" json:"number_phone"`
	BranchName  string `gorm:"column:branch_name" json:"branch_name"`
	Linkfb      string `gorm:"column:linkfb" json:"linkfb"`
}

type ObjectFooter struct {
	ID          uint   `gorm:"column:id" json:"id"`
	CompanyName string `gorm:"column:company_name" json:"company_name"`
	Address     string `gorm:"column:address" json:"address"`
	Fax         string `gorm:"column:fax" json:"fax"`
	Mst         string `gorm:"column:mst" json:"mst"`

	Email       string `gorm:"column:email" json:"email"`
	NumberPhone string `gorm:"column:number_phone" json:"number_phone"`
	BranchName  string `gorm:"column:branch_name" json:"branch_name"`
	Linkfb      string `gorm:"column:linkfb" json:"linkfb"`
}

func GetFooter_Model(limit int, page int, showHidden bool, name string) (*[]ObjectFooter, int64, error) {
	var results []ObjectFooter
	var totalRecords int64
	var err error

	// Tính offset dựa trên limit và page
	offset := (page - 1) * limit

	// Tạo truy vấn cơ bản
	query := db.Table("footers").Where("deleted_at IS NULL")

	if name != "undefined" && name != "" {
		query = query.Where("company_name LIKE ?", "%"+name+"%")
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

func CreateFooter_Model(data map[string]interface{}) error {
	// Create a footer object using the provided data
	footer := Footer{
		CompanyName: data["company_name"].(string),
		Address:     data["address"].(string),
		Fax:         data["fax"].(string),
		Mst:         data["mst"].(string),
		Email:       data["email"].(string),
		NumberPhone: data["number_phone"].(string),
		BranchName:  data["branch_name"].(string),
		Linkfb:      data["linkfb"].(string),
	}

	// Insert into the database
	result := db.Create(&footer)

	if err := result.Error; err != nil {
		return err
	}
	return nil
}

func UpdateFooter_Model(id string, data map[string]interface{}) error {
	footer := Footer{
		CompanyName: data["company_name"].(string),
		Address:     data["address"].(string),
		Fax:         data["fax"].(string),
		Mst:         data["mst"].(string),
		Email:       data["email"].(string),
		NumberPhone: data["number_phone"].(string),
		BranchName:  data["branch_name"].(string),
		Linkfb:      data["linkfb"].(string),
	}

	if err := db.Model(&footer).Where("id = ?", id).Updates(map[string]interface{}{
		"company_name": footer.CompanyName,
		"address":      footer.Address,
		"fax":          footer.Fax,
		"mst":          footer.Mst,
		"email":        footer.Email,
		"number_phone": footer.NumberPhone,
		"branch_name":  footer.BranchName,
		"linkfb":       footer.Linkfb,
		"updated_at":   time.Now(),
	}).Error; err != nil {
		return err
	}
	return nil
}

func DeleteFooter_Model(id string) error {
	var footer Footer
	if err := db.Select("deleted_at").Where("id = ?", id).First(&footer).Error; err != nil {
		// Nếu không tìm thấy bản ghi, trả về lỗi
		return err
	}

	// Tiến hành xóa bản ghi
	if err := db.Delete(&Footer{}, id).Error; err != nil {
		return err
	}

	return nil
}
