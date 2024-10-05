package models

import (
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
	NavigatorId uint   `gorm:"column:userid" json:"userid"`
	Rank        string `gorm:"column:rank" json:"rank"`
	Image       string `gorm:"column:image" json:"image"`
	Lastname    string `gorm:"column:lastname" json:"lastname"`
	Firstname   string `gorm:"column:firstname" json:"firstname"`
}

func GetAllNavigator_Model(limit int, page int) (*[]ObjectHoaTieu, int64, error) {
	var results []ObjectHoaTieu
	var totalRecords int64

	// Tính offset dựa trên limit và page
	offset := (page - 1) * limit

	// Truy vấn để đếm tổng số bản ghi có status = 1
	err := db.Table("hoatieus").Where("status = ?", 1).Count(&totalRecords).Error
	if err != nil {
		return nil, 0, err
	}

	// Truy vấn dữ liệu dựa trên limit, offset và điều kiện status = 1
	err = db.Table("hoatieus").
		Where("status = ?", 1).
		Order("created_at DESC").
		Offset(offset).
		Limit(limit).
		Find(&results).Error

	if err != nil {
		return nil, 0, err
	}

	return &results, totalRecords, nil
}

// func CreateHoaTieu_Model(data map[string]interface{}) error {
// 	_, err := GetCodeCheck[Hoatieu](data["code"].(string))
// 	if err != nil {
// 		return err
// 	}
// 	item := Hoatieu{
// 		Code:              data["code"].(string),
// 		Status:            data["status"].(bool),
// 		UserId:            data["userid"].(uint),
// 		HangId:            data["hangid"].(uint),
// 		BPCTac:            data["bpctac"].(uint),
// 		NgaySinh:          data["ngaysinh"].(string),
// 		Phone:             data["phone"].(string),
// 		Phone2:            data["phone2"].(string),
// 		Sex:               data["sex"].(string),
// 		Tuoi:              data["tuoi"].(string),
// 		ChungChiHoaTieuID: data["chungchihoatieuid"].(uint),
// 	}

// 	result := db.Create(&item)

// 	if err := result.Error; err != nil {
// 		return err
// 	}
// 	return nil
// }

// func UpdateHoaTieu_Model(id string, data map[string]interface{}) error {
// 	_, err := CheckDuplicateCode[Hoatieu](id, data["code"].(string))
// 	if err != nil {
// 		return err
// 	}
// 	item := Hoatieu{
// 		Code:              data["code"].(string),
// 		Status:            data["status"].(bool),
// 		HangId:            data["hangid"].(uint),
// 		BPCTac:            data["bpctac"].(uint),
// 		NgaySinh:          data["ngaysinh"].(string),
// 		Phone:             data["phone"].(string),
// 		Phone2:            data["phone2"].(string),
// 		Sex:               data["sex"].(string),
// 		Tuoi:              data["tuoi"].(string),
// 		ChungChiHoaTieuID: data["chungchihoatieuid"].(uint),
// 	}
// 	if err := db.Model(&item).Where("id = ?", id).Updates(map[string]interface{}{
// 		"code":              item.Code,
// 		"status":            item.Status,
// 		"hangid":            item.HangId,
// 		"bpctac":            item.BPCTac,
// 		"ngaysinh":          item.NgaySinh,
// 		"phone":             item.Phone,
// 		"phone2":            item.Phone2,
// 		"sex":               item.Sex,
// 		"tuoi":              item.Tuoi,
// 		"chungchihoatieuid": item.ChungChiHoaTieuID,
// 	}).Error; err != nil {
// 		return err
// 	}
// 	return nil
// }

// get all
// func GetAllHoaTieu_Model() (*[]Hoatieu, error) {
// 	item := []Hoatieu{}
// 	err := db.Debug().Find(&item).Error

// 	if err != nil && err != gorm.ErrRecordNotFound {
// 		return nil, err
// 	}

//		return &item, nil
//	}

// func GetAllHoaTieuUserID_Model(userid string) (uint, error) {
// 	var results uint

// 	// Execute the JOIN query
// 	err := db.Debug().
// 		Table("hoatieus").
// 		Select(`hoatieus.ID`).
// 		Where("hoatieus.userid = ? ", userid).
// 		Scan(&results).Error

// 	if err != nil {
// 		return 0, err
// 	}
// 	return results, nil
// }
// func GetByHoaTieuUserID_Model(userid string) (string, error) {
// 	var results string

