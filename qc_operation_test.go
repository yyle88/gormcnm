package gormcnm

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yyle88/gormcnm/internal/utils"
)

func TestColumnQcOperation_AND(t *testing.T) {
	columnName := ColumnName[string]("name")
	columnType := ColumnName[string]("type")

	{
		var one Example
		require.NoError(t, caseDB.Where(columnName.Qc("=?").AND(columnType.Qc("=?")).Qs(), "abc", "xyz").First(&one).Error)
		require.Equal(t, "abc", one.Name)
		t.Log(utils.SoftNeatString(one))
	}
	{
		var res []*Example
		require.NoError(t, caseDB.Where(
			columnName.Qc("=?").
				OR(
					columnName.Qc("=?"),
					columnName.Qc("=?"),
				).
				AND(
					columnName.Qc("IS NOT NULL"),
					columnType.Qc("IS NOT NULL"),
				).Qs(), "abc", "aaa", "bbb").Find(&res).Error)
		require.Contains(t, []string{"abc", "aaa"}, res[0].Name)
		require.Contains(t, []string{"abc", "aaa"}, res[1].Name)
		t.Log(utils.SoftNeatString(res))
	}
	{
		var one Example
		require.NoError(t, caseDB.Where(columnName.Qc("=?").NOT().Qs(), "abc").First(&one).Error)
		require.NotEqual(t, "abc", one.Name)
		t.Log(utils.SoftNeatString(one))
	}
}

func TestColumnQcOperation_AND_2(t *testing.T) {
	columnName := ColumnName[string]("name")
	columnType := ColumnName[string]("type")

	{
		var one Example
		require.NoError(t, caseDB.Where(columnName.Qc("=?").AND(columnType.Qc("=?")).Qs(), "abc", "xyz").First(&one).Error)
		require.Equal(t, "abc", one.Name)
		t.Log(utils.SoftNeatString(one))
	}
}
