package gormcnm

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yyle88/gormcnm/internal/utils"
	"github.com/yyle88/neatjson/neatjsons"
	"gorm.io/gorm"
)

func TestOrderByBottle_Scope(t *testing.T) {
	type Example struct {
		Name string `gorm:"primary_key;type:varchar(100);"`
		Type string `gorm:"column:type;"`
	}

	const (
		columnName = ColumnName[string]("name")
		columnType = ColumnName[string]("type")
	)

	utils.CaseRunInSqliteMemDB(func(db *gorm.DB) {
		require.NoError(t, db.AutoMigrate(&Example{}))
		require.NoError(t, db.Save(&Example{Name: "abc", Type: "xyz"}).Error)
		require.NoError(t, db.Save(&Example{Name: "aaa", Type: "xxx"}).Error)

		{
			var res []*Example
			require.NoError(t, db.Where(columnName.In([]string{"abc", "aaa"})).
				Scopes(columnName.OrderByBottle("asc").
					Ob(columnType.Ob("desc")).
					Scope()).
				Find(&res).Error)
			require.Equal(t, "aaa", res[0].Name)
			require.Equal(t, "abc", res[1].Name)
			t.Log(neatjsons.S(res))
		}
		{
			var res []*Example
			require.NoError(t, db.Where(columnName.In([]string{"abc", "aaa"})).
				Scopes(columnName.Ob("desc").
					OrderByBottle(columnType.OrderByBottle("asc")).
					Scope()).
				Find(&res).Error)
			require.Equal(t, "abc", res[0].Name)
			require.Equal(t, "aaa", res[1].Name)
			t.Log(neatjsons.S(res))
		}
	})
}

func TestQxConjunction_Scope(t *testing.T) {
	type Example struct {
		Name string `gorm:"primary_key;type:varchar(100);"`
		Type string `gorm:"column:type;"`
	}

	const (
		columnName = ColumnName[string]("name")
		columnType = ColumnName[string]("type")
	)

	utils.CaseRunInSqliteMemDB(func(db *gorm.DB) {
		require.NoError(t, db.AutoMigrate(&Example{}))
		require.NoError(t, db.Save(&Example{Name: "abc", Type: "xyz"}).Error)
		require.NoError(t, db.Save(&Example{Name: "aaa", Type: "xxx"}).Error)

		{
			var one Example
			require.NoError(t, db.
				Scopes(Qx(columnName.BetweenAND("aba", "abd")).
					AND(Qx(columnType.IsNotNULL())).
					Scope()).
				First(&one).Error)
			require.Equal(t, "abc", one.Name)
			t.Log(neatjsons.S(one))
		}
	})
}

func TestQxConjunction_Scope_Example2(t *testing.T) {
	type Example struct {
		Name string `gorm:"primary_key;type:varchar(100);"`
		Type string `gorm:"column:type;"`
		Rank int    `gorm:"column:rank;"`
	}

	const (
		columnType = ColumnName[string]("type")
		columnRank = ColumnName[int]("rank")
	)

	utils.CaseRunInSqliteMemDB(func(db *gorm.DB) {
		require.NoError(t, db.AutoMigrate(&Example{}))
		require.NoError(t, db.Save(&Example{Name: "abc", Type: "xyz", Rank: 25}).Error)
		require.NoError(t, db.Save(&Example{Name: "aaa", Type: "xxx", Rank: 30}).Error)
		require.NoError(t, db.Save(&Example{Name: "def", Type: "yyy", Rank: 20}).Error)

		{
			var one Example
			require.NoError(t, db.
				Scopes(NewQx(columnRank.Gt(22)).
					AND(Qx(columnType.Ne("xxx"))).
					Scope()).
				First(&one).Error)
			require.Equal(t, "abc", one.Name)
			t.Log(neatjsons.S(one))
		}
	})
}

func TestSelectStatement_Scope(t *testing.T) {
	type Example struct {
		Name string `gorm:"primary_key;type:varchar(100);"`
		Rank int    `gorm:"column:rank;"`
	}

	const (
		columnName = ColumnName[string]("name")
		columnRank = ColumnName[int]("rank")
	)

	utils.CaseRunInSqliteMemDB(func(db *gorm.DB) {
		require.NoError(t, db.AutoMigrate(&Example{}))
		require.NoError(t, db.Save(&Example{Name: "abc", Rank: 100}).Error)
		require.NoError(t, db.Save(&Example{Name: "aaa", Rank: 101}).Error)

		type Result struct {
			Name string
			Mark int
		}

		var results []*Result
		//这是因为 rank 是您查询中的实际字段，而 mark 是 rank 字段的别名。在 ORDER BY 子句中，您应该使用实际的列名（即 rank），而不是别名（即 mark）。
		require.NoError(t, db.Model(&Example{}).
			Scopes(NewSx(columnName.Name()).
				Combine(NewSx(columnRank.AsAlias("mark"))).
				Scope()).
			Order(columnRank.Ob("desc").Ox()).Find(&results).Error)
		t.Log(neatjsons.S(results))

		require.Len(t, results, 2)
		require.Equal(t, results[0].Name, "aaa")
		require.Equal(t, results[0].Mark, 101)
		require.Equal(t, results[1].Name, "abc")
		require.Equal(t, results[1].Mark, 100)
	})
}
