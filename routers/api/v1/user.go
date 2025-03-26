package v1

import (
	"be-hoatieu/pkg/app"
	user_service "be-hoatieu/services/user"
	"fmt"
	"strconv"

	"be-hoatieu/pkg/utils"

	"github.com/gofiber/fiber/v2"
)

type LoginForm struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type FormGetUserName struct {
	Username string `json:"username" validate:"required"`
}

func SignIn(c *fiber.Ctx) error {
	form := &LoginForm{}

	// Check, if received JSON data is valid.
	if err := c.BodyParser(form); err != nil {
		// Return status 400 and error message.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":    400,
			"errors":  err.Error(),
			"message": err.Error(),
		})
	}

	// Validate
	errors := app.ValidateStruct(*form)

	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"errors":  errors,
			"message": errors,
		})

	}

	// JWT
	check, role, id := ProcessUserLogin(form.Username, form.Password)

	// ip, ip_inter, db_inner, user_inner, pass_inter := ProcessUserLogin_ConnectAPI(form.Username, form.Password)

	if check && (len(role) > 0) {
		token, err := utils.GenerateToken(form.Username, role, id)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"code":    500,
				"errors":  err.Error(),
				"message": err.Error(),
			})
		}

		data := make(map[string]string)
		data["token"] = token
		// data["ip"] = ip
		// data["ip_inter"] = ip_inter
		// data["db_inter"] = db_inner
		// data["user_iner"] = user_inner
		// data["pass_inter"] = pass_inter

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"code":    200,
			"errors":  nil,
			"message": "Success login",
			"data":    data,
		})
	}
	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"code":    401,
		"errors":  nil,
		"message": "Error Auth",
	})
}

func ProcessUserLogin(username string, password string) (bool, string, string) {
	UserService := user_service.User{Username: username}
	User, err := UserService.GetUserLogin()
	if err != nil {
		return false, "", ""
	}
	// return utils.PasswordVerify(password, User.UserCredentials.Password), User.Role, strconv.Itoa(User.Id)
	return utils.PasswordVerify(password, User.UserCredentials.Password), User.Role, strconv.FormatUint(uint64(User.Id), 10) // Chuyển đổi uint sang string

}

type UserUpdate struct {
	Username      string `json:"username"`
	Password      string `json:"password"`
	NhanvienId    uint   `json:"nhanvienid"`
	NhanvienIdStr string `json:"nhanvienidstr"`
	Email         string `json:"email" validate:"required,email"`
	Image         string `json:"image"`
	Firstname     string `json:"firstName" validate:"required"`
	Lastname      string `json:"lastName" validate:"required"`
	Phone         string `json:"phone" validate:"required"`
	City          string `json:"city" validate:"required"`
	Role          string `json:"role" validate:"required"`
	Status        bool   `json:"status" validate:"required"`
	Address       string `json:"address" validate:"required"`
}
type UserDetailUpdate struct {
	Email         string `json:"email"`
	Image         string `json:"image"`
	Firstname     string `json:"firstName"`
	Lastname      string `json:"lastName"`
	Phone         string `json:"phone"`
	City          string `json:"city"`
	Address       string `json:"address"`
	Role          string `json:"role"`
	Status        bool   `json:"status"`
	HangId        uint   `form:"hangid" `
	BPCTac        uint   `form:"bpctac" `
	Phone2        string `form:"phone2" `
	Sex           string `form:"sex" `
	PhongBan      uint   `form:"phongban" `
	Tuoi          string `form:"tuoi" `
	NhanvienId    uint   `form:"nhanvienid"`
	NhanvienIdStr string `form:"nhanvienidstr"`
}
type SignUpForm struct {
	Username      string `json:"username" validate:"required"`
	Email         string `json:"email" validate:"required,email"`
	Image         string `json:"image"`
	Firstname     string `json:"firstName"`
	Lastname      string `json:"lastName"`
	Phone         string `json:"phone"`
	City          string `json:"city"`
	Address       string `json:"address"`
	Address2      string `json:"address2"`
	Password      string `json:"password" validate:"required"`
	Status        bool   `json:"status" `
	UserId        uint   `form:"userid" `
	HangId        uint   `form:"hangid" `
	PhongBan      uint   `form:"phongban" `
	Sex           string `form:"sex" `
	Tuoi          string `form:"tuoi" `
	Role          string `form:"role" `
	NhanvienId    uint   `json:"nhanvienid"`
	NhanvienIdStr string `json:"nhanvienidstr"`
}

