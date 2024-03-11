package gormcnm

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yyle88/gormcnm/utilsyyle"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var caseDB *gorm.DB

type Example struct {
	Name string `gorm:"primary_key;type:varchar(100);"`
	Type string `gorm:"column:type;"`
}

func TestMain(m *testing.M) {
	fmt.Println("run_test_main")
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	utilsyyle.AssertDone(err)
	caseDB = db

	utilsyyle.AssertDone(db.AutoMigrate(&Example{}))
	utilsyyle.AssertDone(caseDB.Save(&Example{Name: "abc", Type: "xyz"}).Error)
	utilsyyle.AssertDone(caseDB.Save(&Example{Name: "aaa", Type: "xxx"}).Error)
	m.Run()
	os.Exit(0)
}

func TestColumnName_Op(t *testing.T) {
	columnName := ColumnName[string]("name")

	{
		var one Example
		require.NoError(t, caseDB.Where(columnName.Op("=?", "abc")).First(&one).Error)
		require.Equal(t, "abc", one.Name)
		t.Log(utilsyyle.SoftNeatString(one))
	}
	{
		var one Example
		require.NoError(t, caseDB.Where(columnName.Eq("abc")).First(&one).Error)
		require.Equal(t, "abc", one.Name)
		t.Log(utilsyyle.SoftNeatString(one))
	}
	{
		var one Example
		require.NoError(t, caseDB.Where(columnName.BetweenAND("aba", "abd")).First(&one).Error)
		require.Equal(t, "abc", one.Name)
		t.Log(utilsyyle.SoftNeatString(one))
	}
	{
		var one Example
		require.ErrorIs(t, gorm.ErrRecordNotFound, caseDB.Where(columnName.IsNULL()).First(&one).Error)
		require.Equal(t, "", one.Name)
		t.Log(utilsyyle.SoftNeatString(one))
	}
	{
		var one Example
		require.NoError(t, caseDB.Where(columnName.IsNotNULL()).First(&one).Error)
		require.Contains(t, []string{"abc", "aaa"}, one.Name)
		t.Log(utilsyyle.SoftNeatString(one))
	}
	{
		var res []*Example
		require.NoError(t, caseDB.Where(columnName.In([]string{"abc", "aaa"})).Find(&res).Error)
		require.Contains(t, []string{"abc", "aaa"}, res[0].Name)
		require.Contains(t, []string{"abc", "aaa"}, res[1].Name)
		t.Log(utilsyyle.SoftNeatString(res))
	}
	{
		var res []*Example
		require.NoError(t, caseDB.Where(columnName.NotIn([]string{"aaa", "bbb"})).Find(&res).Error)
		for _, v := range res {
			require.NotEqual(t, "aaa", v.Name)
			require.NotEqual(t, "bbb", v.Name)
		}
		t.Log(utilsyyle.SoftNeatString(res))
	}
}

func TestColumnName_Op2(t *testing.T) {
	columnName := ColumnName[string]("name")
	columnType := ColumnName[string]("type")

	var one Example
	require.NoError(t, caseDB.Where(columnName.Qs("=?")+" AND "+columnType.Qs("=?"), "abc", "xyz").First(&one).Error)
	require.Equal(t, "abc", one.Name)
	t.Log(utilsyyle.SoftNeatString(one))
}

func TestColumnName_Op3(t *testing.T) {
	columnName := ColumnName[string]("name")

	{
		var one Example
		require.NoError(t, caseDB.Where(columnName.Like("%b%")).First(&one).Error)
		require.Equal(t, "abc", one.Name)
		t.Log(utilsyyle.SoftNeatString(one))
	}
	{
		var one Example
		require.NoError(t, caseDB.Where(columnName.NotLike("%b%")).First(&one).Error)
		require.Equal(t, "aaa", one.Name)
		t.Log(utilsyyle.SoftNeatString(one))
	}
}
