package gormcnm

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yyle88/gormcnm/internal/utils"
	"github.com/yyle88/neatjson/neatjsons"
	"gorm.io/gorm"
)

func TestColumnName_SafeCnm(t *testing.T) {
	type Example struct {
		Name   string `gorm:"primary_key;type:varchar(100);"`
		Create string `gorm:"column:create"`
	}

	const columnCreate = ColumnName[string]("create")

	utils.CaseInMemDBRun(func(db *gorm.DB) {
		require.NoError(t, db.AutoMigrate(&Example{}))
		require.NoError(t, db.Save(&Example{
			Name:   "aaa",
			Create: "abc",
		}).Error)
		require.NoError(t, db.Save(&Example{
			Name:   "xxx",
			Create: "xyz",
		}).Error)
		require.NoError(t, db.Save(&Example{
			Name:   "uuu",
			Create: "uvw",
		}).Error)

		{
			var one Example
			require.NoError(t, db.Where(columnCreate.SafeCnm(`""`).Eq("abc")).First(&one).Error)
			require.Equal(t, "aaa", one.Name)
			t.Log(neatjsons.S(one))
		}
		{
			var one Example
			require.NoError(t, db.Where(columnCreate.SafeCnm("`").Eq("xyz")).First(&one).Error)
			require.Equal(t, "xxx", one.Name)
			t.Log(neatjsons.S(one))
		}
		{
			var one Example
			require.NoError(t, db.Where(columnCreate.SafeCnm("[]").Eq("uvw")).First(&one).Error)
			require.Equal(t, "uuu", one.Name)
			t.Log(neatjsons.S(one))
		}
		{
			var one Example
			require.NoError(t, db.Where(columnCreate.SafeCnm("[-quote-]").Eq("uvw")).First(&one).Error)
			require.Equal(t, "uuu", one.Name)
			t.Log(neatjsons.S(one))
		}
	})
}

func TestColumnName_Count(t *testing.T) {
	type Example struct {
		Name string `gorm:"primary_key;type:varchar(100);"`
		Type string `gorm:"column:type;"`
	}

	const columnName = ColumnName[string]("name")

	utils.CaseInMemDBRun(func(db *gorm.DB) {
		require.NoError(t, db.AutoMigrate(&Example{}))
		require.NoError(t, db.Save(&Example{Name: "abc", Type: "xyz"}).Error)
		require.NoError(t, db.Save(&Example{Name: "aaa", Type: "xxx"}).Error)

		{
			var value int
			err := db.Model(&Example{}).Select(columnName.Count("cnt")).First(&value).Error
			require.NoError(t, err)
			require.Equal(t, 2, value)
		}
		{
			type resType struct {
				Cnt int64
			}
			var res resType
			err := db.Model(&Example{}).Select(columnName.CountDistinct("cnt")).First(&res).Error
			require.NoError(t, err)
			require.Equal(t, int64(2), res.Cnt)
		}
	})
}
