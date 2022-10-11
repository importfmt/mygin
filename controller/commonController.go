package controller

import (
	"log"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"mygin.com/mygin/common"
	"mygin.com/mygin/dto"
	"mygin.com/mygin/model"
	"mygin.com/mygin/response"
	"mygin.com/mygin/util"
)

// UserRegister 添加新的用户信息
func UserRegister(context *gin.Context) {
	// 获取请求信息
	var user model.User
	context.BindJSON(&user)

	// 数据验证
	if !util.CheckMobile(user.Mobile) {
		response.ProcessFailed(context, nil, "手机号码不符合规范")
		return
	}

	if !util.CheckEmail(user.Email) {
		response.ProcessFailed(context, nil, "邮箱地址不符合规范")
		return
	}

	if len(user.Password) < 6 {
		response.ProcessFailed(context, nil, "用户密码不能少于6位")
		return
	}

	if len(user.Username) == 0 {
		user.Username = util.RandomString(10)
	}

	if user.Role == "" {
		user.Role = "普通用户"
	}

	if user.City == "" || len(user.City) < 3 {
		user.City = "北京市"
	}

	user.Status = false

	db := common.GetDatabase()

	// 判断手机号是否存在
	if isMobileExist(db, user.Mobile) {
		response.ProcessFailed(context, nil, "手机号码已经存在")
		return
	}

	if isEmailExist(db, user.Email) {
		response.ProcessFailed(context, nil, "邮箱地址已经存在")
		return
	}

	// 寻找软删除的记录中是否有符合邮箱和电话号码的记录
	var unscopedUser model.User
	db.Unscoped().Where("mobile = ?", user.Mobile).First(&unscopedUser)
	if unscopedUser.ID != 0 {
		response.ProcessFailed(context, nil, "电话号码已经注销")
		return
	}

	db.Unscoped().Where("email = ?", user.Email).First(&unscopedUser)
	if unscopedUser.ID != 0 {
		response.ProcessFailed(context, nil, "邮箱地址已经注销")
		return
	}

	// 密码加密
	hasedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		response.ProcessFailed(context, nil, "用户密码加密失败")
		return
	}
	user.Password = string(hasedPassword)

	// 数据库创建用户
	db.Create(&user)

	//返回响应内容
	response.PostSuccessed(context, dto.ToUserDto(user), "用户信息注册成功")
}

// UserLogin 用户登录
func UserLogin(context *gin.Context) {

	var user model.User

	// 获取请求信息
	context.BindJSON(&user)
	// 保存请求中的密码
	password := user.Password

	// 数据验证
	if !util.CheckMobile(user.Mobile) {
		response.ProcessFailed(context, nil, "手机号码不符合规范")
		return
	}

	if len(user.Password) < 6 {
		response.ProcessFailed(context, nil, "用户密码不能少于6位")
		return
	}
	db := common.GetDatabase()

	// 这里需要获取到符合条件的记录，并映射到user
	db.Where("mobile = ?", user.Mobile).First(&user)
	if user.ID == 0 {
		response.ProcessFailed(context, nil, "用户信息不存在")
		return
	}

	// 判断密码是否正确
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		response.ProcessFailed(context, nil, "用户密码错误")
		return
	}

	// 发放 Token
	token, err := common.ReleaseToken(user)
	if err != nil {
		response.ServerFailed(context, nil, "系统异常")
		log.Printf("token generate error: %v", err)
		return
	}

	// 返回结果
	response.PostSuccessed(context, gin.H{"token": token, "user": dto.ToUserDto(user)}, "用户登录成功")
}

