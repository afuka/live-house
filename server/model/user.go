package model

import (
	"golang.org/x/crypto/bcrypt"
)

// User 用户模型
type User struct {
	ID       int    `gorm:"size:10"`
	Username string `gorm:"size:32"`
	Password string `gorm:"size:64"`
	Nickname string `gorm:"size:32"`
	Avatar   string `gorm:"size:128"`
	Email    string `gorm:"size:128"`
	Mobile   string `gorm:"size:11"`
	Status   bool
}

const (
	// PassWordCost 密码加密难度
	PassWordCost = 12
)

// SetPassword 设置密码
func (user *User) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), PassWordCost)
	if err != nil {
		return err
	}
	user.Password = string(bytes)
	return nil
}

// CheckPassword 校验密码
func (user *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	return err == nil
}
