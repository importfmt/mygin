package dto

import (
	"time"

	"mygin.com/mygin/model"
)

// OrderDto 订单的数据传送结构体
type OrderDto struct {
	ID            uint      `json:"id"`
	Number        string    `json:"number"`
	Username      string    `json:"username"`
	License       string    `json:"license"`
	Goodsname     string    `json:"goodsname"`
	GoodsID       uint      `json:"goods_id"`
	CarrierCarID  uint      `json:"carrier_car_id"`
	CarrierUserID uint      `json:"carrier_user_id"`
	Price         uint      `json:"price"`
	Weight        uint      `json:"weight"`
	CourierNumber string    `json:"courier_number"`
	CourierStatus bool      `json:"courier_status"`
	FromCity      string    `json:"from_city"`
	ToCity        string    `json:"to_city"`
	CreatedAt     time.Time `json:"created_at"`
	Status        bool      `json:"status"`
}

// ToOrderDto 将 model.Order 转换为 Orderdto
func ToOrderDto(order model.Order) OrderDto {
	return OrderDto{
		ID:            order.ID,
		Number:        order.Number,
		Username:      order.Username,
		License:       order.License,
		Goodsname:     order.Goodsname,
		GoodsID:       order.GoodsID,
		CarrierCarID:  order.CarrierCarID,
		CarrierUserID: order.CarrierUserID,
		Price:         order.Price,
		Weight:        order.Weight,
		CourierNumber: order.CourierNumber,
		CourierStatus: order.CourierStatus,
		FromCity:      order.FromCity,
		ToCity:        order.ToCity,
		CreatedAt:     order.CreatedAt,
		Status:        order.Status,
	}
}