func SignUp(c *fiber.Ctx) error {
	form := &SignUpForm{}
	// Check, if received JSON data is valid.
	if err := c.BodyParser(form); err != nil {
		// Return status 400 and error message.
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"code":    400,
			"errors":  err.Error(),
			"message": err.Error(),
		})
	}

	// Validate
	errors := app.ValidateStruct(*form)

	if errors != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"code":    400,
			"errors":  errors,
			"message": errors,
		})

	}

	// Checking
	valid, duplicateField := checkUser(c, *form)
	if !valid {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"code":    400,
			"errors":  nil,
			"message": fmt.Sprintf("%s exists.", duplicateField),
		})
	}

	if !valid {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"code":    400,
			"errors":  nil,
			"message": "Username exists",
		})
	}

	UserService := user_service.User{
		Username:      form.Username,
		Email:         form.Email,
		Image:         form.Image,
		Firstname:     form.Firstname,
		Lastname:      form.Lastname,
		Phone:         form.Phone,
		City:          form.City,
		Address:       form.Address,
		Address2:      form.Address2,
		Password:      form.Password,
		Status:        form.Status,
		Role:          form.Role,
		NhanvienId:    form.NhanvienId,
		NhanvienIdStr: form.NhanvienIdStr,
	}

	if err := UserService.RegisterPublicUser(form.Password); err != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"code":    400,
			"errors":  err,
			"message": "Register false",
		})
	}

	data := make(map[string]string)
	data["User_code"] = form.Username

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":    200,
		"errors":  nil,
		"message": "Success",
		"data":    data,
	})

}

func CreateUserAndHoaTieu(c *fiber.Ctx) error {
	form := &SignUpForm{}
	// Check, if received JSON data is valid.
	if err := c.BodyParser(form); err != nil {
		// Return status 400 and error message.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":    400,
			"errors":  err.Error(),
			"message": err.Error(),
		})
	}

	// Validate
	errors := app.ValidateStruct(*form)

	if errors != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"code":    400,
			"errors":  errors,
			"message": errors,
		})

	}

	// Checking
	valid, duplicateField := checkUser(c, *form)
	if !valid {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"code":    400,
			"errors":  nil,
			"message": fmt.Sprintf("%s exists.", duplicateField),
		})
	}

	if !valid {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"code":    400,
			"errors":  nil,
			"message": "Username exists",
		})
	}

	UserService := user_service.User{
		Username:      form.Username,
		Email:         form.Email,
		Firstname:     form.Firstname,
		Lastname:      form.Lastname,
		Phone:         form.Phone,
		City:          form.City,
		Address:       form.Address,
		Password:      form.Password,
		NhanvienId:    form.NhanvienId,
		Role:          form.Role,
		NhanvienIdStr: form.NhanvienIdStr,
	}

	if err := UserService.CreateUserAndHoaTieu_Service(form.Password); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":    400,
			"errors":  err,
			"message": "Register false",
		})
	}

	data := make(map[string]string)
	data["User_code"] = form.Username

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":    200,
		"errors":  nil,
		"message": "Success",
		"data":    data,
	})

}

func checkUser(c *fiber.Ctx, form SignUpForm) (bool, string) {
	isDuplicateUser, duplicateField, err := isDuplicateUser(form.Username, form.Phone, form.Email)
	if err != nil {
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
		})
		return false, duplicateField
	}

	if isDuplicateUser {
		return false, duplicateField
	}

	return true, ""
}

