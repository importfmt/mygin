package response

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

// Response 响应
func Response(context *gin.Context, httpStatus int, code int, data interface{}, msg string) {
	context.JSON(httpStatus,gin.H{"code": code, "data": data, "msg": msg})
}
// GetSuccessed 成功响应
func GetSuccessed(context *gin.Context, data interface{}, msg string) {
	Response(context, http.StatusOK, 200, data, msg)
}

// PostSuccessed 成功响应
func PostSuccessed(context *gin.Context, data interface{}, msg string) {
	Response(context, http.StatusOK, 201, data, msg)
}

// PutSuccessed 成功响应
func PutSuccessed(context *gin.Context, data interface{}, msg string) {
	Response(context, http.StatusOK, 201, data, msg)
}

// PatchSuccessed 成功响应
func PatchSuccessed(context *gin.Context, data interface{}, msg string) {
	Response(context, http.StatusOK, 201, data, msg)
}

// DeleteSuccessed 成功响应
func DeleteSuccessed(context *gin.Context, data interface{}, msg string) {
	Response(context, http.StatusOK, 204, data, msg)
}

// ServerFailed 失败响应
func ServerFailed(context *gin.Context, data interface{}, msg string) {
	Response(context, http.StatusInternalServerError, 500, data, msg)
}

// AuthorizeFailed 认证失败
func AuthorizeFailed(context *gin.Context, data interface{}, msg string) {
	Response(context, http.StatusUnauthorized, 401, data, msg)
}

// ProcessFailed 失败响应
func ProcessFailed(context *gin.Context, data interface{}, msg string) {
	Response(context, http.StatusOK, 422, data, msg)
}

