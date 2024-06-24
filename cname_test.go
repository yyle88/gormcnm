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
		t.Log(utils.SoftNeatString(one))
	}
	{
		var one Example
		require.NoError(t, caseDB.Where(columnName.Eq("abc")).First(&one).Error)
		require.Equal(t, "abc", one.Name)
		t.Log(utils.SoftNeatString(one))
	}
	{
		var one Example
		require.NoError(t, caseDB.Where(columnName.BetweenAND("aba", "abd")).First(&one).Error)
		require.Equal(t, "abc", one.Name)
		t.Log(utils.SoftNeatString(one))
	}
	{
		var one Example
		require.ErrorIs(t, gorm.ErrRecordNotFound, caseDB.Where(columnName.IsNULL()).First(&one).Error)
		require.Equal(t, "", one.Name)
		t.Log(utils.SoftNeatString(one))
	}
	{
		var one Example
		require.NoError(t, caseDB.Where(columnName.IsNotNULL()).First(&one).Error)
		require.Contains(t, []string{"abc", "aaa"}, one.Name)
		t.Log(utils.SoftNeatString(one))
	}
	{
		var res []*Example
		require.NoError(t, caseDB.Where(columnName.In([]string{"abc", "aaa"})).Find(&res).Error)
		require.Contains(t, []string{"abc", "aaa"}, res[0].Name)
		require.Contains(t, []string{"abc", "aaa"}, res[1].Name)
		t.Log(utils.SoftNeatString(res))
	}
	{
		var res []*Example
		require.NoError(t, caseDB.Where(columnName.NotIn([]string{"aaa", "bbb"})).Find(&res).Error)
		for _, v := range res {
			require.NotEqual(t, "aaa", v.Name)
			require.NotEqual(t, "bbb", v.Name)
		}
		t.Log(utils.SoftNeatString(res))
	}
}

func TestColumnName_Op2(t *testing.T) {
	columnName := ColumnName[string]("name")
	columnType := ColumnName[string]("type")

	var one Example
	require.NoError(t, caseDB.Where(columnName.Qs("=?")+" AND "+columnType.Qs("=?"), "abc", "xyz").First(&one).Error)
	require.Equal(t, "abc", one.Name)
	t.Log(utils.SoftNeatString(one))
}

func TestColumnName_Op3(t *testing.T) {
	columnName := ColumnName[string]("name")

	{
		var one Example
		require.NoError(t, caseDB.Where(columnName.Like("%b%")).First(&one).Error)
		require.Equal(t, "abc", one.Name)
		t.Log(utils.SoftNeatString(one))
	}
	{
		var one Example
		require.NoError(t, caseDB.Where(columnName.NotLike("%b%")).First(&one).Error)
		require.Equal(t, "aaa", one.Name)
		t.Log(utils.SoftNeatString(one))
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
		require.NoError(t, caseDB.Where(columnCreate.SafeCnm("\"\"").Eq("abc")).First(&one).Error)
		require.Equal(t, "aaa", one.Name)
		t.Log(utils.SoftNeatString(one))
	}
	{
		var one ExampleSafeCnm
		require.NoError(t, caseDB.Where(columnCreate.SafeCnm("`").Eq("xyz")).First(&one).Error)
		require.Equal(t, "xxx", one.Name)
		t.Log(utils.SoftNeatString(one))
	}
	{
		var one ExampleSafeCnm
		require.NoError(t, caseDB.Where(columnCreate.SafeCnm("[]").Eq("uvw")).First(&one).Error)
		require.Equal(t, "uuu", one.Name)
		t.Log(utils.SoftNeatString(one))
	}
	{
		var one ExampleSafeCnm
		require.NoError(t, caseDB.Where(columnCreate.SafeCnm("[-quote-]").Eq("uvw")).First(&one).Error)
		require.Equal(t, "uuu", one.Name)
		t.Log(utils.SoftNeatString(one))
	}
}

func TestColumnName_CoalesceStmt(t *testing.T) {
	type ExampleCoalesceStmtValue struct {
		Name string `gorm:"primary_key;type:varchar(100);"`
		Rank int    `gorm:"column:rank;"`
	}

	require.NoError(t, caseDB.AutoMigrate(&ExampleCoalesceStmtValue{}))
	require.NoError(t, caseDB.Save(&ExampleCoalesceStmtValue{
		Name: "aaa",
		Rank: 123,
	}).Error)
	require.NoError(t, caseDB.Save(&ExampleCoalesceStmtValue{
		Name: "bbb",
		Rank: 456,
	}).Error)

	columnRank := ColumnName[int]("rank")

	{
		var value int
		err := caseDB.Model(&ExampleCoalesceStmtValue{}).Select(columnRank.CoalesceMaxStmt("")).First(&value).Error
		require.NoError(t, err)
		require.Equal(t, 456, value)
	}
	{
		type resType struct {
			Value int
		}
		var res resType
		err := caseDB.Model(&ExampleCoalesceStmtValue{}).Select(columnRank.CoalesceMinStmt("value")).First(&res).Error
		require.NoError(t, err)
		require.Equal(t, 123, res.Value)
	}
	{
		type resType struct {
			Value int
		}
		var res resType
		err := caseDB.Model(&ExampleCoalesceStmtValue{}).Select(columnRank.CoalesceSumStmt("value")).First(&res).Error
		require.NoError(t, err)
		require.Equal(t, 579, res.Value)
	}
	{
		var value float64
		err := caseDB.Model(&ExampleCoalesceStmtValue{}).Select(columnRank.CoalesceAvgStmt("alias")).First(&value).Error
		require.NoError(t, err)
		require.Equal(t, 289.5, value)
	}
}
