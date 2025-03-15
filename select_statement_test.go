package gormcnm

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yyle88/gormcnm/internal/utils"
	"github.com/yyle88/neatjson/neatjsons"
	"gorm.io/gorm"
)

func TestSelectStatement_Combine(t *testing.T) {
	type Example struct {
		Name string `gorm:"primary_key;type:varchar(100);"`
		Rank int    `gorm:"column:rank;"`
	}

	const (
		columnName = ColumnName[string]("name")
		columnRank = ColumnName[int]("rank")
	)

	utils.CaseInMemDBRun(func(db *gorm.DB) {
		require.NoError(t, db.AutoMigrate(&Example{}))
		require.NoError(t, db.Save(&Example{Name: "abc", Rank: 100}).Error)
		require.NoError(t, db.Save(&Example{Name: "aaa", Rank: 101}).Error)

		type Result struct {
			Name string
			Mark int
		}

		selectStatement := NewSelectStatement(columnName.Name()).Combine(NewSx(columnRank.AsAlias("mark")))

		var results []*Result
		//这是因为 rank 是您查询中的实际字段，而 mark 是 rank 字段的别名。在 ORDER BY 子句中，您应该使用实际的列名（即 rank），而不是别名（即 mark）。
		require.NoError(t, db.Model(&Example{}).Select(selectStatement.Qx0()).Order(columnRank.Ob("desc").Ox()).Find(&results).Error)
		t.Log(neatjsons.S(results))

		require.Len(t, results, 2)
		require.Equal(t, results[0].Name, "aaa")
		require.Equal(t, results[0].Mark, 101)
		require.Equal(t, results[1].Name, "abc")
		require.Equal(t, results[1].Mark, 100)
	})
}
