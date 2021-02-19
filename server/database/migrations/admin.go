package migrations

import (
	"afkser/initialize"
	"afkser/model"
	"os"
	"time"

	"github.com/gookit/color"
	"gorm.io/datatypes"
)

// MigrationAdmin 生成后台管理系列表
func MigrationAdmin() {
	// 权限初始化
	initialize.DB.AutoMigrate(&model.AdminPermissions{})
	// 角色初始
	initialize.DB.AutoMigrate(&model.AdminRoles{})
	initialize.DB.AutoMigrate(&model.AdminRolePermissions{})

	if initialize.DB.Where("code = ?", "admin").Find(&[]model.AdminRoles{}).RowsAffected != 1 {
		role := &model.AdminRoles{Code: "admin", Name: "超级管理员", Slug: "初始化系统超级管理员账号，拥有所有权限"}
		err := initialize.DB.Create(role).Error
		if err != nil {
			color.Warn.Printf("[Mysql]--> admin_roles 表的初始数据失败,err: %v\n", err)
			os.Exit(0)
		}
	}
	// 管理员初始
	initialize.DB.AutoMigrate(&model.AdminUsers{})
	initialize.DB.Model(&model.AdminUsers{}).AddUniqueIndex("idx_admin_user_username", "username")

	if initialize.DB.Where("username = ?", "admin").Find(&[]model.AdminUsers{}).RowsAffected != 1 {
		// 密码 12345678
		admin := &model.AdminUsers{Username: "admin", Password: "$2a$12$G9Pm80QiaRXFjVCiS6nDA..y8yW7OmbdboB48n7OtpgD3MVe1xk/S", Name: "超级管理员", Avatar: "http://qmplusimg.henrongyi.top/gva_header.jpg", RememberToken: "", ExtInfo: datatypes.JSON([]byte(`{"roles":["admin"]}`)), CreatedAt: time.Now(), UpdatedAt: time.Now()}
		err := initialize.DB.Create(admin).Error
		if err != nil {
			color.Warn.Printf("[Mysql]--> admin_users 表的初始数据失败,err: %v\n", err)
			os.Exit(0)
		}
	}
}