func checkUserPut(c *fiber.Ctx, id string, form UserDetailUpdate) (bool, string) {
	isDuplicate, field, err := isCheckUpdate(id, form.Phone, form.Phone2, form.Email)
	if err != nil {
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
		})
		return false, field
	}

	if isDuplicate {
		c.Status(fiber.StatusOK).JSON(fiber.Map{
			"error": true,
			"field": field,
		})
		return false, field
	}

	return true, ""
}

func isDuplicateUser(username, phone, email string) (bool, string, error) {
	UserService := user_service.User{Username: username, Phone: phone, Email: email}

	User, err := UserService.GetUserCheck()
	if err != nil {
		return true, "", err
	}

	if User != nil {
		if User.Username == username {
			return true, "Username", nil
		}
		if User.Phone == phone {
			return true, "Phone", nil
		}
		if User.Email == email {
			return true, "Email", nil
		}
	}

	return false, "", nil
}

func isCheckUpdate(id, phone, phone2, email string) (bool, string, error) {
	UserService := user_service.User{Phone: phone, Email: email}

	User, err := UserService.GetUserCheckUpdate(id)
	if err != nil {
		return false, "", err
	}

	if User != nil {
		if User.Phone == phone {
			return true, "Phone", nil
		}
		if User.Phone2 == phone2 {
			return true, "Phone2", nil
		}
		if User.Email == email {
			return true, "Email", nil
		}
	}

	return false, "", nil
}

// Get All User
func GetAllUser_Router(c *fiber.Ctx) error {
	item := user_service.User{}
	limit := c.Query("limit")
	page := c.Query("page")

	name := c.Query("name")
	role := c.Query("role")
	limitStr, err := strconv.Atoi(limit)
	PageStr, err := strconv.Atoi(page)

	data, totalRecords, err := item.GetAllUser_Service(limitStr, PageStr, role, name)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success":      true,
		"data":         data,
		"totalRecords": totalRecords,
	})
}

func GetAllUserTrue_Router(c *fiber.Ctx) error {
	item := user_service.User{}
	data, err := item.GetAllUserTrue_Service()

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data":    data,
	})
}

// Get ID User
func GetByIdUser_Router(c *fiber.Ctx) error {
	item := user_service.User{}
	data, err := item.GetByIdUser_Service(c.Params("id"))

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data":    data,
	})
}

// Get UserName User
func GetByUserName_Router(c *fiber.Ctx) error {
	item := user_service.User{}
	data, err := item.GetByUsername_Service(c.Params("username"))

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data":    data,
	})
}

// Search Username User
func SearchUser_Router(c *fiber.Ctx) error {
	formSearch := new(LoginForm)
	item := user_service.User{}
	if err := c.BodyParser(formSearch); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	data, err := item.SearchUser_Service(formSearch.Username)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data":    data,
	})
}

// Update User
func PutUsers(c *fiber.Ctx) error {
	form := &UserUpdate{}
	// form := &ItemsUpdate{}
	if err := c.BodyParser(form); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	UserService := user_service.User{
		// Birthday:  form.Birthday,
		Username:      form.Username,
		Email:         form.Email,
		Image:         form.Image,
		Firstname:     form.Firstname,
		Lastname:      form.Lastname,
		Phone:         form.Phone,
		City:          form.City,
		Address:       form.Address,
		Role:          form.Role,
		Status:        form.Status,
		Password:      form.Password,
		NhanvienId:    form.NhanvienId,
		NhanvienIdStr: form.NhanvienIdStr,
	}
	fmt.Print("UserService.Password ====", UserService.Password)
	err := UserService.UpdatUsers(c.Params("id"))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   "Register false",
		})
	}

	data := make(map[string]string)

	return c.Status(fiber.StatusOK).JSON(data)
}

