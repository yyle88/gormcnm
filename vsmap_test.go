package gormcnm

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yyle88/gormcnm/internal/utils"
	"gorm.io/gorm"
)

func TestValuesMap_SetValue(t *testing.T) {
	type Example struct {
		Name string `gorm:"primary_key;type:varchar(100);"`
		Type string `gorm:"column:type;"`
		Rank int    `gorm:"column:rank;"`
	}

	const (
		columnName = ColumnName[string]("name")
		columnType = ColumnName[string]("type")
		columnRank = ColumnName[int]("rank")
	)

	utils.CaseRunInPrivateDB(func(db *gorm.DB) {
		require.NoError(t, db.AutoMigrate(&Example{}))
		require.NoError(t, db.Save(&Example{
			Name: "aaa",
			Type: "xxx",
			Rank: 123,
		}).Error)
		require.NoError(t, db.Save(&Example{
			Name: "bbb",
			Type: "yyy",
			Rank: 456,
		}).Error)

		{
			rst := db.Model(&Example{}).Where(
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
			rst := db.Model(&Example{}).Where(
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
	})
}
