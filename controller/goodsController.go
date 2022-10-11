package controller

import (
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"mygin.com/mygin/common"
	"mygin.com/mygin/dto"
	"mygin.com/mygin/model"
	"mygin.com/mygin/response"
)

// AddGoods 添加货物信息
func AddGoods(context *gin.Context) {

	// 验证用户是否获得权限
	if !IsAuth(context) {
		return
	}

	var goods model.Goods
	context.BindJSON(&goods)

	if goods.Name == "" || goods.Price == 0 || goods.Weight == 0 ||
		goods.FromCity == "" || goods.ToCity == "" || len(goods.FromCity) < 3 || len(goods.ToCity) < 3 {
		response.ProcessFailed(context, nil, "请求参数错误")
		return
	}

	goods.Status = false
	goods.CourierStatus = false

	db := common.GetDatabase()
	db.Create(&goods)
	response.PostSuccessed(context, dto.ToGoodsDto(goods), "新增货物信息成功")
}

// GetGoods 根据 URL 参数查询货物信息
func GetGoods(context *gin.Context) {

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
	var goods []model.Goods
	var total int

	if status == "" {
		if category != "all" && category != "price_asc" && category != "price_desc" && category != "price_more" && category != "price_less" &&
			category != "weight_desc" && category != "weight_asc" && category != "weight_more" && category != "weight_less" &&
			category != "from_city" && category != "to_city" && category != "courier_true" && category != "courier_false" && category != "goods_true" && category != "goods_false" {
			response.ProcessFailed(context, nil, "请求参数错误")
			return
		}

		if category == "all" {
			db = db.Where("name LIKE ?", "%"+query+"%")
			db.Find(&goods)
			total = len(goods)
			db = db.Limit(pageSize).Offset((pageNum - 1) * pageSize)
			db.Where("name LIKE ?", "%"+query+"%").Find(&goods)
		}
		if category != "all" {
			// price_asc price_desc weight_desc weight_asc
			if strings.Split(category, "_")[1] == "asc" || strings.Split(category, "_")[1] == "desc" {
				db = db.Where("name LIKE ?", "%"+query+"%")
				db.Find(&goods)
				total = len(goods)
				db = db.Limit(pageSize).Offset((pageNum - 1) * pageSize)
				db.Order(strings.Split(category, "_")[0] + " " + strings.Split(category, "_")[1]).Find(&goods)
			}

			// price_more price_less weight_more weight_less
			if strings.Split(category, "_")[1] == "more" || strings.Split(category, "_")[1] == "less" {
				sign := ">"
				if strings.Split(category, "_")[1] == "less" {
					sign = "<"
				}

				db = db.Where(strings.Split(category, "_")[0]+" "+sign+" ?", query)
				db.Find(&goods)
				total = len(goods)
				db = db.Limit(pageSize).Offset((pageNum - 1) * pageSize)
				db.Where(strings.Split(category, "_")[0]+" "+sign+" ?", query).Find(&goods)
			}

			// courier_true courier_false goods_true goods_false
			if strings.Split(category, "_")[1] == "true" || strings.Split(category, "_")[1] == "false" {
				field := "status"
				if strings.Split(category, "_")[0] == "courier" {
					field = "courier_status"
				}
				db = db.Where(field + " = " + strings.Split(category, "_")[1])
				db.Find(&goods)
				total = len(goods)
				db = db.Limit(pageSize).Offset((pageNum - 1) * pageSize)
				db.Where(field + " = " + strings.Split(category, "_")[1]).Find(&goods)
			}

			//from_city to_city
			if strings.Split(category, "_")[1] == "city" {
				db = db.Where(category+" LIKE ?", "%"+query+"%")
				db.Find(&goods)
				total = len(goods)
				db = db.Limit(pageSize).Offset((pageNum - 1) * pageSize)
				db = db.Where(category+" LIKE ?", "%"+query+"%").Find(&goods)
			}
		}

	} else if status == "deleted" {
		db = db.Unscoped().Where("name LIKE ? AND deleted_at is not null", "%"+query+"%")
		db.Find(&goods)
		total = len(goods)
		db = db.Limit(pageSize).Offset((pageNum - 1) * pageSize)
		db.Unscoped().Where("deleted_at is not null").Find(&goods)
	}

	var data = make([]dto.GoodsDto, len(goods))
	for i := 0; i < len(goods); i++ {
		data[i] = dto.ToGoodsDto(goods[i])
	}
	response.GetSuccessed(context, gin.H{"total": total, "goods": data}, "获取货物信息成功")
}

// GetGoodsByGoodsID 根据货物 ID 查询货物信息
func GetGoodsByGoodsID(context *gin.Context) {

	// 验证用户是否获得权限
	if !IsAuth(context) {
		return
	}

	goodsID := context.Param("goodsID")
	if goodsID == "" || goodsID == "0" {
		response.ProcessFailed(context, nil, "请求参数错误")
		return
	}

	db := common.GetDatabase()
	var goods model.Goods

	db.Table("goods").Where("id = ?", goodsID).First(&goods)
	if goods.ID == 0 {
		response.ProcessFailed(context, nil, "货物信息不存在")
		return
	}
	response.GetSuccessed(context, dto.ToGoodsDto(goods), "获取货物信息成功")
}

