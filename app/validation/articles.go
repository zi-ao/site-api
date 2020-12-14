package validation

type ArticleStore struct {
	Title      string `form:"title" json:"title" binding:"required,min=2,max=30"`
	CategoryID uint   `form:"cate_id" json:"cate_id" binding:"required,min=1"`
	Content    string `form:"content" json:"content" binding:"required,min=10"`
	TagIDs     []uint `form:"tag_ids" json:"tag_ids" binding:"required,min=1,max=5"`
}
