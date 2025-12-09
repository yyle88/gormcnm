package tests

import (
	"fmt"
	"testing"

	"github.com/google/uuid"
	"github.com/yyle88/rese"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// NewDBRun runs a function with an in-memory SQLite database for testing purposes
// Auto creates a temporary database connection and handles cleanup after function execution
// NewDBRun 在内存数据库中运行函数，用于测试目的
// 自动创建临时数据库连接并在函数执行后处理清理工作
func NewDBRun(t *testing.T, run func(db *gorm.DB)) {
	dsn := fmt.Sprintf("file:db-%s?mode=memory&cache=shared", uuid.New().String())
	db := rese.P1(gorm.Open(sqlite.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	}))
	defer rese.F0(rese.P1(db.DB()).Close)

	t.Log("--- DB BEGIN ---")
	run(db)
	t.Log("--- DB CLOSE ---")
}
