package gormcnm

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yyle88/gormcnm/internal/utils"
	"github.com/yyle88/neatjson/neatjsons"
	"gorm.io/gorm"
)

func TestColumnName_Op(t *testing.T) {
	type Example struct {
		Name string `gorm:"primary_key;type:varchar(100);"`
		Type string `gorm:"column:type;"`
	}

	const columnName = ColumnName[string]("name")

	utils.CaseRunInMemDB(func(db *gorm.DB) {
		require.NoError(t, db.AutoMigrate(&Example{}))
		require.NoError(t, db.Save(&Example{Name: "abc", Type: "xyz"}).Error)
		require.NoError(t, db.Save(&Example{Name: "aaa", Type: "xxx"}).Error)

		{
			var one Example
			require.NoError(t, db.Where(columnName.Op("=?", "abc")).First(&one).Error)
			require.Equal(t, "abc", one.Name)
			t.Log(neatjsons.S(one))
		}
		{
			var one Example
			require.NoError(t, db.Where(columnName.Eq("abc")).First(&one).Error)
			require.Equal(t, "abc", one.Name)
			t.Log(neatjsons.S(one))
		}
		{
			var one Example
			require.NoError(t, db.Where(columnName.BetweenAND("aba", "abd")).First(&one).Error)
			require.Equal(t, "abc", one.Name)
			t.Log(neatjsons.S(one))
		}
		{
			var one Example
			require.ErrorIs(t, gorm.ErrRecordNotFound, db.Where(columnName.IsNULL()).First(&one).Error)
			require.Equal(t, "", one.Name)
			t.Log(neatjsons.S(one))
		}
		{
			var one Example
			require.NoError(t, db.Where(columnName.IsNotNULL()).First(&one).Error)
			require.Contains(t, []string{"abc", "aaa"}, one.Name)
			t.Log(neatjsons.S(one))
		}
		{
			var res []*Example
			require.NoError(t, db.Where(columnName.In([]string{"abc", "aaa"})).Find(&res).Error)
			require.Contains(t, []string{"abc", "aaa"}, res[0].Name)
			require.Contains(t, []string{"abc", "aaa"}, res[1].Name)
			t.Log(neatjsons.S(res))
		}
		{
			var res []*Example
			require.NoError(t, db.Where(columnName.NotIn([]string{"aaa", "bbb"})).Find(&res).Error)
			for _, v := range res {
				require.NotEqual(t, "aaa", v.Name)
				require.NotEqual(t, "bbb", v.Name)
			}
			t.Log(neatjsons.S(res))
		}
	})
}

func TestColumnName_Op2(t *testing.T) {
	type Example struct {
		Name string `gorm:"primary_key;type:varchar(100);"`
		Type string `gorm:"column:type;"`
	}

	const (
		columnName = ColumnName[string]("name")
		columnType = ColumnName[string]("type")
	)

	utils.CaseRunInMemDB(func(db *gorm.DB) {
		require.NoError(t, db.AutoMigrate(&Example{}))
		require.NoError(t, db.Save(&Example{Name: "abc", Type: "xyz"}).Error)
		require.NoError(t, db.Save(&Example{Name: "aaa", Type: "xxx"}).Error)

		var one Example
		require.NoError(t, db.Where(columnName.Qs("=?")+" AND "+columnType.Qs("=?"), "abc", "xyz").First(&one).Error)
		require.Equal(t, "abc", one.Name)
		t.Log(neatjsons.S(one))
	})
}

func TestColumnName_Op3(t *testing.T) {
	type Example struct {
		Name string `gorm:"primary_key;type:varchar(100);"`
		Type string `gorm:"column:type;"`
	}

	const columnName = ColumnName[string]("name")

	utils.CaseRunInMemDB(func(db *gorm.DB) {
		require.NoError(t, db.AutoMigrate(&Example{}))
		require.NoError(t, db.Save(&Example{Name: "abc", Type: "xyz"}).Error)
		require.NoError(t, db.Save(&Example{Name: "aaa", Type: "xxx"}).Error)

		{
			var one Example
			require.NoError(t, db.Where(columnName.Like("%b%")).First(&one).Error)
			require.Equal(t, "abc", one.Name)
			t.Log(neatjsons.S(one))
		}
		{
			var one Example
			require.NoError(t, db.Where(columnName.NotLike("%b%")).First(&one).Error)
			require.Equal(t, "aaa", one.Name)
			t.Log(neatjsons.S(one))
		}
	})
}