// 	// Execute the JOIN query
// 	err := db.Debug().
// 		Table("hoatieus").
// 		Select(`hangs.name as code`).
// 		Joins("LEFT JOIN hangs on hoatieus.hangid = hangs.id").
// 		Where("hoatieus.userid = ? ", userid).
// 		Scan(&results).Error

// 	if err != nil {
// 		return "", err
// 	}
// 	return results, nil
// }

// func SearchHoaTieu_Model(name string) ([]ObjectHoaTieu, error) {
// 	var HoaTieu []ObjectHoaTieu

// 	query := db.Debug().
// 		Table("hoatieus").
// 		Select(`hoatieus.ID,
// 		        hoatieus.status,
// 		        hoatieus.userid,
// 						hoatieus.code,
// 		        hoatieus.ngaysinh,
// 		        hoatieus.hangid,
// 				CONCAT(users.first_name,' ',users.last_name ) AS hoatieuname,
// 		        users.first_name as firstname, users.last_name as lastname ,
// 		        users.username AS username,
// 				users.[image] as [image],
// 		        hangs.name AS hangname,
// 		        hoatieus.bpctac,
// 				chuc_vus.name as bpctacname,
// 		        chung_chi_hoa_tieus.id as chungchihoatieuid,
// 		        chung_chi_hoa_tieus.gcnkncm as gcnkncm,
// 				chung_chi_hoa_tieus.gcnvhdht as gcnvhdht,
// 		        chung_chi_hoa_tieus.ngaycap as ngaycap,
// 		        chung_chi_hoa_tieus.ngayhethan as ngayhethan,
// 		        hoatieus.ngaysinh,
// 		        hoatieus.sex,
// 		        hoatieus.tuoi,
// 		        hoatieus.phone,
// 				hoatieus.phone2,
// 				hoatieus.ID as hoatieuid`).
// 		Joins("LEFT JOIN users ON users.id = hoatieus.userid").
// 		Joins("LEFT JOIN hangs ON hangs.id = hoatieus.hangid").
// 		Joins("LEFT JOIN chuc_vus ON chuc_vus.id = hoatieus.bpctac").
// 		Joins("LEFT JOIN chung_chi_hoa_tieus ON hoatieus.chungchihoatieuid = chung_chi_hoa_tieus.id").
// 		Where("(users.first_name + ' ' + users.last_name) LIKE ?", "%"+name+"%").
// 		Order("hoatieus.created_at desc").
// 		Limit(10).
// 		Scan(&HoaTieu)

// 	if query.Error != nil {
// 		return nil, query.Error
// 	}

// 	return HoaTieu, nil
// }
// func SearchUpDateHoaTieu_Model(iddonhang, ngay, name string) ([]ObjectHoaTieu, error) {
// 	var hangid int

// 	hoatieucunghang := db.Debug().
// 		Table("don_hangs").
// 		Select("taus.hangid as hang").
// 		Where("don_hangs.ID = ?", iddonhang).
// 		Joins("LEFT JOIN taus ON taus.id = don_hangs.tauid").
// 		Scan(&hangid)
// 	if hoatieucunghang.Error != nil {
// 		return nil, fmt.Errorf("Lỗi khi lấy hoa tieu hoạt động: %v", hoatieucunghang.Error)
// 	}

// 	var hoatieudan []int
// 	hoatieuhangdan := db.Debug().
// 		Table("phan_cong_hang_dans").
// 		Select("phan_cong_hang_dan_details.hangphancong").
// 		Where("phan_cong_hang_dans.idhang = ? AND phan_cong_hang_dan_details.deleted_at IS NULL", hangid).
// 		Joins("LEFT JOIN phan_cong_hang_dan_details ON phan_cong_hang_dan_details.idhang = phan_cong_hang_dans.id").
// 		Scan(&hoatieudan)
// 	if hoatieuhangdan.Error != nil {
// 		return nil, fmt.Errorf("Lỗi khi lấy hoa tieu dẫn đầu: %v", hoatieuhangdan.Error)
// 	}
// 	var hoatieuCongTac []int
// 	err := db.Debug().
// 		Table("phieu_cong_tacs").
// 		Select("phieu_cong_tacs.userid").
// 		Where("CONVERT (varchar(10), ?, 103) BETWEEN CONVERT (varchar(10), phieu_cong_tacs.tungay, 103) AND CONVERT (varchar(10), phieu_cong_tacs.denngay, 103)", ngay).
// 		Where("phieu_cong_tacs.approvalstatus = ?", true).
// 		Scan(&hoatieuCongTac).Error
// 	if err != nil {
// 		return nil, fmt.Errorf("Lỗi khi lấy phiếu công tác: %v", err)
// 	}

