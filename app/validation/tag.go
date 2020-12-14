package validation

type TagStore struct {
	TagUpdate
	Slug string `form:"slug" json:"slug" binding:"max=30"`
}

type TagUpdate struct {
	Name string `form:"name" json:"name" binding:"required,min=2,max=20"`
}
