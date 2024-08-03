package example3

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yyle88/done"
	"github.com/yyle88/gormcnm"
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

	done.Done(db.AutoMigrate(&User{}, &Order{}))

	users := []User{
		{ID: 1, Name: "Alice"},
		{ID: 2, Name: "Bob"},
	}
	orders := []Order{
		{ID: 1, UserID: 1, Amount: 100},
		{ID: 2, UserID: 1, Amount: 200},
		{ID: 3, UserID: 2, Amount: 300},
	}
	done.Done(db.Create(&users).Error)
	done.Done(db.Create(&orders).Error)

	caseDB = db
	m.Run()
}

func TestColumnName_TcAlias(t *testing.T) {
	type UserOrder struct {
		UserID      uint
		UserName    string
		OrderID     uint
		OrderAmount float64
	}

	//这是中间结果类的字段
	//通常中间结果定义为局部类型，也可以不配置列名的，直接用 raw string 是更方便的
	//这里配个用于演示效果
	const columnOrderID = gormcnm.ColumnName[string]("order_id")

	var expectedResult string
	{ //这是比较常规的逻辑
		var results []*UserOrder
		require.NoError(t, caseDB.Table("users").
			Select("users.id as user_id, users.name as user_name, orders.id as order_id, orders.amount as order_amount").
			Joins("left join orders on orders.user_id = users.id").
			Order("users.id asc, orders.id asc").
			Scan(&results).Error)

		expectedResult = utils.Neat(results)
		t.Log(expectedResult)
	}
	{ //这是使用名称的逻辑
		var results []*UserOrder

		u := &User{}  //需要通过它来得到表名
		o := &Order{} //需要通过它来得到表名
		//这里需要个类型的实体来使用公共函数, 相当于新的名字空间 namespace，在 gormcngen 项目中已经将这个公共函数给写到了公共类里
		operation := &gormcnm.ColumnOperationClass{}
		require.NoError(t, caseDB.Table(u.TableName()).
			Select(operation.MergeStmts(
				columnID.TC(u).AsName(columnUserID),
				columnName.TC(u).AsAlias("user_name"), //认为使用 raw string 也是可以的，毕竟中间类型是局部类型
				columnID.TC(o).AsName(columnOrderID),
				columnAmount.TC(o).AsAlias("order_amount"), //认为使用 raw string 也是可以的，毕竟中间类型是局部类型
			)).
			Joins(operation.LEFTJOIN(o.TableName()).On(columnUserID.TC(o).Eq(columnID.TC(u)))).
			Order(columnID.TC(u).Ob("asc").Ob(columnID.TC(o).Ob("asc")).Ox()).
			Scan(&results).Error)

		require.Equal(t, expectedResult, utils.Neat(results))
	}
}