// 	// Kiểm tra nếu hoatieuCongTac không có phần tử nào, gán mặc định là 0
// 	if len(hoatieuCongTac) == 0 {
// 		hoatieuCongTac = []int{0}
// 	}

// 	// Tiếp tục xử lý hoatieuCongTac sau khi đã đảm bảo không nil và có giá trị mặc định là 0

// 	var hoaTieuHoatDong []struct {
// 		Hoatieuid int `gorm:"column:hoatieuid"`
// 	}
// 	// Lấy danh sách hoa tiêu hoạt động theo ngày và hạng
// 	hoatieuHD := db.Debug().
// 		Table("hoat_dong_tours").
// 		Select("tua_hoa_tieu_details.hoatieuid as hoatieuid").
// 		Where("CONVERT (varchar(10), ?, 103) BETWEEN CONVERT (varchar(10), hoat_dong_tours.tungay, 103) AND CONVERT (varchar(10), hoat_dong_tours.denngay, 103)", ngay).
// 		Joins("LEFT JOIN tua_hoa_tieus ON tua_hoa_tieus.id = hoat_dong_tours.tourid").
// 		Joins("LEFT JOIN tua_hoa_tieu_details ON tua_hoa_tieu_details.idtuor = tua_hoa_tieus.id").
// 		Joins("LEFT JOIN hoatieus ON hoatieus.id = tua_hoa_tieu_details.hoatieuid").
// 		Where("hoatieus.hangid IN ?", hoatieudan).
// 		Where("tua_hoa_tieu_details.hoatieuid NOT IN ?", hoatieuCongTac).
// 		Where("tua_hoa_tieus.status = ?", true).
// 		Scan(&hoaTieuHoatDong)
// 	if hoatieuHD.Error != nil {
// 		return nil, fmt.Errorf("Lỗi khi lấy hoa tieu hoạt động: %v", hoatieuHD.Error)
// 	}

// 	// Sử dụng map để loại bỏ các hoa tiêu trùng lặp
// 	hoaTieuHoatDongMap := make(map[int]struct{})
// 	hoaTieuHoatDongMapStr := make(map[string]struct{})
// 	for _, record := range hoaTieuHoatDong {
// 		hoaTieuHoatDongMap[record.Hoatieuid] = struct{}{}
// 	}

// 	var hoatieuScores []struct {
// 		HoaTieuID1   int    `gorm:"column:hoa_tieu_id1"`
// 		HoaTieuID2   string `gorm:"column:hoa_tieu_id2"`
// 		HoaTieuIDSid string `gorm:"column:hoa_tieu_id_sid"`
// 	}
// 	// Lấy danh sách hoa tiêu từ đơn hàng
// 	hoatieuUIDs := db.Debug().
// 		Table("don_hangs").
// 		Select("don_hangs.hoatieu1 as hoa_tieu_id1, don_hangs.hoatieu2 as hoa_tieu_id2, don_hangs.hoatieutsid as hoa_tieu_id_sid").
// 		Where("CONVERT(varchar(10), don_hangs.ngaydichchuyen, 103) = ? OR CONVERT(varchar(10), don_hangs.ngayvao, 103) = ? OR CONVERT(varchar(10), don_hangs.ngayra, 103) = ?", ngay, ngay, ngay).
// 		Where("don_hangs.status = ?", false).
// 		Scan(&hoatieuScores)
// 	if hoatieuUIDs.Error != nil {
// 		return nil, fmt.Errorf("Lỗi khi lấy hoa tieu từ đơn hàng: %v", hoatieuUIDs.Error)
// 	}

// 	// Xóa hoa tiêu có trong đơn hàng khỏi danh sách hoa tiêu hoạt động
// 	for _, score := range hoatieuScores {
// 		if score.HoaTieuID1 != 0 {
// 			delete(hoaTieuHoatDongMap, score.HoaTieuID1)
// 		}
// 		if score.HoaTieuID2 != "" {
// 			delete(hoaTieuHoatDongMapStr, score.HoaTieuID2)
// 		}
// 		if score.HoaTieuIDSid != "" {
// 			delete(hoaTieuHoatDongMapStr, score.HoaTieuIDSid)
// 		}
// 	}

