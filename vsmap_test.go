package gormcnm

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestValuesMap_SetValue(t *testing.T) {
	type ExampleSetValue struct {
		Name string `gorm:"primary_key;type:varchar(100);"`
		Type string `gorm:"column:type;"`
		Rank int    `gorm:"column:rank;"`
	}

	require.NoError(t, caseDB.AutoMigrate(&ExampleSetValue{}))
	require.NoError(t, caseDB.Save(&ExampleSetValue{
		Name: "aaa",
		Type: "xxx",
		Rank: 123,
	}).Error)
	require.NoError(t, caseDB.Save(&ExampleSetValue{
		Name: "bbb",
		Type: "yyy",
		Rank: 456,
	}).Error)

	columnName := ColumnName[string]("name")
	columnType := ColumnName[string]("type")
	columnRank := ColumnName[int]("rank")

	{
		rst := caseDB.Model(&ExampleSetValue{}).Where(
			Qx(columnName.Eq("aaa")).
				AND(
					Qx(columnType.Eq("xxx")),
					Qx(columnRank.Eq(123)),
				).Qx3(),
		).UpdateColumns(columnRank.Kw(100).Kw(columnType.Kv("zzz")).Kws())
		require.NoError(t, rst.Error)
		require.Equal(t, int64(1), rst.RowsAffected)
	}
	{
		rst := caseDB.Model(&ExampleSetValue{}).Where(
			Qx(
				columnName.Eq("bbb"),
			).AND1(
				columnType.Eq("yyy"),
			).AND1(
				columnRank.Eq(456),
			).Qx3(),
		).UpdateColumns(columnRank.Kw(200).Kw(columnType.Kv("www")).Map())
		require.NoError(t, rst.Error)
		require.Equal(t, int64(1), rst.RowsAffected)
	}
}
