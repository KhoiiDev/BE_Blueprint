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
	ID      uint   `gorm:"column:ID" json:"ID"`
	Content string `gorm:"column:content" json:"content"`
	Status  bool   `gorm:"column:status" json:"status"`
}

func (a *Introduction) GetIntroduction_Service() (*[]models.ObjectIntroduction, error) {
	item, err := models.GetIntroduction_Model()
	if err != nil {
		return nil, err
	}
	return item, nil
}
