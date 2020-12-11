package validation

type Login struct {
	Username string `form:"username" json:"username" binding:"required,min=3,max=20"`
	Password string `form:"password" json:"password" binding:"required,min=6,max=32"`
}

type Register struct {
	Username string `form:"username" json:"username" binding:"required,username,unique=users"`
	Email    string `form:"email" json:"email" binding:"required,email,unique=users"`
	Password string `form:"password" json:"password" binding:"required,min=6,max=32"`
}

type UpdatePassword struct {
	OldPassword string `form:"password" json:"password" binding:"required,min=6,max=32"`
	NewPassword string `form:"new_password" json:"new_password" binding:"required,min=6,max=32,nefield=OldPassword"`
}
