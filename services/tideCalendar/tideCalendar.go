package TideCalendar_service

import (
	"be-hoatieu/models"

	"gorm.io/gorm"
)

type TideCalendar struct {
	gorm.Model
	Pdfurl   string `gorm:"column:pdfurl" json:"pdfurl"`
	PostDate string `gorm:"column:postdate" json:"postdate"`
	Status   bool   `gorm:"column:status" json:"status"`
}

type ObjectTideCalendar struct {
	ID       uint   `gorm:"column:id" json:"id"`
	Pdfurl   string `gorm:"column:pdfurl" json:"pdfurl"`
	PostDate string `gorm:"column:postdate" json:"postdate"`
	Status   bool   `gorm:"column:status" json:"status"`
}

func (a *TideCalendar) GetTideCalendar_Service(limit int, page int, showHidden bool, date string) (*[]models.ObjectTideCalendar, int64, error) {
	item, totalRecords, err := models.GetTideCalendar_Model(limit, page, showHidden, date)
	if err != nil {
		return nil, 0, err
	}
	return item, totalRecords, nil
}

func (n *TideCalendar) CreateTideCalendar_Service() error {
	item := map[string]interface{}{
		"pdfurl":   n.Pdfurl,
		"postdate": n.PostDate,
		"status":   n.Status,
	}
	if err := models.CreateTideCalendar_Model(item); err != nil { // Change function name to match your model's function
		return err
	}
	return nil
}

func (a *TideCalendar) UpdateTideCalendar_Service(id string) error {
	item := map[string]interface{}{
		"pdfurl":   a.Pdfurl,
		"postdate": a.PostDate,
		"status":   a.Status,
	}
	if err := models.UpdateTideCalendar_Model(id, item); err != nil {
		return err
	}
	return nil
}

func (a *TideCalendar) DeleteTideCalendar_Service(id string) (bool, error) {
	if err := models.DeleteTideCalendar_Model(id); err != nil {
		return false, err
	}
	return true, nil
}
