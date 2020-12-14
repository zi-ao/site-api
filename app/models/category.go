package models

import "github.com/zi-ao/site-api/pkg/model"

// Category 分类模型
type Category struct {
	model.Basic
	Name        string `gorm:"unique;not null;comment:名称" json:"name"`
	Slug        string `gorm:"unique;not null;comment:缩写名" json:"slug"`
	Description string `gorm:"type:varchar(300);comment:描述" json:"description"`

	Articles []Article `json:"articles,omitempty"`
}
