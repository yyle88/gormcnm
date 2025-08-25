package utils

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/yyle88/rese"
	"github.com/yyle88/tern"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// GetValuePointer gets a pointer to any value, especially useful for literal values like numbers 0,1,2,3 or strings "a", "b", "c"
// GetValuePointer 给任何值取地址，特别是当参数为数字0，1，2，3或者字符串"a", "b", "c"的时候
func GetValuePointer[T any](v T) *T {
	return &v
}

// GetPointerValue dereferences a pointer and returns the value, returns zero value if pointer is nil
// GetPointerValue 给任何地址取值，当是空地址时返回 zero 即类型默认的零值
func GetPointerValue[T any](v *T) T {
	if v != nil {
		return *v
	} else {
		var zeroValue T
		return zeroValue
	}
}

// CaseRunInSqliteMemDB runs a function with an in-memory SQLite database for testing purposes
// Auto creates a temporary database connection and handles cleanup after function execution
// CaseRunInSqliteMemDB 在内存数据库中运行函数，用于测试目的
// 自动创建临时数据库连接并在函数执行后处理清理工作
func CaseRunInSqliteMemDB(run func(db *gorm.DB)) {
	dsn := fmt.Sprintf("file:db-%s?mode=memory&cache=shared", uuid.New().String())
	db := rese.P1(gorm.Open(sqlite.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	}))
	defer rese.F0(rese.P1(db.DB()).Close)

	run(db)
}

// ApplyAliasToColumn applies an alias to a column statement, returns format like "COUNT(*) as cnt"
// Auto appends alias if provided, otherwise returns original statement
// ApplyAliasToColumn 为列语句设置别名，返回类似 "COUNT(*) as cnt" 的格式
// 如果提供了别名就自动添加，否则返回原始语句
func ApplyAliasToColumn(stmt string, alias string) string {
	return tern.BVV(alias != "", stmt+" as "+alias, stmt)
}
