package gormcnm

import (
	"github.com/yyle88/gormcnm/internal/utils"
	"github.com/yyle88/tern/zerotern"
)

func (columnName ColumnName[TYPE]) COALESCE() *CoalesceNonNullGuardian {
	return NewCoalesceNonNullGuardian("COALESCE", string(columnName)) // COALESCE 是 SQL 标准中的函数，在大多数数据库系统中都支持
}

func (columnName ColumnName[TYPE]) IFNULLFN() *CoalesceNonNullGuardian {
	return NewCoalesceNonNullGuardian("IFNULL", string(columnName)) // IFNULL 是 MySQL 特定的函数，在其他数据库系统中可能不支持
}

type CoalesceNonNullGuardian struct {
	method string
	column string
}

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
