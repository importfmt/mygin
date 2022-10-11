package controller

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"mygin.com/mygin/common"
	"mygin.com/mygin/dto"
	"mygin.com/mygin/model"
	"mygin.com/mygin/response"
	// "mygin.com/mygin/util"
)

// Addwip 添加新的工单信息
func AddWip(context *gin.Context) {
	// 验证用户是否获得权限
	if !IsAuth(context) {
		return
	}

	var wip model.Wips
	context.BindJSON(&wip)

	if wip.Title == "" || wip.Desc == "" || wip.UserID == 0 {
		response.ProcessFailed(context, nil, "请求参数错误")
		return
	}

	db := common.GetDatabase()

	wip.Status = false

	var user model.User
	db.Table("users").Where("id = ?", wip.UserID).First(&user)
	if user.ID == 0 {
		response.ProcessFailed(context, nil, "用户信息不存在")
		return
	}

	wip.Username = user.Username

	db.Create(&wip)

	response.PostSuccessed(context, dto.ToWipDto(wip), "创建工单信息成功")

}

// GetWips 根据 URL 参数查询工单信息
func GetWips(context *gin.Context) {
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
	category := context.Query("category")

	db := common.GetDatabase()
	var wips []model.Wips
	var total int

	if category != "all" && category != "processed" && category != "untreated" && category != "asc" && category != "desc" {
		response.ProcessFailed(context, nil, "请求参数错误")
	}

	if category == "all" {
		db = db.Where("title LIKE ?", "%"+query+"%")
		db.Find(&wips)
		total = len(wips)
		db = db.Limit(pageSize).Offset((pageNum - 1) * pageSize)
		db.Where("title LIKE ?", "%"+query+"%").Find(&wips)
	}

	if category == "processed" || category == "untreated" {
		status := "true"
		if category == "untreated" {
			status = "false"
		}
		db = db.Where("status = "+status+" AND title LIKE ?", "%"+query+"%")
		db.Find(&wips)
		total = len(wips)
		db = db.Limit(pageSize).Offset((pageNum - 1) * pageSize)
		db.Where("status = "+status+" AND title LIKE ?", "%"+query+"%").Find(&wips)
	}

	if category == "desc" || category == "asc" {
		db = db.Where("title LIKE ?", "%"+query+"%")
		db.Find(&wips)
		total = len(wips)
		db = db.Limit(pageSize).Offset((pageNum - 1) * pageSize)
		db.Where("title LIKE ?", "%"+query+"%").Order("id " + category).Find(&wips)
	}

	var data = make([]dto.WipDto, len(wips))
	for i := 0; i < len(wips); i++ {
		data[i] = dto.ToWipDto(wips[i])
	}
	response.GetSuccessed(context, gin.H{"total": total, "wips": data}, "获取工单信息成功")

}

// UpdateReplyByWipID 根据工单 ID 更新回复信息
func UpdateReplyByWipID(context *gin.Context) {
	// 验证用户是否获得权限
	if !IsAuth(context) {
		return
	}

	wipID := context.Param("wipID")

	var wip model.Wips
	var w struct {
		Reply string
	}

	context.BindJSON(&w)
	db := common.GetDatabase()

	db.Table("wips").Where("id = ?", wipID).First(&wip)
	if wip.ID == 0 {
		response.ProcessFailed(context, nil, "工单信息不存在")
		return
	}

	wip.Reply = w.Reply
	wip.Status = true
	db.Save(&wip)
	response.PatchSuccessed(context, dto.ToWipDto(wip), "更新工单信息成功")
}
