package controller

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"mygin.com/mygin/common"
	"mygin.com/mygin/dto"
	"mygin.com/mygin/model"
	"mygin.com/mygin/response"
)

// GetUsers 根据 URL 参数查询用户信息
func GetUsers(context *gin.Context) {
	// 验证用户是否获得权限
	if !IsAuth(context) {
		return
	}

	pageNum, err := strconv.Atoi(context.Query("pagenum"))
	if pageNum == 0 || err != nil {
		response.ProcessFailed(context, nil, "请求参数错误")
		return
	}
	pageSize, err := strconv.Atoi(context.Query("pagesize"))
	if pageSize == 0 || err != nil {
		response.ProcessFailed(context, nil, "请求参数错误")
		return
	}
	query := context.Query("query")
	status := context.Query("status")
	category := context.Query("category")

	if category != "all" && category != "true" && category != "false" && category != "city" && category != "admin" && status != "deleted" && status != "" {
		response.ProcessFailed(context, nil, "请求参数错误")
	}

	db := common.GetDatabase()
	var users []model.User
	var total int

	if status == "" {
		if category == "all" {
			db = db.Where("username LIKE ? AND role LIKE '%用户'", "%"+query+"%")
			db.Find(&users)
			total = len(users)
			db = db.Limit(pageSize).Offset((pageNum - 1) * pageSize)
			db.Where("username LIKE ? AND role LIKE '%用户'", "%"+query+"%").Find(&users)
		}

		if category == "city" {
			db = db.Where("city LIKE ? AND role LIKE '%用户'", "%"+query+"%")
			db.Find(&users)
			total = len(users)
			db = db.Limit(pageSize).Offset((pageNum - 1) * pageSize)
			db.Where("city LIKE ? AND role LIKE '%用户'", "%"+query+"%").Find(&users)
		}

		if category == "true" || category == "false" {
			db = db.Where("username LIKE ? AND role LIKE '%用户' AND status = "+category, "%"+query+"%")
			db.Find(&users)
			total = len(users)
			db = db.Limit(pageSize).Offset((pageNum - 1) * pageSize)
			db.Table("users").Where("username LIKE ? AND role LIKE '%用户' AND status = "+category, "%"+query+"%").Find(&users)
		}
		if category == "admin" {
			db = db.Where("username LIKE ? AND role LIKE '%管理员%'", "%"+query+"%")
			db.Find(&users)
			total = len(users)
			db = db.Limit(pageSize).Offset((pageNum - 1) * pageSize)
			db.Where("username LIKE ? AND role LIKE '%管理员%' ", "%"+query+"%").Find(&users)
		}

	} else if status == "deleted" {
		db = db.Unscoped().Where("username LIKE ? AND role LIKE '%用户' AND deleted_at is not null", "%"+query+"%")
		db.Find(&users)
		total = len(users)
		db = db.Limit(pageSize).Offset((pageNum - 1) * pageSize)
		db.Unscoped().Where("deleted_at is not null").Find(&users)
	}

	var data = make([]dto.UserDto, len(users))
	for i := 0; i < len(users); i++ {
		data[i] = dto.ToUserDto(users[i])
	}

	response.GetSuccessed(context, gin.H{"total": total, "users": data}, "用户信息获取成功")

}

// GetUserByUserID 根据用户 ID 查询用户信息
func GetUserByUserID(context *gin.Context) {
	// 验证用户是否获得权限
	if !IsAuth(context) {
		return
	}

	// 获取 URL 中的 userID 参数
	userID := context.Param("userID")
	if userID == "" || userID == "0" {
		response.ProcessFailed(context, nil, "请求参数错误")
		return
	}

	db := common.GetDatabase()
	var user model.User

	db.Where("id = ?", userID).First(&user)
	if user.ID == 0 {
		response.ProcessFailed(context, nil, "用户信息不存在")
		return
	}

	response.GetSuccessed(context, dto.ToUserDto(user), "用户信息获取成功")
}

