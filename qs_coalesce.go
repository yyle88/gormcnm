// Package gormcnm provides COALESCE and IFNULL operations for NULL-safe SQL queries
// Auto handles NULL values in aggregate functions using COALESCE (standard) or IFNULL (MySQL)
// Supports SUM, COUNT, AVG, MAX, MIN with automatic NULL value protection
//
// gormcnm 提供 COALESCE 和 IFNULL 操作，实现 NULL 安全的 SQL 查询
// 自动使用 COALESCE（标准）或 IFNULL（MySQL）处理聚合函数中的 NULL 值
// 支持 SUM、COUNT、AVG、MAX、MIN，具备自动 NULL 值保护
package gormcnm

import (
	"github.com/yyle88/gormcnm/internal/utils"
	"github.com/yyle88/tern/zerotern"
)

// COALESCE creates a COALESCE function wrapper for handling NULL values in SQL queries
// Auto uses SQL standard COALESCE function, supported by most database systems
// COALESCE 为处理 SQL 查询中的 NULL 值创建 COALESCE 函数包装器
// 自动使用 SQL 标准的 COALESCE 函数，被大多数数据库系统支持
func (columnName ColumnName[TYPE]) COALESCE() *CoalesceNonNullGuardian {
	return NewCoalesceNonNullGuardian("COALESCE", string(columnName))
}

// IFNULLFN creates an IFNULL function wrapper for MySQL-specific NULL handling
// MySQL-specific function, may not be supported in other database systems
// IFNULLFN 为 MySQL 特定的 NULL 处理创建 IFNULL 函数包装器
// MySQL 特定函数，在其他数据库系统中可能不受支持
func (columnName ColumnName[TYPE]) IFNULLFN() *CoalesceNonNullGuardian {
	return NewCoalesceNonNullGuardian("IFNULL", string(columnName))
}

// CoalesceNonNullGuardian provides SQL aggregate functions with NULL value protection
// Auto handles NULL values using COALESCE or IFNULL functions to provide default values
// CoalesceNonNullGuardian 提供带有 NULL 值保护的 SQL 聚合函数
// 自动使用 COALESCE 或 IFNULL 函数处理 NULL 值以提供默认值
type CoalesceNonNullGuardian struct {
	method string // SQL function name (COALESCE or IFNULL) // SQL 函数名（COALESCE 或 IFNULL）
	column string // Column name to apply the function to // 要应用函数的列名
}

// NewCoalesceNonNullGuardian creates a new CoalesceNonNullGuardian with specified method and column
// NewCoalesceNonNullGuardian 使用指定的方法和列名创建新的 CoalesceNonNullGuardian
func NewCoalesceNonNullGuardian(methodName string, columnName string) *CoalesceNonNullGuardian {
	return &CoalesceNonNullGuardian{
		method: methodName,
		column: columnName,
	}
}

// Stmt generates an SQL statement for the COALESCE or IFNULL function with the given function and default value.
// Stmt 生成一个 SQL 语句，包含 COALESCE 或 IFNULL 函数，并指定默认值。
func (qs *CoalesceNonNullGuardian) Stmt(sfn string, dfv string, alias string) string {
	return utils.ApplyAliasToColumn(qs.method+"("+sfn+"("+string(qs.column)+"), "+zerotern.VV(dfv, "0")+")", alias)
}

// SumStmt generates an SQL statement to calculate the sum of the column, using 0 as the default value.
// SumStmt 生成一个 SQL 语句，计算列的总和，默认值为 0。
func (qs *CoalesceNonNullGuardian) SumStmt(alias string) string {
	return qs.Stmt("SUM", "0", alias)
}

// MaxStmt generates an SQL statement to retrieve the maximum value of the column, using 0 as the default value.
// MaxStmt 生成一个 SQL 语句，检索列的最大值，默认值为 0。
func (qs *CoalesceNonNullGuardian) MaxStmt(alias string) string {
	return qs.Stmt("MAX", "0", alias)
}

// MinStmt generates an SQL statement to retrieve the minimum value of the column, using 0 as the default value.
// MinStmt 生成一个 SQL 语句，检索列的最小值，默认值为 0。
func (qs *CoalesceNonNullGuardian) MinStmt(alias string) string {
	return qs.Stmt("MIN", "0", alias)
}

// AvgStmt generates an SQL statement to calculate the average value of the column, using 0 as the default value.
// AvgStmt 生成一个 SQL 语句，计算列的平均值，默认值为 0。
func (qs *CoalesceNonNullGuardian) AvgStmt(alias string) string {
	return qs.Stmt("AVG", "0", alias)
}
