package models

import (
	"be-hoatieu/pkg/setting"
	"fmt"
	"log"
	"time"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

var db *gorm.DB

// gorm.Model definition
type BaseModel struct {
	Id        int            `gorm:"primaryKey;column:id" 		json:"id"`
	CreatedAt time.Time      `gorm:"column:created_at"	  		json:"createdAt"`
	UpdatedAt time.Time      `gorm:"column:updated_at" 			json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index;column:deleted_at" 	json:"deletedAt"`
}

func Setup() {
	var err error

	dns := fmt.Sprintf("%s://%s:%s@%s?database="+setting.DatabaseSetting.Name, setting.DatabaseSetting.Type, setting.DatabaseSetting.User, setting.DatabaseSetting.Password, setting.DatabaseSetting.Host)

	db, err = gorm.Open(sqlserver.Open(dns), &gorm.Config{
		PrepareStmt: true,
	})

	db.AutoMigrate(
		&Hoatieu{}, &Carousel{}, &Introduction{}, &Dichvu{}, &News{}, &ManeuveringDraft{}, &Servicelist{}, &Ship{},
		&TideCalendar{},
	)
	if err != nil {
		log.Fatalf("models.Setup err: %v", err)
	}
}
