package header_service

import (
	"be-hoatieu/models"

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

func (h *Header) GetHeader_Service(limit int, page int, name string) (*[]models.ObjectHeader, int64, error) {
	header, totalRecords, err := models.GetHeader_Model(limit, page, false, name)
	if err != nil {
		return nil, totalRecords, err
	}
	return header, totalRecords, nil
}

func (h *Header) CreateHeader_Service() error {
	header := map[string]interface{}{
		"company_name": h.CompanyName,
		"address":      h.Address,
		"fax":          h.Fax,
		"email":        h.Email,
		"number_phone": h.NumberPhone,
		"branch_name":  h.BranchName,
	}
	if err := models.CreateHeader_Model(header); err != nil {
		return err
	}
	return nil
}

func (h *Header) UpdateHeader_Service(id string) error {
	header := map[string]interface{}{
		"company_name": h.CompanyName,
		"address":      h.Address,
		"fax":          h.Fax,
		"email":        h.Email,
		"number_phone": h.NumberPhone,
		"branch_name":  h.BranchName,
	}
	if err := models.UpdateHeader_Model(id, header); err != nil {
		return err
	}
	return nil
}

func (h *Header) DeleteHeader_Service(id string) (bool, error) {
	if err := models.DeleteHeader_Model(id); err != nil {
		return false, err
	}
	return true, nil
}
