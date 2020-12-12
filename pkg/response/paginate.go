package response

import (
	"github.com/gin-gonic/gin"
	"github.com/zi-ao/site-api/pkg/model"
	"gorm.io/gorm"
	"net/http"
)

type PageFunc func(db *gorm.DB) *gorm.DB

type paginate struct {
	Page     int
	Limit    int
	PageFunc PageFunc
}

// NewPaginate 实例化分页器
func NewPaginate(pageFunc PageFunc) *paginate {
	paginate := &paginate{
		Page:     1,
		Limit:    10,
		PageFunc: pageFunc,
	}
	return paginate
}

// Set 设置页数和每页个数
func (pager *paginate) Set(page, limit int) *paginate {
	pager.Page = page
	pager.Limit = limit
	return pager
}

// 使用分页器分页
func (pager *paginate) Paginate(context *gin.Context, list interface{}) {
	db := model.DB.Offset((pager.Page - 1) * pager.Limit).Limit(pager.Limit)
	if pager.PageFunc != nil {
		db = pager.PageFunc(db)
	}
	if err := model.Result(db.Find(list)); err != nil {
		FAIL(context, http.StatusInternalServerError, nil)
		return
	}
	var count int64
	db.Table(db.Statement.Table).Count(&count)
	SUCCESS(context, map[string]interface{}{
		"list": list,
		"paginate": map[string]interface{}{
			"page":  pager.Page,
			"limit": pager.Limit,
			"count": count,
		},
	})
}
