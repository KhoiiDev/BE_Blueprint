package introduction_service

import (
	"be-hoatieu/models"

	"gorm.io/gorm"
)

type Introduction struct {
	gorm.Model
	Content string `gorm:"column:content" json:"content"`
	Status  bool   `gorm:"column:status" json:"status"`
}

type ObjectIntroduction struct {
	ID      uint   `gorm:"column:id" json:"id"`
	Content string `gorm:"column:content" json:"content"`
	Status  bool   `gorm:"column:status" json:"status"`
}

func (a *Introduction) GetIntroduction_Service(limit int, page int, showHidden bool) (*[]models.ObjectIntroduction, int64, error) {
	item, totalRecords, err := models.GetIntroduction_Model(limit, page, showHidden)
	if err != nil {
		return nil, 0, err
	}
	return item, totalRecords, nil
}

func (n *Introduction) CreateIntroduction_Service() error {
	item := map[string]interface{}{
		"content": n.Content,
		"status":  n.Status,
	}
	if err := models.CreateIntroduction_Model(item); err != nil { // Change function name to match your model's function
		return err
	}
	return nil
}

func (a *Introduction) UpdateIntroduction_Service(id string) error {
	item := map[string]interface{}{
		"content": a.Content,
		"status":  a.Status,
	}
	if err := models.UpdateIntroduction_Model(id, item); err != nil {
		return err
	}
	return nil
}

func (a *Introduction) DeleteIntroduction_Service(id string) (bool, error) {
	if err := models.DeleteIntroduction_Model(id); err != nil {
		return false, err
	}
	return true, nil
}
