package gormcnm

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yyle88/gormcnm/internal/utils"
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

	utils.CaseRunInSqliteMemDB(func(db *gorm.DB) {
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