// DeleteGoodsByGoodsID 根据货物 ID 删除货物信息
func DeleteGoodsByGoodsID(context *gin.Context) {

	// 验证用户是否获得权限
	if !IsAuth(context) {
		return
	}

	goodsID := context.Param("goodsID")
	if goodsID == "" || goodsID == "0" {
		response.ProcessFailed(context, nil, "请求参数错误")
		return
	}

	db := common.GetDatabase()
	var goods model.Goods

	db.Table("goods").Where("id = ?", goodsID).First(&goods)
	if goods.ID == 0 {
		response.ProcessFailed(context, nil, "货物信息不存在")
		return
	}

	if goods.CourierStatus == true {
		response.ProcessFailed(context, nil, "货物正在运输中")
		return
	}

	// 判断是否该货物是否有订单正在作业，如果作业未完成则不能删除
	var order model.Order
	db.Table("orders").Where("goods_id = ?", goodsID).First(&order)
	if order.ID != 0 && order.CourierStatus == true {
		response.ProcessFailed(context, nil, "货物正在运输中")
		return
	}

	db.Delete(&goods)

	response.DeleteSuccessed(context, nil, "删除货物信息成功")
}

// UpdateGoodsInfoByGoodsID 根据货物 ID 更新货物信息
func UpdateGoodsInfoByGoodsID(context *gin.Context) {

	// 验证用户是否获得权限
	if !IsAuth(context) {
		return
	}

	goodsID := context.Param("goodsID")
	if goodsID == "" || goodsID == "0" {
		response.ProcessFailed(context, nil, "请求参数错误")
		return
	}

	var g struct {
		Name     string
		Price    uint
		Weight   uint
		FromCity string
		ToCity   string
	}
	context.BindJSON(&g)

	if g.Name == "" || g.Price == 0 || g.Weight == 0 || g.FromCity == "" ||
		len(g.FromCity) < 3 || g.ToCity == "" || len(g.ToCity) < 3 {
		response.ProcessFailed(context, nil, "请求参数错误")
	}

	db := common.GetDatabase()
	var goods model.Goods

	db.Table("goods").Where("id = ?", goodsID).First(&goods)
	if goods.ID == 0 {
		response.ProcessFailed(context, nil, "货物信息不存在")
		return
	}
	if goods.CourierStatus == true || goods.Status == true {
		response.ProcessFailed(context, nil, "货物已经完成运输")
		return
	}
	goods.Name = g.Name
	goods.Price = g.Price
	goods.Weight = g.Weight
	goods.FromCity = g.FromCity
	goods.ToCity = g.ToCity

	db.Save(&goods)
	response.PutSuccessed(context, dto.ToGoodsDto(goods), "更新货物信息成功")
}

// RecoverGoodsByGoodsID 根据货物 ID 恢复货物信息
func RecoverGoodsByGoodsID(context *gin.Context) {
	// 验证用户是否获得权限
	if !IsAuth(context) {
		return
	}

	goodsID := context.Param("goodsID")
	if goodsID == "" || goodsID == "0" {
		response.ProcessFailed(context, nil, "请求参数错误")
		return
	}

	db := common.GetDatabase()
	var goods model.Goods

	db.Table("goods").Unscoped().Where("id = ?", goodsID).First(&goods)
	if goods.ID == 0 {
		response.ProcessFailed(context, nil, "货物信息不存在")
		return
	}

	db.Unscoped().Model(&goods).Update("deleted_at", nil)
	response.PutSuccessed(context, dto.ToGoodsDto(goods), "恢复货物信息成功")
}

// UpdateGoodsStatusByGoodsID 根据货物 ID 更新货物状态
func UpdateGoodsStatusByGoodsID(context *gin.Context) {
	// 验证用户是否获得权限
	if !IsAuth(context) {
		return
	}

	requestUser, _ := context.Get("user")
	if requestUser == nil {
		response.ProcessFailed(context, nil, "用户权限不足")
		return
	}

	goodsID := context.Param("goodsID")
	if goodsID == "" || goodsID == "0" {
		response.ProcessFailed(context, nil, "请求参数错误")
		return
	}

	var g struct{ Status bool }
	context.BindJSON(&g)

	db := common.GetDatabase()
	var goods model.Goods

	db.Table("goods").Where("id = ?", goodsID).First(&goods)
	if goods.ID == 0 {
		response.ProcessFailed(context, nil, "货物信息不存在")
		return
	}
	db.Model(&goods).Update("status", g.Status)
	if g.Status == true {
		db.Model(&goods).Update("courier_status", false)
	}
	response.PutSuccessed(context, dto.ToGoodsDto(goods), "更新货物完成状态成功")
}

// UpdateGoodsCourierStatusByGoodsID 根据货物 ID 更新货物运输状态
func UpdateGoodsCourierStatusByGoodsID(context *gin.Context) {
	// 验证用户是否获得权限
	if !IsAuth(context) {
		return
	}

	goodsID := context.Param("goodsID")
	if goodsID == "" || goodsID == "0" {
		response.ProcessFailed(context, nil, "请求参数错误")
		return
	}

	var g struct{ CourierStatus bool }
	context.BindJSON(&g)

	db := common.GetDatabase()
	var goods model.Goods

	db.Table("goods").Where("id = ?", goodsID).First(&goods)
	if goods.ID == 0 {
		response.ProcessFailed(context, nil, "货物信息不存在")
		return
	}
	if goods.Status == true {
		response.ProcessFailed(context, nil, "货物已经完成运输")
		return
	}

	db.Model(&goods).Update("courier_status", g.CourierStatus)
	response.PutSuccessed(context, dto.ToGoodsDto(goods), "更新货物运输状态成功")
}
