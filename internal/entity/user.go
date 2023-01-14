package entity

type User struct {
	Id       int    `json:"-"`
	Username string `json:"username" validate:"required"`
	Email    string `json:"email"    validate:"required,email"`
	Password string `json:"password" validate:"required"`
}
