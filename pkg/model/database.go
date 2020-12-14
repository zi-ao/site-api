package model

import (
	"gorm.io/gorm"
)

func Result(result *gorm.DB) error {
	if result.RowsAffected > 0 && result.Error == nil {
		return nil
	}
	return result.Error
}

// Insert 插入一行数据
func Insert(model interface{}) error {
	result := DB.Create(model)
	return Result(result)
}

// First 根据 ID 查找一条数据
func First(model interface{}, id interface{}, preload ...string) error {
	db := DB
	for _, query := range preload {
		db = db.Preload(query)
	}
	result := db.First(model, id)
	return Result(result)
}

// Update 更新但个字段
func Update(model interface{}, fieldName string, value interface{}) error {
	result := DB.Model(model).Update(fieldName, value)
	return Result(result)
}

// Updates 更新多个字段
func Updates(model interface{}, values interface{}) error {
	result := DB.Model(model).Omit("id", "created_at", "updated_at", "deleted_at").Updates(values)
	return Result(result)
}

// Delete 删除
func Delete(model interface{}, conds ...interface{}) error {
	result := DB.Delete(model, conds...)
	return Result(result)
}

// Where 数据条件
func Where(query interface{}, args ...interface{}) *gorm.DB {
	return DB.Where(query, args...)
}

// Select 搜索字段
func Select(query interface{}, args ...interface{}) *gorm.DB {
	return DB.Select(query, args...)
}

// Omit 搜索字段
func Omit(columns ...string) *gorm.DB {
	return DB.Omit(columns...)
}

// Paginate 分页器
func Paginate(page, limit int) *gorm.DB {
	return DB.Offset((page - 1) * limit).Limit(limit)
}
