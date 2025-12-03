// Package gormcnm provides table JOIN operations to build multi-table data relationships
// Auto creates LEFT JOIN, INNER JOIN, and custom JOIN operations with ON conditions
// Supports building complex table relationships with type-safe join clause construction
//
// gormcnm 提供表 JOIN 操作，用于构建多表查询关系
// 自动创建 LEFT JOIN、INNER JOIN 和自定义 JOIN 操作，包含 ON 条件
// 支持构建复杂的表关系，具备类型安全的连接子句构建
package gormcnm

import (
	"strings"

	"gorm.io/gorm/clause"
)

// LEFTJOIN creates a left join operation on the specified table.
// LEFTJOIN 给指定表创建一个左连接操作。
func (common *ColumnOperationClass) LEFTJOIN(tableName string) *TableJoin {
	return newTableJoin(clause.LeftJoin, tableName)
}

// RIGHTJOIN creates a right join operation on the specified table.
// RIGHTJOIN 给指定表创建一个右连接操作。
func (common *ColumnOperationClass) RIGHTJOIN(tableName string) *TableJoin {
	return newTableJoin(clause.RightJoin, tableName)
}

// INNERJOIN creates an INNER join operation on the specified table.
// INNERJOIN 给指定表创建一个内连接操作。
func (common *ColumnOperationClass) INNERJOIN(tableName string) *TableJoin {
	return newTableJoin(clause.InnerJoin, tableName)
}

// CROSSJOIN creates a cross join operation on the specified table.
// CROSSJOIN 给指定表创建一个交叉连接操作。
func (common *ColumnOperationClass) CROSSJOIN(tableName string) *TableJoin {
	return newTableJoin(clause.CrossJoin, tableName)
}

// TableJoin represents a join operation on a table, including its type and name
// Supports every standard SQL join type: LEFT, RIGHT, INNER, and CROSS joins
// Auto generates correct SQL JOIN syntax with ON clause conditions
//
// TableJoin 表示表上的连接操作，包括连接类型和表名
// 支持所有标准 SQL 连接类型：LEFT、RIGHT、INNER 和 CROSS 连接
// 自动生成适当的 SQL JOIN 语法和 ON 子句条件
type TableJoin struct {
	whichJoin clause.JoinType // Type of join (LEFT, RIGHT, INNER, CROSS) // 连接类型（LEFT、RIGHT、INNER、CROSS）
	tableName string          // Name of the table involved in the join // 参与连接的表名
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