// 	// Chuyển đổi map thành slice để sử dụng trong truy vấn cuối cùng
// 	var filteredHoaTieuIDs []int
// 	for hoaTieuID := range hoaTieuHoatDongMap {
// 		filteredHoaTieuIDs = append(filteredHoaTieuIDs, hoaTieuID)
// 	}

// 	var HoaTieu []ObjectHoaTieu
// 	query := db.Debug().
// 		Table("hoatieus").
// 		Select(`hoatieus.ID,
//                 hoatieus.status,
//                 hoatieus.userid,
//                 users.birthday AS ngaysinh,
//                 hoatieus.hangid,
//                 CONCAT(users.first_name,' ',users.last_name ) AS hoatieuname,
//                 users.first_name as firstname, users.last_name as lastname,
//                 users.username AS username,
//                 users.[image] as [image],
//                 hangs.name AS hangname,
//                 hoatieus.bpctac,
//                 chuc_vus.name as bpctacname,
//                 chung_chi_hoa_tieus.id as chungchihoatieuid,
//                 chung_chi_hoa_tieus.gcnkncm as gcnkncm,
//                 chung_chi_hoa_tieus.gcnvhdht as gcnvhdht,
//                 chung_chi_hoa_tieus.ngaycap as ngaycap,
//                 chung_chi_hoa_tieus.ngayhethan as ngayhethan,
//                 hoatieus.ngaysinh,
//                 hoatieus.sex,
//                 hoatieus.tuoi,
//                 hoatieus.phone,
//                 hoatieus.phone2,
//                 hoatieus.ID as hoatieuid,
// 								hoatieus.code`).
// 		Joins("LEFT JOIN users ON users.id = hoatieus.userid").
// 		Joins("LEFT JOIN hangs ON hangs.id = hoatieus.hangid").
// 		Joins("LEFT JOIN chuc_vus ON chuc_vus.id = hoatieus.bpctac").
// 		Joins("LEFT JOIN chung_chi_hoa_tieus ON hoatieus.chungchihoatieuid = chung_chi_hoa_tieus.id").
// 		Where("(users.first_name + ' ' + users.last_name) LIKE ?", "%"+name+"%").
// 		Where("hoatieus.ID IN (?)", filteredHoaTieuIDs).
// 		Order("hoatieus.created_at desc").
// 		Limit(10).
// 		Scan(&HoaTieu)

// 	fmt.Println("Hoa Tieu Hoat Dong:", hoaTieuHoatDong)
// 	fmt.Println("Hoa Tieu Scores:", hoatieuScores)
// 	fmt.Println("Filtered Hoa Tieu IDs:", filteredHoaTieuIDs)

// 	if query.Error != nil {
// 		return nil, query.Error
// 	}

// 	return HoaTieu, nil
// }

// func SearchHT2UpDateHoaTieu_Model(iddonhang, ngay, name string) ([]ObjectHoaTieu, error) {
// 	// var hangid int

// 	// hoatieucunghang := db.Debug().
// 	// 	Table("don_hangs").
// 	// 	Select("taus.hangid as hang").
// 	// 	Where("don_hangs.ID = ?", iddonhang).
// 	// 	Joins("LEFT JOIN taus ON taus.id = don_hangs.tauid").
// 	// 	Scan(&hangid)
// 	// if hoatieucunghang.Error != nil {
// 	// 	return nil, fmt.Errorf("Lỗi khi lấy hoa tieu hoạt động: %v", hoatieucunghang.Error)
// 	// }

// 	// var hoatieudan []int
// 	// hoatieuhangdan := db.Debug().
// 	// 	Table("phan_cong_hang_dans").
// 	// 	Select("phan_cong_hang_dan_details.hangphancong").
// 	// 	Where("phan_cong_hang_dans.idhang = ? AND phan_cong_hang_dan_details.deleted_at IS NULL", hangid).
// 	// 	Joins("LEFT JOIN phan_cong_hang_dan_details ON phan_cong_hang_dan_details.idhang = phan_cong_hang_dans.id").
// 	// 	Scan(&hoatieudan)
// 	// if hoatieuhangdan.Error != nil {
// 	// 	return nil, fmt.Errorf("Lỗi khi lấy hoa tieu dẫn đầu: %v", hoatieuhangdan.Error)
// 	// }
// 	var hoatieuCongTac []int
// 	err := db.Debug().
// 		Table("phieu_cong_tacs").
// 		Select("phieu_cong_tacs.userid").
// 		// Where("CONVERT (varchar(10), ?, 103) BETWEEN CONVERT (varchar(10), phieu_cong_tacs.tungay, 103) AND CONVERT (varchar(10), phieu_cong_tacs.denngay, 103)", ngay).
// 		Where("CONVERT (date, ?, 103) BETWEEN CONVERT (date, phieu_cong_tacs.tungay, 103) AND CONVERT (date, phieu_cong_tacs.denngay, 103)", ngay).
// 		Where("phieu_cong_tacs.approvalstatus = ?", true).
// 		Scan(&hoatieuCongTac).Error
// 	if err != nil {
// 		return nil, fmt.Errorf("Lỗi khi lấy phiếu công tác: %v", err)
// 	}

