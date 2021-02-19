package model

import (
	"encoding/json"
	"time"

	"afkser/facades"
	"afkser/utils"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/datatypes"
)

// AdminUsers 后台管理模型
type AdminUsers struct {
	ID            int            `gorm:"size:10"`
	Username      string         `gorm:"size:32"`
	Password      string         `gorm:"size:64" json:"-"`
	Name          string         `gorm:"size:32"`
	Avatar        string         `gorm:"size:128"`
	RememberToken string         `gorm:"size:100"`
	ExtInfo       datatypes.JSON `gorm:"column:ext_info" sql:"type:json"`
	CreatedAt     time.Time      `gorm:"null" json:"-"`
	UpdatedAt     time.Time      `gorm:"null" json:"-"`
	Status        bool           `gorm:"default:1"`
}

// AdminUsersExtInfo 序列华中存在的内容
type AdminUsersExtInfo struct {
	Roles []string `json:roles`
	// Province string   `json:province`
}

// GetUserByID 通过id 来查询用户
func (admin *AdminUsers) GetUserByID(id int) error {
	err := facades.DB.Where("id = ?", id).First(&admin).Error
	return err
}

// GetUserByUsername 通过id 来查询用户
func (admin *AdminUsers) GetUserByUsername(username string) error {
	err := facades.DB.Where("username = ?", username).First(&admin).Error
	return err
}

// GenerateToken 更新用户Token
func (admin *AdminUsers) GenerateToken() error {
	// 生成随机串
	token := utils.RandStringRunes(32)
	// 保存
	err := facades.DB.Where("id=?", admin.ID).First(&admin).Update("remember_token", token).Error

	return err
}

// SetPassword 设置密码
func (admin *AdminUsers) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return err
	}
	admin.Password = string(bytes)
	return nil
}

// CheckPassword 校验密码
func (admin *AdminUsers) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(password))
	return err == nil
}

// GetExtInfo 获取extInfo
func (admin *AdminUsers) GetExtInfo() (AdminUsersExtInfo, error) {
	extInfo := AdminUsersExtInfo{}
	s, err := admin.ExtInfo.Value()
	if err != nil {
		return extInfo, err
	}
	err = json.Unmarshal([]byte(s.(string)), &extInfo)
	if err != nil {
		return extInfo, err
	}

	return extInfo, err
}
