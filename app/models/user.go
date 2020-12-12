package models

import (
	"errors"
	"github.com/zi-ao/site-api/pkg/model"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type User struct {
	model.Basic
	Name     string `gorm:"not null;comment:名称" json:"name"`
	Username string `gorm:"unique;not null;comment:用户名" json:"username"`
	Password string `gorm:"not null;comment:密码" json:"-"`
	Email    string `gorm:"unique;comment:邮箱" json:"email"`
	Phone    string `gorm:"unique;comment:手机" json:"phone"`
	RealName string `gorm:"comment:真实姓名" json:"real_name"`
	Avatar   string `gorm:"comment:头像" json:"avatar"`

	IsAdmin bool `gorm:"default:false;comment:是管理员" json:"is_admin"`

	ActivatedAt    *time.Time `gorm:"comment:激活时间" json:"activated_at"`
	LastActivateAt *time.Time `gorm:"comment:最后活动时间" json:"last_activate_at"`
	RememberToken  string     `gorm:"type:char(13);comment:登录随机令牌" json:"remember_token"`

	Articles []Article `gorm:"foreignKey:OwnerID"`
}

// UpdatePassword 更改密码
func UpdatePassword(user *User, newPassword string) error {
	password, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err == nil && model.Update(user, "password", string(password)) == nil {
		return nil
	}
	return errors.New("密码修改失败")
}
