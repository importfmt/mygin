package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"mygin.com/mygin/model"
	"mygin.com/mygin/response"
)

// isMobileExist 判断手机号码是否存在
func isMobileExist(db *gorm.DB, mobile string) bool {
	var user model.User
	db.Where("mobile = ?", mobile).First(&user)


	if user.ID != 0 {
		return true
	}
	return false

}

// isEmailExist 判断邮箱是否存在
func isEmailExist(db *gorm.DB, email string) bool {
	var user model.User

	db.Where("email = ?", email).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}

// isIDExist 判断用户是否存在
func isIDExist(db *gorm.DB, id string) bool {
	var user model.User

	db.Where("ID = ?", id).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}

// isOwnerExist 判断车辆的拥有者是否存在
// func isOwnerExist(db *gorm.DB, ownerid string) bool {
// 	var user model.User
// 	db.Table("users").Where("id = ?", ownerid).First(&user)
// 	if user.ID != 0 {
// 		return true
// 	}
// 	return false
// }

// isLicenseExist 判断车辆车牌号码是否存在
func isLicenseExist(db *gorm.DB, license string) bool {
	var car model.Car
	db.Where("license = ?", license).First(&car)
	if car.ID != 0 {
		return true
	}
	return false

}

// IsAuth 判断用户是否获得权限
func IsAuth(context *gin.Context) bool {
	requestUser, _ := context.Get("user")
	if requestUser == nil {
		response.ProcessFailed(context, nil, "用户权限不足")
		return false
	}
	return true
}