// GetOrdersByUserID 根据用户 ID 查询订单信息
func GetOrdersByUserID(context *gin.Context) {
	// 验证用户是否获得权限
	if !IsAuth(context) {
		return
	}

	userID := context.Param("userID")
	if userID == "" || userID == "0" {
		response.ProcessFailed(context, nil, "请求参数错误")
		return
	}

	db := common.GetDatabase()
	var orders []model.Order
	db.Table("orders").Where("carrier_user_id = ?", userID).Find(&orders)

	var data = make([]dto.OrderDto, len(orders))

	for i := 0; i < len(orders); i++ {
		if orders[i].ID == 0 {
			response.ProcessFailed(context, nil, "订单信息不存在")
			return
		}
		data[i] = dto.ToOrderDto(orders[i])
	}

	response.GetSuccessed(context, data, "获取订单信息成功")
}

// GetWipsByUserID 根据用户 ID 获取用户工单信息
func GetWipsByUserID(context *gin.Context) {
	// 验证用户是否获得权限
	if !IsAuth(context) {
		return
	}

	userID := context.Param("userID")
	if userID == "" || userID == "0" {
		response.ProcessFailed(context, nil, "请求参数错误")
		return
	}

	db := common.GetDatabase()
	var wips []model.Wips
	db.Table("wips").Where("user_id = ?", userID).Find(&wips)

	var data = make([]dto.WipDto, len(wips))

	for i := 0; i < len(wips); i++ {
		if wips[i].ID == 0 {
			response.ProcessFailed(context, nil, "工单信息不存在")
			return
		}
		data[i] = dto.ToWipDto(wips[i])
	}

	response.GetSuccessed(context, data, "获取工单信息成功")
}

// GetDayByUserID 根据用户 ID 查询用户的工作天数
func GetDayByUserID(context *gin.Context) {
	// 验证用户是否获得权限
	if !IsAuth(context) {
		return
	}
	userID := context.Param("userID")
	if userID == "" || userID == "0" {
		response.ProcessFailed(context, nil, "请求参数错误")
		return
	}

	db := common.GetDatabase()
	var user model.User

	db.Where("id = ?", userID).Find(&user)
	if user.ID == 0 {
		response.ProcessFailed(context, nil, "用户信息不存在")
		return
	}

	createAt := user.CreatedAt.Unix()
	currentData := time.Now().Unix()

	response.GetSuccessed(context, gin.H{"day": (currentData - createAt) / 86400}, "获取天数信息成功")

}

// DeleteUserByUserID 根据用户 ID 删除用户信息
func DeleteUserByUserID(context *gin.Context) {
	// 验证用户是否获得权限
	if !IsAuth(context) {
		return
	}

	userID := context.Param("userID")
	if userID == "" || userID == "0" {
		response.ProcessFailed(context, nil, "请求参数错误")
		return
	}

	db := common.GetDatabase()
	var user model.User

	db.Where("id = ?", userID).First(&user)
	if user.ID == 0 {
		response.ProcessFailed(context, nil, "用户信息不存在")
		return
	}

	if user.Status == true {
		response.ProcessFailed(context, nil, "用户正在进行作业")
		return
	}

	db.Delete(&user)

	response.DeleteSuccessed(context, nil, "删除用户信息成功")
}