// 	// Kiểm tra nếu hoatieuCongTac không có phần tử nào, gán mặc định là 0
// 	if len(hoatieuCongTac) == 0 {
// 		hoatieuCongTac = []int{0}
// 	}

// 	// Tiếp tục xử lý hoatieuCongTac sau khi đã đảm bảo không nil và có giá trị mặc định là 0

// 	var hoaTieuHoatDong []struct {
// 		Hoatieuid int `gorm:"column:hoatieuid"`
// 	}
// 	// Lấy danh sách hoa tiêu hoạt động theo ngày và hạng
// 	hoatieuHD := db.Debug().
// 		Table("hoat_dong_tours").
// 		Select("tua_hoa_tieu_details.hoatieuid as hoatieuid").
// 		// Where("CONVERT (varchar(10), ?, 103) BETWEEN CONVERT (varchar(10), hoat_dong_tours.tungay, 103) AND CONVERT (varchar(10), hoat_dong_tours.denngay, 103)", ngay).
// 		Where("CONVERT (date, ?, 103) BETWEEN CONVERT (date, hoat_dong_tours.tungay, 103) AND CONVERT (date, hoat_dong_tours.denngay, 103)", ngay).
// 		Joins("LEFT JOIN tua_hoa_tieus ON tua_hoa_tieus.id = hoat_dong_tours.tourid").
// 		Joins("LEFT JOIN tua_hoa_tieu_details ON tua_hoa_tieu_details.idtuor = tua_hoa_tieus.id").
// 		Joins("LEFT JOIN hoatieus ON hoatieus.id = tua_hoa_tieu_details.hoatieuid").
// 		// Where("hoatieus.hangid IN ?", hoatieudan).
// 		Where("tua_hoa_tieu_details.hoatieuid NOT IN ?", hoatieuCongTac).
// 		Where("tua_hoa_tieus.status = ?", true).
// 		Scan(&hoaTieuHoatDong)
// 	if hoatieuHD.Error != nil {
// 		return nil, fmt.Errorf("Lỗi khi lấy hoa tieu hoạt động: %v", hoatieuHD.Error)
// 	}
// 	fmt.Errorf("Lỗi khi lấy hoa tieu hoạt động: %v", hoatieuHD)

// 	// Sử dụng map để loại bỏ các hoa tiêu trùng lặp
// 	hoaTieuHoatDongMap := make(map[int]struct{})
// 	hoaTieuHoatDongMapStr := make(map[string]struct{})
// 	for _, record := range hoaTieuHoatDong {
// 		hoaTieuHoatDongMap[record.Hoatieuid] = struct{}{}
// 	}

// 	var hoatieuScores []struct {
// 		HoaTieuID1   int    `gorm:"column:hoa_tieu_id1"`
// 		HoaTieuID2   string `gorm:"column:hoa_tieu_id2"`
// 		HoaTieuIDSid string `gorm:"column:hoa_tieu_id_sid"`
// 	}
// 	// Lấy danh sách hoa tiêu từ đơn hàng
// 	hoatieuUIDs := db.Debug().
// 		Table("don_hangs").
// 		Select("don_hangs.hoatieu1 as hoa_tieu_id1, don_hangs.hoatieu2 as hoa_tieu_id2, don_hangs.hoatieutsid as hoa_tieu_id_sid").
// 		Where("CONVERT(varchar(10), don_hangs.ngaydichchuyen, 103) = ? OR CONVERT(varchar(10), don_hangs.ngayvao, 103) = ? OR CONVERT(varchar(10), don_hangs.ngayra, 103) = ?", ngay, ngay, ngay).
// 		Where("don_hangs.status = ?", false).
// 		Scan(&hoatieuScores)
// 	if hoatieuUIDs.Error != nil {
// 		return nil, fmt.Errorf("Lỗi khi lấy hoa tieu từ đơn hàng: %v", hoatieuUIDs.Error)
// 	}

