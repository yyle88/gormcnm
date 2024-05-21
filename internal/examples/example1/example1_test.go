package example1

import (
	"fmt"
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
	"github.com/yyle88/done"
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
	_ = db.Save(&Example{Name: "abc", Type: "xyz", Rank: 123}).Error
	_ = db.Save(&Example{Name: "aaa", Type: "xxx", Rank: 456}).Error

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
		require.Equal(t, 123, res.Rank)
	}
	{ //select an example data
		var res Example
		if err := db.Where(columnName.Eq("abc")).
			Where(columnType.Eq("xyz")).
			Where(columnRank.Gt(100)).
			Where(columnRank.Lt(200)).
			First(&res).Error; err != nil {
			panic(errors.WithMessage(err, "wrong"))
		}
		fmt.Println(res)
		require.Equal(t, 123, res.Rank)
	}
}
