package users

import (
	"errors"
	"github.com/zi-ao/site-api/app/validation"
	"github.com/zi-ao/site-api/pkg/model"
	"github.com/zi-ao/site-api/pkg/utils"
	"golang.org/x/crypto/bcrypt"
)

// Login 用户登录
func Login(form *validation.Login) *User {
	var user = &User{}
	tx := model.DB.Where("username = ? OR email = ?", form.Username, form.Username).First(user)
	if tx.RowsAffected == 1 && CheckPassword(user.Password, form.Password) {
		return user
	}
	return nil
}

// Register 用户注册
func Register(form *validation.Register) (*User, error) {
	password, err := bcryptPassword(form.Password)
	if err != nil {
		return nil, err
	}
	user := &User{
		Name:          form.Username,
		Username:      form.Username,
		Password:      password,
		Email:         form.Email,
		RememberToken: rememberToken(),
	}
	tx := model.DB.Debug().Create(user)
	if tx.RowsAffected == 1 && user.ID > 0 {
		return user, nil
	}
	return nil, errors.New("用户注册失败")
}

// UpdatePassword 更改密码
func UpdatePassword(user *User, newPassword string) error {
	password, err := bcryptPassword(newPassword)
	if err == nil {
		tx := model.DB.Model(user).Where("id = ?", user.ID).Update("password", password)
		if tx.RowsAffected == 1 {
			return nil
		}
	}
	return errors.New("密码修改失败")
}

// Find 根据 ID 查找用户
func Find(id uint) *User {
	var user = &User{}
	tx := model.DB.Where("id = ?", id).First(user)
	if tx.RowsAffected == 1 && user.ID != 0 {
		return user
	}
	return nil
}

// CheckPassword 检测两次密码是否一致
func CheckPassword(hashPassword, password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password)) == nil
}

func bcryptPassword(password string) (string, error) {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashPassword), nil
}

func rememberToken() string {
	return utils.RandString(13)
}
