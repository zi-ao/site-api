package validation

type ArticleStore struct {
	Title   string `form:"title" json:"title" binding:"required,min=2,max=30"`
	Content string `form:"content" json:"content" binding:"required,min=10"`
}
