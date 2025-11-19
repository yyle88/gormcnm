// Package gormcnm tests validate column value map operations for updates
// Auto verifies ColumnValueMap functionality with GORM Updates and Save operations
// Tests cover key-value mapping, batch updates, and value management
//
// gormcnm 测试包验证用于更新的列值映射操作
// 自动验证 ColumnValueMap 功能与 GORM Updates 和 Save 操作
// 测试涵盖键值映射、批量更新和值管理
package gormcnm

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yyle88/gormcnm/internal/utils"
	"gorm.io/gorm"
)

func TestValuesMap_SetValue(t *testing.T) {
	type Example struct {
		Name string `gorm:"primary_key;type:varchar(100);"`
		Type string `gorm:"column:type;"`
		Rank int    `gorm:"column:rank;"`
	}

	const (
		columnName = ColumnName[string]("name")
		columnType = ColumnName[string]("type")
		columnRank = ColumnName[int]("rank")
	)

	utils.InMemDB(func(db *gorm.DB) {
		require.NoError(t, db.AutoMigrate(&Example{}))
		require.NoError(t, db.Save(&Example{
			Name: "aaa",
			Type: "xxx",
			Rank: 123,
		}).Error)
		require.NoError(t, db.Save(&Example{
			Name: "bbb",
			Type: "yyy",
			Rank: 456,
		}).Error)

		{
			result := db.Model(&Example{}).Where(
				Qx(columnName.Eq("aaa")).
					AND(
						Qx(columnType.Eq("xxx")),
						Qx(columnRank.Eq(123)),
					).Qx3(),
			).UpdateColumns(columnRank.Kw(100).Kw(columnType.Kv("zzz")).Kws())
			require.NoError(t, result.Error)
			require.Equal(t, int64(1), result.RowsAffected)
		}
		{
			result := db.Model(&Example{}).Where(
				Qx(
					columnName.Eq("bbb"),
				).AND1(
					columnType.Eq("yyy"),
				).AND1(
					columnRank.Eq(456),
				).Qx3(),
			).UpdateColumns(columnRank.Kw(200).Kw(columnType.Kv("www")).Map())
			require.NoError(t, result.Error)
			require.Equal(t, int64(1), result.RowsAffected)
		}
	})
}
