package models

import (
	"fmt"

	"gorm.io/gorm"
)

type User struct {
	BaseModel
	Username          string          `gorm:"column:username" json:"username"`
	Birthday          string          `gorm:"column:birthday" json:"birthday"`
	Email             string          `gorm:"column:email" json:"email"`
	Image             string          `gorm:"column:image" json:"image"`
	Firstname         string          `gorm:"column:first_name" json:"firstName"`
	Lastname          string          `gorm:"column:last_name" json:"lastName"`
	Phone             string          `gorm:"column:phone" json:"phone"`
	Phone2            string          `gorm:"column:phone2" json:"phone2"`
	City              string          `gorm:"column:city" json:"city"`
	Address           string          `gorm:"column:address" json:"address"`
	Address2          string          `gorm:"column:address2" json:"address2"`
	Role              string          `gorm:"column:role;" json:"role"`
	ResetKey          string          `gorm:"column:reset_key" json:"resetKey"`
	ResetCount        int             `gorm:"column:reset_count" json:"resetCount"`
	ResetTimestamp    string          `gorm:"column:reset_timestamp" json:"resetTimestamp"`
	ResetKeyTimestamp string          `gorm:"column:reset_key_timestamp" json:"resetKeyTimestamp"`
	PhongBan          string          `gorm:"column:phongban" json:"phongban"`
	PhongBanid        string          `gorm:"column:phongbanid" json:"phongbanid"`
	PhongBanlist      string          `gorm:"column:phongbanlist" json:"phongbanlist"`
	Isduyet           bool            `gorm:"column:isduyet" json:"isduyet"`
	NhanvienId        uint            `gorm:"column:nhanvienid" json:"nhanvienid"`
	NhanvienIdStr     string          `gorm:"column:nhanvienidstr" json:"nhanvienidstr"`
	UserCredentials   UserCredentials `gorm:"foreignKey:User_id;" json:"UserCredentials"`
	Status            bool            `gorm:"column:status" json:"status"`
}

type UserPassword struct {
	Password string `gorm:"column:password" json:"password"`
	UserId   uint   `gorm:"column:User_id" json:"UserId"`
}
type UserChangePassword struct {
	Password string `gorm:"column:password" json:"password"`
}

func (a *User) GetUsers(maps interface{}) (*User, error) {
	var user User

	if maps != nil {
		err := db.Where("id = ?", a.Id).Where(maps).First(&user).Error
		if err != nil && err != gorm.ErrRecordNotFound {
			return nil, err
		}
	} else {
		err := db.Where("id = ?", a.Id).First(&user).Error
		if err != nil && err != gorm.ErrRecordNotFound {
			return nil, err
		}
	}
	return &user, nil
}

