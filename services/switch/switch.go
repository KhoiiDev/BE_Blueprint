package switch_service

import (
	"be-hoatieu/models"

	"gorm.io/gorm"
)

type Switch struct {
	gorm.Model
	Flag bool `gorm:"column:flag;default:false" json:"flag"`
}

// GetSwitch_Service retrieves a switch by ID
func (s *Switch) GetSwitch_Service(id uint) (*models.Switch, error) {
	switchItem, err := models.GetSwitch_Model(id)
	if err != nil {
		return nil, err
	}
	return switchItem, nil
}

// CreateSwitch_Service creates a new switch
func (s *Switch) CreateSwitch_Service() error {
	if err := models.CreateSwitch_Model(s.Flag); err != nil {
		return err
	}
	return nil
}

// UpdateSwitch_Service updates an existing switch
func (s *Switch) UpdateSwitch_Service(id uint) error {
	if err := models.UpdateSwitch_Model(id, s.Flag); err != nil {
		return err
	}
	return nil
}

// DeleteSwitch_Service deletes a switch by ID
func (s *Switch) DeleteSwitch_Service(id uint) (bool, error) {
	if err := models.DeleteSwitch_Model(id); err != nil {
		return false, err
	}
	return true, nil
}
