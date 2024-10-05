package hoatieu_service

import (
	"be-hoatieu/models"

	"gorm.io/gorm"
)

type Hoatieu struct {
	gorm.Model
	Code        string `gorm:"column:code;" json:"code"`
	Status      bool   `gorm:"column:status;" json:"status"`
	NavigatorId uint   `gorm:"column:userid" json:"userid"`
	Rank        string `gorm:"column:rank" json:"rank"`
	Image       string `gorm:"column:image" json:"image"`
	Lastname    string `gorm:"column:lastname" json:"lastname"`
	Firstname   string `gorm:"column:firstname" json:"firstname"`
}

type ObjectHoaTieu struct {
	ID          uint   `gorm:"column:ID" json:"ID"`
	Code        string `gorm:"column:code;" json:"code"`
	Status      bool   `gorm:"column:status;" json:"status"`
	NavigatorId uint   `gorm:"column:navigatorid" json:"navigatorid"`
	Rank        string `gorm:"column:rank" json:"rank"`
	Image       string `gorm:"column:image" json:"image"`
	Lastname    string `gorm:"column:lastname" json:"lastname"`
	Firstname   string `gorm:"column:firstname" json:"firstname"`
}

func (a *Hoatieu) GetAllNavigator_Service(limit int, page int) (*[]models.ObjectHoaTieu, int64, error) {
	item, totalRecords, err := models.GetAllNavigator_Model(limit, page)
	if err != nil {
		return nil, 0, err
	}
	return item, totalRecords, nil
}

// func (a *Hoatieu) CreateHoaTieu_Service() error {
// 	item := map[string]interface{}{
// 		"code":              a.Code,
// 		"status":            a.Status,
// 		"userid":            a.UserId,
// 		"hangid":            a.HangId,
// 		"bpctac":            a.BPCTac,
// 		"ngaysinh":          a.NgaySinh,
// 		"phone":             a.Phone,
// 		"phone2":            a.Phone2,
// 		"sex":               a.Sex,
// 		"tuoi":              a.Tuoi,
// 		"chungchihoatieuid": a.ChungChiHoaTieuID,
// 	}
// 	if err := models.CreateHoaTieu_Model(item); err != nil {
// 		return err
// 	}
// 	return nil
// }

// func (a *Hoatieu) UpdateHoaTieu_Service(id string) error {
// 	item := map[string]interface{}{
// 		"code":              a.Code,
// 		"status":            a.Status,
// 		"userid":            a.UserId,
// 		"hangid":            a.HangId,
// 		"bpctac":            a.BPCTac,
// 		"ngaysinh":          a.NgaySinh,
// 		"phone":             a.Phone,
// 		"phone2":            a.Phone2,
// 		"sex":               a.Sex,
// 		"tuoi":              a.Tuoi,
// 		"chungchihoatieuid": a.ChungChiHoaTieuID,
// 	}
// 	if err := models.UpdateHoaTieu_Model(id, item); err != nil {
// 		return err
// 	}
// 	return nil
// }

// func (a *Hoatieu) GetAllHoaTieuUserID_Service(userid string) (uint, error) {
// 	item, err := models.GetAllHoaTieuUserID_Model(userid)
// 	if err != nil {
// 		return 0, err
// 	}
// 	return item, nil
// }

// func (a *Hoatieu) GetByHoaTieuUserID_Service(userid string) (string, error) {
// 	item, err := models.GetByHoaTieuUserID_Model(userid)
// 	if err != nil {
// 		return "", err
// 	}
// 	return item, nil
// }
// func SearchHoaTieu_Service(name string) ([]models.ObjectHoaTieu, error) {
// 	return models.SearchHoaTieu_Model(name)
// }
// func SearchUpDateHoaTieu_Service(iddonhang, ngay, name string) ([]models.ObjectHoaTieu, error) {
// 	return models.SearchUpDateHoaTieu_Model(iddonhang, ngay, name)
// }
// func SearchHT2UpDateHoaTieu_Service(iddonhang, ngay, name string) ([]models.ObjectHoaTieu, error) {
// 	return models.SearchHT2UpDateHoaTieu_Model(iddonhang, ngay, name)
// }
// func SearchHTBangKeSanLuongUpDateHoaTieu_Service(ngay, name string) ([]models.ObjectHoaTieu, error) {
// 	return models.SearchHTBangKeSanLuongUpDateHoaTieu_Model(ngay, name)
// }
// func SearchHoaTieuRole_Service(name, role string) ([]models.ObjectHoaTieu, error) {
// 	return models.SearchHoaTieuRole_Model(name, role)
// }
// func (a *Hoatieu) GetMultipleHoaTieuTrue_Service(name string) (*[]models.ObjectHoaTieu, error) {
// 	item, err := models.GetMultipleHoaTieuTrue_Model(name)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return item, nil
// }
