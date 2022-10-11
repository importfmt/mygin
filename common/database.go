package common

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql" // 依赖库
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	"mygin.com/mygin/model"
)

var db *gorm.DB

// InitDB 初始化数据库
func InitDB() *gorm.DB {
	var driverName string = viper.GetString("datasource.driverName")
	var host string = viper.GetString("datasource.host")
	var port string = viper.GetString("datasource.port")
	var database string = viper.GetString("datasource.database")
	var username string = viper.GetString("datasource.username")
	var password string = viper.GetString("datasource.password")
	var charset string = viper.GetString("datasource.charset")

	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true",
		username, password, host, port, database, charset)

	// 建立数据库连接
	var err error
	db, err = gorm.Open(driverName, args)
	if err != nil {
		panic("faild to connect database, err: " + err.Error())
	}

	// 创建数据表
	if !db.HasTable(&model.User{}) {
		db.CreateTable(&model.User{})
	} else {
		db.AutoMigrate(&model.User{})
	}

	if !db.HasTable(&model.Car{}) {
		db.CreateTable(&model.Car{})
	} else {
		db.AutoMigrate(&model.Car{})
	}

	if !db.HasTable(&model.Order{}) {
		db.CreateTable(&model.Order{})
	} else {
		db.AutoMigrate(&model.Order{})
	}

	if !db.HasTable(&model.Goods{}) {
		db.CreateTable(&model.Goods{})
	} else {
		db.AutoMigrate(&model.Goods{})
	}

	if !db.HasTable(&model.Wips{}) {
		db.CreateTable(&model.Wips{})
	} else {
		db.AutoMigrate(&model.Wips{})
	}

	return db
}

// GetDatabase 获取数据库实例
func GetDatabase() *gorm.DB {
	return db
}
