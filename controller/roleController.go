package controller

import (
	"github.com/gin-gonic/gin"
	"mygin.com/mygin/response"
	"mygin.com/mygin/dto"
)


// GetRoles 获取权限菜单列表
func GetRoles(context *gin.Context) {
	// 验证用户是否获得权限
	if !IsAuth(context) {
		return
	}

	array := make([]dto.RoleDto, 8)
	array[0] = dto.RoleDto{Name: "超级管理员", Desc:"拥有所有功能的权限"}
	array[1] = dto.RoleDto{Name: "用户管理员", Desc:"拥有用户功能的权限"}
	array[2] = dto.RoleDto{Name: "车辆管理员", Desc:"拥有车辆功能的权限"}
	array[3] = dto.RoleDto{Name: "货物管理员", Desc:"拥有货物功能的权限"}
	array[4] = dto.RoleDto{Name: "订单管理员", Desc:"拥有订单功能的权限"}
	array[5] = dto.RoleDto{Name: "权限管理员", Desc:"拥有权限功能的权限"}
	array[6] = dto.RoleDto{Name: "测试用户", Desc:"拥有权限功能的权限"}
	array[7] = dto.RoleDto{Name: "普通用户", Desc:"不拥有任何功能权限"}

	response.GetSuccessed(context, array, "获取权限菜单列表成功")
}
