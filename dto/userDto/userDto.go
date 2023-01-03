package userDto

import "login-vue/models"

type UserDto struct {
	Id       uint   `json:"id"`
	Username string `json:"username"`
	Tel      string `json:"tel"`
}

func NewUserDto(user *models.User) *UserDto {
	userdto := &UserDto{
		Id:       user.ID,
		Username: user.Username,
		Tel:      user.Tel,
	}
	return userdto
}
