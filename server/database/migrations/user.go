package migrations

import (
	"afkser/initialize"
	"afkser/model"
)

// MigrationUser 生成用户表
func MigrationUser() {
	initialize.DB.AutoMigrate(&model.User{})
	initialize.DB.Model(&model.User{}).AddIndex("idx_user_mobile", "mobile")
}
