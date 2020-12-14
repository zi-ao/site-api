package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/zi-ao/site-api/app/models"
	"github.com/zi-ao/site-api/app/validation"
	"github.com/zi-ao/site-api/pkg/auth"
	"github.com/zi-ao/site-api/pkg/model"
	"github.com/zi-ao/site-api/pkg/response"
	"github.com/zi-ao/site-api/pkg/utils"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

// ArticleIndexEndpoint 文章列表
func ArticleIndexEndpoint(context *gin.Context) {
	page, _ := strconv.Atoi(context.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(context.DefaultQuery("limit", "10"))
	pager := response.NewPaginate(func(db *gorm.DB) *gorm.DB {
		return db.Preload("Category").Preload("Tags").Omit("content").Order("created_at desc,view desc")
	})
	pager.Set(page, limit)
	pager.Paginate(context, &[]models.Article{})
}

// ArticleShowEndpoint 文章详情
func ArticleShowEndpoint(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id"))
	article := &models.Article{}
	if err != nil || model.First(article, id, "Category", "Tags") != nil {
		response.FAIL(context, http.StatusNotFound, nil)
		return
	}
	response.SUCCESS(context, article)
}

// ArticleStoreEndpoint 保存文章
func ArticleStoreEndpoint(context *gin.Context) {
	var form validation.ArticleStore
	if err := context.ShouldBind(&form); err != nil {
		response.FAIL(context, http.StatusBadRequest, err.Error())
		return
	}
	article := createArticle(&form)
	article.OwnerID = auth.ID(context)
	err := model.Insert(article)
	if err != nil {
		response.FAIL(context, http.StatusInternalServerError, "文章保存失败")
		return
	}
	response.SUCCESS(context, article)
}

// ArticleUpdateEndpoint 修改文章
func ArticleUpdateEndpoint(context *gin.Context) {
	var form validation.ArticleStore
	if err := context.ShouldBind(&form); err != nil {
		response.FAIL(context, http.StatusBadRequest, err.Error())
		return
	}
	article := getArticle(context)
	values := createArticle(&form)
	if model.Updates(article, values) != nil {
		response.FAIL(context, http.StatusInternalServerError, nil)
		return
	}
	response.SUCCESS(context, nil)
}

// ArticleDeleteEndpoint 删除文章
func ArticleDeleteEndpoint(context *gin.Context) {
	article := getArticle(context)
	if article == nil {
		response.FAIL(context, http.StatusNotFound, nil)
		return
	}
	if model.Delete(article) != nil {
		response.FAIL(context, http.StatusInternalServerError, nil)
		return
	}
	response.SUCCESS(context, nil)
}

func createArticle(form *validation.ArticleStore) *models.Article {
	tags := make([]*models.Tag, len(form.TagIDs))
	for index, id := range form.TagIDs {
		tag := &models.Tag{}
		tag.ID = id
		tags[index] = tag
	}
	return &models.Article{
		Title:      form.Title,
		CategoryID: form.CategoryID,
		Abstract:   form.Content[:utils.Min(200, len(form.Content))],
		Content:    form.Content,
		Tags:       tags,
	}
}

func getArticle(context *gin.Context) *models.Article {
	id := context.Param("id")
	if id == "" {
		return nil
	}
	article := &models.Article{}
	ownerID := auth.ID(context)
	if model.Result(model.Where("owner_id = ?", ownerID).First(article, id)) != nil {
		return nil
	}
	return article
}
