package gormcnm

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yyle88/gormcnm/internal/utils"
	"github.com/yyle88/neatjson/neatjsons"
	"gorm.io/gorm"
)

func TestQsConjunction_AND(t *testing.T) {
	type Example struct {
		Name string `gorm:"primary_key;type:varchar(100);"`
		Type string `gorm:"column:type;"`
	}

	const (
		columnName = ColumnName[string]("name")
		columnType = ColumnName[string]("type")
	)

	utils.CaseInMemDBRun(func(db *gorm.DB) {
		require.NoError(t, db.AutoMigrate(&Example{}))
		require.NoError(t, db.Save(&Example{Name: "abc", Type: "xyz"}).Error)
		require.NoError(t, db.Save(&Example{Name: "aaa", Type: "xxx"}).Error)

		{
			var one Example
			require.NoError(t, db.Where(columnName.Qc("=?").AND(columnType.Qc("=?")).Qs(), "abc", "xyz").First(&one).Error)
			require.Equal(t, "abc", one.Name)
			t.Log(neatjsons.S(one))
		}
		{
			var res []*Example
			require.NoError(t, db.Where(
				columnName.Qc("=?").
					OR(
						columnName.Qc("=?"),
						columnName.Qc("=?"),
					).
					AND(
						columnName.Qc("IS NOT NULL"),
						columnType.Qc("IS NOT NULL"),
					).Qs(), "abc", "aaa", "bbb").Find(&res).Error)
			require.Contains(t, []string{"abc", "aaa"}, res[0].Name)
			require.Contains(t, []string{"abc", "aaa"}, res[1].Name)
			t.Log(neatjsons.S(res))
		}
		{
			var one Example
			require.NoError(t, db.Where(columnName.Qc("=?").NOT().Qs(), "abc").First(&one).Error)
			require.NotEqual(t, "abc", one.Name)
			t.Log(neatjsons.S(one))
		}
	})
}

func TestQsConjunction_AND_2(t *testing.T) {
	type Example struct {
		Name string `gorm:"primary_key;type:varchar(100);"`
		Type string `gorm:"column:type;"`
	}

	const (
		columnName = ColumnName[string]("name")
		columnType = ColumnName[string]("type")
	)

	utils.CaseInMemDBRun(func(db *gorm.DB) {
		require.NoError(t, db.AutoMigrate(&Example{}))
		require.NoError(t, db.Save(&Example{Name: "abc", Type: "xyz"}).Error)
		require.NoError(t, db.Save(&Example{Name: "aaa", Type: "xxx"}).Error)

		{
			var one Example
			require.NoError(t, db.Where(columnName.Qc("=?").AND(columnType.Qc("=?")).Qs(), "abc", "xyz").First(&one).Error)
			require.Equal(t, "abc", one.Name)
			t.Log(neatjsons.S(one))
		}
	})
}
