package gormcnm

import "gorm.io/gorm/clause"

// Clause creates a ClauseColumn of the ColumnName instance.
// Clause 给 ColumnName 实例创建一个 ClauseColumn。
func (columnName ColumnName[TYPE]) Clause() *ClauseColumn[TYPE] {
	return &ClauseColumn[TYPE]{
		Table: "",
		Name:  columnName.Name(),
		Alias: "",
		Raw:   false, // 非 raw 的会被继续加工为合理的语句，比如增加表名，增加转义符号等，因此这里推荐使用 false（默认值）
	}
}

// ClauseWithTable creates a ClauseColumn with a specified table name.
// ClauseWithTable 创建一个带有指定表名的 ClauseColumn。
func (columnName ColumnName[TYPE]) ClauseWithTable(tableName string) *ClauseColumn[TYPE] {
	return &ClauseColumn[TYPE]{
		Table: tableName,
		Name:  columnName.Name(),
		Alias: "",
		Raw:   false, // 非 raw 的会被继续加工为合理的语句，比如增加表名，增加转义符号等，因此这里推荐使用 false（默认值）
	}
}

// ClauseColumn represents a column with additional properties such as table, alias, and raw flag.
// ClauseColumn 表示一个具有额外属性（如表名、别名和 raw 标志）的列。
type ClauseColumn[TYPE any] clause.Column

// WithTable sets the table name for the ClauseColumn and returns the updated ClauseColumn.
// WithTable 为 ClauseColumn 设置表名并返回更新后的 ClauseColumn。
func (clauseColumn *ClauseColumn[TYPE]) WithTable(tableName string) *ClauseColumn[TYPE] {
	clauseColumn.Table = tableName
	return clauseColumn
}

// WithAlias sets the alias for the ClauseColumn and returns the updated ClauseColumn.
// WithAlias 为 ClauseColumn 设置别名并返回更新后的 ClauseColumn。
func (clauseColumn *ClauseColumn[TYPE]) WithAlias(alias string) *ClauseColumn[TYPE] {
	clauseColumn.Alias = alias
	return clauseColumn
}

// WithRaw sets the raw flag for the ClauseColumn and returns the updated ClauseColumn.
// WithRaw 为 ClauseColumn 设置 raw 标志并返回更新后的 ClauseColumn。
func (clauseColumn *ClauseColumn[TYPE]) WithRaw(raw bool) *ClauseColumn[TYPE] {
	clauseColumn.Raw = raw
	return clauseColumn
}

// Column returns a GORM clause.Column, which represents a column in a SQL statement.
// Column 返回一个 GORM 的 clause.Column，表示 SQL 语句中的一个列。
func (clauseColumn *ClauseColumn[TYPE]) Column() clause.Column {
	return clause.Column(*clauseColumn)
}

// Assignment returns a GORM clause.Assignment, which is used to assign a value to a column.
// Assignment 返回一个 GORM 的 clause.Assignment，用于将值赋给列。
func (clauseColumn *ClauseColumn[TYPE]) Assignment(value TYPE) clause.Assignment {
	return clause.Assignment{Column: clauseColumn.Column(), Value: value}
}
