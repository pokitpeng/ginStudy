package dto

import "ginStudy/model"

// UserDto 返回的用户数据对象
type UserDto struct {
	Name      string `json:"name"`
	Telephone string `json:"telephone"`
}

// ToUserDto 返回的用户数据对象工厂函数
func ToUserDto(user model.User) UserDto {
	return UserDto{
		Name:      user.Name,
		Telephone: user.Telephone,
	}
}