// UpdateUserInfoByUserID 根据用户 ID 更新用户信息
func UpdateUserInfoByUserID(context *gin.Context) {
	// 验证用户是否获得权限
	if !IsAuth(context) {
		return
	}

	userID := context.Param("userID")
	if userID == "" || userID == "0" {
		response.ProcessFailed(context, nil, "请求参数错误")
		return
	}

	db := common.GetDatabase()
	var user model.User
	var u struct {
		Mobile string
		Email  string
		City   string
	}
	context.BindJSON(&u)

	db.Where("ID = ?", userID).First(&user)
	if user.ID == 0 {
		response.ProcessFailed(context, nil, "用户信息不存在")
		return
	}

	var unscopedUser model.User
	db.Table("users").Unscoped().Where("(mobile = ? OR email = ?)AND deleted_at is not null", u.Mobile, u.Email).First(&unscopedUser)
	if unscopedUser.ID != 0 {
		response.ProcessFailed(context, nil, "用户手机号码或者邮箱地址已经注销")
		return
	}

	if isMobileExist(db, u.Mobile) && u.Mobile != user.Mobile {
		response.ProcessFailed(context, nil, "手机号码已经存在")
		return
	}

	if isEmailExist(db, u.Email) && u.Email != user.Email {
		response.ProcessFailed(context, nil, "邮箱地址已经存在")
		return
	}

	if len(u.City) < 3 {
		response.ProcessFailed(context, nil, "请求参数错误")
		return
	}

	user.Mobile = u.Mobile
	user.Email = u.Email
	user.City = u.City

	db.Save(&user)
	response.PutSuccessed(context, dto.ToUserDto(user), "更新用户信息成功")
}

// UpdateUserEmailByUserID 根据用户 ID 更新用户信息
func UpdateUserEmailByUserID(context *gin.Context) {
	// 验证用户是否获得权限
	if !IsAuth(context) {
		return
	}

	userID := context.Param("userID")
	if userID == "" || userID == "0" {
		response.ProcessFailed(context, nil, "请求参数错误")
		return
	}

	db := common.GetDatabase()
	var user model.User
	var u struct{ Email string }
	context.BindJSON(&u)

	db.Where("ID = ?", userID).First(&user)
	if user.ID == 0 {
		response.ProcessFailed(context, nil, "用户信息不存在")
		return
	}

	var unscopedUser model.User
	db.Table("users").Unscoped().Where("email = ? AND deleted_at is not null", u.Email).First(&unscopedUser)
	if unscopedUser.ID != 0 {
		response.ProcessFailed(context, nil, "用户邮箱地址已经注销")
		return
	}

	if isEmailExist(db, u.Email) && u.Email != user.Email {
		response.ProcessFailed(context, nil, "邮箱地址已经存在")
		return
	}
	user.Email = u.Email

	db.Save(&user)
	response.PutSuccessed(context, dto.ToUserDto(user), "更新用户信息成功")
}

// UpdateUserMobileByUserID 根据用户 ID 更新用户信息
func UpdateUserMobileByUserID(context *gin.Context) {
	// 验证用户是否获得权限
	if !IsAuth(context) {
		return
	}

	userID := context.Param("userID")
	if userID == "" || userID == "0" {
		response.ProcessFailed(context, nil, "请求参数错误")
		return
	}

	db := common.GetDatabase()
	var user model.User
	var u struct{ Mobile string }
	context.BindJSON(&u)

	db.Where("ID = ?", userID).First(&user)
	if user.ID == 0 {
		response.ProcessFailed(context, nil, "用户信息不存在")
		return
	}

	var unscopedUser model.User
	db.Table("users").Unscoped().Where("mobile = ? AND deleted_at is not null", u.Mobile).First(&unscopedUser)
	if unscopedUser.ID != 0 {
		response.ProcessFailed(context, nil, "用户手机号码已经注销")
		return
	}

	if isMobileExist(db, u.Mobile) && u.Mobile != user.Mobile {
		response.ProcessFailed(context, nil, "手机号码已经存在")
		return
	}

	user.Mobile = u.Mobile

	db.Save(&user)
	response.PutSuccessed(context, dto.ToUserDto(user), "更新用户信息成功")
}

// RecoverUsersByUserID 根据用户 ID 恢复用户信息
func RecoverUsersByUserID(context *gin.Context) {
	// 验证用户是否获得权限
	if !IsAuth(context) {
		return
	}

	userID := context.Param("userID")
	if userID == "" || userID == "0" {
		response.ProcessFailed(context, nil, "请求参数错误")
		return
	}

	var user model.User
	db := common.GetDatabase()
	db.Unscoped().Where("id = ? AND deleted_at is not null", userID).First(&user)
	if user.ID == 0 {
		response.ProcessFailed(context, nil, "用户信息不存在")
		return
	}

	db.Unscoped().Model(&user).Update("deleted_at", nil)

	response.PatchSuccessed(context, dto.ToUserDto(user), "恢复用户信息成功")
}

