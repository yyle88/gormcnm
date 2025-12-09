// Package gormcnm tests validate statement and arguments tuple operations
// Auto verifies statementArgumentsTuple features with SQL statement and argument management
// Tests examine driver.Valuer implementation, argument binding, and GORM queries integration
//
// gormcnm 测试包验证语句和参数元组操作
// 自动验证 statementArgumentsTuple 功能，包含 SQL 语句和参数管理
// 测试涵盖 driver.Valuer 实现、参数绑定和 GORM 查询集成
package gormcnm

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yyle88/gormcnm/internal/tests"
	"gorm.io/gorm"
)

func Test_statementArgumentsTuple_Qx1(t *testing.T) {
	oneItem := newStatementArgumentsTuple("name=?", []any{"abc"})
	stmt, arg0 := oneItem.Qx1()
	require.Equal(t, "name=?", stmt)
	require.Equal(t, "abc", arg0)
}

func Test_statementArgumentsTuple_Qx2(t *testing.T) {
	type Example struct {
		Name string `gorm:"primary_key;type:varchar(100);"`
		Rank int    `gorm:"column:rank;"`
	}

	const (
		columnName = ColumnName[string]("name")
		columnRank = ColumnName[int]("rank")
	)

	tests.NewDBRun(t, func(db *gorm.DB) {
		require.NoError(t, db.AutoMigrate(&Example{}))
		require.NoError(t, db.Save(&Example{Name: "abc", Rank: 100}).Error)
		require.NoError(t, db.Save(&Example{Name: "aaa", Rank: 101}).Error)

		stmtItem := newStatementArgumentsTuple(columnName.Name()+"=?"+" AND "+columnRank.Name()+"=?", []any{"abc", 100})
		stmt, arg1, arg2 := stmtItem.Qx2()
		require.Equal(t, "name=? AND rank=?", stmt)
		require.Equal(t, "abc", arg1)
		require.Equal(t, 100, arg2)

		var count int64
		require.NoError(t, db.Model(&Example{}).Where(stmtItem.Qx2()).Count(&count).Error)
		require.Equal(t, int64(1), count)

		var example Example
		require.NoError(t, db.Where(stmtItem.Qx2()).First(&example).Error)
		require.Equal(t, "abc", example.Name)
		require.Equal(t, 100, example.Rank)
	})
}
