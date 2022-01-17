package model

import (
	"crm/pkg/model"
	"crm/utils"
	"time"
)

type User struct {
	Id       uint64 `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Status   uint `json:"status"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

func (u *User) TableName() string {
	return "users"
}

// GetUserByUsernameAndPassword 根据账号和密码查询用户
func (u *User) GetUserByUsernameAndPassword(parameters map[string]string) (User, error) {
	var user User

	if err := model.Database.Where("username = ? AND deleted_at IS NULL", parameters["username"]).First(&user).Error; err != nil {
		return user, err
	}

	//验证密码是否正确
	if !utils.ComparePasswords(user.Password,[]byte(parameters["password"])) {
		user.Id = 0

		return user,nil
	}

	return user, nil
}

// GetUserStatusByEmail 根据email 查询是否存在用户
func (u *User) GetUserStatusByEmail(email string) bool {
	var user User

	model.Database.Select("id").Where("email = ? ", email).First(&user)

	if user.Id > 0 {
		return false
	}

	return true
}

// GetUserStatusByPhone 根据phone 查询是否存在用户
func (u *User) GetUserStatusByPhone(phone string) bool {
	var user User

	model.Database.Select("id").Where("phone = ?", phone).First(&user)

	if user.Id > 0 {
		return false
	}

	return true
}

// GetUserStatusByUsername 根据username 查询是否存在用户
func (u *User) GetUserStatusByUsername(username string) bool {
	var user User

	model.Database.Select("id").Where("username = ?", username).First(&user)

	if user.Id > 0 {
		return false
	}

	return true
}

// GetUserByPhone 根据手机号查询用户
func (u *User) GetUserByPhone(phone string) (User, error) {
	var user User

	if err := model.Database.Where("phone = ?", phone).First(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

// CreateUser 创建用户
func (u *User) CreateUser(parameters map[string]string) (User,error) {
	timeDate := time.Now()

	var user User = User{
		Name: parameters["name"],
		Username: parameters["username"],
		Password: utils.HashAndSalt(parameters["password"]),
		Email : parameters["email"],
		Phone : parameters["phone"],
		Status: 0,
		CreatedAt : &timeDate,
		UpdatedAt : &timeDate,
		DeletedAt: nil,
	}

	if err := model.Database.Create(&user).Error; err != nil {
		return user,err
	}

	return user, nil
}

// UpdateUserStatus 更改user状态
func (u *User) UpdateUserStatus(userId int) (User,error){
	var user User = User{
		Id: uint64(userId),
	}

	result := model.Database.Model(&user).Update("status",1)

	if result.Error != nil {
		return user,result.Error
	}

	return user,nil
}