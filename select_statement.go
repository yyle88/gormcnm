// Package gormcnm provides SELECT statement building operations for custom column selection
// Auto constructs SELECT clauses with column combinations, aliases, and aggregate functions
// Supports building complex SELECT queries with type-safe column management and GORM integration
//
// gormcnm 提供 SELECT 语句构建操作，用于自定义列选择
// 自动构建 SELECT 子句，包含列组合、别名和聚合函数
// 支持构建复杂的 SELECT 操作，具备类型安全的列管理和 GORM 集成
package gormcnm

import (
	"strings"

	"gorm.io/gorm"
)

// SxType is an alias when using SelectStatement as a short type name
// SxType 是 SelectStatement 的别名，用作简短的类型名称
type SxType = SelectStatement

// NewSx creates a new SelectStatement instance with the provided statement and arguments
// NewSx 使用提供的语句和参数创建一个新的 SelectStatement 实例
func NewSx(stmt string, args ...interface{}) *SxType {
	return &SxType{
		statementArgumentsTuple: newStatementArgumentsTuple(stmt, args),
	}
}

// SelectStatement represents a SELECT statement with arguments used in GORM db.Select operations
// Handles complex SELECT scenarios where conditions and arguments are needed
// Auto combines multiple select statements with comma separation to form multi-column queries
//
// Usage scenarios:
// - 99% of cases db.Select requires no parameters
// - However, with complex queries such as: SELECT COUNT(CASE WHEN condition THEN 1 END) as cnt
// - Need to merge multiple column select statements and corresponding parameters
// - Columns are separated using commas
//
// SelectStatement 表示用于 GORM db.Select 操作的 SELECT 语句及参数
// 处理需要条件和参数的复杂 SELECT 场景
// 自动使用逗号分隔合并多个 SELECT 语句进行多列查询
//
// 使用场景说明：
// - 99%的情况下 db.Select 不需要参数
// - 但对于复杂查询如：SELECT COUNT(CASE WHEN condition THEN 1 END) as cnt
// - 需要合并多个列的 SELECT 语句和对应参数
// - 各列之间使用逗号分隔
type SelectStatement struct {
	*statementArgumentsTuple // Embedded statement-arguments tuple // 嵌入的语句-参数元组
}

// NewSelectStatement creates a new SelectStatement with the provided query string and arguments.
// NewSelectStatement 使用提供的查询字符串和参数创建一个新的 SelectStatement 实例。
func NewSelectStatement(stmt string, args ...interface{}) *SelectStatement {
	return &SelectStatement{
		statementArgumentsTuple: newStatementArgumentsTuple(stmt, args),
	}
}

// Combine combines the current SelectStatement with other SelectStatements by merging their query strings and arguments.
// Combine 将当前的 SelectStatement 与其他 SelectStatement 合并，通过合并它们的查询字符串和参数。
func (sx *SelectStatement) Combine(cs ...*SelectStatement) *SelectStatement {
	var qsVs []string
	qsVs = append(qsVs, sx.Qs())
	var args []any
	args = append(args, sx.Args()...)
	for _, c := range cs {
		qsVs = append(qsVs, c.Qs())
		args = append(args, c.Args()...)
	}
	var stmt = strings.Join(qsVs, ", ")      //得到的就是gorm db.Select() 的要选中的列信息，因此使用逗号分隔
	return NewSelectStatement(stmt, args...) //得到的就是 gorm db.Select() 的选中信息和附带的参数信息，比如 COUNT(CASE WHEN condition THEN 1 END) 里 condition 的参数信息
}

// Scope converts the SelectStatement to a GORM ScopeFunction used with db.Scopes().
// It applies the SELECT operation defined by SelectStatement to the GORM select.
// Scope 将 SelectStatement 转换为 GORM 的 ScopeFunction，以便于被 db.Scopes() 调用。
// 它将 SelectStatement 定义的查询语句应用于 GORM 查询。
func (sx *SelectStatement) Scope() ScopeFunction {
	return func(db *gorm.DB) *gorm.DB {
		return db.Select(sx.Qs(), sx.args...)
	}
}
