package footer_service

import (
	"be-hoatieu/models"

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

func (f *Footer) GetFooter_Service(limit int, page int, name string) (*[]models.ObjectFooter, int64, error) {
	footer, totalRecords, err := models.GetFooter_Model(limit, page, false, name)
	if err != nil {
		return nil, totalRecords, err
	}
	return footer, totalRecords, nil
}

func (f *Footer) CreateFooter_Service() error {
	footer := map[string]interface{}{
		"company_name": f.CompanyName,
		"address":      f.Address,
		"fax":          f.Fax,
		"mst":          f.Mst,
		"email":        f.Email,
		"number_phone": f.NumberPhone,
		"branch_name":  f.BranchName,
		"linkfb":       f.Linkfb,
	}
	if err := models.CreateFooter_Model(footer); err != nil {
		return err
	}
	return nil
}

func (f *Footer) UpdateFooter_Service(id string) error {
	footer := map[string]interface{}{
		"company_name": f.CompanyName,
		"address":      f.Address,
		"fax":          f.Fax,
		"mst":          f.Mst,
		"email":        f.Email,
		"number_phone": f.NumberPhone,
		"branch_name":  f.BranchName,
		"linkfb":       f.Linkfb,
	}
	if err := models.UpdateFooter_Model(id, footer); err != nil {
		return err
	}
	return nil
}

func (f *Footer) DeleteFooter_Service(id string) (bool, error) {
	if err := models.DeleteFooter_Model(id); err != nil {
		return false, err
	}
	return true, nil
}
