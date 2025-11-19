// Package gormcnm provides table-qualified column operations for multi-table queries
// Auto creates columns with table prefixes for JOIN operations and complex queries
// Supports building fully-qualified column names with table association
//
// gormcnm 提供表限定列操作，用于多表查询
// 自动创建带表前缀的列，用于 JOIN 操作和复杂查询
// 支持构建完全限定的列名，具备表关联功能
package gormcnm

import "github.com/yyle88/gormcnm/internal/utils"

// TB creates a TableColumn by associating this column with a table interface
// TB 通过将此列与表接口关联创建 TableColumn
func (columnName ColumnName[TYPE]) TB(tab utils.GormTableNameFace) *TableColumn[TYPE] {
	return columnName.WithTable(tab)
}

// TC creates a TableColumn by associating this column with a table interface (alias for TB)
// TC 通过将此列与表接口关联创建 TableColumn（TB 的别名）
func (columnName ColumnName[TYPE]) TC(tab utils.GormTableNameFace) *TableColumn[TYPE] {
	return columnName.WithTable(tab)
}

// TN creates a TableColumn by associating this column with a table name string
// TN 通过将此列与表名字符串关联创建 TableColumn
func (columnName ColumnName[TYPE]) TN(tableName string) *TableColumn[TYPE] {
	return columnName.WithTable(utils.NewTableNameImp(tableName))
}

func (columnName ColumnName[TYPE]) WithTable(tab utils.GormTableNameFace) *TableColumn[TYPE] {
	return &TableColumn[TYPE]{
		tab: tab,
		cnm: columnName,
	}
}

func (columnName ColumnName[TYPE]) WithTableName(tableName string) *TableColumn[TYPE] {
	return columnName.WithTable(utils.NewTableNameImp(tableName))
}

// TableColumn represents a combination of a table and a column.
// TableColumn 表示表和列的组合。
type TableColumn[TYPE any] struct {
	tab utils.GormTableNameFace
	cnm ColumnName[TYPE]
}

// Eq generates an equality condition in SQL format, ensuring type consistency between two columns.
// Eq 生成 SQL 格式的相等条件，确保两列之间的类型一致。
func (tc *TableColumn[TYPE]) Eq(xc *TableColumn[TYPE]) string {
	return tc.Name() + " = " + xc.Name()
}

// Ne generates a SQL inequality condition, ensuring type consistency between two columns.
// Ne 生成 SQL 格式的不等条件，确保两列之间的类型一致。
func (tc *TableColumn[TYPE]) Ne(xc *TableColumn[TYPE]) string {
	return tc.Name() + " != " + xc.Name()
}

// Op generates a custom SQL operation between two columns using the specified operator.
// Op 使用指定的操作符生成两列之间的自定义 SQL 操作。
func (tc *TableColumn[TYPE]) Op(op string, xc *TableColumn[TYPE]) string {
	return tc.Name() + " " + op + " " + xc.Name()
}

// Name returns the fully qualified name of the column in the format "table.column".
// Name 返回列的完全限定名称，格式为 "table.column"。
func (tc *TableColumn[TYPE]) Name() string {
	return tc.tab.TableName() + "." + tc.cnm.Name()
}

// ColumnName retrieves the column name in a ColumnName format, representing the combination of the table and column.
// ColumnName 获取以 ColumnName 格式表示的列名，代表表和列的组合。
func (tc *TableColumn[TYPE]) ColumnName() ColumnName[TYPE] {
	return ColumnName[TYPE](tc.Name())
}

// Cnm retrieves the column name in a ColumnName format, representing the combination of the table and column.
// Cnm 获取以 ColumnName 格式表示的列名，代表表和列的组合。
func (tc *TableColumn[TYPE]) Cnm() ColumnName[TYPE] {
	return ColumnName[TYPE](tc.Name())
}

// Ob creates an OrderByBottle object for specifying ordering based on the column name and direction.
// Ob 基于列名和方向创建一个 OrderByBottle 对象用于指定排序。
func (tc *TableColumn[TYPE]) Ob(direction string) OrderByBottle {
	return tc.Cnm().Ob(direction)
}

// AsAlias generates a SQL alias for the column in the format "table.column AS alias".
// AsAlias 生成列的 SQL 别名，格式为 "table.column AS alias"。
func (tc *TableColumn[TYPE]) AsAlias(alias string) string {
	return utils.ApplyAliasToColumn(tc.Name(), alias)
}

// AsName generates a SQL alias for the column using another ColumnName as the alias.
// AsName 使用另一个 ColumnName 作为别名生成列的 SQL 别名。
func (tc *TableColumn[TYPE]) AsName(alias ColumnName[TYPE]) string {
	return utils.ApplyAliasToColumn(tc.Name(), alias.Name())
}
