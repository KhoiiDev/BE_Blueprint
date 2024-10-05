package models

import (
	"gorm.io/gorm"
)

type ManeuveringDraft struct {
	gorm.Model
	Image  string `gorm:"column:image" json:"image"`
	Status bool   `gorm:"column:status" json:"status"`
}

type ObjectManeuveringDraft struct {
	ID     uint   `gorm:"column:id" json:"id"`
	Image  string `gorm:"column:image" json:"image"`
	Status bool   `gorm:"column:status" json:"status"`
}

func GetManeuveringDraft_Model() (*[]ObjectManeuveringDraft, error) {
	var results []ObjectManeuveringDraft

	// Truy vấn dữ liệu dựa trên limit và điều kiện status = 1
	err := db.Table("maneuvering_drafts").
		Where("status = ?", 1).
		Order("created_at DESC").
		Limit(1).
		Find(&results).Error

	if err != nil {
		return nil, err
	}

	return &results, nil
}
