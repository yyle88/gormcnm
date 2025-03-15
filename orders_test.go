package gormcnm

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yyle88/gormcnm/internal/utils"
	"github.com/yyle88/neatjson/neatjsons"
	"gorm.io/gorm"
)

func TestOrderByBottle_Ob(t *testing.T) {
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
			var res []*Example
			require.NoError(t, db.Where(columnName.In([]string{"abc", "aaa"})).
				Order(columnName.Ob("asc").
					Ob(columnType.Ob("desc")).
					Ox()).
				Find(&res).Error)
			require.Equal(t, "aaa", res[0].Name)
			require.Equal(t, "abc", res[1].Name)
			t.Log(neatjsons.S(res))
		}
		{
			var res []*Example
			require.NoError(t, db.Where(columnName.In([]string{"abc", "aaa"})).
				Order(columnName.Ob("desc").
					Ob(columnType.Ob("asc")).
					Ox()).
				Find(&res).Error)
			require.Equal(t, "abc", res[0].Name)
			require.Equal(t, "aaa", res[1].Name)
			t.Log(neatjsons.S(res))
		}
	})
}
