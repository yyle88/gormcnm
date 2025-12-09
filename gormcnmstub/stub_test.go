// Package gormcnmstub_test validates stub implementation with integration tests
// Auto verifies stub wrapping functions with actual GORM database operations
// Tests examine condition logic, select statements, and column operations
//
// gormcnmstub_test 包通过集成测试验证存根实现
// 自动使用真实 GORM 数据库操作验证存根包装函数
// 测试涵盖查询条件、select 语句和列操作
package gormcnmstub_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yyle88/done"
	"github.com/yyle88/gormcnm"
	"github.com/yyle88/gormcnm/gormcnmstub"
	"github.com/yyle88/gormcnm/internal/tests"
	"github.com/yyle88/neatjson/neatjsons"
	"gorm.io/gorm"
)

func TestExample000(t *testing.T) {
	type Example struct {
		Name string `gorm:"primary_key;type:varchar(100);"`
		Type string `gorm:"column:type;"`
		Rank int    `gorm:"column:rank;"`
	}

	const (
		columnName = gormcnm.ColumnName[string]("name")
		columnType = gormcnm.ColumnName[string]("type")
		columnRank = gormcnm.ColumnName[int]("rank")
	)

	tests.NewDBRun(t, func(db *gorm.DB) {
		done.Done(db.AutoMigrate(&Example{}))
		done.Done(db.Save(&Example{Name: "abc", Type: "xyz", Rank: 123}).Error)
		done.Done(db.Save(&Example{Name: "aaa", Type: "xxx", Rank: 456}).Error)

		{
			result := db.Model(&Example{}).Where(
				gormcnmstub.NewQx(
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
			stmt := db.Model(&Example{})
			stmt = gormcnmstub.Where(stmt, gormcnmstub.Qx(columnName.Eq("aaa")).
				AND(
					gormcnmstub.Qx(columnType.Eq("xxx")),
					gormcnmstub.Qx(columnRank.Eq(456)),
				).
				OR1(columnName.Eq("abc")),
			)
			var results []*Example
			require.NoError(t, stmt.Find(&results).Error)
			t.Log(neatjsons.S(results))
		}
		{
			stmt := db.Model(&Example{})
			stmt = gormcnmstub.Where(stmt, gormcnmstub.Qx(columnName.Eq("aaa")).
				AND(
					gormcnmstub.Qx(columnType.Eq("xxx")),
					gormcnmstub.Qx(columnRank.Eq(456)),
				),
			)
			result := gormcnmstub.UpdateColumns(stmt, gormcnmstub.NewKw().Kw(columnRank.Kv(100)).Kw(columnType.Kv("uvw")))
			require.NoError(t, result.Error)
			require.Equal(t, int64(1), result.RowsAffected)
		}
		{
			stmt := db.Model(&Example{})
			stmt = gormcnmstub.Where(stmt, gormcnmstub.Qx(columnName.Eq("abc")).OR(gormcnmstub.Qx(columnName.Eq("aaa"))))
			stmt = gormcnmstub.OrderByColumns(stmt, columnRank.Ob("asc"))

			var examples []*Example
			result := stmt.Find(&examples)
			require.NoError(t, result.Error)
			require.Equal(t, 2, len(examples))
			require.Equal(t, 100, examples[0].Rank)
			require.Equal(t, 200, examples[1].Rank)
			t.Log(neatjsons.S(examples))
		}
	})
}

func TestExample001(t *testing.T) {
	type Example struct {
		Name string `gorm:"primary_key;type:varchar(100);"`
		Type string `gorm:"column:type;"`
		Rank int    `gorm:"column:rank;"`
	}

	const columnName = gormcnm.ColumnName[string]("name")

	tests.NewDBRun(t, func(db *gorm.DB) {
		done.Done(db.AutoMigrate(&Example{}))
		done.Done(db.Save(&Example{Name: "abc", Type: "xyz", Rank: 123}).Error)
		done.Done(db.Save(&Example{Name: "aaa", Type: "xxx", Rank: 456}).Error)

		type resType struct {
			Who string
			Cnt int64
		}

		{
			var results []resType
			require.NoError(t, db.Model(&Example{}).
				Group(columnName.Name()).
				Select(gormcnmstub.MergeStmts(
					columnName.AsAlias("who"),
					gormcnmstub.CountStmt("cnt"),
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
			require.NoError(t, db.Model(&Example{}).
				Group(columnName.Name()).
				Select(gormcnmstub.MergeStmts(
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
			require.NoError(t, db.Model(&Example{}).
				Group(columnName.Name()).
				Select(gormcnmstub.MergeStmts(
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
	})
}

func TestExample002(t *testing.T) {
	type Example struct {
		Name string `gorm:"primary_key;type:varchar(100);"`
		Type string `gorm:"column:type;"`
		Rank int    `gorm:"column:rank;"`
	}

	tests.NewDBRun(t, func(db *gorm.DB) {
		done.Done(db.AutoMigrate(&Example{}))
		done.Done(db.Save(&Example{Name: "abc", Type: "xyz", Rank: 123}).Error)
		done.Done(db.Save(&Example{Name: "aaa", Type: "xxx", Rank: 456}).Error)

		type resType struct {
			Cnt int64
		}

		var res resType
		require.NoError(t, db.Model(&Example{}).Select(gormcnmstub.CountStmt("cnt")).Find(&res).Error)
		require.Equal(t, int64(2), res.Cnt)
	})
}

func TestExample003(t *testing.T) {
	type Example struct {
		Name string `gorm:"primary_key;type:varchar(100);"`
		Type string `gorm:"column:type;"`
		Rank int    `gorm:"column:rank;"`
	}

	const (
		columnName = gormcnm.ColumnName[string]("name")
		columnType = gormcnm.ColumnName[string]("type")
	)

	tests.NewDBRun(t, func(db *gorm.DB) {
		done.Done(db.AutoMigrate(&Example{}))
		done.Done(db.Save(&Example{Name: "abc", Type: "xyz", Rank: 123}).Error)
		done.Done(db.Save(&Example{Name: "aaa", Type: "xxx", Rank: 456}).Error)

		type resType struct {
			Cnt int64
		}

		{
			var res resType
			stmt := gormcnmstub.CountCaseWhenStmt(columnName.Name()+"="+"'aaa'", "cnt")
			t.Log(stmt)
			require.NoError(t, db.Model(&Example{}).Select(stmt).Find(&res).Error)
			require.Equal(t, int64(1), res.Cnt)
		}

		{
			var res resType
			stmt := gormcnmstub.CountCaseWhenStmt(columnName.Name()+"="+"'aaa'"+" AND "+columnType.Name()+"="+"'xxx'", "cnt")
			t.Log(stmt)
			require.NoError(t, db.Model(&Example{}).Select(stmt).Find(&res).Error)
			require.Equal(t, int64(1), res.Cnt)
		}

		{
			var results []*Example
			var qx = gormcnmstub.NewQx(columnName.Eq("aaa")).AND1(columnType.Eq("xxx"))
			t.Log(qx.Qs())
			require.NoError(t, db.Model(&Example{}).Where(qx.Qx2()).Find(&results).Error)
			t.Log(len(results))
		}

		{
			var res resType
			var qx = gormcnmstub.NewQx(columnName.Eq("aaa")).AND1(columnType.Eq("xxx"))
			t.Log(qx.Qs())
			var sx = gormcnmstub.CountCaseWhenQxSx(qx, "cnt")
			t.Log(sx.Qs())
			require.NoError(t, db.Model(&Example{}).Select(sx.Qx2()).Find(&res).Error)
			require.Equal(t, int64(1), res.Cnt)
		}

		{
			var res resType
			var qx = gormcnmstub.NewQx(columnName.Eq("aaa")).AND1(columnType.Eq("xxx"))
			t.Log(qx.Qs())
			var sx = gormcnmstub.CountCaseWhenQxSx(qx, "cnt")
			t.Log(sx.Qs())
			db = db.Model(&Example{})
			db = gormcnmstub.Select(db, sx)
			require.NoError(t, db.Find(&res).Error)
			require.Equal(t, int64(1), res.Cnt)
		}
	})
}
