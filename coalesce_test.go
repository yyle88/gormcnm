package gormcnm

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCoalesceQs_Stmt(t *testing.T) {
	type ExampleCoalesceType struct {
		Name string `gorm:"primary_key;type:varchar(100);"`
		Rank int    `gorm:"column:rank;"`
	}

	require.NoError(t, caseDB.AutoMigrate(&ExampleCoalesceType{}))
	require.NoError(t, caseDB.Save(&ExampleCoalesceType{
		Name: "aaa",
		Rank: 123,
	}).Error)
	require.NoError(t, caseDB.Save(&ExampleCoalesceType{
		Name: "bbb",
		Rank: 456,
	}).Error)

	columnRank := ColumnName[int]("rank")

	{
		var value int
		err := caseDB.Model(&ExampleCoalesceType{}).Select(columnRank.IFNULLFn().MaxStmt("max_rank")).First(&value).Error
		require.NoError(t, err)
		require.Equal(t, 456, value)
	}

	{
		var value int
		err := caseDB.Model(&ExampleCoalesceType{}).Select(columnRank.COALESCE().MaxStmt("")).First(&value).Error
		require.NoError(t, err)
		require.Equal(t, 456, value)
	}
	{
		type resType struct {
			Value int
		}
		var res resType
		err := caseDB.Model(&ExampleCoalesceType{}).Select(columnRank.COALESCE().MinStmt("value")).First(&res).Error
		require.NoError(t, err)
		require.Equal(t, 123, res.Value)
	}
	{
		type resType struct {
			Value int
		}
		var res resType
		err := caseDB.Model(&ExampleCoalesceType{}).Select(columnRank.COALESCE().SumStmt("value")).First(&res).Error
		require.NoError(t, err)
		require.Equal(t, 579, res.Value)
	}
	{
		var value float64
		err := caseDB.Model(&ExampleCoalesceType{}).Select(columnRank.COALESCE().AvgStmt("alias")).First(&value).Error
		require.NoError(t, err)
		require.Equal(t, 289.5, value)
	}
}
