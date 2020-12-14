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

// TagIndexEndpoint 标签列表
func TagIndexEndpoint(context *gin.Context) {
	page, _ := strconv.Atoi(context.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(context.DefaultQuery("limit", "10"))
	pager := response.NewPaginate(func(db *gorm.DB) *gorm.DB {
		return db.Order("created_at desc")
	})
	pager.Set(page, limit)
	pager.Paginate(context, &[]models.Tag{})
}

// TagShowEndpoint 标签详情
func TagShowEndpoint(context *gin.Context) {
	Tag := getTag(context)
	if Tag == nil {
		response.FAIL(context, http.StatusNotFound, nil)
		return
	}
	response.SUCCESS(context, Tag)
}

// TagStoreEndpoint 保存标签
func TagStoreEndpoint(context *gin.Context) {
	var form validation.TagStore
	if err := context.ShouldBind(&form); err != nil {
		fmt.Println(err.Error())
		response.FAIL(context, http.StatusBadRequest, err.Error())
		return
	}
	Tag := &models.Tag{
		Name: form.Name,
		Slug: form.Slug,
	}
	if model.Insert(Tag) != nil {
		response.FAIL(context, http.StatusInternalServerError, "标签创建失败")
		return
	}
	response.SUCCESS(context, nil)
}

// TagUpdateEndpoint 修改标签
func TagUpdateEndpoint(context *gin.Context) {
	var form validation.TagUpdate
	if err := context.ShouldBind(&form); err != nil {
		fmt.Println(err.Error())
		response.FAIL(context, http.StatusBadRequest, err.Error())
		return
	}
	Tag := getTag(context)
	if Tag == nil {
		response.FAIL(context, http.StatusNotFound, nil)
		return
	}
	values := map[string]interface{}{
		"name": form.Name,
	}
	if model.Updates(Tag, &values) != nil {
		response.FAIL(context, http.StatusInternalServerError, "标签修改失败")
		return
	}
	response.SUCCESS(context, nil)
}

// TagDeleteEndpoint 删除标签
func TagDeleteEndpoint(context *gin.Context) {
	Tag := getTag(context)
	if Tag == nil {
		response.FAIL(context, http.StatusNotFound, nil)
		return
	}
	if model.Delete(Tag) != nil {
		response.FAIL(context, http.StatusInternalServerError, nil)
		return
	}
	response.SUCCESS(context, nil)
}

func getTag(context *gin.Context) *models.Tag {
	id := context.Param("id")
	if id == "" {
		return nil
	}
	Tag := &models.Tag{}
	if model.First(Tag, id) != nil {
		return nil
	}
	return Tag
}
