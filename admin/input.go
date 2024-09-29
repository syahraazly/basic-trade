package admin

type RegisterInput struct {
	Name     string `json:"name" validate:"required" form:"name"`
	Email    string `json:"email" validate:"required,email" form:"email"`
	Password string `json:"password" validate:"required" form:"password"`
}

type LoginInput struct {
	Email    string `json:"email" validate:"required,email" form:"email"`
	Password string `json:"password" validate:"required" form:"password"`
}