package initialize

import (
	"afkser/utils"
	"time"

	"github.com/jinzhu/gorm"

	//
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// NewDatabase 在中间件中初始化mysql链接
func NewDatabase(user, password, path, dbname, config string) *gorm.DB {

	connString := user + ":" + password + "@tcp(" + path + ")/" + dbname + "?" + config
	db, err := gorm.Open("mysql", connString)
	db.LogMode(true)
	// Error
	if err != nil {
		utils.Log().Panic("连接数据库不成功", err)
	}
	//设置连接池
	//空闲
	db.DB().SetMaxIdleConns(50)
	//打开
	db.DB().SetMaxOpenConns(100)
	//超时
	db.DB().SetConnMaxLifetime(time.Second * 30)

	return db
}
