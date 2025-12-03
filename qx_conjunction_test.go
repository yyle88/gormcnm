// Package gormcnm tests validate conjunction operations with parameter binding
// Auto verifies QxConjunction functionality with statement and arguments management
// Tests cover AND, OR combinations, parameter binding, and GORM WHERE clause integration
//
// gormcnm 测试包验证查询连接词操作，具备参数绑定功能
// 自动验证 QxConjunction 功能，包含语句和参数管理
// 测试涵盖 AND、OR 组合、参数绑定和 GORM WHERE 子句集成
package gormcnm

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yyle88/gormcnm/internal/utils"
	"github.com/yyle88/neatjson/neatjsons"
	"gorm.io/gorm"
)

func TestColumnQx_AND(t *testing.T) {
	type Example struct {
		Name string `gorm:"primary_key;type:varchar(100);"`
		Type string `gorm:"column:type;"`
	}

	const (
		columnName = ColumnName[string]("name")
		columnType = ColumnName[string]("type")
	)

	utils.InMemDB(func(db *gorm.DB) {
		require.NoError(t, db.AutoMigrate(&Example{}))
		require.NoError(t, db.Save(&Example{Name: "abc", Type: "xyz"}).Error)
		require.NoError(t, db.Save(&Example{Name: "aaa", Type: "xxx"}).Error)

		{
			var one Example
			qx := columnName.Qx("=?", "abc").AND(columnType.Qx("=?", "xyz"))
			t.Log(qx.Qs())
			t.Log(qx.Args())
			require.NoError(t, db.Where(qx.Qs(), qx.Args()...).First(&one).Error)
			require.Equal(t, "abc", one.Name)
			t.Log(neatjsons.S(one))
		}
		{
			var res []*Example
			qx := columnName.Qx("=?", "abc").OR(columnName.Qx("=?", "aaa"))
			t.Log(qx.Qs())
			t.Log(qx.Args())
			require.NoError(t, db.Where(qx.Qs(), qx.Args()...).Find(&res).Error)
			require.Contains(t, []string{"abc", "aaa"}, res[0].Name)
			require.Contains(t, []string{"abc", "aaa"}, res[1].Name)
			t.Log(neatjsons.S(res))
		}
		{
			var one Example
			qx := columnName.Qx("=?", "abc").NOT()
			t.Log(qx.Qs())
			t.Log(qx.Args())
			require.NoError(t, db.Where(qx.Qs(), qx.Args()).First(&one).Error)
			require.NotEqual(t, "abc", one.Name)
			t.Log(neatjsons.S(one))
		}
	})
}

func TestColumnQx_AND_2(t *testing.T) {
	type Example struct {
		Name string `gorm:"primary_key;type:varchar(100);"`
		Type string `gorm:"column:type;"`
	}

	const (
		columnName = ColumnName[string]("name")
		columnType = ColumnName[string]("type")
	)

	utils.InMemDB(func(db *gorm.DB) {
		require.NoError(t, db.AutoMigrate(&Example{}))
		require.NoError(t, db.Save(&Example{Name: "abc", Type: "xyz"}).Error)
		require.NoError(t, db.Save(&Example{Name: "aaa", Type: "xxx"}).Error)

		{
			var one Example
			require.NoError(t, db.Where(columnName.Qx("=?", "abc").AND(columnType.Qx("=?", "xyz")).Qx2()).First(&one).Error)
			require.Equal(t, "abc", one.Name)
			t.Log(neatjsons.S(one))
		}
		{
			var res []*Example
			require.NoError(t, db.Where(columnName.Qx("=?", "abc").OR(columnName.Qx("=?", "aaa")).Qx2()).Find(&res).Error)
			require.Contains(t, []string{"abc", "aaa"}, res[0].Name)
			require.Contains(t, []string{"abc", "aaa"}, res[1].Name)
			t.Log(neatjsons.S(res))
		}
		{
			var one Example
			require.NoError(t, db.Where(columnName.Qx("=?", "abc").NOT().Qx1()).First(&one).Error)
			require.NotEqual(t, "abc", one.Name)
			t.Log(neatjsons.S(one))
		}
	})
}

