// Package gormcnm tests validate COALESCE and IFNULL operations for NULL-safe aggregates
// Auto verifies CoalesceNonNullGuardian functionality with SUM, AVG, MAX, MIN operations
// Tests cover NULL value protection, default value handling, and MySQL/standard SQL compatibility
//
// gormcnm 测试包验证 COALESCE 和 IFNULL 操作，实现 NULL 安全的聚合函数
// 自动验证 CoalesceNonNullGuardian 功能，包含 SUM、AVG、MAX、MIN 操作
// 测试涵盖 NULL 值保护、默认值处理和 MySQL/标准 SQL 兼容性
package gormcnm

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yyle88/gormcnm/internal/tests"
	"github.com/yyle88/rese"
)

func TestCoalesceStmt(t *testing.T) {
	type Example struct {
		Name string `gorm:"primary_key;type:varchar(100);"`
		Rank int    `gorm:"column:rank;"`
	}

	const columnRank = ColumnName[int]("rank")

	db := tests.NewMemDB(t)
	defer rese.F0(rese.P1(db.DB()).Close)

	require.NoError(t, db.AutoMigrate(&Example{}))
	require.NoError(t, db.Save(&Example{
		Name: "aaa",
		Rank: 123,
	}).Error)
	require.NoError(t, db.Save(&Example{
		Name: "bbb",
		Rank: 456,
	}).Error)

	t.Run("case-1", func(t *testing.T) {
		var value int
		err := db.Model(&Example{}).Select(columnRank.IFNULLFN().MaxStmt("max_rank")).First(&value).Error
		require.NoError(t, err)
		require.Equal(t, 456, value)
	})

	t.Run("case-2", func(t *testing.T) {
		var value int
		err := db.Model(&Example{}).Select(columnRank.COALESCE().MaxStmt("")).First(&value).Error
		require.NoError(t, err)
		require.Equal(t, 456, value)
	})
	t.Run("case-3", func(t *testing.T) {
		type resType struct {
			Value int
		}
		var res resType
		err := db.Model(&Example{}).Select(columnRank.COALESCE().MinStmt("value")).First(&res).Error
		require.NoError(t, err)
		require.Equal(t, 123, res.Value)
	})
	t.Run("case-4", func(t *testing.T) {
		type resType struct {
			Value int
		}
		var res resType
		err := db.Model(&Example{}).Select(columnRank.COALESCE().SumStmt("value")).First(&res).Error
		require.NoError(t, err)
		require.Equal(t, 579, res.Value)
	})
	t.Run("case-5", func(t *testing.T) {
		var value float64
		err := db.Model(&Example{}).Select(columnRank.COALESCE().AvgStmt("alias")).First(&value).Error
		require.NoError(t, err)
		require.Equal(t, 289.5, value)
	})
}
