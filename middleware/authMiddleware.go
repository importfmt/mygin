package middleware

import (
	"strings"
	// "fmt"
	"github.com/gin-gonic/gin"
	"mygin.com/mygin/common"
	"mygin.com/mygin/model"
	"mygin.com/mygin/response"
)

// AuthMiddleware 验证 token 有效性
func AuthMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		tokenString := context.GetHeader("Authorization")
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer") {
			response.AuthorizeFailed(context, nil, "用户认证失败")
			context.Abort()
			return
		}

		// 将携带的 Bearer 字段截取，以获得真实的 token
		tokenString = tokenString[7:]
		token, claims, err := common.ParseToken(tokenString)
		if err != nil || !token.Valid {
			response.AuthorizeFailed(context, nil, "用户认证失败")
			context.Abort()
			return
		}

		// 获取 token 中的 userID
		userID := claims.UserID
		// 验证 userID 是否存在
		db := common.GetDatabase()
		var user model.User
		db.First(&user, userID)

		if user.ID == 0 {
			response.AuthorizeFailed(context, nil, "用户认证失败")
			context.Abort()
			return
		}

		// 实现用户实现的功能请求接口的认证放行

		if strings.Split(context.FullPath(), "/")[3] == "menus" {
			context.Set("user", user)
		}

		if strings.Split(context.FullPath(), "/")[3] == "info" {
			context.Set("user", user)
		}

		if strings.Split(context.FullPath(), "/")[3] == "users" {
			if user.Role == "超级管理员" || user.Role == "用户管理员" {
				context.Set("user", user)
			}
		}

		if strings.Split(context.FullPath(), "/")[3] == "cars" {
			if user.Role == "超级管理员" || user.Role == "车辆管理员" {
				context.Set("user", user)
			}
		}

		if strings.Split(context.FullPath(), "/")[3] == "goods" {
			if user.Role == "超级管理员" || user.Role == "货物管理员" {
				context.Set("user", user)
			}
		}

		if strings.Split(context.FullPath(), "/")[3] == "orders" {
			if user.Role == "超级管理员" || user.Role == "订单管理员" {
				context.Set("user", user)
			}
		}

		if strings.Split(context.FullPath(), "/")[3] == "roles" {
			if user.Role == "超级管理员" || user.Role == "权限管理员" {
				context.Set("user", user)
			}
		}

		if strings.Split(context.FullPath(), "/")[3] == "wips" {
			if user.Role == "超级管理员" || user.Role == "工单管理员" {
				context.Set("user", user)
			}
		}

		context.Next()
	}
}
