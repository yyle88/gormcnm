// Package gormcnm tests validate ordering operations and sort statement construction
// Auto verifies OrderByBottle features with GORM Order clauses
// Tests examine basic ordering, combined sorting, and SQL execution
//
// gormcnm 测试包验证排序操作和排序语句构建
// 自动验证 OrderByBottle 功能与 GORM Order 子句
// 测试涵盖基础排序、组合排序和查询执行
package gormcnm

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yyle88/gormcnm/internal/tests"
	"github.com/yyle88/neatjson/neatjsons"
	"gorm.io/gorm"
)

func TestOrderByBottle_Ob(t *testing.T) {
	type Example struct {
		Name string `gorm:"primary_key;type:varchar(100);"`
		Type string `gorm:"column:type;"`
	}

	const (
		columnName = ColumnName[string]("name")
		columnType = ColumnName[string]("type")
	)

	tests.NewDBRun(t, func(db *gorm.DB) {
		require.NoError(t, db.AutoMigrate(&Example{}))
		require.NoError(t, db.Save(&Example{Name: "abc", Type: "xyz"}).Error)
		require.NoError(t, db.Save(&Example{Name: "aaa", Type: "xxx"}).Error)

		{
			var res []*Example
			require.NoError(t, db.Where(columnName.In([]string{"abc", "aaa"})).
				Order(columnName.OrderByBottle("asc").
					Ob(columnType.Ob("desc")).
					Ox()).
				Find(&res).Error)
			require.Equal(t, "aaa", res[0].Name)
			require.Equal(t, "abc", res[1].Name)
			t.Log(neatjsons.S(res))
		}
		{
			var res []*Example
			require.NoError(t, db.Where(columnName.In([]string{"abc", "aaa"})).
				Order(columnName.Ob("desc").
					OrderByBottle(columnType.Ob("asc")).
					Orders()).
				Find(&res).Error)
			require.Equal(t, "abc", res[0].Name)
			require.Equal(t, "aaa", res[1].Name)
			t.Log(neatjsons.S(res))
		}
	})
}
