package model

import (
	"github.com/jinzhu/gorm"
)

// AdminRolePermissions 用户角色权限表
type AdminRolePermissions struct {
	gorm.Model
	RoleID       int `grom:"size:10"`
	PermissionID int `grom:"size:10"`
}
