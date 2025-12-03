// Package example3 demonstrates namespace pattern and manual column definition with gormcnm
// Auto shows User and Order models with Columns() methods returning column name collections
// Supports avoiding column name conflicts across multiple tables
//
// example3 包演示使用 gormcnm 的命名空间模式和手动列定义
// 自动展示 User 和 Order 模型，通过 Columns() 方法返回列名集合
// 支持避免多个表之间的列名冲突
package example3

import "github.com/yyle88/gormcnm"

// 当模型/表的数量比较多时，也可以使用名字空间把列名包裹起来
// 比如下面的 UserColumns 和 OrderColumns 两个名字空间
// 这样就能在不同表有重名的列时就能避免混淆
// 区分得很清楚
// 当然这个代码其实可以通过语法分析自动得到
// 在项目 https://github.com/yyle88/gormcngen 里有自动生成列名的逻辑，就能自动配置啦，里面有各种示例，非常便捷

type User struct {
	ID   uint
	Name string
}

func (*User) TableName() string {
	return "users"
}

func (*User) Columns() *UserColumns {
	return &UserColumns{
		ID:   "id",
		Name: "name",
	}
}

type UserColumns struct {
	gormcnm.ColumnOperationClass //继承操作函数，让查询更便捷
	// 模型各个列名和类型:
	ID   gormcnm.ColumnName[uint]
	Name gormcnm.ColumnName[string]
}

type Order struct {
	ID     uint
	UserID uint
	Amount float64
}

func (*Order) TableName() string {
	return "orders"
}

func (*Order) Columns() *OrderColumns {
	return &OrderColumns{
		ID:     "id",
		UserID: "user_id",
		Amount: "amount",
	}
}

type OrderColumns struct {
	gormcnm.ColumnOperationClass //继承操作函数，让查询更便捷
	// 模型各个列名和类型:
	ID     gormcnm.ColumnName[uint]
	UserID gormcnm.ColumnName[uint]
	Amount gormcnm.ColumnName[float64]
}
