package gormcnm

func New[T any](name string) ColumnName[T] {
	return ColumnName[T](name)
}

func Cnm[T any](v T, name string) ColumnName[T] {
	return ColumnName[T](name)
}
