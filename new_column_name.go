package gormcnm

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

func NewPlainDecoration() *PlainDecoration {
	return &PlainDecoration{}
}

func (options *PlainDecoration) DecorateColumnName(name string) string {
	return name
}

type TableDecoration struct {
	tableName string
}

func NewTableDecoration(tableName string) *TableDecoration {
	return &TableDecoration{tableName: tableName}
}

func (options *TableDecoration) DecorateColumnName(name string) string {
	if options.tableName != "" {
		return options.tableName + "." + name
	}
	return name
}

type CustomDecoration struct {
	decorateFunc func(string) string
}

func NewCustomDecoration(decorateFunc func(name string) string) *CustomDecoration {
	return &CustomDecoration{decorateFunc: decorateFunc}
}

func (d *CustomDecoration) DecorateColumnName(name string) string {
	return d.decorateFunc(name)
}