// UpdateUserCityByUserID 根据用户 ID 更新用户信息
func UpdateUserCityByUserID(context *gin.Context) {
	// 验证用户是否获得权限
	if !IsAuth(context) {
		return
	}

	userID := context.Param("userID")
	if userID == "" || userID == "0" {
		response.ProcessFailed(context, nil, "请求参数错误")
		return
	}

	db := common.GetDatabase()
	var user model.User
	var u struct{ City string }
	context.BindJSON(&u)

	db.Where("ID = ?", userID).First(&user)
	if user.ID == 0 {
		response.ProcessFailed(context, nil, "用户信息不存在")
		return
	}

	user.City = u.City

	db.Save(&user)
	response.PutSuccessed(context, dto.ToUserDto(user), "更新用户信息成功")
}

// UpdateUserPasswordByUserID 根据用户 ID 更新用户密码
func UpdateUserPasswordByUserID(context *gin.Context) {
	// 验证用户是否获得权限
	if !IsAuth(context) {
		return
	}

	userID := context.Param("userID")
	if userID == "" || userID == "0" {
		response.ProcessFailed(context, nil, "请求参数错误")
		return
	}

	var pw struct {
		OldPassword    string
		NewPassword    string
		verifyPassword string
	}

	context.BindJSON(&pw)

	db := common.GetDatabase()
	var user model.User

	db.Where("ID = ?", userID).First(&user)
	if user.ID == 0 {
		response.ProcessFailed(context, nil, "用户信息不存在")
		return
	}

	// 判断密码是否正确
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(pw.OldPassword)); err != nil {
		response.ProcessFailed(context, nil, "用户密码错误")
		return
	}

	hasedPassword, err := bcrypt.GenerateFromPassword([]byte(pw.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		response.ProcessFailed(context, nil, "用户密码加密失败")
	}

	user.Password = string(hasedPassword)

	db.Save(&user)
	response.PutSuccessed(context, nil, "更新用户密码成功")
}

// UpdateUserStatusByUserID 根据用户 ID 更新用户状态
func UpdateUserStatusByUserID(context *gin.Context) {
	// 验证用户是否获得权限
	if !IsAuth(context) {
		return
	}

	db := common.GetDatabase()
	var user model.User

	userID := context.Param("userID")
	if userID == "" || userID == "0" {
		response.ProcessFailed(context, nil, "请求参数错误")
		return
	}

	var u struct{ Status bool }
	context.BindJSON(&u)

	db.Where("id = ?", userID).First(&user)
	if user.ID == 0 {
		response.ProcessFailed(context, nil, "用户信息不存在")
		return
	}

	db.Model(&user).Update("status", u.Status)
	// 更新该用户下所有车辆状态
	// db.Table("cars").Where("owner_id = ?", userID).Update("status", u.Status)

	response.PatchSuccessed(context, dto.ToUserDto(user), "修改用户信息成功")
}

// UpdateUserRoleByUserID 根据用户 ID 修改用户角色
func UpdateUserRoleByUserID(context *gin.Context) {
	// 验证用户是否获得权限
	if !IsAuth(context) {
		return
	}

	userID := context.Param("userID")
	if userID == "" || userID == "0" {
		response.ProcessFailed(context, nil, "请求参数错误")
		return
	}

	var u struct{ Role string }
	context.BindJSON(&u)
	if u.Role == "" {
		response.ProcessFailed(context, nil, "请求参数错误")
		return
	}

	db := common.GetDatabase()
	var user model.User
	db.Where("ID = ?", userID).First(&user)
	if user.ID == 0 {
		response.ProcessFailed(context, nil, "用户信息不存在")
		return
	}
	db.Model(&user).Update("role", u.Role)

	response.PatchSuccessed(context, dto.ToUserDto(user), "用户信息更新成功")
}