// 	// Xóa hoa tiêu có trong đơn hàng khỏi danh sách hoa tiêu hoạt động
// 	for _, score := range hoatieuScores {
// 		if score.HoaTieuID1 != 0 {
// 			delete(hoaTieuHoatDongMap, score.HoaTieuID1)
// 		}
// 		if score.HoaTieuID2 != "" {
// 			delete(hoaTieuHoatDongMapStr, score.HoaTieuID2)
// 		}
// 		if score.HoaTieuIDSid != "" {
// 			delete(hoaTieuHoatDongMapStr, score.HoaTieuIDSid)
// 		}
// 	}

// 	// Chuyển đổi map thành slice để sử dụng trong truy vấn cuối cùng
// 	var filteredHoaTieuIDs []int
// 	for hoaTieuID := range hoaTieuHoatDongMap {
// 		filteredHoaTieuIDs = append(filteredHoaTieuIDs, hoaTieuID)
// 	}

// 	var HoaTieu []ObjectHoaTieu
// 	query := db.Debug().
// 		Table("hoatieus").
// 		Select(`hoatieus.ID,
//                 hoatieus.status,
//                 hoatieus.userid,
//                 users.birthday AS ngaysinh,
//                 hoatieus.hangid,
//                 CONCAT(users.first_name,' ',users.last_name ) AS hoatieuname,
//                 users.first_name as firstname, users.last_name as lastname,
//                 users.username AS username,
//                 users.[image] as [image],
//                 hangs.name AS hangname,
//                 hoatieus.bpctac,
//                 chuc_vus.name as bpctacname,
//                 chung_chi_hoa_tieus.id as chungchihoatieuid,
//                 chung_chi_hoa_tieus.gcnkncm as gcnkncm,
//                 chung_chi_hoa_tieus.gcnvhdht as gcnvhdht,
//                 chung_chi_hoa_tieus.ngaycap as ngaycap,
//                 chung_chi_hoa_tieus.ngayhethan as ngayhethan,
//                 hoatieus.ngaysinh,
//                 hoatieus.sex,
//                 hoatieus.tuoi,
//                 hoatieus.phone,
//                 hoatieus.phone2,
//                 hoatieus.ID as hoatieuid,
// 								hoatieus.code`).
// 		Joins("LEFT JOIN users ON users.id = hoatieus.userid").
// 		Joins("LEFT JOIN hangs ON hangs.id = hoatieus.hangid").
// 		Joins("LEFT JOIN chuc_vus ON chuc_vus.id = hoatieus.bpctac").
// 		Joins("LEFT JOIN chung_chi_hoa_tieus ON hoatieus.chungchihoatieuid = chung_chi_hoa_tieus.id").
// 		Where("(users.first_name + ' ' + users.last_name) LIKE ?", "%"+name+"%").
// 		Where("hoatieus.ID IN (?)", filteredHoaTieuIDs).
// 		Order("hoatieus.created_at desc").
// 		Limit(10).
// 		Scan(&HoaTieu)

// 	fmt.Println("Hoa Tieu Hoat Dong:", hoaTieuHoatDong)
// 	fmt.Println("Hoa Tieu Scores:", hoatieuScores)
// 	fmt.Println("Filtered Hoa Tieu IDs:", filteredHoaTieuIDs)

// 	if query.Error != nil {
// 		return nil, query.Error
// 	}

// 	return HoaTieu, nil
// }
// func SearchHTBangKeSanLuongUpDateHoaTieu_Model(ngay, name string) ([]ObjectHoaTieu, error) {
// 	var hoatieuCongTac []int
// 	err := db.Debug().
// 		Table("phieu_cong_tacs").
// 		Select("phieu_cong_tacs.userid").
// 		Where("CONVERT (varchar(10), ?, 103) BETWEEN CONVERT (varchar(10), phieu_cong_tacs.tungay, 103) AND CONVERT (varchar(10), phieu_cong_tacs.denngay, 103)", ngay).
// 		Where("phieu_cong_tacs.approvalstatus = ?", true).
// 		Scan(&hoatieuCongTac).Error
// 	if err != nil {
// 		return nil, fmt.Errorf("Lỗi khi lấy phiếu công tác: %v", err)
// 	}

// 	// Kiểm tra nếu hoatieuCongTac không có phần tử nào, gán mặc định là 0
// 	if len(hoatieuCongTac) == 0 {
// 		hoatieuCongTac = []int{0}
// 	}