func GetUser(nickName string) (*User, error) {
	var user User

	err := db.Where("username = ?", nickName).First(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return &user, nil
}

// Edit User modify a single User
func EditUser(id int, data interface{}) error {
	if err := db.Model(&User{}).Where("id = ?", id).Updates(data).Error; err != nil {
		return err
	}
	return nil

}

func EditUserByCondition(condition interface{}, data interface{}) error {
	if err := db.Model(&User{}).Where(condition).Updates(data).Error; err != nil {
		return err
	}
	return nil
}

func AddUser(data map[string]interface{}) error {
	user := User{
		Username:          data["username"].(string),
		Email:             data["email"].(string),
		Image:             data["image"].(string),
		Firstname:         data["firstName"].(string),
		Lastname:          data["lastName"].(string),
		Status:            data["status"].(bool),
		Phone:             data["phone"].(string),
		City:              data["city"].(string),
		Address:           data["address"].(string),
		Address2:          data["address2"].(string),
		Role:              data["role"].(string),
		ResetKey:          data["resetKey"].(string),
		ResetKeyTimestamp: data["resetKeyTimestamp"].(string),
		ResetTimestamp:    data["resetTimestamp"].(string),
		ResetCount:        data["resetCount"].(int),
		UserCredentials:   UserCredentials{Password: data["password"].(string)},
		NhanvienId:        data["nhanvienid"].(uint),
		NhanvienIdStr:     data["nhanvienidstr"].(string),
	}
	result := db.Create(&user)
	if err := result.Error; err != nil {
		return err
	}

	return nil
}

func AddUserAndHoaTieu(data map[string]interface{}) error {
	user := User{
		Username:          data["username"].(string),
		Birthday:          data["ngaysinh"].(string),
		Email:             data["email"].(string),
		Image:             data["image"].(string),
		Firstname:         data["firstName"].(string),
		Lastname:          data["lastName"].(string),
		Status:            data["status"].(bool),
		Phone:             data["phone"].(string),
		City:              data["city"].(string),
		Address:           data["address"].(string),
		Address2:          data["address2"].(string),
		Role:              data["role"].(string),
		ResetKey:          data["resetKey"].(string),
		ResetCount:        data["resetCount"].(int),
		ResetTimestamp:    data["resetTimestamp"].(string),
		ResetKeyTimestamp: data["resetKeyTimestamp"].(string),
		UserCredentials:   UserCredentials{Password: data["password"].(string)},
		NhanvienId:        data["nhanvienid"].(uint),
		NhanvienIdStr:     data["nhanvienidstr"].(string),
	}
	result := db.Create(&user)
	if err := result.Error; err != nil {
		return err
	}

	return nil
}

// user := User{
// 	UserCredentials: UserCredentials{Password: data["Password"].(string)},
// }
// fmt.Print("===================model==================", user)

//	if err := db.Exec("UPDATE user_credentials SET password = ? WHERE User_id = ?", user.UserCredentials.Password, UserId).Error; err != nil {
//		return err
//	}
func UpdateUsers_Model(id string, data map[string]interface{}) error {
	user := User{
		Email:         data["Email"].(string),
		Image:         data["Image"].(string),
		Firstname:     data["Firstname"].(string),
		Lastname:      data["Lastname"].(string),
		NhanvienId:    data["NhanvienId"].(uint),
		NhanvienIdStr: data["NhanvienIdStr"].(string),
		Phone:         data["Phone"].(string),
		City:          data["City"].(string),
		Address:       data["Address"].(string),
		Status:        data["Status"].(bool),
		Role:          data["Role"].(string),
	}

	// Cập nhật các trường của User trước
	if err := db.Model(&user).Where("id = ?", id).Updates(map[string]interface{}{
		"Email":         user.Email,
		"Image":         user.Image,
		"Firstname":     user.Firstname,
		"Lastname":      user.Lastname,
		"NhanvienId":    user.NhanvienId,
		"NhanvienIdStr": user.NhanvienIdStr,
		"Status":        user.Status,
		"Phone":         user.Phone,
		"City":          user.City,
		"Address":       user.Address,
		"Role":          user.Role,
	}).Error; err != nil {
		return err
	}
	// Nếu có mật khẩu trong dữ liệu và mật khẩu không phải chuỗi rỗng, thì cập nhật UserCredentials
	if password, ok := data["Password"].(string); ok && password != "" {
		userCredentials := UserCredentials{
			Password: password, // Cập nhật mật khẩu mới
		}
		// Cập nhật trường password trong bảng UserCredentials cho user với id tương ứng
		if err := db.Model(&userCredentials).Where("user_id = ?", id).Update("password", userCredentials.Password).Error; err != nil {
			return err
		}
	}

	return nil
}

func UpdateImageUsers_Model(id string, data map[string]interface{}) error {
	user := User{
		Image: data["Image"].(string),
	}

	if err := db.Model(&user).Where("id = ?", id).Updates(map[string]interface{}{
		"Image": user.Image,
	}).Error; err != nil {
		return err
	}
	return nil
}

func UpdateDetailUsers_Model(id string, data map[string]interface{}) error {
	user := User{
		Email:     data["Email"].(string),
		Firstname: data["Firstname"].(string),
		Lastname:  data["Lastname"].(string),
		Status:    data["Status"].(bool),
		Phone:     data["Phone"].(string),
		City:      data["City"].(string),
		Address:   data["Address"].(string),
		Role:      data["Role"].(string),
	}

	if err := db.Model(&user).Where("id = ?", id).Updates(map[string]interface{}{
		"Email":     user.Email,
		"Firstname": user.Firstname,
		"Lastname":  user.Lastname,
		"Status":    user.Status,
		"Phone":     user.Phone,
		"City":      user.City,
		"Address":   user.Address,
		"Role":      user.Role,
	}).Error; err != nil {
		return err
	}
	return nil
}

func GetUserUserName(username string) (*User, error) {
	var user User
	err := db.Table("users").
		Select("users.*").
		// Joins("LEFT JOIN nhan_viens ON users.nhanvienid = nhan_viens.id").
		Where("users.username = ? AND users.status = ?", username, 1).
		First(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return &user, nil
}
func GetUserCheck(username, phone, email string) (*User, error) {
	var user User
	err := db.Model(&User{}).
		Where("username = ? OR phone = ? OR email = ?", username, phone, email).
		First(&user).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}

	return &user, nil
}

func GetUserCheckUpdate(id, phone, email string) (*User, error) {
	var user User
	err := db.Model(&User{}).
		Where("id != ? AND (phone = ? OR email = ?)", id, phone, email).
		First(&user).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}

	return &user, nil
}

