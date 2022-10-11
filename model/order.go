package model

import (
	"github.com/jinzhu/gorm"
)

// Order 订单
type Order struct {
	gorm.Model
	Number        string `gorm:"type:varchar(20);not null;unique"`
	Username      string `gorm:"type:varchar(16);not null"`
	License       string `gorm:"type:varchar(10);not null"`
	Goodsname     string `gorm:"type:varchar(20);not null"`
	GoodsID       uint   `gorm:"type:int unsigned;not null"`
	CarrierCarID  uint   `gorm:"type:int unsigned"`
	CarrierUserID uint   `gorm:"type:int unsigned"`
	Price         uint   `gorm:"int unsigned;not null"`
	Weight        uint   `gorm:"int unsigned;not null"`
	CourierNumber string `gorm:"type:varchar(20);unique"`
	CourierStatus bool   `gorm:"type:bool;not null"`
	FromCity      string `grom:"type:varchar(30);not null"`
	ToCity        string `grom:"type:varchar(30);not null"`
	Status        bool   `gorm:"type:bool;not null"`
}
