package gormcnm

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yyle88/done"
	"github.com/yyle88/gormcnm/internal/utils"
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
	db := done.VCE(gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})).Nice()
	defer func() {
		done.Done(done.VCE(db.DB()).Nice().Close())
	}()

	done.Done(db.AutoMigrate(&Example{}))
	done.Done(db.Save(&Example{Name: "abc", Type: "xyz"}).Error)
	done.Done(db.Save(&Example{Name: "aaa", Type: "xxx"}).Error)

	caseDB = db
	m.Run()
}

func TestColumnName_Op(t *testing.T) {
	columnName := ColumnName[string]("name")

	{
		var one Example
		require.NoError(t, caseDB.Where(columnName.Op("=?", "abc")).First(&one).Error)
		require.Equal(t, "abc", one.Name)
		t.Log(utils.Neat(one))
	}
	{
		var one Example
		require.NoError(t, caseDB.Where(columnName.Eq("abc")).First(&one).Error)
		require.Equal(t, "abc", one.Name)
		t.Log(utils.Neat(one))
	}
	{
		var one Example
		require.NoError(t, caseDB.Where(columnName.BetweenAND("aba", "abd")).First(&one).Error)
		require.Equal(t, "abc", one.Name)
		t.Log(utils.Neat(one))
	}
	{
		var one Example
		require.ErrorIs(t, gorm.ErrRecordNotFound, caseDB.Where(columnName.IsNULL()).First(&one).Error)
		require.Equal(t, "", one.Name)
		t.Log(utils.Neat(one))
	}
	{
		var one Example
		require.NoError(t, caseDB.Where(columnName.IsNotNULL()).First(&one).Error)
		require.Contains(t, []string{"abc", "aaa"}, one.Name)
		t.Log(utils.Neat(one))
	}
	{
		var res []*Example
		require.NoError(t, caseDB.Where(columnName.In([]string{"abc", "aaa"})).Find(&res).Error)
		require.Contains(t, []string{"abc", "aaa"}, res[0].Name)
		require.Contains(t, []string{"abc", "aaa"}, res[1].Name)
		t.Log(utils.Neat(res))
	}
	{
		var res []*Example
		require.NoError(t, caseDB.Where(columnName.NotIn([]string{"aaa", "bbb"})).Find(&res).Error)
		for _, v := range res {
			require.NotEqual(t, "aaa", v.Name)
			require.NotEqual(t, "bbb", v.Name)
		}
		t.Log(utils.Neat(res))
	}
}

func TestColumnName_Op2(t *testing.T) {
	columnName := ColumnName[string]("name")
	columnType := ColumnName[string]("type")

	var one Example
	require.NoError(t, caseDB.Where(columnName.Qs("=?")+" AND "+columnType.Qs("=?"), "abc", "xyz").First(&one).Error)
	require.Equal(t, "abc", one.Name)
	t.Log(utils.Neat(one))
}

func TestColumnName_Op3(t *testing.T) {
	columnName := ColumnName[string]("name")

	{
		var one Example
		require.NoError(t, caseDB.Where(columnName.Like("%b%")).First(&one).Error)
		require.Equal(t, "abc", one.Name)
		t.Log(utils.Neat(one))
	}
	{
		var one Example
		require.NoError(t, caseDB.Where(columnName.NotLike("%b%")).First(&one).Error)
		require.Equal(t, "aaa", one.Name)
		t.Log(utils.Neat(one))
	}
}

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
		t.Log(utils.Neat(one))
	}
	{
		var one ExampleSafeCnm
		require.NoError(t, caseDB.Where(columnCreate.SafeCnm("`").Eq("xyz")).First(&one).Error)
		require.Equal(t, "xxx", one.Name)
		t.Log(utils.Neat(one))
	}
	{
		var one ExampleSafeCnm
		require.NoError(t, caseDB.Where(columnCreate.SafeCnm("[]").Eq("uvw")).First(&one).Error)
		require.Equal(t, "uuu", one.Name)
		t.Log(utils.Neat(one))
	}
	{
		var one ExampleSafeCnm
		require.NoError(t, caseDB.Where(columnCreate.SafeCnm("[-quote-]").Eq("uvw")).First(&one).Error)
		require.Equal(t, "uuu", one.Name)
		t.Log(utils.Neat(one))
	}
}
