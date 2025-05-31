package utils

import (
	"github.com/yyle88/rese"
	"github.com/yyle88/tern"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// GetPointer 给任何值取地址，特别是当参数为数字0，1，2，3或者字符串"a", "b", "c"的时候
func GetPointer[T any](v T) *T {
	return &v
}

// VOr0 给任何地址取值，当是空地址时返回 zero 即类型默认的零值
func VOr0[T any](v *T) T {
	if v != nil {
		return *v
	} else {
		var zero T
		return zero
	}
}

// CaseRunInSqliteMemDB 在内存数据库中运行函数(但由于函数已经足够短，其实也没有封装的必要，但还是留着吧)
func CaseRunInSqliteMemDB(run func(db *gorm.DB)) {
	db := rese.P1(gorm.Open(sqlite.Open("file::memory:?cache=private"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	}))
	defer rese.F0(rese.P1(db.DB()).Close)

	run(db)
}

// ApplyAliasToColumn 设置别名，返回类似 COUNT(*) as cnt 这样的
func ApplyAliasToColumn(stmt string, alias string) string {
	return tern.BVV(alias != "", stmt+" as "+alias, stmt)
}
