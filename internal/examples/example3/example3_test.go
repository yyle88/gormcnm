// Package example3 demonstrates advanced table join operations using gormcnm
// Auto generates complex SQL joins with type-safe column operations
// Tests compare traditional SQL with gormcnm-generated equivalent queries
//
// example3 包演示了使用 gormcnm 的高级表连接操作
// 自动生成具有类型安全列操作的复杂 SQL 连接
// 测试比较传统 SQL 与 gormcnm 生成的等效查询
package example3

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yyle88/done"
	"github.com/yyle88/gormcnm"
	"github.com/yyle88/gormcnm/gormcnmstub"
	"github.com/yyle88/gormcnm/internal/tests"
	"github.com/yyle88/neatjson/neatjsons"
	"gorm.io/gorm"
)

type UserOrder struct {
	UserID      uint
	UserName    string
	OrderID     uint
	OrderAmount float64
}

func TestExample(t *testing.T) {
	tests.NewDBRun(t, func(db *gorm.DB) {
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
		//使用第三种方案，确保两者结果相同
		require.Equal(t, expectedText, neatjsons.S(selectFunc3(t, db)))
	})
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

	var results []*UserOrder
	require.NoError(t, db.Table(user.TableName()).
		Select(gormcnmstub.MergeStmts(
			userColumns.ID.TB(user).AsAlias("user_id"),
			userColumns.Name.TB(user).AsAlias("user_name"),
			orderColumns.ID.TB(order).AsAlias("order_id"),
			orderColumns.Amount.TB(order).AsAlias("order_amount"),
		)).
		Joins(userColumns.LEFTJOIN(order.TableName()).On(orderColumns.UserID.TB(order).Eq(userColumns.ID.TB(user)))).
		Order(userColumns.ID.TB(user).Ob("asc").Ob(orderColumns.ID.TB(order).Ob("asc")).Ox()).
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

	//这是使用名称的逻辑
	var results []*UserOrder
	require.NoError(t, db.Table(user.TableName()).
		Select(gormcnmstub.MergeStmts(
			userColumns.ID.WithTable(user).AsName(gormcnm.New[uint]("user_id")),
			userColumns.Name.WithTable(user).AsName("user_name"),
			orderColumns.ID.WithTable(order).AsName(gormcnm.New[uint]("order_id")),
			orderColumns.Amount.WithTable(order).AsName("order_amount"),
		)).
		Joins(userColumns.LEFTJOIN(order.TableName()).On(orderColumns.UserID.WithTable(order).Eq(userColumns.ID.WithTable(user)))).
		Order(userColumns.ID.WithTable(user).Ob("asc").Ob(orderColumns.ID.WithTable(order).Ob("asc")).Ox()).
		Scan(&results).Error)
	t.Log(neatjsons.S(results))
	return results
}

func selectFunc3(t *testing.T, db *gorm.DB) []*UserOrder {
	user := &User{}
	userColumns := user.Columns()
	order := &Order{}
	orderColumns := order.Columns()

	userOrder := &UserOrder{}

	//这是使用名称的逻辑
	var results []*UserOrder
	require.NoError(t, db.Table(user.TableName()).
		Select(gormcnmstub.MergeStmts(
			userColumns.ID.TC(user).AsName(gormcnm.Cnm(userOrder.UserID, "user_id")),
			userColumns.Name.TC(user).AsName(gormcnm.Cnm(userOrder.UserName, "user_name")),
			orderColumns.ID.TC(order).AsName(gormcnm.Cnm(userOrder.OrderID, "order_id")),
			orderColumns.Amount.TC(order).AsName(gormcnm.Cnm(userOrder.OrderAmount, "order_amount")),
		)).
		Joins(userColumns.LEFTJOIN(order.TableName()).On(orderColumns.UserID.TC(order).Eq(userColumns.ID.TC(user)))).
		Order(userColumns.ID.TC(user).Ob("asc").Ob(orderColumns.ID.TC(order).Ob("asc")).Ox()).
		Scan(&results).Error)
	t.Log(neatjsons.S(results))
	return results
}
