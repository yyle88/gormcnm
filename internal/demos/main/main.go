package main

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/yyle88/done"
	"github.com/yyle88/gormcnm"
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
	db := done.VCE(gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})).Nice()

	//create example data
	_ = db.AutoMigrate(&Example{})
	_ = db.Save(&Example{Name: "abc", Type: "xyz", Rank: 123}).Error
	_ = db.Save(&Example{Name: "aaa", Type: "xxx", Rank: 456}).Error

	{
		var res Example
		err := db.Where("name=?", "abc").First(&res).Error
		done.Done(err)
		fmt.Println(res)
	}
	{ //select an example data
		var res Example
		if err := db.Where(columnName.Eq("abc")).
			Where(columnType.Eq("xyz")).
			Where(columnRank.Gt(100)).
			Where(columnRank.Lt(200)).
			First(&res).Error; err != nil {
			panic(errors.WithMessage(err, "wrong"))
		}
		fmt.Println(res)
	}
}
