package migrations

import (
	"afkser/initialize"
	"afkser/model"
)

// MigrationUser 生成用户表
func MigrationUser() {
	initialize.DB.AutoMigrate(&model.Users{})
	initialize.DB.Model(&model.Users{}).AddUniqueIndex("idx_user_mobile", "mobile")
}
