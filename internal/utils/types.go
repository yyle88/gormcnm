package utils

type ColumnNameInterface interface {
	Name() string
}

type GormTableNameFace interface {
	TableName() string
}

type TableNameImp struct {
	tableName string
}

func NewTableNameImp(tableName string) *TableNameImp {
	return &TableNameImp{tableName: tableName}
}

func (X *TableNameImp) TableName() string {
	return X.tableName
}
