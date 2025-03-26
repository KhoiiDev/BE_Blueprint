package user_service

import (
	"be-hoatieu/models"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// type User struct {
// 	Username          string
// 	Birthday          string
// 	Email             string
// 	Image             string
// 	Firstname         string
// 	Lastname          string
// 	Status            string
// 	Phone             string
// 	City              string
// 	Address           string
// 	Address2          string
// 	Role              string
// 	ResetKey          string
// 	ResetCount        int
// 	ResetTimestamp    string
// 	ResetKeyTimestamp string
// 	Password          string
// 	Id                string
// 	UserId            uint
// 	HangId            uint
// 	PhongBan          uint
// 	Sex               string
// 	Tuoi              string
// 	NhanvienId        uint
// }

type User struct {
	Username          string `gorm:"column:username" json:"username"`
	Password          string `gorm:"column:password" json:"password"`
	Birthday          string `gorm:"column:birthday" json:"birthday"`
	Email             string `gorm:"column:email" json:"email"`
	Image             string `gorm:"column:image" json:"image"`
	Firstname         string `gorm:"column:first_name" json:"firstName"`
	Lastname          string `gorm:"column:last_name" json:"lastName"`
	Phone             string `gorm:"column:phone" json:"phone"`
	Phone2            string `gorm:"column:phone2" json:"phone2"`
	City              string `gorm:"column:city" json:"city"`
	Address           string `gorm:"column:address" json:"address"`
	Address2          string `gorm:"column:address2" json:"address2"`
	Role              string `gorm:"column:role;" json:"role"`
	ResetKey          string `gorm:"column:reset_key" json:"resetKey"`
	ResetCount        int    `gorm:"column:reset_count" json:"resetCount"`
	ResetTimestamp    string `gorm:"column:reset_timestamp" json:"resetTimestamp"`
	ResetKeyTimestamp string `gorm:"column:reset_key_timestamp" json:"resetKeyTimestamp"`
	PhongBan          string `gorm:"column:phongban" json:"phongban"`
	NhanvienId        uint   `gorm:"column:nhanvienid" json:"nhanvienid"`
	NhanvienIdStr     string `gorm:"column:nhanvienidstr" json:"nhanvienidstr"`
	Status            bool   `gorm:"column:status" json:"status"`
}

type UserImage struct {
	Image string
}
type UserPassword struct {
	Username string
	Password string
}

type ConnectAPI struct {
	USERNAME   string
	IP         string
	IP_INTER   string
	DB_INTER   string
	USER_INNER string
	PASS_INNER string
}

func (m *User) RegisterPublicUser(password string) error {
	salt := 14
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), salt)

	if err != nil {
		return err
	}

	User := map[string]interface{}{
		"username":          m.Username,
		"email":             m.Email,
		"image":             m.Image,
		"firstName":         m.Firstname,
		"lastName":          m.Lastname,
		"status":            m.Status,
		"phone":             m.Phone,
		"city":              m.City,
		"address":           m.Address,
		"nhanvienid":        m.NhanvienId,
		"nhanvienidstr":     m.NhanvienIdStr,
		"address2":          "",
		"role":              m.Role,
		"resetKey":          "",
		"resetKeyTimestamp": "",
		"resetTimestamp":    "",
		"resetCount":        0,
		"password":          string(hashedPassword),
	}

	if err := models.AddUser(User); err != nil {
		return err
	}

	return nil
}

func (m *User) CreateUserAndHoaTieu_Service(password string) error {
	salt := 14
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), salt)

	if err != nil {
		return err
	}

	User := map[string]interface{}{
		"username":          m.Username,
		"email":             m.Email,
		"image":             "",
		"firstName":         m.Firstname,
		"lastName":          m.Lastname,
		"status":            "1",
		"phone":             m.Phone,
		"city":              m.City,
		"address":           m.Address,
		"address2":          "",
		"role":              "hoatieu",
		"resetKey":          "",
		"resetKeyTimestamp": "",
		"resetTimestamp":    "",
		"resetCount":        0,
		"password":          string(hashedPassword),
		"phongban":          m.PhongBan,
		"nhanvienid":        m.NhanvienId,
		"nhanvienidstr":     m.NhanvienIdStr,
	}

	if err := models.AddUserAndHoaTieu(User); err != nil {
		return err
	}
	return nil
}

func (m *User) GetUserUsername() (*models.User, error) {
	User, err := models.GetUserUserName(m.Username)

	if err != nil {
		return nil, err
	}

	return User, nil
}

func (m *User) GetUserCheck() (*models.User, error) {
	User, err := models.GetUserCheck(m.Username, m.Phone, m.Email)

	if err != nil {
		return nil, err
	}

	return User, nil
}

func (m *User) GetUserCheckUpdate(id string) (*models.User, error) {
	User, err := models.GetUserCheckUpdate(id, m.Phone, m.Email)
	if err != nil {
		return nil, err
	}
	return User, nil
}

