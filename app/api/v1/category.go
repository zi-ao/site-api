package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/zi-ao/site-api/app/models"
	"github.com/zi-ao/site-api/app/validation"
	"github.com/zi-ao/site-api/pkg/model"
	"github.com/zi-ao/site-api/pkg/response"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

// CategoryIndexEndpoint 分类列表
func CategoryIndexEndpoint(context *gin.Context) {
	page, _ := strconv.Atoi(context.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(context.DefaultQuery("limit", "10"))
	pager := response.NewPaginate(func(db *gorm.DB) *gorm.DB {
		return db.Order("created_at desc")
	})
	pager.Set(page, limit)
	pager.Paginate(context, &[]models.Category{})
}

// CategoryShowEndpoint 分类详情
func CategoryShowEndpoint(context *gin.Context) {
	category := getCategory(context)
	if category == nil {
		response.FAIL(context, http.StatusNotFound, nil)
		return
	}
	response.SUCCESS(context, category)
}

// CategoryStoreEndpoint 保存分类
func CategoryStoreEndpoint(context *gin.Context) {
	var form validation.CategoryStore
	if err := context.ShouldBind(&form); err != nil {
		fmt.Println(err.Error())
		response.FAIL(context, http.StatusBadRequest, err.Error())
		return
	}
	category := &models.Category{
		Name:        form.Name,
		Slug:        form.Slug,
		Description: form.Description,
	}
	if model.Insert(category) != nil {
		response.FAIL(context, http.StatusInternalServerError, "分类创建失败")
		return
	}
	response.SUCCESS(context, nil)
}

// CategoryUpdateEndpoint 修改分类
func CategoryUpdateEndpoint(context *gin.Context) {
	var form validation.CategoryStore
	if err := context.ShouldBind(&form); err != nil {
		fmt.Println(err.Error())
		response.FAIL(context, http.StatusBadRequest, err.Error())
		return
	}
	category := getCategory(context)
	if category == nil {
		response.FAIL(context, http.StatusNotFound, nil)
		return
	}
	values := map[string]interface{}{
		"name":        form.Name,
		"description": form.Description,
	}
	if model.Updates(category, &values) != nil {
		response.FAIL(context, http.StatusInternalServerError, nil)
		return
	}
	response.SUCCESS(context, nil)
}

// CategoryDeleteEndpoint 删除分类
func CategoryDeleteEndpoint(context *gin.Context) {
	category := getCategory(context)
	if category == nil {
		response.FAIL(context, http.StatusNotFound, nil)
		return
	}
	if model.Delete(category) != nil {
		response.FAIL(context, http.StatusInternalServerError, nil)
		return
	}
	response.SUCCESS(context, nil)
}

func getCategory(context *gin.Context) *models.Category {
	id := context.Param("id")
	if id == "" {
		return nil
	}
	category := &models.Category{}
	if model.First(category, id) != nil {
		return nil
	}
	return category
}
