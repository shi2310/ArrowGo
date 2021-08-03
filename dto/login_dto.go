package dto

// LoginDTO ...
type LoginDTO struct {
	UserName string `binding:"required"`
	Pwd      string `binding:"required"`
}
