package rule

import (
	"github.com/go-playground/validator/v10"
	"github.com/zi-ao/site-api/pkg/model"
	"strings"
)

// UniqueRule 数据库唯一规则
func UniqueRule(field validator.FieldLevel) bool {
	value := field.Field().String()
	params := strings.Split(field.Param(), ",")
	if len(params) == 1 {
		params = append(params, field.FieldName())
	}
	var count int64
	model.DB.Table(params[0]).Where(params[1]+" = ?", value).Count(&count)
	return count == 0
}
