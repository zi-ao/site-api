package auth

import (
	"errors"

	"github.com/zi-ao/site-api/app/models"
	"github.com/zi-ao/site-api/app/validation"
	"github.com/zi-ao/site-api/pkg/model"
	"github.com/zi-ao/site-api/pkg/utils"
	"golang.org/x/crypto/bcrypt"
)

// Login 用户登录
func Login(form *validation.Login) *models.User {
	var user = &models.User{}
	tx := model.DB.Where("username = ? OR email = ?", form.Username, form.Username).First(user)
	if tx.RowsAffected == 1 && CheckPassword(user.Password, form.Password) {
		return user
	}
	return nil
}

// Register 用户注册
func Register(form *validation.Register) (*models.User, error) {
	password, err := bcrypt.GenerateFromPassword([]byte(form.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	user := &models.User{
		Name:          form.Username,
		Username:      form.Username,
		Password:      string(password),
		Email:         form.Email,
		RememberToken: rememberToken(),
	}
	tx := model.DB.Create(user)
	if tx.RowsAffected == 1 && user.ID > 0 {
		return user, nil
	}
	return nil, errors.New("用户注册失败")
}

// CheckPassword 检测两次密码是否一致
func CheckPassword(hashPassword, password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password)) == nil
}

func rememberToken() string {
	return utils.RandString(13)
}
