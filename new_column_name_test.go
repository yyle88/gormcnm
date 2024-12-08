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
