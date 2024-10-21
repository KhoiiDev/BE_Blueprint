package ManeuveringDraft_service

import (
	"be-hoatieu/models"

	"gorm.io/gorm"
)

type ManeuveringDraft struct {
	gorm.Model
	PostDate string `gorm:"column:postdate" json:"postdate"`
	Dataurl  string `gorm:"column:dataurl" json:"dataurl"`
	Status   bool   `gorm:"column:status" json:"status"`
}

type ObjectManeuveringDraft struct {
	ID       uint   `gorm:"column:id" json:"id"`
	PostDate string `gorm:"column:postdate" json:"postdate"`
	Dataurl  string `gorm:"column:dataurl" json:"dataurl"`
	Status   bool   `gorm:"column:status" json:"status"`
}

func (a *ManeuveringDraft) GetManeuveringDraft_Service(limit int, page int, showHidden bool, date string) (*[]models.ObjectManeuveringDraft, int64, error) {
	item, totalRecords, err := models.GetManeuveringDraft_Model(limit, page, showHidden, date)
	if err != nil {
		return nil, 0, err
	}
	return item, totalRecords, nil
}

func (n *ManeuveringDraft) CreateManeuveringDraft_Service() error {
	item := map[string]interface{}{
		"dataurl":  n.Dataurl,
		"postdate": n.PostDate,
		"status":   n.Status,
	}
	if err := models.CreateManeuveringDraft_Model(item); err != nil { // Change function name to match your model's function
		return err
	}
	return nil
}

func (a *ManeuveringDraft) UpdateManeuveringDraft_Service(id string) error {
	item := map[string]interface{}{
		"dataurl":  a.Dataurl,
		"postdate": a.PostDate,
		"status":   a.Status,
	}
	if err := models.UpdateManeuveringDraft_Model(id, item); err != nil {
		return err
	}
	return nil
}

func (a *ManeuveringDraft) DeleteManeuveringDraft_Service(id string) (bool, error) {
	if err := models.DeleteManeuveringDraft_Model(id); err != nil {
		return false, err
	}
	return true, nil
}
