package model

import (
	"github.com/jinzhu/gorm"
)

// AdminPermissions 权限表
type AdminPermissions struct {
	gorm.Model
	Name       string `grom:"size:32"`
	Slug       string `grom:"size:64"`
	HTTPMethod string `grom:"size:128"`
	HTTPPath   string `grom:"type:text"`
}