func (m *User) GetUserEmail() (*models.User, error) {
	User, err := models.GetUserEmail(m.Email)

	if err != nil {
		return nil, err
	}

	return User, nil
}

func (m *User) GetUserLogin() (*models.User, error) {
	User, err := models.GetUserLogin(m.Username)

	if err != nil {
		return nil, err
	}

	return User, nil
}

// Get All User
func (a *User) GetAllUser_Service(limit int, page int, role, name string) (*[]models.User, int64, error) {
	item, totalRecords, err := models.GetAllUser_Model(limit, page, role, name)

	if err != nil {
		return nil, 0, err
	}
	return item, totalRecords, nil
}

func (a *User) GetAllUserTrue_Service() (*[]models.User, error) {
	item, err := models.GetAllUserTrue_Model()
	if err != nil {
		return nil, err
	}
	return item, nil
}

// Get ID User
func (a *User) GetByIdUser_Service(id string) (*models.User, error) {
	item, err := models.GetByIdUser_Model(id)
	if err != nil {
		return nil, err
	}

	return item, nil
}

// Get username User
func (a *User) GetByUsername_Service(username string) (*models.User, error) {
	item, err := models.GetUserUserName(username)
	if err != nil {
		return nil, err
	}

	return item, nil
}

// Search Username User
func (a *User) SearchUser_Service(username string) (*[]models.User, error) {
	item, err := models.SearchUser_Model(username)
	if err != nil {
		return nil, err
	}
	return item, nil
}

// Update User
func (a *User) UpdatUsers(id string) error {
	var hashedPassword string
	if a.Password != "" {
		// Nếu password không rỗng, mã hóa
		salt := 14
		hashed, err := bcrypt.GenerateFromPassword([]byte(a.Password), salt)
		if err != nil {
			return err
		}
		hashedPassword = string(hashed)
	}

	// Tạo map user và chỉ thêm password nếu đã mã hóa
	user := map[string]interface{}{
		"Username":      a.Username,
		"NhanvienId":    a.NhanvienId,
		"NhanvienIdStr": a.NhanvienIdStr,
		"Email":         a.Email,
		"Image":         a.Image,
		"Firstname":     a.Firstname,
		"Lastname":      a.Lastname,
		"Phone":         a.Phone,
		"City":          a.City,
		"Address":       a.Address,
		"Status":        a.Status,
		"Role":          a.Role,
	}

	// Nếu password không rỗng và đã được mã hóa, thêm vào map user
	if hashedPassword != "" {
		user["Password"] = hashedPassword
	}

	// Gọi models.UpdateUsers_Model
	if err := models.UpdateUsers_Model(id, user); err != nil {
		return err
	}

	return nil
}

// Update Image User
func (a *User) UpdatImageUsers(id string) error {

	user := map[string]interface{}{
		"Image": a.Image,
	}

	fmt.Print(user)
	if err := models.UpdateImageUsers_Model(id, user); err != nil {
		return err
	}

	return nil
}

func (a *User) UpdateDetailUsers(id string) error {

	user := map[string]interface{}{
		"Username":  a.Username,
		"Email":     a.Email,
		"Firstname": a.Firstname,
		"Lastname":  a.Lastname,
		"Status":    a.Status,
		"Phone":     a.Phone,
		"City":      a.City,
		"Address":   a.Address,
		"Role":      a.Role,
		"status":    a.Status,
	}
	if err := models.UpdateDetailUsers_Model(id, user); err != nil {
		return err
	}
	return nil
}

func (a *User) UpdateUserAndHoaTieu(id string) error {

	user := map[string]interface{}{
		"Username":      a.Username,
		"Email":         a.Email,
		"Firstname":     a.Firstname,
		"Lastname":      a.Lastname,
		"Status":        a.Status,
		"Phone":         a.Phone,
		"City":          a.City,
		"Address":       a.Address,
		"Role":          a.Role,
		"status":        a.Status,
		"phongban":      a.PhongBan,
		"phone":         a.Phone,
		"nhanvienid":    a.NhanvienId,
		"nhanvienidstr": a.NhanvienIdStr,
	}
	if err := models.UpdateDetailUsers_Model(id, user); err != nil {
		return err
	}
	return nil
}

// Update Password User
func (a *UserPassword) UpdatUsersPassword(id string) error {
	fmt.Print("===================A==================", a)

	salt := 14
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(a.Password), salt)

	if err != nil {
		return err
	}

	user := map[string]interface{}{
		"Username": a.Username,
		"Password": string(hashedPassword),
	}
	if err := models.UpdateUsersPassword(id, user); err != nil {
		return err
	}

	return nil
}

// Hàm xóa người dùng theo ID
func (a *User) DeleteUser(id string) error {
	// Gọi hàm xóa trong models để xóa người dùng
	err := models.DeleteUserById(id)
	if err != nil {
		// Nếu có lỗi khi xóa, trả về lỗi
		return fmt.Errorf("failed to delete user: %v", err)
	}
	// Nếu xóa thành công, trả về nil
	return nil
}
