package main

import (
	"github.com/gin-gonic/gin"
	"mygin.com/mygin/controller"
	"mygin.com/mygin/middleware"
)

// CollectRoute 路由分发
func CollectRoute(engine *gin.Engine) {
	// 以下接口不需要进行权限认证
	engine.POST("/api/v1/register", controller.UserRegister) // 用户注册
	engine.POST("/api/v1/login", controller.UserLogin)       // 用户登录
	// 以下接口为客户端所调用
	engine.GET("/api/v1/menus/user", middleware.UserAuthMiddleware(), controller.GetUserMenu)                                            // 获取用户菜单列表
	engine.GET("/api/v1/info", middleware.UserAuthMiddleware(), controller.GetMyselfInfo)                                                // 获取自己的信息
	engine.GET("/api/v1/users/:userID/wips", middleware.UserAuthMiddleware(), controller.GetWipsByUserID)                                // 根据用户 ID 查询用户的工单信息
	engine.GET("/api/v1/users/:userID/day", middleware.UserAuthMiddleware(), controller.GetDayByUserID)                                  // 根据用户 ID 查询用户的工作天数
	engine.GET("/api/v1/users/:userID/orders", middleware.UserAuthMiddleware(), controller.GetOrdersByUserID)                            // 根据用户 ID 查询用户的订单信息
	engine.PATCH("/api/v1/users/:userID/email", middleware.UserAuthMiddleware(), controller.UpdateUserEmailByUserID)                     // 根据用户 ID 更新用户邮箱
	engine.PATCH("/api/v1/users/:userID/mobile", middleware.UserAuthMiddleware(), controller.UpdateUserMobileByUserID)                   // 根据用户 ID 更新用户手机号码
	engine.PATCH("/api/v1/users/:userID/password", middleware.UserAuthMiddleware(), controller.UpdateUserPasswordByUserID)               // 根据用户 ID 更新用户密码
	engine.PATCH("/api/v1/users/:userID/city", middleware.UserAuthMiddleware(), controller.UpdateUserCityByUserID)                       // 根据用户 ID 更新用户城市
	engine.PATCH("/api/v1/orders/:orderID/status", middleware.UserAuthMiddleware(), controller.UpdateOrderStatusByOrderID)               // 根据订单 ID 修改订单物流状态
	engine.PATCH("/api/v1/orders/:orderID/courierstatus", middleware.UserAuthMiddleware(), controller.UpdateOrderCourierStatusByOrderID) // 根据订单 ID 修改订单完成状态
	engine.POST("/api/v1/wips", middleware.UserAuthMiddleware(), controller.AddWip)                                                      // 添加新的工单信息
	// 以下接口为后台管理系统所调用
	engine.GET("/api/v1/menus", middleware.AuthMiddleware(), controller.GetAllMenu) // 获取菜单列表

	userGroup := engine.Group("/api/v1/users")                              // 用户路由组
	userGroup.Use(middleware.AuthMiddleware())                              // 为 userGroup 设置中间件
	userGroup.GET("", controller.GetUsers)                                  // 根据 URL 参数查询用户信息
	userGroup.GET("/:userID", controller.GetUserByUserID)                   // 根据用户 ID 查询用户信息
	userGroup.DELETE("/:userID", controller.DeleteUserByUserID)             // 根据用户 ID 删除用户信息
	userGroup.PUT("/:userID", controller.UpdateUserInfoByUserID)            // 根据用户 ID 更新用户信息
	userGroup.PATCH("/:userID/recover", controller.RecoverUsersByUserID)    // 根据用户 ID 恢复用户信息
	userGroup.PATCH("/:userID/status", controller.UpdateUserStatusByUserID) // 根据用户 ID 更新用户状态
	userGroup.PATCH("/:userID/role", controller.UpdateUserRoleByUserID)     // 根据用户 ID 更新用户角色

	carGroup := engine.Group("/api/v1/cars")                            // 车辆路由组
	carGroup.Use(middleware.AuthMiddleware())                           // 为 carGroup 设置中间件
	carGroup.POST("", controller.AddCar)                                // 添加新的车辆信息
	carGroup.GET("", controller.GetCars)                                // 根据 URL 参数查询车辆信息
	carGroup.GET("/:carID", controller.GetCarsByCarID)                  // 根据车辆 ID 查询车辆信息 || 根据车辆车牌号码查询车辆信息
	carGroup.GET("/:carID/orders", controller.GetOrdersByCarID)         // 根据车辆 ID 查询订单信息
	carGroup.DELETE("/:carID", controller.DeleteCarByCarID)             // 根据车辆 ID 删除车辆信息
	carGroup.DELETE("", controller.DeleteCarByCarLicense)               // 根据车辆车牌号码删除车辆信息
	carGroup.PUT("/:carID", controller.UpdateCarInfoByCarID)            // 根据车辆 ID 更新车辆信息
	carGroup.PATCH("/:carID/recover", controller.RecoverCarsByCarID)    // 根据车辆 ID 恢复车辆信息
	carGroup.PATCH("/:carID/status", controller.UpdateCarStatusByCarID) // 根据车辆 ID 更新车辆状态

	goodsGroup := engine.Group("/api/v1/goods")                                               // 获取路由组
	goodsGroup.Use(middleware.AuthMiddleware())                                               // 为 goodsGroup 设置中间件
	goodsGroup.POST("", controller.AddGoods)                                                  // 添加新的货物信息
	goodsGroup.GET("", controller.GetGoods)                                                   // 根据 URL 参数查询货物信息
	goodsGroup.GET("/:goodsID", controller.GetGoodsByGoodsID)                                 // 根据货物 ID 查询货物信息
	goodsGroup.DELETE("/:goodsID", controller.DeleteGoodsByGoodsID)                           // 根据货物 ID 删除货物信息
	goodsGroup.PUT("/:goodsID", controller.UpdateGoodsInfoByGoodsID)                          // 根据货物 ID 更新货物信息
	goodsGroup.PATCH("/:goodsID/recover", controller.RecoverGoodsByGoodsID)                   // 根据货物 ID 恢复货物信息
	goodsGroup.PATCH("/:goodsID/status", controller.UpdateGoodsStatusByGoodsID)               // 根据货物 ID 更新货物状态
	goodsGroup.PATCH("/:goodsID/courierstatus", controller.UpdateGoodsCourierStatusByGoodsID) // 根据货物 ID 更新货物运输状态

	orderGroup := engine.Group("/api/v1/orders")                             // 订单路由组
	orderGroup.Use(middleware.AuthMiddleware())                              // 为 orderGroup 设置中间件
	orderGroup.POST("", controller.AddOrder)                                 // 添加新的订单信息
	orderGroup.GET("", controller.GetOrders)                                 // 根据 URL 参数查询订单信息
	orderGroup.GET("/:orderID", controller.GetOrdersByOrderID)               // 根据订单 ID 查询订单信息
	orderGroup.DELETE("", controller.DeleteOrderByOrderNumber)               // 根据订单编号删除订单信息
	orderGroup.DELETE("/:orderID", controller.DeleteOrderByOrderID)          // 根据订单 ID 删除订单信息
	orderGroup.PUT("/:orderID", controller.UpdateOrderInfoByOrderID)         // 根据订单 ID 修改订单信息
	orderGroup.PATCH("/:orderID/recover", controller.RecoverOrdersByOrderID) // 根据订单 ID 恢复订单信息

	roleGroup := engine.Group("/api/v1/roles") // 角色路由组
	roleGroup.Use(middleware.AuthMiddleware()) // 为 roleGroup 设置中间件
	roleGroup.GET("", controller.GetRoles)     // 获取权限列表

	wipGroup := engine.Group("/api/v1/wips")                       // 工单路由组
	wipGroup.Use(middleware.AuthMiddleware())                      // 为 roleGroup 设置中间件
	wipGroup.GET("", controller.GetWips)                           // 根据 URL 参数查询工单信息
	wipGroup.PATCH("/:wipID/reply", controller.UpdateReplyByWipID) // 根据工单 ID 修改回复信息

}
