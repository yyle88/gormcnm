package utils

import (
	"github.com/yyle88/done"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// PtrX 给任何值取地址，特别是当参数为数字0，1，2，3或者字符串"a", "b", "c"的时候
func PtrX[T any](v T) *T {
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

func CaseRunInPrivateDB(run func(db *gorm.DB)) {
	db := done.VCE(gorm.Open(sqlite.Open("file::memory:?cache=private"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})).Nice()
	defer func() {
		done.Done(done.VCE(db.DB()).Nice().Close())
	}()
	run(db)
}
