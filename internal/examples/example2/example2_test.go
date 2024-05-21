package example2

import (
	"fmt"
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
	"github.com/yyle88/done"
	"github.com/yyle88/gormcnm/internal/utils"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var caseDB *gorm.DB

func TestMain(m *testing.M) {
	//new db connection
	db := done.VCE(gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})).Nice()
	defer func() {
		done.Done(done.VCE(db.DB()).Nice().Close())
	}()

	//create example data
	_ = db.AutoMigrate(&Example{})
	_ = db.Save(&Example{Name: "abc", Type: utils.PtrX("xyz"), Rank: utils.PtrX(123)}).Error
	_ = db.Save(&Example{Name: "aaa", Type: utils.PtrX("xxx"), Rank: utils.PtrX(456)}).Error

	caseDB = db
	m.Run()
}

func TestExample(t *testing.T) {
	db := caseDB
	{
		var res Example
		err := db.Where("name=?", "abc").First(&res).Error
		done.Done(err)
		fmt.Println(res)
		require.Equal(t, 123, utils.VOr0(res.Rank))
	}
	{ //select an example data
		var res Example
		if err := db.Where(columnName.Eq("abc")).
			Where(columnType.Eq(utils.PtrX("xyz"))).
			Where(columnRank.Gt(utils.PtrX(100))).
			Where(columnRank.Lt(utils.PtrX(200))).
			First(&res).Error; err != nil {
			panic(errors.WithMessage(err, "wrong"))
		}
		fmt.Println(res)
		require.Equal(t, 123, utils.VOr0(res.Rank))
	}
}
