package example2

import (
	"fmt"
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
	"github.com/yyle88/done"
	"github.com/yyle88/gormcnm"
	"github.com/yyle88/gormcnm/internal/utils"
	"gorm.io/gorm"
)

func TestExample(t *testing.T) {
	// Example is a gorm model define 3 fields(name, type, rank)
	type Example struct {
		Name string  `gorm:"primary_key;type:varchar(100);"`
		Type *string `gorm:"column:type;"` //通常不建议把字段定义为指针类型，但不建议不表示不能，因此要试试这种场景
		Rank *int    `gorm:"column:rank;"` //通常不建议把字段定义为指针类型，但不建议不表示不能，因此要试试这种场景
	}

	// Now define the fields enum vars(name, type rank)
	const (
		columnName = gormcnm.ColumnName[string]("name")
		columnType = gormcnm.ColumnName[*string]("type")
		columnRank = gormcnm.ColumnName[*int]("rank")
	)

	utils.CaseRunInSqliteMemDB(func(db *gorm.DB) {
		//create example data
		done.Done(db.AutoMigrate(&Example{}))
		done.Done(db.Save(&Example{Name: "abc", Type: utils.GetValuePointer("xyz"), Rank: utils.GetValuePointer(123)}).Error)
		done.Done(db.Save(&Example{Name: "aaa", Type: utils.GetValuePointer("xxx"), Rank: utils.GetValuePointer(456)}).Error)

		{
			var res Example
			err := db.Where("name=?", "abc").First(&res).Error
			done.Done(err)
			fmt.Println(res)
			require.Equal(t, 123, utils.GetPointerValue(res.Rank))
		}
		{ //select an example data
			var res Example
			if err := db.Where(columnName.Eq("abc")).
				Where(columnType.Eq(utils.GetValuePointer("xyz"))).
				Where(columnRank.Gt(utils.GetValuePointer(100))).
				Where(columnRank.Lt(utils.GetValuePointer(200))).
				First(&res).Error; err != nil {
				panic(errors.WithMessage(err, "wrong"))
			}
			fmt.Println(res)
			require.Equal(t, 123, utils.GetPointerValue(res.Rank))
		}
	})
}
