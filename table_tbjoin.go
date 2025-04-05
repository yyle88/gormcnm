package gormcnm

import (
	"strings"

	"gorm.io/gorm/clause"
)

// LEFTJOIN creates a left join operation for the specified table.
// LEFTJOIN 给指定表创建一个左连接操作。
func (common *ColumnOperationClass) LEFTJOIN(tableName string) *TableJoin {
	return newTableJoin(clause.LeftJoin, tableName)
}

// RIGHTJOIN creates a right join operation for the specified table.
// RIGHTJOIN 给指定表创建一个右连接操作。
func (common *ColumnOperationClass) RIGHTJOIN(tableName string) *TableJoin {
	return newTableJoin(clause.RightJoin, tableName)
}

// INNERJOIN creates an inner join operation for the specified table.
// INNERJOIN 给指定表创建一个内连接操作。
func (common *ColumnOperationClass) INNERJOIN(tableName string) *TableJoin {
	return newTableJoin(clause.InnerJoin, tableName)
}

// CROSSJOIN creates a cross join operation for the specified table.
// CROSSJOIN 给指定表创建一个交叉连接操作。
func (common *ColumnOperationClass) CROSSJOIN(tableName string) *TableJoin {
	return newTableJoin(clause.CrossJoin, tableName)
}

// TableJoin represents a join operation on a table, including its type and name.
// TableJoin 表示表上的连接操作，包括连接类型和表名。
type TableJoin struct {
	whichJoin clause.JoinType // Type of join (e.g., LEFT, RIGHT, INNER, CROSS)
	tableName string          // Name of the table involved in the join
}

// newTableJoin creates a new TableJoin instance with the specified table name and join type.
// newTableJoin 使用指定的表名和连接类型创建一个新的 TableJoin 实例。
func newTableJoin(whichJoin clause.JoinType, tableName string) *TableJoin {
	return &TableJoin{
		whichJoin: whichJoin,
		tableName: tableName,
	}
}

// On generates the SQL ON clause for the join, combining multiple statements with "AND".
// On 给连接生成 SQL 的 ON 子句，将多个语句用 "AND" 组合。
func (op *TableJoin) On(stmts ...string) string {
	return string(op.whichJoin) + " JOIN " + op.tableName + " ON " + strings.Join(stmts, " AND ")
}
