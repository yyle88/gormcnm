package gormcnm

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yyle88/neatjson"
)

func TestColumnName_Count(t *testing.T) {
	columnName := ColumnName[string]("name")
	{
		var value int
		err := caseDB.Model(&Example{}).Select(columnName.Count("cnt")).First(&value).Error
		require.NoError(t, err)
		require.Equal(t, 2, value)
	}
	{
		type resType struct {
			Cnt int64
		}
		var res resType
		err := caseDB.Model(&Example{}).Select(columnName.CountDistinct("cnt")).First(&res).Error
		require.NoError(t, err)
		require.Equal(t, int64(2), res.Cnt)
	}
}

func TestColumnName_SafeCnm(t *testing.T) {
	type ExampleSafeCnm struct {
		Name   string `gorm:"primary_key;type:varchar(100);"`
		Create string `gorm:"column:create"`
	}

	require.NoError(t, caseDB.AutoMigrate(&ExampleSafeCnm{}))
	require.NoError(t, caseDB.Save(&ExampleSafeCnm{
		Name:   "aaa",
		Create: "abc",
	}).Error)
	require.NoError(t, caseDB.Save(&ExampleSafeCnm{
		Name:   "xxx",
		Create: "xyz",
	}).Error)
	require.NoError(t, caseDB.Save(&ExampleSafeCnm{
		Name:   "uuu",
		Create: "uvw",
	}).Error)

	columnCreate := ColumnName[string]("create")

	{
		var one ExampleSafeCnm
		require.NoError(t, caseDB.Where(columnCreate.SafeCnm(`""`).Eq("abc")).First(&one).Error)
		require.Equal(t, "aaa", one.Name)
		t.Log(neatjson.TAB.Soft().S(one))
	}
	{
		var one ExampleSafeCnm
		require.NoError(t, caseDB.Where(columnCreate.SafeCnm("`").Eq("xyz")).First(&one).Error)
		require.Equal(t, "xxx", one.Name)
		t.Log(neatjson.TAB.Soft().S(one))
	}
	{
		var one ExampleSafeCnm
		require.NoError(t, caseDB.Where(columnCreate.SafeCnm("[]").Eq("uvw")).First(&one).Error)
		require.Equal(t, "uuu", one.Name)
		t.Log(neatjson.TAB.Soft().S(one))
	}
	{
		var one ExampleSafeCnm
		require.NoError(t, caseDB.Where(columnCreate.SafeCnm("[-quote-]").Eq("uvw")).First(&one).Error)
		require.Equal(t, "uuu", one.Name)
		t.Log(neatjson.TAB.Soft().S(one))
	}
}
