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

func TestColumnName_AsAlias(t *testing.T) {
	utils.CaseRunInMemDB(func(db *gorm.DB) {
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

		expectedText := neatjsons.S(selectFunc(t, db))
		//使用第一种方案，确保两者结果相同
		require.Equal(t, expectedText, neatjsons.S(selectFunc1(t, db)))
		//使用第二种方案，确保两者结果相同
		require.Equal(t, expectedText, neatjsons.S(selectFunc2(t, db)))
	})
}

type UserOrder struct {
	UserID      uint
	UserName    string
	OrderID     uint
	OrderAmount float64
}

// 这是比较常规的逻辑
func selectFunc(t *testing.T, db *gorm.DB) []*UserOrder {
	var results []*UserOrder
	require.NoError(t, db.Table("users").
		Select("users.id as user_id, users.name as user_name, orders.id as order_id, orders.amount as order_amount").
		Joins("left join orders on orders.user_id = users.id").
		Order("users.id asc, orders.id asc").
		Scan(&results).Error)
	t.Log(neatjsons.S(results))
	return results
}

// 这是使用名称的逻辑
func selectFunc1(t *testing.T, db *gorm.DB) []*UserOrder {
	user := &User{}
	userColumns := user.Columns()
	order := &Order{}
	orderColumns := order.Columns()
	operation := &gormcnm.ColumnOperationClass{}

	var results []*UserOrder
	require.NoError(t, db.Table(user.TableName()).
		Select(operation.MergeStmts(
			userColumns.ID.TC(user).AsAlias("user_id"),
			userColumns.Name.TC(user).AsAlias("user_name"),
			orderColumns.ID.TC(order).AsAlias("order_id"),
			orderColumns.Amount.TC(order).AsAlias("order_amount"),
		)).
		Joins(userColumns.LEFTJOIN(order.TableName()).On(orderColumns.UserID.TC(order).Eq(userColumns.ID.TC(user)))).
		Order(userColumns.ID.TC(user).Ob("asc").Ob(orderColumns.ID.TC(order).Ob("asc")).Ox()).
		Scan(&results).Error)
	t.Log(neatjsons.S(results))
	return results
}

// 这是使用名称的逻辑
func selectFunc2(t *testing.T, db *gorm.DB) []*UserOrder {
	user := &User{}
	userColumns := user.Columns()
	order := &Order{}
	orderColumns := order.Columns()
	operation := &gormcnm.ColumnOperationClass{}

	const (
		columnUserID      = gormcnm.ColumnName[uint]("user_id")
		columnUserName    = gormcnm.ColumnName[string]("user_name")
		columnOrderID     = gormcnm.ColumnName[uint]("order_id")
		columnOrderAmount = gormcnm.ColumnName[float64]("order_amount")
	)

	//这是使用名称的逻辑
	var results []*UserOrder
	require.NoError(t, db.Table(user.TableName()).
		Select(operation.MergeStmts(
			userColumns.ID.TC(user).AsName(columnUserID),
			userColumns.Name.TC(user).AsName(columnUserName),
			orderColumns.ID.TC(order).AsName(columnOrderID),
			orderColumns.Amount.TC(order).AsName(columnOrderAmount),
		)).
		Joins(userColumns.LEFTJOIN(order.TableName()).On(orderColumns.UserID.TC(order).Eq(userColumns.ID.TC(user)))).
		Order(userColumns.ID.TC(user).Ob("asc").Ob(orderColumns.ID.TC(order).Ob("asc")).Ox()).
		Scan(&results).Error)
	t.Log(neatjsons.S(results))
	return results
}
