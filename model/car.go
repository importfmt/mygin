package model

import (
	"github.com/jinzhu/gorm"
)

// Car 车辆
type Car struct {
	gorm.Model
	Brand      string `gorm:"type:varchar(20);not null"`
	License    string `gorm:"type:varchar(10);not null;unique"`
	DeadWeight uint   `gorm:"type:int unsigned;not null"`
	City       string `gorm:"type:varchar(30);not null"`
	Status     bool   `gorm:"type:bool;not null"`
}
