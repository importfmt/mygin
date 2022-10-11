package model

import (
	"github.com/jinzhu/gorm"
)

// User 用户
type User struct {
	gorm.Model 
	Username string `gorm:"type:varchar(16);not null;unique"`
	Password string `gorm:"size:255;not null"`
	Email    string `gorm:"type:varchar(20);not null;unique"`
	Mobile   string `gorm:"type:varchar(11);not null;unique"`
	City     string `gorm:"type:varchar(30);not null"`
	Role     string `gorm:"type:varchar(10);not null"`
	Status   bool `gorm:"type:bool;not null"`
}
