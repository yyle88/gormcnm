// Package main demonstrates advanced gormcnm usage with multi-condition WHERE queries
// Auto shows comparison with traditional GORM queries vs type-safe column operations
// Runs with SQLite IN-MEMORY database to showcase Gt, Lt, Eq operations
//
// main 包演示 gormcnm 多条件 WHERE 查询的高级用法
// 自动展示传统 GORM 查询与类型安全列操作的对比
// 使用 SQLite 内存数据库运行以展示 Gt、Lt、Eq 操作
package main

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/yyle88/done"
	"github.com/yyle88/gormcnm"
	"github.com/yyle88/must"
	"github.com/yyle88/rese"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Example is a gorm model define 3 fields(name, type, rank)
type Example struct {
	Name string `gorm:"primary_key;type:varchar(100);"`
	Type string `gorm:"column:type;"`
	Rank int    `gorm:"column:rank;"`
}

// Now define the fields enum vars(name, type rank)
const (
	columnName = gormcnm.ColumnName[string]("name")
	columnType = gormcnm.ColumnName[string]("type")
	columnRank = gormcnm.ColumnName[int]("rank")
)

func main() {
	//new db connection
	dsn := fmt.Sprintf("file:db-%s?mode=memory&cache=shared", uuid.New().String())
	db := rese.P1(gorm.Open(sqlite.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	}))
	defer rese.F0(rese.P1(db.DB()).Close)

	//create example data
	done.Done(db.AutoMigrate(&Example{}))
	done.Done(db.Save(&Example{Name: "abc", Type: "xyz", Rank: 123}).Error)
	done.Done(db.Save(&Example{Name: "aaa", Type: "xxx", Rank: 456}).Error)

	{
		//SELECT * FROM `examples` WHERE name="abc" ORDER BY `examples`.`name` LIMIT 1
		var res Example
		must.Done(db.Where("name=?", "abc").First(&res).Error)
		fmt.Println(res)
	}
	{
		//SELECT * FROM `examples` WHERE name="abc" AND type="xyz" AND rank>100 AND rank<200 ORDER BY `examples`.`name` LIMIT 1
		var res Example
		must.Done(db.Where(columnName.Eq("abc")).
			Where(columnType.Eq("xyz")).
			Where(columnRank.Gt(100)).
			Where(columnRank.Lt(200)).
			First(&res).Error)
		fmt.Println(res)
	}
}
