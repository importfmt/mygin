package model

import (
	"github.com/jinzhu/gorm"
)

// Goods 订单
type Goods struct {
	gorm.Model 
	Name string `gorm:"type:varchar(20);not null"`
	Price uint `gorm:"int unsigned;not null"`
	Weight uint `gorm:"int unsigned;not null"`
	CourierStatus   bool `gorm:"type:bool;not null"`
	CourierNumber string `gorm:"type:varchar(20)"`
	FromCity string `gorm:"type:varchar(30);not null"`
	ToCity string `gorm:"type:varchar(30);not null"`
	Status bool `gorm:"type:bool"`
}