// 	// Tiếp tục xử lý hoatieuCongTac sau khi đã đảm bảo không nil và có giá trị mặc định là 0

// 	var hoaTieuHoatDong []struct {
// 		Hoatieuid int `gorm:"column:hoatieuid"`
// 	}
// 	// Lấy danh sách hoa tiêu hoạt động theo ngày và hạng
// 	hoatieuHD := db.Debug().
// 		Table("hoat_dong_tours").
// 		Select("tua_hoa_tieu_details.hoatieuid as hoatieuid").
// 		Where("CONVERT (varchar(10), ?, 103) BETWEEN CONVERT (varchar(10), hoat_dong_tours.tungay, 103) AND CONVERT (varchar(10), hoat_dong_tours.denngay, 103)", ngay).
// 		Joins("LEFT JOIN tua_hoa_tieus ON tua_hoa_tieus.id = hoat_dong_tours.tourid").
// 		Joins("LEFT JOIN tua_hoa_tieu_details ON tua_hoa_tieu_details.idtuor = tua_hoa_tieus.id").
// 		Joins("LEFT JOIN hoatieus ON hoatieus.id = tua_hoa_tieu_details.hoatieuid").
// 		// Where("hoatieus.hangid IN ?", hoatieudan).
// 		Where("tua_hoa_tieu_details.hoatieuid NOT IN ?", hoatieuCongTac).
// 		Where("tua_hoa_tieus.status = ?", true).
// 		Scan(&hoaTieuHoatDong)
// 	if hoatieuHD.Error != nil {
// 		return nil, fmt.Errorf("Lỗi khi lấy hoa tieu hoạt động: %v", hoatieuHD.Error)
// 	}

// 	// Sử dụng map để loại bỏ các hoa tiêu trùng lặp
// 	hoaTieuHoatDongMap := make(map[int]struct{})
// 	hoaTieuHoatDongMapStr := make(map[string]struct{})
// 	for _, record := range hoaTieuHoatDong {
// 		hoaTieuHoatDongMap[record.Hoatieuid] = struct{}{}
// 	}

// 	var hoatieuScores []struct {
// 		HoaTieuID1   int    `gorm:"column:hoa_tieu_id1"`
// 		HoaTieuID2   string `gorm:"column:hoa_tieu_id2"`
// 		HoaTieuIDSid string `gorm:"column:hoa_tieu_id_sid"`
// 	}

// 	// Xóa hoa tiêu có trong đơn hàng khỏi danh sách hoa tiêu hoạt động
// 	for _, score := range hoatieuScores {
// 		if score.HoaTieuID1 != 0 {
// 			delete(hoaTieuHoatDongMap, score.HoaTieuID1)
// 		}
// 		if score.HoaTieuID2 != "" {
// 			delete(hoaTieuHoatDongMapStr, score.HoaTieuID2)
// 		}
// 		if score.HoaTieuIDSid != "" {
// 			delete(hoaTieuHoatDongMapStr, score.HoaTieuIDSid)
// 		}
// 	}

// 	// Chuyển đổi map thành slice để sử dụng trong truy vấn cuối cùng
// 	var filteredHoaTieuIDs []int
// 	for hoaTieuID := range hoaTieuHoatDongMap {
// 		filteredHoaTieuIDs = append(filteredHoaTieuIDs, hoaTieuID)
// 	}

// 	var HoaTieu []ObjectHoaTieu
// 	query := db.Debug().
// 		Table("hoatieus").
// 		Select(`hoatieus.ID,
//                 hoatieus.status,
//                 hoatieus.userid,
//                 users.birthday AS ngaysinh,
//                 hoatieus.hangid,
//                 CONCAT(users.first_name,' ',users.last_name ) AS hoatieuname,
//                 users.first_name as firstname, users.last_name as lastname,
//                 users.username AS username,
//                 users.[image] as [image],
//                 hangs.name AS hangname,
//                 hoatieus.bpctac,
//                 chuc_vus.name as bpctacname,
//                 chung_chi_hoa_tieus.id as chungchihoatieuid,
//                 chung_chi_hoa_tieus.gcnkncm as gcnkncm,
//                 chung_chi_hoa_tieus.gcnvhdht as gcnvhdht,
//                 chung_chi_hoa_tieus.ngaycap as ngaycap,
//                 chung_chi_hoa_tieus.ngayhethan as ngayhethan,
//                 hoatieus.ngaysinh,
//                 hoatieus.sex,
//                 hoatieus.tuoi,
//                 hoatieus.phone,
//                 hoatieus.phone2,
//                 hoatieus.ID as hoatieuid,
// 								hoatieus.code`).
// 		Joins("LEFT JOIN users ON users.id = hoatieus.userid").
// 		Joins("LEFT JOIN hangs ON hangs.id = hoatieus.hangid").
// 		Joins("LEFT JOIN chuc_vus ON chuc_vus.id = hoatieus.bpctac").
// 		Joins("LEFT JOIN chung_chi_hoa_tieus ON hoatieus.chungchihoatieuid = chung_chi_hoa_tieus.id").
// 		Where("(users.first_name + ' ' + users.last_name) LIKE ?", "%"+name+"%").
// 		Where("hoatieus.ID IN (?)", filteredHoaTieuIDs).
// 		Order("hoatieus.created_at desc").
// 		Limit(10).
// 		Scan(&HoaTieu)

