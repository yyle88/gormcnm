package example3

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yyle88/done"
	"github.com/yyle88/gormcnm"
	"github.com/yyle88/gormcnm/internal/utils"
	"github.com/yyle88/neatjson/neatjsons"
	"gorm.io/gorm"
)

func TestColumnName_TcAlias(t *testing.T) {
	utils.CaseRunInPrivateDB(func(db *gorm.DB) {
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

		type UserOrder struct {
			UserID      uint
			UserName    string
			OrderID     uint
			OrderAmount float64
		}

		var expected string
		{ //这是比较常规的逻辑
			var results []*UserOrder
			require.NoError(t, db.Table("users").
				Select("users.id as user_id, users.name as user_name, orders.id as order_id, orders.amount as order_amount").
				Joins("left join orders on orders.user_id = users.id").
				Order("users.id asc, orders.id asc").
				Scan(&results).Error)

			expected = neatjsons.S(results)
			t.Log(expected)
		}
		{ //这是使用名称的逻辑
			var results []*UserOrder

			user := &User{} //需要通过它来得到表名
			userColumns := user.Columns()
			order := &Order{} //需要通过它来得到表名
			orderColumns := order.Columns()

			//这里需要个类型的实体来使用公共函数, 相当于新的名字空间 namespace，在 gormcngen 项目中已经将这个公共函数给写到了公共类里
			operation := &gormcnm.ColumnOperationClass{}
			require.NoError(t, db.Table(user.TableName()).
				Select(operation.MergeStmts(
					userColumns.ID.TC(user).AsAlias("user_id"),
					userColumns.Name.TC(user).AsAlias("user_name"), //认为使用 raw string 也是可以的，毕竟中间类型是局部类型
					orderColumns.ID.TC(order).AsAlias("order_id"),
					orderColumns.Amount.TC(order).AsAlias("order_amount"), //认为使用 raw string 也是可以的，毕竟中间类型是局部类型
				)).
				Joins(operation.LEFTJOIN(order.TableName()).On(orderColumns.UserID.TC(order).Eq(userColumns.ID.TC(user)))).
				Order(userColumns.ID.TC(user).Ob("asc").Ob(orderColumns.ID.TC(order).Ob("asc")).Ox()).
				Scan(&results).Error)

			require.Equal(t, expected, neatjsons.S(results))
		}
	})
}
