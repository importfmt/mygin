package controller

import (
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"mygin.com/mygin/common"
	"mygin.com/mygin/dto"
	"mygin.com/mygin/model"
	"mygin.com/mygin/response"
	"mygin.com/mygin/util"
)

// AddOrder 添加订单信息
func AddOrder(context *gin.Context) {
	// 验证用户是否获得权限
	if !IsAuth(context) {
		return
	}

	var order model.Order
	context.BindJSON(&order)

	if order.CarrierUserID == 0 || order.CarrierCarID == 0 || order.GoodsID == 0 {
		response.ProcessFailed(context, nil, "请求参数错误")
		return
	}

	db := common.GetDatabase()
	var goods model.Goods
	db.Table("goods").Where("id = ?", order.GoodsID).First(&goods)
	if goods.ID == 0 {
		response.ProcessFailed(context, nil, "货物信息不存在")
		return
	}
	order.Goodsname = goods.Name

	var user model.User
	db.Table("users").Where("id = ?", order.CarrierUserID).First(&user)
	if user.ID == 0 {
		response.ProcessFailed(context, nil, "用户信息不存在")
		return
	}
	order.Username = user.Username

	var car model.Car
	db.Table("cars").Where("id = ?", order.CarrierCarID).First(&car)
	if car.ID == 0 {
		response.ProcessFailed(context, nil, "车辆信息不存在")
		return
	}
	order.License = car.License

	if user.City != car.City || user.City != goods.FromCity || car.City != goods.FromCity {
		response.ProcessFailed(context, nil, "用户|车辆|货物不在同一个地区")
		return
	}

	if user.Status == true || car.Status == true || goods.Status == true {
		response.ProcessFailed(context, nil, "用户|车辆|货物正在作业")
		return
	}

	if goods.CourierStatus == true {
		response.ProcessFailed(context, nil, "货物已经完成运输")
		return
	}

	if car.DeadWeight < goods.Weight {
		response.ProcessFailed(context, nil, "货物重量超过车辆的载重")
		return
	}

	order.Price = goods.Price
	order.Weight = goods.Weight
	order.FromCity = goods.FromCity
	order.ToCity = goods.ToCity
	order.Number = util.RandomString(20)

	// 生成 20 位的物流单号
	var courierNumber = util.RandomNumber(20)
	var o model.Order
	for {
		db.Where("courier_number = ?", courierNumber).First(&o)
		if o.ID == 0 {
			break
		}
		courierNumber = util.RandomNumber(20)
	}
	order.CourierNumber = courierNumber

	db.Create(&order)
	db.Model(&user).Update("status", true)
	db.Model(&car).Update("status", true)
	db.Model(&goods).Updates(map[string]interface{}{"courier_number": courierNumber})
	response.PostSuccessed(context, dto.ToOrderDto(order), "新建订单信息成功")
}

