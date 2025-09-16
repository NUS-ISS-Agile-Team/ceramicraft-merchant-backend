package db

import (
	"fmt"

	"github.com/cerami-craft-shop/merchant-backend/config"
	"github.com/cerami-craft-shop/merchant-backend/repository/db/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func Init() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.Config.MySQL.UserName,
		config.Config.MySQL.Password,
		config.Config.MySQL.Host,
		config.Config.MySQL.Port,
		config.Config.MySQL.DBName,
	)
	DB, err = gorm.Open(mysql.Open(dsn),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)
	if err != nil {
		panic(err)
	}
	err = DB.AutoMigrate(
		&model.Item{},
	)
	if err != nil {
		panic(err)
	}
}
