package gormcnm

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yyle88/gormcnm/internal/utils"
	"gorm.io/gorm"
)

func TestCoalesceStmt(t *testing.T) {
	type Example struct {
		Name string `gorm:"primary_key;type:varchar(100);"`
		Rank int    `gorm:"column:rank;"`
	}

	const columnRank = ColumnName[int]("rank")

	utils.CaseRunInSqliteMemDB(func(db *gorm.DB) {
		require.NoError(t, db.AutoMigrate(&Example{}))
		require.NoError(t, db.Save(&Example{
			Name: "aaa",
			Rank: 123,
		}).Error)
		require.NoError(t, db.Save(&Example{
			Name: "bbb",
			Rank: 456,
		}).Error)

		{
			var value int
			err := db.Model(&Example{}).Select(columnRank.IFNULLFN().MaxStmt("max_rank")).First(&value).Error
			require.NoError(t, err)
			require.Equal(t, 456, value)
		}

		{
			var value int
			err := db.Model(&Example{}).Select(columnRank.COALESCE().MaxStmt("")).First(&value).Error
			require.NoError(t, err)
			require.Equal(t, 456, value)
		}
		{
			type resType struct {
				Value int
			}
			var res resType
			err := db.Model(&Example{}).Select(columnRank.COALESCE().MinStmt("value")).First(&res).Error
			require.NoError(t, err)
			require.Equal(t, 123, res.Value)
		}
		{
			type resType struct {
				Value int
			}
			var res resType
			err := db.Model(&Example{}).Select(columnRank.COALESCE().SumStmt("value")).First(&res).Error
			require.NoError(t, err)
			require.Equal(t, 579, res.Value)
		}
		{
			var value float64
			err := db.Model(&Example{}).Select(columnRank.COALESCE().AvgStmt("alias")).First(&value).Error
			require.NoError(t, err)
			require.Equal(t, 289.5, value)
		}
	})
}