// 	// fmt.Println("Hoa Tieu Hoat Dong:", hoaTieuHoatDong)
// 	// fmt.Println("Hoa Tieu Scores:", hoatieuScores)
// 	// fmt.Println("Filtered Hoa Tieu IDs:", filteredHoaTieuIDs)

// 	if query.Error != nil {
// 		return nil, query.Error
// 	}

// 	return HoaTieu, nil
// }
// func GetMultipleHoaTieuTrue_Model(name string) (*[]ObjectHoaTieu, error) {
// 	item := []ObjectHoaTieu{}
// 	// Xây dựng câu lệnh SQL từ chuỗi ID
// 	query := `
//         SELECT
//             ht.id as ID,
//             CONCAT(u.first_name, ' ', u.last_name) AS hoatieuname
//         FROM hoatieus ht
//         LEFT JOIN users u ON u.id = ht.userid
//         WHERE ht.id IN (
//             SELECT CAST(value AS INT) FROM STRING_SPLIT(?, ',')
//         );
//     `

// 	// Thực hiện truy vấn với GORM
// 	if err := db.Raw(query, name).Scan(&item).Error; err != nil {
// 		if err == gorm.ErrRecordNotFound {
// 			return &item, nil
// 		}
// 		return nil, err
// 	}

// 	return &item, nil
// }

// func SearchHoaTieuRole_Model(name, role string) ([]ObjectHoaTieu, error) {
// 	var HoaTieu []ObjectHoaTieu
// 	query := db.Debug().
// 		Table("hoatieus").
// 		Select(`hoatieus.ID,
//                 hoatieus.status,
//                 hoatieus.userid,
//                 users.birthday AS ngaysinh,
//                 hoatieus.hangid,
//                 CONCAT(users.first_name,' ',users.last_name ) AS hoatieuname,
//                 users.first_name as firstname, users.last_name as lastname,
//                 users.username AS username,
//                 users.[image] as [image],
//                 hangs.name AS hangname,
//                 hoatieus.bpctac,
//                 chuc_vus.name as bpctacname,
//                 chung_chi_hoa_tieus.id as chungchihoatieuid,
//                 chung_chi_hoa_tieus.gcnkncm as gcnkncm,
//                 chung_chi_hoa_tieus.gcnvhdht as gcnvhdht,
//                 chung_chi_hoa_tieus.ngaycap as ngaycap,
//                 chung_chi_hoa_tieus.ngayhethan as ngayhethan,
//                 hoatieus.ngaysinh,
//                 hoatieus.sex,
//                 hoatieus.tuoi,
//                 hoatieus.phone,
//                 hoatieus.phone2,
//                 hoatieus.ID as hoatieuid`).
// 		Joins("LEFT JOIN users ON users.id = hoatieus.userid").
// 		Joins("LEFT JOIN hangs ON hangs.id = hoatieus.hangid").
// 		Joins("LEFT JOIN chuc_vus ON chuc_vus.id = hoatieus.bpctac").
// 		Joins("LEFT JOIN chung_chi_hoa_tieus ON hoatieus.chungchihoatieuid = chung_chi_hoa_tieus.id").
// 		Where("(users.first_name + ' ' + users.last_name) LIKE ?", "%"+name+"%").
// 		Where("users.role = ?", role).
// 		// Where("hoatieus.ID IN (?) AND users.role = ?", name, role).
// 		Order("hoatieus.created_at desc").
// 		Limit(10).
// 		Scan(&HoaTieu)

// 	if query.Error != nil {
// 		return nil, query.Error
// 	}

// 	return HoaTieu, nil
// }