// Update Password User
func PutPasswordUsers(c *fiber.Ctx) error {
	form := &LoginForm{}
	fmt.Print("===================form==================", form)
	if err := c.BodyParser(form); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	UserService := user_service.UserPassword{
		Username: form.Username,
		Password: form.Password,
	}

	err := UserService.UpdatUsersPassword(c.Params("id"))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   "Register false",
		})
	}

	data := make(map[string]string)

	return c.Status(fiber.StatusOK).JSON(data)
}

// Update Image User
func PutIamgeUsers(c *fiber.Ctx) error {
	form := &UserUpdate{}
	// form := &ItemsUpdate{}
	if err := c.BodyParser(form); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	UserService := user_service.User{
		Image: form.Image,
	}

	err := UserService.UpdatImageUsers(c.Params("id"))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   "Update image false",
		})
	}

	data := make(map[string]string)

	return c.Status(fiber.StatusOK).JSON(data)
}

func PutDetailUsers(c *fiber.Ctx) error {
	form := &UserDetailUpdate{}
	// form := &ItemsUpdate{}
	if err := c.BodyParser(form); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	valid, field := checkUserPut(c, c.Params("id"), *form)
	if !valid {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"code":    400,
			"errors":  nil,
			"message": fmt.Sprintf("%s exists.", field),
		})
	}
	if !valid {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"code":    400,
			"errors":  nil,
			"message": "Username exists",
		})
	}
	// var statusStr string = strconv.FormatBool(form.Status)
	UserService := user_service.User{
		Email:     form.Email,
		Firstname: form.Firstname,
		Lastname:  form.Lastname,
		Phone:     form.Phone,
		City:      form.City,
		Address:   form.Address,
		Role:      form.Role,
		Status:    form.Status,
	}

	err := UserService.UpdateDetailUsers(c.Params("id"))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   "Register false",
		})
	}

	data := make(map[string]string)

	return c.Status(fiber.StatusOK).JSON(data)
}

func PutUserAndHoaTieu(c *fiber.Ctx) error {
	form := &UserDetailUpdate{}
	// form := &ItemsUpdate{}
	if err := c.BodyParser(form); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	valid, field := checkUserPut(c, c.Params("id"), *form)
	if !valid {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"code":    400,
			"errors":  nil,
			"message": fmt.Sprintf("%s exists.", field),
		})
	}
	if !valid {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"code":    400,
			"errors":  nil,
			"message": "Username exists.",
		})
	}
	// var statusStr string = strconv.FormatBool(form.Status)
	UserService := user_service.User{
		Email:         form.Email,
		Firstname:     form.Firstname,
		Lastname:      form.Lastname,
		Phone:         form.Phone,
		City:          form.City,
		Address:       form.Address,
		Role:          form.Role,
		Status:        form.Status,
		NhanvienId:    form.NhanvienId,
		NhanvienIdStr: form.NhanvienIdStr,
	}

	err := UserService.UpdateUserAndHoaTieu(c.Params("id"))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   "Register false",
		})
	}

	data := UserService

	return c.Status(fiber.StatusOK).JSON(data)
}

// DeleteUser - Hàm xử lý API xóa người dùng theo ID
func DeleteUser(c *fiber.Ctx) error {
	// Lấy ID người dùng từ đường dẫn URL
	id := c.Params("id")

	// Tạo một đối tượng UserService để gọi hàm xóa người dùng
	UserService := user_service.User{}

	// Gọi hàm DeleteUser để xóa người dùng
	err := UserService.DeleteUser(id)
	if err != nil {
		// Nếu có lỗi khi xóa người dùng, trả về mã lỗi 400 với thông báo lỗi
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": fmt.Sprintf("Không thể xóa người dùng: %v", err),
		})
	}

	// Nếu xóa thành công, trả về mã 200 với thông báo thành công
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "Xóa người dùng thành công",
	})
}
