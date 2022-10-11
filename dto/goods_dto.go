package dto

import(
	"mygin.com/mygin/model"
)

// GoodsDto 用户的数据传送结构体
type GoodsDto struct{
	ID uint `json:"id"`
	Name string `json:"name"`
	Price uint `json:"price"`
	Weight uint `json:"weight"`
	CourierStatus bool `json:"courier_status"`
	CourierNumber string `json:"courier_number"`
	FromCity string `json:"from_city"`
	ToCity string `json:"to_city"`
	Status bool `json:"status"`
}

// ToGoodsDto 将 model.Goods 转换为 GoodsDto
func ToGoodsDto(goods model.Goods) GoodsDto  {
	return GoodsDto {
		ID: goods.ID,
		Name: goods.Name,
		Price: goods.Price,
		Weight: goods.Weight,
		CourierStatus: goods.CourierStatus,
		CourierNumber: goods.CourierNumber,
		FromCity: goods.FromCity,
		ToCity: goods.ToCity,
		Status: goods.Status,
	}
}