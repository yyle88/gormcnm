package gormcnm

import "github.com/yyle88/must"

// New creates a new ColumnName with the specified type and name
// New 创建一个带有指定类型和名称的新 ColumnName
func New[T any](name string) ColumnName[T] {
	return ColumnName[T](name)
}

// Cnm creates a ColumnName using the type inferred from the value parameter
// Cnm 使用从值参数推断出的类型创建 ColumnName
func Cnm[T any](v T, name string) ColumnName[T] {
	return ColumnName[T](name)
}

// Cmn creates a ColumnName with decoration applied to the name
// Cmn 创建一个对名称应用装饰的 ColumnName
func Cmn[T any](v T, name string, decoration ColumnNameDecoration) ColumnName[T] {
	return ColumnName[T](decoration.DecorateColumnName(name))
}

// ColumnNameDecoration defines an interface for decorating column names
// ColumnNameDecoration 定义装饰列名的接口
type ColumnNameDecoration interface {
	DecorateColumnName(name string) string
}

// PlainDecoration provides simple decoration that returns the name unchanged
// PlainDecoration 提供简单的装饰，返回不变的名称
type PlainDecoration struct{}

// NewPlainDecoration creates a new PlainDecoration instance
// NewPlainDecoration 创建一个新的 PlainDecoration 实例
func NewPlainDecoration() ColumnNameDecoration {
	return &PlainDecoration{}
}

// DecorateColumnName returns the column name without any modification
// DecorateColumnName 返回未经任何修改的列名
func (D *PlainDecoration) DecorateColumnName(name string) string {
	return name
}

// TableDecoration adds table prefix to column names
// TableDecoration 为列名添加表前缀
type TableDecoration struct {
	tableName string
}

// NewTableDecoration creates a new TableDecoration with the specified table name
// NewTableDecoration 使用指定的表名创建一个新的 TableDecoration
func NewTableDecoration(tableName string) ColumnNameDecoration {
	return &TableDecoration{tableName: tableName}
}

// DecorateColumnName adds table prefix to the column name if table name is not empty
// DecorateColumnName 如果表名不为空，则为列名添加表前缀
func (D *TableDecoration) DecorateColumnName(name string) string {
	if D.tableName != "" {
		return D.tableName + "." + name
	}
	return name
}

// CustomDecoration allows custom decoration logic via a function
// CustomDecoration 允许通过函数实现自定义装饰逻辑
type CustomDecoration struct {
	decorateFunc func(string) string
}

// NewCustomDecoration creates a new CustomDecoration with the provided function
// NewCustomDecoration 使用提供的函数创建一个新的 CustomDecoration
func NewCustomDecoration(decorateFunc func(name string) string) ColumnNameDecoration {
	return &CustomDecoration{decorateFunc: decorateFunc}
}

// DecorateColumnName applies the custom decoration function to the column name
// DecorateColumnName 将自定义装饰函数应用到列名上
func (D *CustomDecoration) DecorateColumnName(name string) string {
	return must.Nice(D.decorateFunc(name))
}
