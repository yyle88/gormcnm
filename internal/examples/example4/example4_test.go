package example4

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yyle88/done"
	"github.com/yyle88/gormcnm"
	"github.com/yyle88/gormcnm/gormcnmstub"
	"github.com/yyle88/gormcnm/internal/utils"
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
	utils.CaseRunInSqliteMemDB(func(db *gorm.DB) {
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
	userColumns := user.TableColumns(gormcnm.NewTableDecoration(user.TableName()))
	order := &Order{}
	orderColumns := order.TableColumns(gormcnm.NewTableDecoration(order.TableName()))

	var results []*UserOrder
	require.NoError(t, db.Table(user.TableName()).
		Select(gormcnmstub.MergeStmts(
			userColumns.ID.AsAlias("user_id"),
			userColumns.Name.AsAlias("user_name"),
			orderColumns.ID.AsAlias("order_id"),
			orderColumns.Amount.AsAlias("order_amount"),
		)).
		Joins(userColumns.LEFTJOIN(order.TableName()).On(orderColumns.UserID.OnEq(userColumns.ID))).
		Order(userColumns.ID.Ob("asc").Ob(orderColumns.ID.Ob("asc")).Ox()).
		Scan(&results).Error)
	t.Log(neatjsons.S(results))
	return results
}

// 这是使用名称的逻辑
func selectFunc2(t *testing.T, db *gorm.DB) []*UserOrder {
	user := &User{}
	userColumns := user.TableColumns(gormcnm.NewTableDecoration(user.TableName()))
	order := &Order{}
	orderColumns := order.TableColumns(gormcnm.NewTableDecoration(order.TableName()))

	//这是使用名称的逻辑
	var results []*UserOrder
	require.NoError(t, db.Table(user.TableName()).
		Select(gormcnmstub.MergeStmts(
			userColumns.ID.AsName(gormcnm.New[uint]("user_id")),
			userColumns.Name.AsName("user_name"),
			orderColumns.ID.AsName(gormcnm.New[uint]("order_id")),
			orderColumns.Amount.AsName("order_amount"),
		)).
		Joins(userColumns.LEFTJOIN(order.TableName()).On(orderColumns.UserID.OnEq(userColumns.ID))).
		Order(userColumns.ID.Ob("asc").Ob(orderColumns.ID.Ob("asc")).Ox()).
		Scan(&results).Error)
	t.Log(neatjsons.S(results))
	return results
}

func selectFunc3(t *testing.T, db *gorm.DB) []*UserOrder {
	user := &User{}
	userColumns := user.TableColumns(gormcnm.NewTableDecoration(user.TableName()))
	order := &Order{}
	orderColumns := order.TableColumns(gormcnm.NewTableDecoration(order.TableName()))

	userOrder := &UserOrder{}

	//这是使用名称的逻辑
	var results []*UserOrder
	require.NoError(t, db.Table(user.TableName()).
		Select(gormcnmstub.MergeStmts(
			userColumns.ID.AsName(gormcnm.Cnm(userOrder.UserID, "user_id")),
			userColumns.Name.AsName(gormcnm.Cnm(userOrder.UserName, "user_name")),
			orderColumns.ID.AsName(gormcnm.Cnm(userOrder.OrderID, "order_id")),
			orderColumns.Amount.AsName(gormcnm.Cnm(userOrder.OrderAmount, "order_amount")),
		)).
		Joins(userColumns.LEFTJOIN(order.TableName()).On(orderColumns.UserID.OnEq(userColumns.ID))).
		Order(userColumns.ID.Ob("asc").Ob(orderColumns.ID.Ob("asc")).Ox()).
		Scan(&results).Error)
	t.Log(neatjsons.S(results))
	return results
}
