package controller

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"mygin.com/mygin/common"
	"mygin.com/mygin/dto"
	"mygin.com/mygin/model"
	"mygin.com/mygin/response"
)

// AddCar 添加新的车辆信息
func AddCar(context *gin.Context) {
	// 验证用户是否获得权限
	if !IsAuth(context) {
		return
	}
	var car model.Car
	context.BindJSON(&car)
	db := common.GetDatabase()
	if car.Brand == "" || car.License == "" || car.DeadWeight == 0 {
		response.ProcessFailed(context, nil, "请求参数错误")
		return
	}

	// 初始化车辆状态为 false
	car.Status = false

	if len(car.City) < 3 || car.City == "" {
		car.City = "北京市"
	}

	// 如果 License 存在
	if isLicenseExist(db, car.License) {
		response.ProcessFailed(context, nil, "车牌号码已经存在")
		return
	}

	// 寻找软删除的记录中是否有符合车牌号码的记录
	var unscopedCar model.Car
	db.Unscoped().Where("license = ?", car.License).First(&unscopedCar)
	if unscopedCar.ID != 0 {
		response.ProcessFailed(context, nil, "车牌号码已经注销")
		return
	}

	db.Create(&car)

	// 验证创建新的车辆是否成功
	if !isLicenseExist(db, car.License) {
		response.ProcessFailed(context, nil, "车辆信息已经注销")
		return
	}

	response.PostSuccessed(context, dto.ToCarDto(car), "新增车辆信息成功")
}

// GetCars 根据 URL 参数查询车辆信息
func GetCars(context *gin.Context) {
	if context.Query("pagenum") == "0" || context.Query("pagenum") == "" ||
		context.Query("pagesize") == "0" || context.Query("pagesize") == "" {
		GetCarsByCarLicense(context)
		return
	}

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

	if category != "all" && category != "true" && category != "false" && category != "city" &&
		category != "brand" && category != "asc" && category != "desc" && category != "more" && category != "less" &&
		status != "deleted" && status != "" {
		response.ProcessFailed(context, nil, "请求参数错误")
	}

	db := common.GetDatabase()
	var cars []model.Car
	var total int

	if status == "" {
		if category == "all" {
			db = db.Where("license LIKE ?", "%"+query+"%")
			db.Find(&cars)
			total = len(cars)
			db = db.Limit(pageSize).Offset((pageNum - 1) * pageSize)
			db.Where("license LIKE ?", "%"+query+"%").Find(&cars)
		}

		if category == "true" || category == "false" {
			db = db.Where("license LIKE ? AND status = "+category, "%"+query+"%")
			db.Find(&cars)
			total = len(cars)
			db = db.Limit(pageSize).Offset((pageNum - 1) * pageSize)
			db.Table("cars").Where("license LIKE ? AND status = "+category, "%"+query+"%").Find(&cars)
		}

		if category == "city" {
			db = db.Where("city LIKE ?", "%"+query+"%")
			db.Find(&cars)
			total = len(cars)
			db = db.Limit(pageSize).Offset((pageNum - 1) * pageSize)
			db.Where("city LIKE ?", "%"+query+"%").Find(&cars)
		}

		if category == "desc" || category == "asc" {
			db = db.Where("license LIKE ?", "%"+query+"%")
			db.Find(&cars)
			total = len(cars)
			db = db.Limit(pageSize).Offset((pageNum - 1) * pageSize)
			db.Order("dead_weight " + category).Find(&cars)
		}

		if category == "brand" {
			db = db.Where("brand LIKE ?", "%"+query+"%")
			db.Find(&cars)
			total = len(cars)
			db = db.Limit(pageSize).Offset((pageNum - 1) * pageSize)
			db.Where("brand LIKE ?", "%"+query+"%").Find(&cars)
		}

		if category == "more" || category == "less" {
			sign := ">"
			if category == "less" {
				sign = "<"
			}
			db = db.Where("dead_weight "+sign+" ?", query)
			db.Find(&cars)
			total = len(cars)
			db = db.Limit(pageSize).Offset((pageNum - 1) * pageSize)
			db.Where("dead_weight "+sign+" ?", query).Find(&cars)
		}

	} else if status == "deleted" {
		db = db.Unscoped().Where("license LIKE ? AND deleted_at is not null", "%"+query+"%")
		db.Find(&cars)
		total = len(cars)
		db = db.Limit(pageSize).Offset((pageNum - 1) * pageSize)
		db.Unscoped().Where("deleted_at is not null").Find(&cars)
	}

	var data = make([]dto.CarDto, len(cars))
	for i := 0; i < len(cars); i++ {
		data[i] = dto.ToCarDto(cars[i])
	}

	response.GetSuccessed(context, gin.H{"total": total, "cars": data}, "查询车辆信息成功")
}

// GetCarsByCarLicense 根据车辆车牌号获取车辆信息
func GetCarsByCarLicense(context *gin.Context) {
	// 验证用户是否获得权限
	if !IsAuth(context) {
		return
	}

	var c struct{ License string }
	context.BindJSON(&c)

	if c.License == "" {
		response.ProcessFailed(context, nil, "请求参数错误")
		return
	}

	db := common.GetDatabase()
	var car model.Car

	db.Where("license = ?", c.License).First(&car)
	if car.ID == 0 {
		response.ProcessFailed(context, nil, "车辆信息不存在")
		return
	}

	response.GetSuccessed(context, dto.ToCarDto(car), "获取车辆信息成功")
}

