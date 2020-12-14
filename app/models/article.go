package models

import (
	"github.com/zi-ao/site-api/pkg/model"
)

type Article struct {
	model.Basic
	Title      string `gorm:"not null;index;comment:标题" json:"title"`
	OwnerID    uint   `gorm:"comment:所有者 ID" json:"owner_id"`
	CategoryID uint   `gorm:"comment:分类 ID" json:"category_id"`
	Abstract   string `gorm:"type:varchar(500);not null;comment:内容摘要" json:"abstract"`
	Content    string `gorm:"type:mediumtext;not null;comment:内容" json:"content"`
	View       uint   `gorm:"default:0;comment:被浏览次数" json:"view"`

	Owner    *User     `gorm:"foreignKey:ID;references:OwnerID" json:"owner,omitempty"`
	Category *Category `gorm:"foreignKey:CategoryID" json:"category,omitempty"`
	Tags     []*Tag    `gorm:"many2many:article_tags;" json:"tags,omitempty"`
}
