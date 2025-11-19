// Package gormcnm tests validate column name creation and decoration operations
// Auto verifies New, Cnm, Cmn constructors with PlainDecoration, TableDecoration, CustomDecoration
// Tests cover column name instantiation, table prefix decoration, and custom transformation logic
//
// gormcnm 测试包验证列名创建和装饰操作
// 自动验证 New、Cnm、Cmn 构造函数，包含 PlainDecoration、TableDecoration、CustomDecoration
// 测试涵盖列名实例化、表前缀装饰和自定义转换逻辑
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
