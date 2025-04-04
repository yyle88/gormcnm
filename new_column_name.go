package gormcnm

import "github.com/yyle88/must"

func New[T any](name string) ColumnName[T] {
	return ColumnName[T](name)
}

func Cnm[T any](v T, name string) ColumnName[T] {
	return ColumnName[T](name)
}

func Cmn[T any](v T, name string, decoration ColumnNameDecoration) ColumnName[T] {
	return ColumnName[T](decoration.DecorateColumnName(name))
}

type ColumnNameDecoration interface {
	DecorateColumnName(name string) string
}

type PlainDecoration struct{}

func NewPlainDecoration() ColumnNameDecoration {
	return &PlainDecoration{}
}

func (D *PlainDecoration) DecorateColumnName(name string) string {
	return name
}

type TableDecoration struct {
	tableName string
}

func NewTableDecoration(tableName string) ColumnNameDecoration {
	return &TableDecoration{tableName: tableName}
}

func (D *TableDecoration) DecorateColumnName(name string) string {
	if D.tableName != "" {
		return D.tableName + "." + name
	}
	return name
}

type CustomDecoration struct {
	decorateFunc func(string) string
}

func NewCustomDecoration(decorateFunc func(name string) string) ColumnNameDecoration {
	return &CustomDecoration{decorateFunc: decorateFunc}
}

func (D *CustomDecoration) DecorateColumnName(name string) string {
	return must.Nice(D.decorateFunc(name))
}
