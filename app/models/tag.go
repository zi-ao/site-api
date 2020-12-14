package models

import "github.com/zi-ao/site-api/pkg/model"

type Tag struct {
	model.Basic
	Name string `gorm:"unique;not null;comment:名称" json:"name"`
	Slug string `gorm:"unique;not null;comment:缩写名" json:"slug"`

	Articles []*Article `gorm:"many2many:article_tags;" json:"articles,omitempty"`
}
