package gormcnm_test

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yyle88/gormcnm"
	"github.com/yyle88/gormcnm/utilsgormcnm"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type ExampleOutPackage struct {
	Name string `gorm:"primary_key;type:varchar(100);"`
	Type string `gorm:"column:type;"`
	Rank int    `gorm:"column:rank;"`
}

var caseDB *gorm.DB
var onceIt sync.Once

// 这还是一个问题
// 我想测试在其他包调用工具的场景，因此我在这个测试包名称增加_test的后缀
// 这时很明显，再定义的重名的全局变量，已经不受影响
// 而且我再次定义 TestMain 时 IDE 也没报错，但当我运行测试时却报错
// 错误信息时：multiple definitions of TestMain
// 啊哈！由此推测go在编译测试时应该是遍历一个目录找 TestMain 的，假如发现两个就会报错
// 因此我们需要做一些设计
// 当你运行这个测试文件时，你会发现另一个 TestMain 里的逻辑也被执行
// 当你测试整个包的时候，在另一个包的 TestMain 确实只会被运行一次，而不是两次
// 这个问题是无关痛痒的，因此也没必要给go官方提问题
func initOnce() {
	onceIt.Do(func() {
		db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		})
		utilsgormcnm.AssertDone(err)
		caseDB = db

		utilsgormcnm.AssertDone(db.AutoMigrate(&ExampleOutPackage{}))
		utilsgormcnm.AssertDone(caseDB.Save(&ExampleOutPackage{Name: "abc", Type: "xyz", Rank: 123}).Error)
		utilsgormcnm.AssertDone(caseDB.Save(&ExampleOutPackage{Name: "aaa", Type: "xxx", Rank: 456}).Error)
	})
}

func TestFunctionOutPackage(t *testing.T) {
	initOnce()

	c := &gormcnm.ColumnBaseFuncClass{}
	columnName := gormcnm.ColumnName[string]("name")
	columnType := gormcnm.ColumnName[string]("type")
	columnRank := gormcnm.ColumnName[int]("rank")

	{
		result := caseDB.Model(&ExampleOutPackage{}).Where(
			c.NewQx(
				columnName.Eq("abc"),
			).AND1(
				columnType.Eq("xyz"),
			).AND1(
				columnRank.Eq(123),
			).Qx3(),
		).UpdateColumns(columnRank.Kw(200).Kw(columnType.Kv("www")).Map())
		require.NoError(t, result.Error)
		require.Equal(t, int64(1), result.RowsAffected)
	}
	{
		stmt := caseDB.Model(&ExampleOutPackage{})
		stmt = c.Where(stmt, c.Qx(columnName.Eq("aaa")).
			AND(
				c.Qx(columnType.Eq("xxx")),
				c.Qx(columnRank.Eq(456)),
			),
		)
		result := c.UpdateColumns(stmt, c.NewKw().Kw(columnRank.Kv(100)).Kw(columnType.Kv("uvw")))
		require.NoError(t, result.Error)
		require.Equal(t, int64(1), result.RowsAffected)
	}
	{
		stmt := caseDB.Model(&ExampleOutPackage{})
		stmt = c.Where(stmt, c.Qx(columnName.Eq("abc")).OR(c.Qx(columnName.Eq("aaa"))))
		stmt = c.Order(stmt, columnRank.Ob("asc"))

		var examples []*ExampleOutPackage
		result := stmt.Find(&examples)
		require.NoError(t, result.Error)
		require.Equal(t, 2, len(examples))
		require.Equal(t, 100, examples[0].Rank)
		require.Equal(t, 200, examples[1].Rank)
		t.Log(utilsgormcnm.SoftNeatString(examples))
	}
}
