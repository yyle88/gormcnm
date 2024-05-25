package gormcnm_test

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yyle88/done"
	"github.com/yyle88/gormcnm"
	"github.com/yyle88/gormcnm/internal/utils"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var caseDB2 *gorm.DB
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
func onceNewGorm() *gorm.DB {
	onceIt.Do(func() {
		db := done.VCE(gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		})).Nice()

		caseDB2 = db
	})
	return caseDB2
}

func TestExample000(t *testing.T) {
	db := onceNewGorm()

	type Example000 struct {
		Name string `gorm:"primary_key;type:varchar(100);"`
		Type string `gorm:"column:type;"`
		Rank int    `gorm:"column:rank;"`
	}

	done.Done(db.AutoMigrate(&Example000{}))
	done.Done(db.Save(&Example000{Name: "abc", Type: "xyz", Rank: 123}).Error)
	done.Done(db.Save(&Example000{Name: "aaa", Type: "xxx", Rank: 456}).Error)

	c := &gormcnm.ColumnOperationClass{}
	columnName := gormcnm.ColumnName[string]("name")
	columnType := gormcnm.ColumnName[string]("type")
	columnRank := gormcnm.ColumnName[int]("rank")

	{
		result := db.Model(&Example000{}).Where(
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
		stmt := db.Model(&Example000{})
		stmt = c.Where(stmt, c.Qx(columnName.Eq("aaa")).
			AND(
				c.Qx(columnType.Eq("xxx")),
				c.Qx(columnRank.Eq(456)),
			).
			OR1(columnName.Eq("abc")),
		)
		var results []*Example000
		require.NoError(t, stmt.Find(&results).Error)
		t.Log(utils.SoftNeatString(results))
	}
	{
		stmt := db.Model(&Example000{})
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
		stmt := db.Model(&Example000{})
		stmt = c.Where(stmt, c.Qx(columnName.Eq("abc")).OR(c.Qx(columnName.Eq("aaa"))))
		stmt = c.OrderByColumns(stmt, columnRank.Ob("asc"))

		var examples []*Example000
		result := stmt.Find(&examples)
		require.NoError(t, result.Error)
		require.Equal(t, 2, len(examples))
		require.Equal(t, 100, examples[0].Rank)
		require.Equal(t, 200, examples[1].Rank)
		t.Log(utils.SoftNeatString(examples))
	}
}

func TestExample001(t *testing.T) {
	db := onceNewGorm()

	type Example001 struct {
		Name string `gorm:"primary_key;type:varchar(100);"`
		Type string `gorm:"column:type;"`
		Rank int    `gorm:"column:rank;"`
	}

	done.Done(db.AutoMigrate(&Example001{}))
	done.Done(db.Save(&Example001{Name: "abc", Type: "xyz", Rank: 123}).Error)
	done.Done(db.Save(&Example001{Name: "aaa", Type: "xxx", Rank: 456}).Error)

	c := &gormcnm.ColumnOperationClass{}
	columnName := gormcnm.ColumnName[string]("name")

	type resType struct {
		Who string
		Cnt int64
	}

	{
		var results []resType
		require.NoError(t, db.Model(&Example001{}).
			Group(columnName.Name()).
			Select(c.MergeStmts(
				columnName.AsAlias("who"),
				c.CountStmt("cnt"),
			)).
			Find(&results).Error)
		require.Equal(t, 2, len(results))
		for _, one := range results {
			t.Log(one.Who, one.Cnt)
			require.NotEmpty(t, one.Who)
			require.Positive(t, one.Cnt)
		}
	}
	{
		var results []resType
		require.NoError(t, db.Model(&Example001{}).
			Group(columnName.Name()).
			Select(c.MergeStmts(
				columnName.AsAlias("who"),
				columnName.Count("cnt"),
			)).
			Find(&results).Error)
		require.Equal(t, 2, len(results))
		for _, one := range results {
			t.Log(one.Who, one.Cnt)
			require.NotEmpty(t, one.Who)
			require.Positive(t, one.Cnt)
		}
	}
	{
		var results []resType
		require.NoError(t, db.Model(&Example001{}).
			Group(columnName.Name()).
			Select(c.MergeStmts(
				columnName.AsAlias("who"),
				columnName.CountDistinct("cnt"),
			)).
			Find(&results).Error)
		require.Equal(t, 2, len(results))
		for _, one := range results {
			t.Log(one.Who, one.Cnt)
			require.NotEmpty(t, one.Who)
			require.Positive(t, one.Cnt)
		}
	}
}

func TestExample002(t *testing.T) {
	db := onceNewGorm()

	type Example002 struct {
		Name string `gorm:"primary_key;type:varchar(100);"`
		Type string `gorm:"column:type;"`
		Rank int    `gorm:"column:rank;"`
	}

	done.Done(db.AutoMigrate(&Example002{}))
	done.Done(db.Save(&Example002{Name: "abc", Type: "xyz", Rank: 123}).Error)
	done.Done(db.Save(&Example002{Name: "aaa", Type: "xxx", Rank: 456}).Error)

	c := &gormcnm.ColumnOperationClass{}

	type resType struct {
		Cnt int64
	}

	var res resType
	require.NoError(t, db.Model(&Example002{}).Select(c.CountStmt("cnt")).Find(&res).Error)
	require.Equal(t, int64(2), res.Cnt)
}

func TestExample003(t *testing.T) {
	db := onceNewGorm()

	type Example003 struct {
		Name string `gorm:"primary_key;type:varchar(100);"`
		Type string `gorm:"column:type;"`
		Rank int    `gorm:"column:rank;"`
	}

	done.Done(db.AutoMigrate(&Example003{}))
	done.Done(db.Save(&Example003{Name: "abc", Type: "xyz", Rank: 123}).Error)
	done.Done(db.Save(&Example003{Name: "aaa", Type: "xxx", Rank: 456}).Error)

	c := &gormcnm.ColumnOperationClass{}
	columnName := gormcnm.ColumnName[string]("name")
	columnType := gormcnm.ColumnName[string]("type")

	type resType struct {
		Cnt int64
	}

	{
		var res resType
		stmt := c.CountCaseWhenStmt(columnName.Name()+"="+"'aaa'", "cnt")
		t.Log(stmt)
		require.NoError(t, db.Model(&Example003{}).Select(stmt).Find(&res).Error)
		require.Equal(t, int64(1), res.Cnt)
	}

	{
		var res resType
		stmt := c.CountCaseWhenStmt(columnName.Name()+"="+"'aaa'"+" AND "+columnType.Name()+"="+"'xxx'", "cnt")
		t.Log(stmt)
		require.NoError(t, db.Model(&Example003{}).Select(stmt).Find(&res).Error)
		require.Equal(t, int64(1), res.Cnt)
	}

	{
		var results []*Example003
		var qx = c.NewQx(columnName.Eq("aaa")).AND1(columnType.Eq("xxx"))
		t.Log(qx.Qs())
		require.NoError(t, db.Model(&Example003{}).Where(qx.Qx2()).Find(&results).Error)
		t.Log(len(results))
	}

	{
		var res resType
		var qx *gormcnm.QxType = c.NewQx(columnName.Eq("aaa")).AND1(columnType.Eq("xxx"))
		t.Log(qx.Qs())
		var sx *gormcnm.SxType = c.CountCaseWhenQxSx(qx, "cnt")
		t.Log(sx.Qs())
		require.NoError(t, db.Model(&Example003{}).Select(sx.Qx2()).Find(&res).Error)
		require.Equal(t, int64(1), res.Cnt)
	}

	{
		var res resType
		var qx *gormcnm.QxType = c.NewQx(columnName.Eq("aaa")).AND1(columnType.Eq("xxx"))
		t.Log(qx.Qs())
		var sx *gormcnm.SxType = c.CountCaseWhenQxSx(qx, "cnt")
		t.Log(sx.Qs())
		db = db.Model(&Example003{})
		db = c.Select(db, sx)
		require.NoError(t, db.Find(&res).Error)
		require.Equal(t, int64(1), res.Cnt)
	}
}
