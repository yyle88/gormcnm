package gormcnm

import "gorm.io/gorm"

// ColumnBaseFuncClass 让外部能够继承它，这样就能继承操作函数，让查询更便捷
// 自动生成的代码，不仅要包含各个列的基本信息，能进行简单查询，还要能够有便捷的操作函数，以便于复杂的查询
type ColumnBaseFuncClass struct{}

func (c *ColumnBaseFuncClass) OK() bool {
	return true //这个函数有奇效，让你把变量的创建放在if{}代码块里
}

func (c *ColumnBaseFuncClass) NewQx(qs string, args ...interface{}) *QxType {
	return NewQx(qs, args...)
}

func (c *ColumnBaseFuncClass) Qx(qs string, args ...interface{}) *QxType {
	return NewQx(qs, args...)
}

func (c *ColumnBaseFuncClass) NewKw() KeywordArguments {
	return NewKw()
}

func (c *ColumnBaseFuncClass) Kw(columnName string, value interface{}) KeywordArguments {
	return Kw(columnName, value)
}

// Where 设置查询条件
// 很明显这样做会破坏gorm链式操作的写法，但这样也是可行的，也能简化些代码
func (c *ColumnBaseFuncClass) Where(db *gorm.DB, qxs ...*QxType) *gorm.DB {
	stmt := db
	for _, qx := range qxs {
		stmt = stmt.Where(qx.Qs(), qx.args...)
	}
	return stmt
}

// Order 设置排序方向
// 很明显这样做会破坏gorm链式操作的写法，但这样也是可行的，也能简化些代码
func (c *ColumnBaseFuncClass) Order(db *gorm.DB, obs ...ColumnOrderByAscDesc) *gorm.DB {
	stmt := db
	for _, ob := range obs {
		stmt = stmt.Order(ob.Ox())
	}
	return stmt
}

// UpdateColumns 根据字典更新数据
// 很明显这样做会破坏gorm链式操作的写法，但这样也是可行的，也能简化些代码
func (c *ColumnBaseFuncClass) UpdateColumns(db *gorm.DB, kws ...KeywordArguments) *gorm.DB {
	mp := map[string]interface{}{}
	for _, kw := range kws {
		for k, v := range kw.AsMap() {
			mp[k] = v
		}
	}
	return db.UpdateColumns(mp)
}
