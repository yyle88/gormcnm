package gormcnm

import (
	"strings"

	"gorm.io/gorm/clause"
)

func (c *ColumnOperationClass) LEFTJOIN(tableName string) *tbjoin {
	return newTbjoin(tableName, clause.LeftJoin)
}

func (c *ColumnOperationClass) RIGHTJOIN(tableName string) *tbjoin {
	return newTbjoin(tableName, clause.RightJoin)
}

func (c *ColumnOperationClass) INNERJOIN(tableName string) *tbjoin {
	return newTbjoin(tableName, clause.InnerJoin)
}

func (c *ColumnOperationClass) CROSSJOIN(tableName string) *tbjoin {
	return newTbjoin(tableName, clause.CrossJoin)
}

type tbjoin struct {
	whichJoin clause.JoinType
	tableName string
}

func newTbjoin(tableName string, whichJoin clause.JoinType) *tbjoin {
	return &tbjoin{
		whichJoin: whichJoin,
		tableName: tableName,
	}
}

func (op *tbjoin) On(stmts ...string) string {
	return string(op.whichJoin) + " JOIN " + op.tableName + " ON " + strings.Join(stmts, " AND ")
}
