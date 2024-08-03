package gormcnm

type classImplementsTableName struct {
	tableName string
}

func (X *classImplementsTableName) TableName() string {
	return X.tableName
}
