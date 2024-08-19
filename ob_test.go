package gormcnm

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yyle88/neatjson"
)

func TestObColumnAscDesc_Ob(t *testing.T) {
	columnName := ColumnName[string]("name")
	columnType := ColumnName[string]("type")

	{
		var res []*Example
		require.NoError(t, caseDB.Where(columnName.In([]string{"abc", "aaa"})).
			Order(columnName.Ob("asc").
				Ob(columnType.Ob("desc")).
				Ox()).
			Find(&res).Error)
		require.Equal(t, "aaa", res[0].Name)
		require.Equal(t, "abc", res[1].Name)
		t.Log(neatjson.TAB.Soft().S(res))
	}
	{
		var res []*Example
		require.NoError(t, caseDB.Where(columnName.In([]string{"abc", "aaa"})).
			Order(columnName.Ob("desc").
				Ob(columnType.Ob("asc")).
				Ox()).
			Find(&res).Error)
		require.Equal(t, "abc", res[0].Name)
		require.Equal(t, "aaa", res[1].Name)
		t.Log(neatjson.TAB.Soft().S(res))
	}
}
