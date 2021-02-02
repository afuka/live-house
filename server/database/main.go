package database

import (
	"afkser/database/migrations"
	"fmt"
)

// Migration 生成表
func Migration() {
	fmt.Println("Migration------开始执行")
	migrations.MigrationUser()
	fmt.Println("Migration------运行结束")
}

// Seed 生成假数据
func Seed() {

}
