package dto

import(
	"mygin.com/mygin/model"
)

// UserDto 用户的数据传送结构体
type UserDto struct{
	ID uint `json:"id"`
	Username string `json:"username"`
	Email string `json:"email"`
	Mobile string `json:"mobile"`
	Role string `json:"role"`
	City string `json:"city"`
	Status bool `json:"status"`
}

// ToUserDto 将 model.User 转换为 UserDto
func ToUserDto(user model.User) UserDto  {
	return UserDto{
		ID: user.ID,
		Username: user.Username,
		Email: user.Email,
		Mobile: user.Mobile,
		Role: user.Role,
		City: user.City,
		Status: user.Status,
	}
}