func TestColumnQx_AND_3(t *testing.T) {
	type Example struct {
		Name string `gorm:"primary_key;type:varchar(100);"`
		Type string `gorm:"column:type;"`
	}

	const (
		columnName = ColumnName[string]("name")
		columnType = ColumnName[string]("type")
	)

	utils.InMemDB(func(db *gorm.DB) {
		require.NoError(t, db.AutoMigrate(&Example{}))
		require.NoError(t, db.Save(&Example{Name: "abc", Type: "xyz"}).Error)
		require.NoError(t, db.Save(&Example{Name: "aaa", Type: "xxx"}).Error)

		{
			var one Example
			require.NoError(t, db.Where(NewQx(columnName.Eq("abc")).AND(NewQx(columnType.Eq("xyz"))).Qx2()).First(&one).Error)
			require.Equal(t, "abc", one.Name)
			t.Log(neatjsons.S(one))
		}
		{
			var res []*Example
			require.NoError(t, db.Where(NewQx(columnName.Eq("abc")).OR(NewQx(columnName.Eq("aaa"))).Qx2()).Find(&res).Error)
			require.Contains(t, []string{"abc", "aaa"}, res[0].Name)
			require.Contains(t, []string{"abc", "aaa"}, res[1].Name)
			t.Log(neatjsons.S(res))
		}
		{
			var one Example
			require.NoError(t, db.Where(NewQx(columnName.Eq("abc")).NOT().Qx1()).First(&one).Error)
			require.NotEqual(t, "abc", one.Name)
			t.Log(neatjsons.S(one))
		}
	})
}

func TestColumnQx_AND_4(t *testing.T) {
	type Example struct {
		Name string `gorm:"primary_key;type:varchar(100);"`
		Type string `gorm:"column:type;"`
	}

	const (
		columnName = ColumnName[string]("name")
		columnType = ColumnName[string]("type")
	)

	utils.InMemDB(func(db *gorm.DB) {
		require.NoError(t, db.AutoMigrate(&Example{}))
		require.NoError(t, db.Save(&Example{Name: "abc", Type: "xyz"}).Error)
		require.NoError(t, db.Save(&Example{Name: "aaa", Type: "xxx"}).Error)

		{
			var one Example
			require.NoError(t, db.Where(Qx(columnName.Eq("abc")).AND(Qx(columnType.Eq("xyz"))).Qx2()).First(&one).Error)
			require.Equal(t, "abc", one.Name)
			t.Log(neatjsons.S(one))
		}
		{
			var res []*Example
			require.NoError(t, db.Where(Qx(columnName.Eq("abc")).OR(Qx(columnName.Eq("aaa"))).Qx2()).Find(&res).Error)
			require.Contains(t, []string{"abc", "aaa"}, res[0].Name)
			require.Contains(t, []string{"abc", "aaa"}, res[1].Name)
			t.Log(neatjsons.S(res))
		}
		{
			var one Example
			require.NoError(t, db.Where(Qx(columnName.Eq("abc")).NOT().Qx1()).First(&one).Error)
			require.NotEqual(t, "abc", one.Name)
			t.Log(neatjsons.S(one))
		}
	})
}

func TestColumnQx_AND_5(t *testing.T) {
	type Example struct {
		Name string `gorm:"primary_key;type:varchar(100);"`
		Type string `gorm:"column:type;"`
	}

	const (
		columnName = ColumnName[string]("name")
		columnType = ColumnName[string]("type")
	)

	utils.InMemDB(func(db *gorm.DB) {
		require.NoError(t, db.AutoMigrate(&Example{}))
		require.NoError(t, db.Save(&Example{Name: "abc", Type: "xyz"}).Error)
		require.NoError(t, db.Save(&Example{Name: "aaa", Type: "xxx"}).Error)

		{
			var one Example
			require.NoError(t, db.Where(Qx(columnName.BetweenAND("aba", "abd")).
				AND(Qx(columnType.IsNotNULL())).Qx2()).
				First(&one).Error)
			require.Equal(t, "abc", one.Name)
			t.Log(neatjsons.S(one))
		}
	})
}
