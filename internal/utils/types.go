package utils

type ColumnNameInterface interface {
	Name() string
}

type GormTableNameFace interface {
	TableName() string
}

type ClassImpTableName struct {
	tableName string
}

func NewClassImpTableName(tableName string) *ClassImpTableName {
	return &ClassImpTableName{tableName: tableName}
}

func (X *ClassImpTableName) TableName() string {
	return X.tableName
}