func GetUserEmail(email string) (*User, error) {
	var user User

	err := db.Model(&User{}).Where("email = ?", email).First(&user).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return &user, nil
}

func GetUserLogin(username string) (*User, error) {
	var user User
	// err := db.Model(&User{}).Preload("UserCredentials").Where("username = ?", username).First(&user).Error
	err := db.Model(&User{}).Preload("UserCredentials").Where("username = ? AND status = ?", username, 1).First(&user).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &user, nil
}
func handleEmptyString(value string) *string {
	if value == "" || value == "undefined" || value == "Invalid%20date" || value == " " {
		return nil
	}
	return &value
}

// Get All user
func GetAllUser_Model(limit int, page int, role string, name string) (*[]User, int64, error) {
	item := []User{}
	var totalRecords int64
	if name == "undefined" {
		name = ""
	}
	var roleHandle = handleEmptyString(role)
	// Đếm tổng số phần tử
	err := db.Model(&User{}).
		Where("CONCAT(first_name, ' ', last_name) LIKE ? OR username LIKE ? ", "%"+name+"%", "%"+name+"%").
		// Where("username LIKE ?", "%"+name+"%").

		// Where("first_name LIKE ?", "%"+name+"%").
		Where("deleted_at IS NULL").
		Where("? IS NULL OR role = ?", roleHandle, roleHandle).
		Count(&totalRecords).Error
	if err != nil {
		return nil, 0, err
	}
	offset := (page - 1) * limit

	err = db.Model(&User{}).
		Where("CONCAT(first_name, ' ', last_name) LIKE ? OR username LIKE ? ", "%"+name+"%", "%"+name+"%").

		// Where("first_name LIKE ?", "%"+name+"%").
		Where("deleted_at IS NULL").
		Where("? IS NULL OR role = ?", roleHandle, roleHandle).
		Offset(offset).
		Limit(limit).
		// Order("users.created_at desc").
		Find(&item).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, 0, err
	}

	return &item, totalRecords, nil
}

func GetAllUserTrue_Model() (*[]User, error) {

	item := []User{}
	// check loi database
	// err := db.Find(&item).Error
	err := db.Where("status = ? ", true).Find(&item).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return &item, nil
}

// Get ID User
func GetByIdUser_Model(id string) (*User, error) {
	var item User

	err := db.Where("id = ?", id).First(&item).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &item, nil
}

// Search Username user
func SearchUser_Model(username string) (*[]User, error) {
	item := []User{}

	err := db.Where("username LIKE ?", "%"+username+"%").Find(&item).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return &item, nil
}

// Update Password
func UpdateUsersPassword(UserId string, data map[string]interface{}) error {
	user := User{
		UserCredentials: UserCredentials{Password: data["Password"].(string)},
	}
	fmt.Print("===================model==================", user)

	if err := db.Exec("UPDATE user_credentials SET password = ? WHERE User_id = ?", user.UserCredentials.Password, UserId).Error; err != nil {
		return err
	}
	// if err := db.Model(&user).Where("id = ?", UserId).Joins("user_credentials").Updates(map[string]interface{}{}).Error; err != nil {
	// 	return err
	// }
	return nil
}

// DeleteUserById xóa người dùng theo ID
func DeleteUserById(id string) error {
	// Xóa user theo id
	if err := db.Where("id = ?", id).Delete(&User{}).Error; err != nil {
		return err
	}
	return nil
}
