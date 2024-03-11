package gormcnm

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yyle88/gormcnm/utilsgormcnm"
)

func TestColumnQx_AND(t *testing.T) {
	columnName := ColumnName[string]("name")
	columnType := ColumnName[string]("type")

	{
		var one Example
		qx := columnName.Qx("=?", "abc").AND(columnType.Qx("=?", "xyz"))
		t.Log(qx.Qs())
		t.Log(qx.Args())
		require.NoError(t, caseDB.Where(qx.Qs(), qx.Args()...).First(&one).Error)
		require.Equal(t, "abc", one.Name)
		t.Log(utilsgormcnm.SoftNeatString(one))
	}
	{
		var res []*Example
		qx := columnName.Qx("=?", "abc").OR(columnName.Qx("=?", "aaa"))
		t.Log(qx.Qs())
		t.Log(qx.Args())
		require.NoError(t, caseDB.Where(qx.Qs(), qx.Args()...).Find(&res).Error)
		require.Contains(t, []string{"abc", "aaa"}, res[0].Name)
		require.Contains(t, []string{"abc", "aaa"}, res[1].Name)
		t.Log(utilsgormcnm.SoftNeatString(res))
	}
	{
		var one Example
		qx := columnName.Qx("=?", "abc").NOT()
		t.Log(qx.Qs())
		t.Log(qx.Args())
		require.NoError(t, caseDB.Where(qx.Qs(), qx.Args()).First(&one).Error)
		require.NotEqual(t, "abc", one.Name)
		t.Log(utilsgormcnm.SoftNeatString(one))
	}
}

func TestColumnQx_AND_2(t *testing.T) {
	columnName := ColumnName[string]("name")
	columnType := ColumnName[string]("type")

	{
		var one Example
		require.NoError(t, caseDB.Where(columnName.Qx("=?", "abc").AND(columnType.Qx("=?", "xyz")).Qx2()).First(&one).Error)
		require.Equal(t, "abc", one.Name)
		t.Log(utilsgormcnm.SoftNeatString(one))
	}
	{
		var res []*Example
		require.NoError(t, caseDB.Where(columnName.Qx("=?", "abc").OR(columnName.Qx("=?", "aaa")).Qx2()).Find(&res).Error)
		require.Contains(t, []string{"abc", "aaa"}, res[0].Name)
		require.Contains(t, []string{"abc", "aaa"}, res[1].Name)
		t.Log(utilsgormcnm.SoftNeatString(res))
	}
	{
		var one Example
		require.NoError(t, caseDB.Where(columnName.Qx("=?", "abc").NOT().Qx1()).First(&one).Error)
		require.NotEqual(t, "abc", one.Name)
		t.Log(utilsgormcnm.SoftNeatString(one))
	}
}

func TestColumnQx_AND_3(t *testing.T) {
	columnName := ColumnName[string]("name")
	columnType := ColumnName[string]("type")

	{
		var one Example
		require.NoError(t, caseDB.Where(NewQx(columnName.Eq("abc")).AND(NewQx(columnType.Eq("xyz"))).Qx2()).First(&one).Error)
		require.Equal(t, "abc", one.Name)
		t.Log(utilsgormcnm.SoftNeatString(one))
	}
	{
		var res []*Example
		require.NoError(t, caseDB.Where(NewQx(columnName.Eq("abc")).OR(NewQx(columnName.Eq("aaa"))).Qx2()).Find(&res).Error)
		require.Contains(t, []string{"abc", "aaa"}, res[0].Name)
		require.Contains(t, []string{"abc", "aaa"}, res[1].Name)
		t.Log(utilsgormcnm.SoftNeatString(res))
	}
	{
		var one Example
		require.NoError(t, caseDB.Where(NewQx(columnName.Eq("abc")).NOT().Qx1()).First(&one).Error)
		require.NotEqual(t, "abc", one.Name)
		t.Log(utilsgormcnm.SoftNeatString(one))
	}
}

func TestColumnQx_AND_4(t *testing.T) {
	columnName := ColumnName[string]("name")
	columnType := ColumnName[string]("type")

	{
		var one Example
		require.NoError(t, caseDB.Where(Qx(columnName.Eq("abc")).AND(Qx(columnType.Eq("xyz"))).Qx2()).First(&one).Error)
		require.Equal(t, "abc", one.Name)
		t.Log(utilsgormcnm.SoftNeatString(one))
	}
	{
		var res []*Example
		require.NoError(t, caseDB.Where(Qx(columnName.Eq("abc")).OR(Qx(columnName.Eq("aaa"))).Qx2()).Find(&res).Error)
		require.Contains(t, []string{"abc", "aaa"}, res[0].Name)
		require.Contains(t, []string{"abc", "aaa"}, res[1].Name)
		t.Log(utilsgormcnm.SoftNeatString(res))
	}
	{
		var one Example
		require.NoError(t, caseDB.Where(Qx(columnName.Eq("abc")).NOT().Qx1()).First(&one).Error)
		require.NotEqual(t, "abc", one.Name)
		t.Log(utilsgormcnm.SoftNeatString(one))
	}
}

func TestColumnQx_AND_5(t *testing.T) {
	columnName := ColumnName[string]("name")
	columnType := ColumnName[string]("type")

	{
		var one Example
		require.NoError(t, caseDB.Where(Qx(columnName.BetweenAND("aba", "abd")).AND(Qx(columnType.IsNotNULL())).Qx2()).First(&one).Error)
		require.Equal(t, "abc", one.Name)
		t.Log(utilsgormcnm.SoftNeatString(one))
	}
}
