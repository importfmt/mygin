package model

import (
	"github.com/jinzhu/gorm"
)

type Wips struct {
	gorm.Model
	UserID   uint   `gorm:"type:int(11);not null"`
	Username string `gorm:"type:varchar(16);not null"`
	Title    string `gorm:"type:varchar(20);not null"`
	Desc     string `gorm:"type:varchar(255);not null"`
	Reply    string `gorm:"type:varchar(255)`
	Status   bool   `gorm:"type:bool;not null"`
}
