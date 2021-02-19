package model

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

// User 用户模型
type Users struct {
	ID        int       `gorm:"size:10"`
	Username  string    `gorm:"size:32"`
	Password  string    `gorm:"size:64"`
	Nickname  string    `gorm:"size:32"`
	Avatar    string    `gorm:"size:128"`
	Email     string    `gorm:"size:128"`
	Mobile    string    `gorm:"size:11"`
	ExtInfo   string    `gorm:"type:text"`
	CreatedAt time.Time `gorm:"null"`
	UpdatedAt time.Time `gorm:"null"`
	Status    bool      `gorm:"default:1"`
}

// SetPassword 设置密码
func (user *Users) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return err
	}
	user.Password = string(bytes)
	return nil
}

// CheckPassword 校验密码
func (user *Users) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	return err == nil
}
