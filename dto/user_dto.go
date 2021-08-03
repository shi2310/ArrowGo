package dto

// UserDTO ...
type UserDTO struct {
	UserName string `binding:"required"`
	Pwd      string `binding:"required"`
	Photo    string
}
