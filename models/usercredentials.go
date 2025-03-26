package models

import "gorm.io/gorm"

type UserCredentials struct {
	BaseModel
	Password string `gorm:"column:password" json:"password"`
	UserId   uint   `gorm:"column:User_id" json:"UserId"`
}

func (u *UserCredentials) FindCredentials(UserId string) (*UserCredentials, error) {
	var userCredentials UserCredentials
	err := db.Where("User_id = ?", UserId).First(&userCredentials).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &userCredentials, err
}
