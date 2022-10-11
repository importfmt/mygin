package dto

import (
	"mygin.com/mygin/model"
)

// CarDto 用户的数据传送结构体
type CarDto struct {
	ID         uint   `json:"id"`
	Brand      string `json:"brand"`
	License    string `json:"license"`
	DeadWeight uint   `json:"dead_weight"`
	City       string `json:"city"`
	Status     bool   `json:"status"`
}

// ToCarDto 将 model.Car 转换为 CarDto
func ToCarDto(car model.Car) CarDto {
	return CarDto{
		ID:         car.ID,
		Brand:      car.Brand,
		License:    car.License,
		DeadWeight: car.DeadWeight,
		City:       car.City,
		Status:     car.Status,
	}
}
