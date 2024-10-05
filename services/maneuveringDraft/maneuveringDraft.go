package ManeuveringDraft_service

import (
	"be-hoatieu/models"

	"gorm.io/gorm"
)

type ManeuveringDraft struct {
	gorm.Model
	Image  string `gorm:"column:image" json:"image"`
	Status bool   `gorm:"column:status" json:"status"`
}

type ObjectManeuveringDraft struct {
	ID     uint   `gorm:"column:ID" json:"ID"`
	Image  string `gorm:"column:image" json:"image"`
	Status bool   `gorm:"column:status" json:"status"`
}

func (a *ManeuveringDraft) GetManeuveringDraft_Service() (*[]models.ObjectManeuveringDraft, error) {
	item, err := models.GetManeuveringDraft_Model()
	if err != nil {
		return nil, err
	}
	return item, nil
}
