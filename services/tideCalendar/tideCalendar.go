package TideCalendar_service

import (
	"be-hoatieu/models"

	"gorm.io/gorm"
)

type TideCalendar struct {
	gorm.Model
	Pdfuri   string `gorm:"column:pdfuri" json:"pdfuri"`
	PostDate string `gorm:"column:postdate" json:"postdate"`
	Status   bool   `gorm:"column:status" json:"status"`
}

type ObjectTideCalendar struct {
	ID       uint   `gorm:"column:id" json:"id"`
	Pdfuri   string `gorm:"column:pdfuri" json:"pdfuri"`
	PostDate string `gorm:"column:postdate" json:"postdate"`
	Status   bool   `gorm:"column:status" json:"status"`
}

func (a *TideCalendar) GetTideCalendar_Service(limit int, page int, showHidden bool) (*[]models.ObjectTideCalendar, int64, error) {
	item, totalRecords, err := models.GetTideCalendar_Model(limit, page, showHidden)
	if err != nil {
		return nil, 0, err
	}
	return item, totalRecords, nil
}

func (n *TideCalendar) CreateTideCalendar_Service() error {
	item := map[string]interface{}{
		"pdfuri":   n.Pdfuri,
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
		"pdfuri":   a.Pdfuri,
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
