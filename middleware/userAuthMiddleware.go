package middleware

import (
	"strings"
	// "fmt"
	"github.com/gin-gonic/gin"
	"mygin.com/mygin/common"
	"mygin.com/mygin/model"
	"mygin.com/mygin/response"
)

// UserAuthMiddleware 验证 token 有效性
func UserAuthMiddleware() gin.HandlerFunc {
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
		context.Set("user", user)

		context.Next()
	}
}
