package gormcnm

import (
	"strings"

	"gorm.io/gorm/clause"
)

func (c *ColumnOperationClass) LEFTJOIN(tableName string) *OperationTABLEJOIN {
	return newOperationTABLEJOIN(tableName, clause.LeftJoin)
}

func (c *ColumnOperationClass) RIGHTJOIN(tableName string) *OperationTABLEJOIN {
	return newOperationTABLEJOIN(tableName, clause.RightJoin)
}

func (c *ColumnOperationClass) INNERJOIN(tableName string) *OperationTABLEJOIN {
	return newOperationTABLEJOIN(tableName, clause.InnerJoin)
}

func (c *ColumnOperationClass) CROSSJOIN(tableName string) *OperationTABLEJOIN {
	return newOperationTABLEJOIN(tableName, clause.CrossJoin)
}

type OperationTABLEJOIN struct {
	joinType  clause.JoinType
	tableName string
}

func newOperationTABLEJOIN(tableName string, joinType clause.JoinType) *OperationTABLEJOIN {
	return &OperationTABLEJOIN{
		joinType:  joinType,
		tableName: tableName,
	}
}

func (op *OperationTABLEJOIN) On(stmts ...string) string {
	return string(op.joinType) + " JOIN " + op.tableName + " ON " + strings.Join(stmts, " AND ")
}