// GetCarsByCarID 根据车辆 ID 获取车辆信息
func GetCarsByCarID(context *gin.Context) {
	// 验证用户是否获得权限
	if !IsAuth(context) {
		return
	}

	carID := context.Param("carID")
	if carID == "" || carID == "0" {
		response.ProcessFailed(context, nil, "请求参数错误")
		return
	}

	db := common.GetDatabase()
	var car model.Car

	db.Where("id = ?", carID).First(&car)
	if car.ID == 0 {
		response.ProcessFailed(context, nil, "车辆信息不存在")
		return
	}

	response.GetSuccessed(context, dto.ToCarDto(car), "获取车辆信息成功")
}

// DeleteCarByCarLicense 根据车辆车牌号码删除车辆信息
func DeleteCarByCarLicense(context *gin.Context) {
	// 验证用户是否获得权限
	if !IsAuth(context) {
		return
	}

	var c struct{ License string }
	context.BindJSON(&c)
	if c.License == "" {
		response.ProcessFailed(context, nil, "请求参数错误")
		return
	}

	var car model.Car
	db := common.GetDatabase()

	db.Where("license = ?", c.License).First(&car)
	if car.ID == 0 {
		response.ProcessFailed(context, nil, "车辆信息不存在")
		return
	}
	db.Delete(&car)
	response.DeleteSuccessed(context, nil, "删除车辆信息成功")
}

// DeleteCarByCarID 根据车辆 ID 删除车辆信息
func DeleteCarByCarID(context *gin.Context) {
	// 验证用户是否获得权限
	if !IsAuth(context) {
		return
	}

	carID := context.Param("carID")
	if carID == "" || carID == "0" {
		response.ProcessFailed(context, nil, "请求参数错误")
		return
	}

	var car model.Car
	db := common.GetDatabase()
	db.Where("id = ?", carID).First(&car)
	if car.ID == 0 {
		response.ProcessFailed(context, nil, "车辆信息不存在")
		return
	}
	db.Delete(&car)

	response.DeleteSuccessed(context, nil, "删除车辆信息成功")
}

// UpdateCarInfoByCarID 根据车辆 ID 更新车辆信息
func UpdateCarInfoByCarID(context *gin.Context) {
	// 验证用户是否获得权限
	if !IsAuth(context) {
		return
	}

	carID := context.Param("carID")
	if carID == "" || carID == "0" {
		response.ProcessFailed(context, nil, "请求参数错误")
		return
	}

	var c struct {
		Brand      string
		License    string
		DeadWeight uint
		City       string
	}
	context.BindJSON(&c)

	if c.Brand == "" || c.License == "" || c.DeadWeight == 0 || c.City == "" || len(c.City) < 3 {
		response.ProcessFailed(context, nil, "请求参数错误")
		return
	}

	db := common.GetDatabase()
	var car model.Car

	db.Where("id = ?", carID).First(&car)
	if car.ID == 0 {
		response.ProcessFailed(context, nil, "车辆信息不存在")
		return
	}

	car.Brand = c.Brand
	car.License = c.License
	car.DeadWeight = c.DeadWeight
	car.City = c.City

	// 判断车辆车牌号码是否有过注册记录
	var unscopedCar model.Car
	db.Unscoped().Where("license = ? AND deleted is not null", car.License).First(&unscopedCar)
	if unscopedCar.ID != 0 {
		response.ProcessFailed(context, nil, "车辆车牌号码已经注销")
		return
	}

	db.Save(&car)
	response.PutSuccessed(context, dto.ToCarDto(car), "更新车辆信息成功")

}

// RecoverCarsByCarID 根据用户 ID 恢复用户信息
func RecoverCarsByCarID(context *gin.Context) {

	// 验证用户是否获得权限
	if !IsAuth(context) {
		return
	}

	carID := context.Param("carID")
	if carID == "" || carID == "0" {
		response.ProcessFailed(context, nil, "请求参数错误")
		return
	}

	var car model.Car
	db := common.GetDatabase()
	db.Unscoped().Where("id = ? AND deleted_at is not null", carID).First(&car)
	if car.ID == 0 {
		response.ProcessFailed(context, nil, "车辆信息不存在")
		return
	}

	db.Unscoped().Model(&car).Update("deleted_at", nil)

	response.PatchSuccessed(context, dto.ToCarDto(car), "恢复车辆信息成功")
}

// UpdateCarStatusByCarID 根据车辆 ID 更新车辆状态
func UpdateCarStatusByCarID(context *gin.Context) {

	// 验证用户是否获得权限
	if !IsAuth(context) {
		return
	}

	carID := context.Param("carID")
	if carID == "" || carID == "0" {
		response.ProcessFailed(context, nil, "请求参数错误")
		return
	}

	var c struct{ Status bool }
	context.BindJSON(&c)

	db := common.GetDatabase()
	var car model.Car

	db.Where("id = ?", carID).First(&car)
	if car.ID == 0 {
		response.ProcessFailed(context, nil, "车辆信息不存在")
		return
	}

	db.Model(&car).Update("status", c.Status)

	response.PatchSuccessed(context, dto.ToCarDto(car), "更新车辆状态成功")
}
