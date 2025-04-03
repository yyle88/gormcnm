package gormcnm

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {
	columnName := New[string]("name")

	qs, value := columnName.Eq("abc")
	t.Log(qs, value)

	require.Equal(t, "name=?", qs)
	require.Equal(t, "abc", value)
}

func TestCnm(t *testing.T) {
	type Example struct {
		Name string `gorm:"column:name;primary_key;type:varchar(100);"`
	}

	res := Example{}
	columnName := Cnm(res.Name, "name")

	qs, value := columnName.Eq("abc")
	t.Log(qs, value)

	require.Equal(t, "name=?", qs)
	require.Equal(t, "abc", value)
}

func TestCmn(t *testing.T) {
	type Example struct {
		Name string `gorm:"column:name;primary_key;type:varchar(100);"`
	}

	decoration := NewPlainDecoration()

	res := Example{}
	columnName := Cmn(res.Name, "name", decoration)

	qs, value := columnName.Eq("abc")
	t.Log(qs, value)

	require.Equal(t, "name=?", qs)
	require.Equal(t, "abc", value)
}

func TestCmn_WithTableNameDecoration(t *testing.T) {
	type Example struct {
		Name string `gorm:"column:name;primary_key;type:varchar(100);"`
	}

	decoration := NewTableDecoration("examples")

	res := Example{}
	columnName := Cmn(res.Name, "name", decoration)

	qs, value := columnName.Eq("abc")
	t.Log(qs, value)

	require.Equal(t, "examples.name=?", qs)
	require.Equal(t, "abc", value)
}

func TestCmn_WithCustomDecoration(t *testing.T) {
	type Example struct {
		Name string `gorm:"column:name;primary_key;type:varchar(100);"`
	}

	decoration := NewCustomDecoration(func(name string) string {
		return "examples" + "." + name
	})

	res := Example{}
	columnName := Cmn(res.Name, "name", decoration)

	qs, value := columnName.Eq("abc")
	t.Log(qs, value)

	require.Equal(t, "examples.name=?", qs)
	require.Equal(t, "abc", value)
}