// GetOrders 根据 URL 参数查询订单信息
func GetOrders(context *gin.Context) {
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

	if status != "deleted" && status != "" {
		response.ProcessFailed(context, nil, "请求参数错误")
		return
	}


	db := common.GetDatabase()
	var orders []model.Order
	var total int

	if status == "deleted" {

		db = db.Unscoped().Where("number LIKE ? AND deleted_at is not null", "%"+query+"%")
		db.Find(&orders)
		total = len(orders)
		db = db.Limit(pageSize).Offset((pageNum - 1) * pageSize)
		db.Unscoped().Where("deleted_at is not null").Find(&orders)

	} else if status == "" {

		if category != "all" && category != "username" && category != "license" && category != "goodsname" &&
			category != "goods_true" && category != "goods_false" && category != "from_city" && category != "to_city" &&
			category != "courier_true" && category != "courier_false" {
			response.ProcessFailed(context, nil, "请求参数错误")
			return
		}

		if category == "all" {
			db = db.Where("number LIKE ?", "%"+query+"%")
			db.Find(&orders)
			total = len(orders)
			db = db.Limit(pageSize).Offset((pageNum - 1) * pageSize)
			db.Where("number LIKE ?", "%"+query+"%").Find(&orders)
		}
		if category == "username" {
			db = db.Where("username LIKE ?", "%"+query+"%")
			db.Find(&orders)
			total = len(orders)
			db = db.Limit(pageSize).Offset((pageNum - 1) * pageSize)
			db.Where("username LIKE ?", "%"+query+"%").Find(&orders)
		}
		if category == "license" {
			db = db.Where("license LIKE ?", "%"+query+"%")
			db.Find(&orders)
			total = len(orders)
			db = db.Limit(pageSize).Offset((pageNum - 1) * pageSize)
			db.Where("license LIKE ?", "%"+query+"%").Find(&orders)
		}
		if category == "goodsname" {
			db = db.Where("goodsname LIKE ?", "%"+query+"%")
			db.Find(&orders)
			total = len(orders)
			db = db.Limit(pageSize).Offset((pageNum - 1) * pageSize)
			db.Where("goodsname LIKE ?", "%"+query+"%").Find(&orders)
		}
		if category != "goodsname" && category != "license" && category != "username" && category != "all" {
			// courier_true courier_false goods_true goods_false
			if strings.Split(category, "_")[1] == "true" || strings.Split(category, "_")[1] == "false" {
				field := "status"
				if strings.Split(category, "_")[0] == "courier" {
					field = "courier_status"
				}
				db = db.Where(field + " = " + strings.Split(category, "_")[1])
				db.Find(&orders)
				total = len(orders)
				db = db.Limit(pageSize).Offset((pageNum - 1) * pageSize)
				db.Where(field + " = " + strings.Split(category, "_")[1]).Find(&orders)
			}
			//from_city to_city
			if strings.Split(category, "_")[1] == "city" {
				db = db.Where(category+" LIKE ?", "%"+query+"%")
				db.Find(&orders)
				total = len(orders)
				db = db.Limit(pageSize).Offset((pageNum - 1) * pageSize)
				db = db.Where(category+" LIKE ?", "%"+query+"%").Find(&orders)
			}
		}
	}
	var data = make([]dto.OrderDto, len(orders))
	for i := 0; i < len(orders); i++ {
		data[i] = dto.ToOrderDto(orders[i])
	}

	response.GetSuccessed(context, gin.H{"total": total, "orders": data}, "获取订单信息成功")
}

// GetOrdersByOrderID 根据订单 ID 查询订单信息
func GetOrdersByOrderID(context *gin.Context) {
	// 验证用户是否获得权限
	if !IsAuth(context) {
		return
	}

	orderID := context.Param("orderID")
	if orderID == "" || orderID == "0" {
		response.ProcessFailed(context, nil, "请求参数错误")
		return
	}

	db := common.GetDatabase()
	var order model.Order

	db.Table("orders").Where("id = ?", orderID).First(&order)
	if order.ID == 0 {
		response.ProcessFailed(context, nil, "订单信息不存在")
		return
	}

	response.GetSuccessed(context, dto.ToOrderDto(order), "查询订单信息成功")
}

