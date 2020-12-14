package validation

type CategoryStore struct {
	CategoryUpdate
	Slug string `form:"slug" json:"slug" binding:"max=30"`
}

type CategoryUpdate struct {
	Name        string `form:"name" json:"name" binding:"required,min=2,max=20"`
	Description string `form:"description" json:"description" binding:"max=300"`
}
