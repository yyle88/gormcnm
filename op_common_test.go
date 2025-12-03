// Package gormcnm_test provides comprehensive integration tests with gormcnm features
// Auto validates column operations, where generation, and GORM integration patterns
// Tests demonstrate real-world usage scenarios with an SQLite memory-based database
//
// gormcnm_test 包为 gormcnm 功能提供全面的集成测试
// 自动验证列操作、查询生成和 GORM 集成模式
// 测试展示了使用 SQLite 内存数据库的实际使用场景
package gormcnm_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yyle88/done"
	"github.com/yyle88/gormcnm"
	"github.com/yyle88/gormcnm/internal/utils"
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

	operation := &gormcnm.ColumnOperationClass{}

	utils.InMemDB(func(db *gorm.DB) {
		done.Done(db.AutoMigrate(&Example{}))
		done.Done(db.Save(&Example{Name: "abc", Type: "xyz", Rank: 123}).Error)
		done.Done(db.Save(&Example{Name: "aaa", Type: "xxx", Rank: 456}).Error)

		{
			result := db.Model(&Example{}).Where(
				operation.NewQx(
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
			stmt = operation.Where(stmt, operation.Qx(columnName.Eq("aaa")).
				AND(
					operation.Qx(columnType.Eq("xxx")),
					operation.Qx(columnRank.Eq(456)),
				).
				OR1(columnName.Eq("abc")),
			)
			var results []*Example
			require.NoError(t, stmt.Find(&results).Error)
			t.Log(neatjsons.S(results))
		}
		{
			stmt := db.Model(&Example{})
			stmt = operation.Where(stmt, operation.Qx(columnName.Eq("aaa")).
				AND(
					operation.Qx(columnType.Eq("xxx")),
					operation.Qx(columnRank.Eq(456)),
				),
			)
			result := operation.UpdateColumns(stmt, operation.NewKw().Kw(columnRank.Kv(100)).Kw(columnType.Kv("uvw")))
			require.NoError(t, result.Error)
			require.Equal(t, int64(1), result.RowsAffected)
		}
		{
			stmt := db.Model(&Example{})
			stmt = operation.Where(stmt, operation.Qx(columnName.Eq("abc")).OR(operation.Qx(columnName.Eq("aaa"))))
			stmt = operation.OrderByColumns(stmt, columnRank.Ob("asc"))

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

	operation := &gormcnm.ColumnOperationClass{}

	utils.InMemDB(func(db *gorm.DB) {
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
				Select(operation.MergeStmts(
					columnName.AsAlias("who"),
					operation.CountStmt("cnt"),
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
				Select(operation.MergeStmts(
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
				Select(operation.MergeStmts(
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

	operation := &gormcnm.ColumnOperationClass{}

	utils.InMemDB(func(db *gorm.DB) {
		done.Done(db.AutoMigrate(&Example{}))
		done.Done(db.Save(&Example{Name: "abc", Type: "xyz", Rank: 123}).Error)
		done.Done(db.Save(&Example{Name: "aaa", Type: "xxx", Rank: 456}).Error)

		type resType struct {
			Cnt int64
		}

		var res resType
		require.NoError(t, db.Model(&Example{}).Select(operation.CountStmt("cnt")).Find(&res).Error)
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

	operation := &gormcnm.ColumnOperationClass{}

	utils.InMemDB(func(db *gorm.DB) {
		done.Done(db.AutoMigrate(&Example{}))
		done.Done(db.Save(&Example{Name: "abc", Type: "xyz", Rank: 123}).Error)
		done.Done(db.Save(&Example{Name: "aaa", Type: "xxx", Rank: 456}).Error)

		type resType struct {
			Cnt int64
		}

		{
			var res resType
			stmt := operation.CountCaseWhenStmt(columnName.Name()+"="+"'aaa'", "cnt")
			t.Log(stmt)
			require.NoError(t, db.Model(&Example{}).Select(stmt).Find(&res).Error)
			require.Equal(t, int64(1), res.Cnt)
		}

		{
			var res resType
			stmt := operation.CountCaseWhenStmt(columnName.Name()+"="+"'aaa'"+" AND "+columnType.Name()+"="+"'xxx'", "cnt")
			t.Log(stmt)
			require.NoError(t, db.Model(&Example{}).Select(stmt).Find(&res).Error)
			require.Equal(t, int64(1), res.Cnt)
		}

		{
			var results []*Example
			var qx = operation.NewQx(columnName.Eq("aaa")).AND1(columnType.Eq("xxx"))
			t.Log(qx.Qs())
			require.NoError(t, db.Model(&Example{}).Where(qx.Qx2()).Find(&results).Error)
			t.Log(len(results))
		}

		{
			var res resType
			var qx = operation.NewQx(columnName.Eq("aaa")).AND1(columnType.Eq("xxx"))
			t.Log(qx.Qs())
			var sx = operation.CountCaseWhenQxSx(qx, "cnt")
			t.Log(sx.Qs())
			require.NoError(t, db.Model(&Example{}).Select(sx.Qx2()).Find(&res).Error)
			require.Equal(t, int64(1), res.Cnt)
		}

		{
			var res resType
			var qx = operation.NewQx(columnName.Eq("aaa")).AND1(columnType.Eq("xxx"))
			t.Log(qx.Qs())
			var sx = operation.CountCaseWhenQxSx(qx, "cnt")
			t.Log(sx.Qs())
			db = db.Model(&Example{})
			db = operation.Select(db, sx)
			require.NoError(t, db.Find(&res).Error)
			require.Equal(t, int64(1), res.Cnt)
		}
	})
}

func TestColumnOperationClass_MergeSlices(t *testing.T) {
	operation := &gormcnm.ColumnOperationClass{}

	slice1 := []string{"a", "b", "c"}
	slice2 := []string{"d", "e", "f"}
	slice3 := []string{"g", "h"}

	results := operation.MergeSlices(slice1, slice2, slice3)
	require.Equal(t, "a, b, c, d, e, f, g, h", results)
}