// GetOrdersByCarID 根据车辆 ID 查询订单信息
func GetOrdersByCarID(context *gin.Context) {
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
        var orders []model.Order
        db.Table("orders").Where("carrier_car_id = ?", carID).Find(&orders)

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

// DeleteOrderByOrderNumber 根据订单编号删除订单信息
func DeleteOrderByOrderNumber(context *gin.Context) {
	// 验证用户是否获得权限
	if !IsAuth(context) {
		return
	}

	var o struct{ Number string }
	context.BindJSON(&o)
	if o.Number == "" || len(o.Number) != 20 {
		response.ProcessFailed(context, nil, "请求参数错误")
		return
	}

	db := common.GetDatabase()
	var order model.Order

	db.Where("number = ?", o.Number).First(&order)
	if order.ID == 0 {
		response.ProcessFailed(context, nil, "订单信息不存在")
		return
	}
	db.Delete(&order)
	response.DeleteSuccessed(context, nil, "删除订单信息成功")

	// 如果订单未完成就被删除
	if order.Status == false {
		// 修改承运用户的状态为 false
		var user model.User
		db.Table("users").Where("id = ?", order.CarrierUserID).First(&user)
		if user.ID == 0 {
			response.ProcessFailed(context, nil, "用户信息不存在")
			return
		}
		db.Model(&user).Update("status", false)

		// 修改承运车辆的状态为 false
		var car model.Car
		db.Table("cars").Where("id = ?", order.CarrierCarID).First(&car)
		if car.ID == 0 {
			response.ProcessFailed(context, nil, "车辆信息不存在")
			return
		}
		db.Model(&car).Update("status", false)

		var goods model.Goods
		db.Table("goods").Where("id = ?", order.GoodsID).First(&goods)
		if goods.ID == 0 {
			response.ProcessFailed(context, nil, "货物信息不存在")
			return
		}
		db.Model(&goods).Updates(map[string]interface{}{"status": false, "courier_number": nil})
	}
}

// DeleteOrderByOrderID 根据订单 ID 删除订单信息
func DeleteOrderByOrderID(context *gin.Context) {
	// 验证用户是否获得权限
	if !IsAuth(context) {
		return
	}

	orderID := context.Param("orderID")
	if orderID == "" || orderID == "0" {
		response.ProcessFailed(context, nil, "请求参数错误")
		return
	}

	db := common.GetDatabase()
	var order model.Order

	db.Table("orders").Where("id = ?", orderID).First(&order)
	if order.ID == 0 {
		response.ProcessFailed(context, nil, "订单信息不存在")
		return
	}

	db.Delete(&order)
	response.DeleteSuccessed(context, nil, "删除订单信息成功")

	// 如果订单未完成就被删除
	if order.Status == false {
		// 修改承运用户的状态为 false
		var user model.User
		db.Table("users").Where("id = ?", order.CarrierUserID).First(&user)
		if user.ID == 0 {
			response.ProcessFailed(context, nil, "用户信息不存在")
			return
		}
		db.Model(&user).Update("status", false)

		// 修改承运车辆的状态为 false
		var car model.Car
		db.Table("cars").Where("id = ?", order.CarrierCarID).First(&car)
		if car.ID == 0 {
			response.ProcessFailed(context, nil, "车辆信息不存在")
			return
		}
		db.Model(&car).Update("status", false)

		// 修改商品运输的状态为 false
		var goods model.Goods
		db.Table("goods").Where("id = ?", order.GoodsID).First(&goods)
		if goods.ID == 0 {
			response.ProcessFailed(context, nil, "货物信息不存在")
			return
		}
		db.Model(&goods).Updates(map[string]interface{}{"courier_status": false, "courier_number": nil})
	}
}

// UpdateOrderInfoByOrderID 根据订单 ID 修改订单信息
func UpdateOrderInfoByOrderID(context *gin.Context) {
	// 验证用户是否获得权限
	if !IsAuth(context) {
		return
	}

	orderID := context.Param("orderID")
	if orderID == "" || orderID == "0" {
		response.ProcessFailed(context, nil, "请求参数错误")
	}

	var o struct {
		CarrierUserID uint
		CarrierCarID  uint
		GoodsID       uint
	}
	context.BindJSON(&o)
	if o.CarrierUserID == 0 || o.CarrierCarID == 0 || o.GoodsID == 0 {
		response.ProcessFailed(context, nil, "请求参数错误")
		return
	}

	db := common.GetDatabase()
	var order model.Order

	db.Table("orders").Where("id = ?", orderID).First(&order)
	if order.ID == 0 {
		response.ProcessFailed(context, nil, "订单信息不存在")
		return
	}

	if order.Status == true {
		response.ProcessFailed(context, nil, "订单已完成")
		return
	}

	var goods model.Goods
	db.Table("goods").Where("id = ?", o.GoodsID).First(&goods)
	if goods.ID == 0 {
		response.ProcessFailed(context, nil, "货物信息不存在")
		return
	}
	if goods.Status == true {
		response.ProcessFailed(context, nil, "货物正在作业")
		return
	}
	if goods.CourierStatus == true {
		response.ProcessFailed(context, nil, "货物已经完成运输")
		return
	}

	var car model.Car
	db.Table("cars").Where("id = ?", o.CarrierCarID).First(&car)
	if car.ID == 0 {
		response.ProcessFailed(context, nil, "车辆信息不存在")
		return
	}
	if car.Status == true {
		response.ProcessFailed(context, nil, "车辆正在作业")
		return
	}

	var user model.User
	db.Table("users").Where("id = ?", o.CarrierUserID).First(&user)
	if user.ID == 0 {
		response.ProcessFailed(context, nil, "用户信息不存在")
		return
	}
	if user.Status == true {
		response.ProcessFailed(context, nil, "用户正在作业")
		return
	}

	// 判断车辆是否超重
	if goods.Weight > car.DeadWeight {
		response.ProcessFailed(context, nil, "货物重量超过车辆载重")
		return
	}

	order.CarrierUserID = o.CarrierUserID
	order.CarrierCarID = o.CarrierCarID
	order.GoodsID = o.GoodsID
	order.Weight = goods.Weight
	order.Price = goods.Price

	db.Save(&order)
}

// RecoverOrdersByOrderID 根据用户 ID 恢复用户信息
func RecoverOrdersByOrderID(context *gin.Context) {
	// 验证用户是否获得权限
	if !IsAuth(context) {
		return
	}

	orderID := context.Param("orderID")
	if orderID == "" || orderID == "0" {
		response.ProcessFailed(context, nil, "请求参数错误")
		return
	}

	var order model.Order
	db := common.GetDatabase()
	db.Table("orders").Unscoped().Where("id = ? AND deleted_at is not null", orderID).First(&order)
	if order.ID == 0 {
		response.ProcessFailed(context, nil, "订单信息不存在")
		return
	}

	db.Unscoped().Model(&order).Update("deleted_at", nil)

	response.PatchSuccessed(context, dto.ToOrderDto(order), "恢复订单信息成功")
}

// UpdateOrderStatusByOrderID 根据订单 ID 修改订单物流状态
func UpdateOrderStatusByOrderID(context *gin.Context) {
	// 验证用户是否获得权限
	if !IsAuth(context) {
		return
	}

	orderID := context.Param("orderID")
	if orderID == "" || orderID == "0" {
		response.ProcessFailed(context, nil, "请求参数错误")
		return
	}
	var o struct{ Status bool }
	context.BindJSON(&o)

	db := common.GetDatabase()
	var order model.Order

	db.Table("orders").Where("id = ?", orderID).First(&order)
	if order.ID == 0 {
		response.ProcessFailed(context, nil, "订单信息不存在")
		return
	}
	db.Model(&order).Update("status", o.Status)

	var goods model.Goods
	db.Table("goods").Where("id = ?", order.GoodsID).First(&goods)
	if goods.ID == 0 {
		response.ProcessFailed(context, nil, "货物信息不存在")
		return
	}
	db.Model(&goods).Updates(map[string]interface{}{"status": o.Status})

	var user model.User
	db.Table("users").Where("id = ?", order.CarrierUserID).First(&user)
	if user.ID == 0 {
		response.ProcessFailed(context, nil, "用户信息不存在")
		return
	}
	db.Model(&user).Updates(map[string]interface{}{"city": goods.ToCity, "status": false})

	var car model.Car
	db.Table("cars").Where("id = ?", order.CarrierCarID).First(&car)
	if car.ID == 0 {
		response.ProcessFailed(context, nil, "车辆信息不存在")
		return
	}
	db.Model(&car).Updates(map[string]interface{}{"city": goods.ToCity, "status": false})

	response.PatchSuccessed(context, dto.ToOrderDto(order), "更新订单状态成功")
}

// UpdateOrderCourierStatusByOrderID 根据订单 ID 修改订单物流状态
func UpdateOrderCourierStatusByOrderID(context *gin.Context) {
	// 验证用户是否获得权限
	if !IsAuth(context) {
		return
	}

	orderID := context.Param("orderID")
	if orderID == "" || orderID == "0" {
		response.ProcessFailed(context, nil, "请求参数错误")
		return
	}
	var o struct{ CourierStatus bool }
	context.BindJSON(&o)

	db := common.GetDatabase()
	var order model.Order

	db.Table("orders").Where("id = ?", orderID).First(&order)
	if order.ID == 0 {
		response.ProcessFailed(context, nil, "订单信息不存在")
		return
	}
	db.Model(&order).Update("courier_status", o.CourierStatus)

	var goods model.Goods
	db.Table("goods").Where("id = ?", order.GoodsID).First(&goods)
	if goods.ID == 0 {
		response.ProcessFailed(context, nil, "货物信息不存在")
		return
	}
	db.Model(&goods).Updates(map[string]interface{}{"courier_status": o.CourierStatus})

	response.PatchSuccessed(context, dto.ToOrderDto(order), "更新订单物流状态成功")
}
