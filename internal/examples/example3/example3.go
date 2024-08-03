package example3

import "github.com/yyle88/gormcnm"

// 这里是用到的各个列的列名，需要手动配置，根据前面的表很容易得到
// 在项目 https://github.com/yyle88/gormcngen 里有自动生成列名的逻辑，就能自动配置啦，里面有各种示例，非常便捷
const (
	//这是用户表
	columnID   = gormcnm.ColumnName[string]("id")
	columnName = gormcnm.ColumnName[string]("name")

	//这是订单表，由于ID字段重复因此不用配它
	columnUserID = gormcnm.ColumnName[string]("user_id")
	columnAmount = gormcnm.ColumnName[string]("amount")
)

type User struct {
	ID   uint
	Name string
}

func (*User) TableName() string {
	return "users"
}

type Order struct {
	ID     uint
	UserID uint
	Amount float64
}

func (*Order) TableName() string {
	return "orders"
}