// GetAllMenu 获取所有菜单列表
func GetAllMenu(context *gin.Context) {
	// 验证用户是否获得权限
	if !IsAuth(context) {
		return
	}


	array := make([]dto.MenuDto, 5)
	// 用户管理菜单
	var userSubSlice []dto.MenuDto = make([]dto.MenuDto, 3)
	userSubSlice[0] = dto.MenuDto{ID: 11, Name: "用户列表", Children: nil, Path: "users"}
	userSubSlice[1] = dto.MenuDto{ID: 12, Name: "恢复用户", Children: nil, Path: "recoverusers"}
	userSubSlice[2] = dto.MenuDto{ID: 13, Name: "用户查询", Children: nil, Path: "queryusers"}
	userMenu := dto.MenuDto{ID: 1, Name: "用户管理", Children: userSubSlice, Path: "users"}

	// 车辆管理菜单
	var carSubSlice []dto.MenuDto = make([]dto.MenuDto, 3)
	carSubSlice[0] = dto.MenuDto{ID: 51, Name: "车辆列表", Children: nil, Path: "cars"}
	carSubSlice[1] = dto.MenuDto{ID: 52, Name: "恢复车辆", Children: nil, Path: "recovercars"}
	carSubSlice[2] = dto.MenuDto{ID: 53, Name: "车辆查询", Children: nil, Path: "querycars"}
	carMenu := dto.MenuDto{ID: 5, Name: "车辆管理", Children: carSubSlice, Path: "cars"}

	// 系统管理菜单
	var systemSubSlice []dto.MenuDto = make([]dto.MenuDto, 3)
	systemSubSlice[0] = dto.MenuDto{ID: 21, Name: "角色列表", Children: nil, Path: "roles"}
	systemSubSlice[1] = dto.MenuDto{ID: 22, Name: "工单列表", Children: nil, Path: "wips"}
	systemSubSlice[2] = dto.MenuDto{ID: 23, Name: "管理员列表", Children: nil, Path: "admins"}
	systemMenu := dto.MenuDto{ID: 2, Name: "系统管理", Children: systemSubSlice, Path: "system"}

	// 商品管理菜单
	var GoodsSubSlice []dto.MenuDto = make([]dto.MenuDto, 3)
	GoodsSubSlice[0] = dto.MenuDto{ID: 31, Name: "商品列表", Children: nil, Path: "goods"}
	GoodsSubSlice[1] = dto.MenuDto{ID: 32, Name: "恢复商品", Children: nil, Path: "recovergoods"}
	GoodsSubSlice[2] = dto.MenuDto{ID: 33, Name: "商品查询", Children: nil, Path: "querygoods"}
	goodsMenu := dto.MenuDto{ID: 3, Name: "商品管理", Children: GoodsSubSlice, Path: "goods"}

	// 订单管理菜单
	var orderSubSlice []dto.MenuDto = make([]dto.MenuDto, 3)
	orderSubSlice[0] = dto.MenuDto{ID: 41, Name: "订单列表", Children: nil, Path: "orders"}
	orderSubSlice[1] = dto.MenuDto{ID: 42, Name: "恢复订单", Children: nil, Path: "recoverorders"}
	orderSubSlice[2] = dto.MenuDto{ID: 43, Name: "订单查询", Children: nil, Path: "queryorders"}
	orderMenu := dto.MenuDto{ID: 4, Name: "订单管理", Children: orderSubSlice, Path: "orders"}

	array[0] = userMenu
	array[1] = carMenu
	array[2] = systemMenu
	array[3] = goodsMenu
	array[4] = orderMenu

	response.GetSuccessed(context, array, "获取菜单信息成功")
}

// GetUserMenu 获取用户菜单列表
func GetUserMenu(context *gin.Context) {
	// 验证用户是否获得权限
	if !IsAuth(context) {
		return
	}

	array := make([]dto.MenuDto, 4)
	array[0] = dto.MenuDto{ID: 1, Name: "主页", Children: nil, Path: "welcome"}
	array[1] = dto.MenuDto{ID: 2, Name: "个人信息", Children: nil, Path: "personinformation"}
	array[2] = dto.MenuDto{ID: 3, Name: "工单提交", Children: nil, Path: "submitwips"}
	array[3] = dto.MenuDto{ID: 4, Name: "工作记录", Children: nil, Path: "workrecord"}

	response.GetSuccessed(context, array, "获取用户菜单信息成功")
}

// GetMyselfInfo 获取自己的信息
func GetMyselfInfo(context *gin.Context) {
	if !IsAuth(context) {
		return
	}
	user, _ := context.Get("user")
	if user == nil {
		response.ProcessFailed(context, nil, "获取自身信息失败")
		return
	}
	response.GetSuccessed(context, dto.ToUserDto(user.(model.User)), "获取自身信息成功")
}
