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

// func Setup() {
// 	var err error

// 	dns := fmt.Sprintf("%s://%s:%s@%s?database="+setting.DatabaseSetting.Name, setting.DatabaseSetting.Type, setting.DatabaseSetting.User, setting.DatabaseSetting.Password, setting.DatabaseSetting.Host)

// 	db, err = gorm.Open(sqlserver.Open(dns), &gorm.Config{
// 		PrepareStmt: true,
// 	})

//		db.AutoMigrate(&Switch{}, &Kehoachdantau{}, &User{}, &UserCredentials{},
//			&Hoatieu{}, &Carousel{}, &Introduction{}, &Dichvu{}, &News{}, &ManeuveringDraft{}, &Servicelist{}, &Ship{},
//			&TideCalendar{}, &Items{},
//		)
//		// Khởi tạo dòng Switch mặc định
//		initializeDefaultSwitch()
//		if err != nil {
//			log.Fatalf("models.Setup err: %v", err)
//		}
//	}
//
// Setup function để khởi tạo database và dòng mặc định
func Setup() {
	var err error

	// Tạo DSN cho SQL Server
	dns := fmt.Sprintf("%s://%s:%s@%s?database=%s",
		setting.DatabaseSetting.Type,
		setting.DatabaseSetting.User,
		setting.DatabaseSetting.Password,
		setting.DatabaseSetting.Host,
		setting.DatabaseSetting.Name,
	)

	// Kết nối database
	db, err = gorm.Open(sqlserver.Open(dns), &gorm.Config{
		PrepareStmt: true,
	})
	if err != nil {
		log.Fatalf("models.Setup err: Failed to connect database: %v", err)
	}

	// Auto migrate các model
	err = db.AutoMigrate(
		&Header{},
		&Footer{},
		&Switch{},
		&Kehoachdantau{},
		&User{},
		&UserCredentials{},
		&Hoatieu{},
		&Carousel{},
		&Introduction{},
		&Dichvu{},
		&News{},
		&ManeuveringDraft{},
		&Servicelist{},
		&Ship{},
		&TideCalendar{},
		&Items{},
	)
	if err != nil {
		log.Fatalf("models.Setup err: Failed to migrate tables: %v", err)
	}

	// Khởi tạo dòng Switch mặc định
	initializeDefaultSwitch()
}
