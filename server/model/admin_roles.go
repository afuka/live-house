package model

import (
	"github.com/jinzhu/gorm"
)

// AdminRoles 后台用户角色表
type AdminRoles struct {
	gorm.Model
	Code string `grom:"size:32"`
	Name string `grom:"size:32"`
	Slug string `grom:"size:64"`
}
