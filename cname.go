package gormcnm

import "github.com/yyle88/gormcnm/internal/utils"

/*
Defines a reusable `ColumnName` type designed to simplify and optimize SQL query construction.
It supports a wide range of SQL operations, including comparisons, equality checks, range queries,
pattern matching, and handling NULL values or boolean expressions. Additionally, it enables the use of column aliases
and retrieval of raw column names, providing flexibility and reducing the risk of manual errors in queries.
This utility improves code readability, reduces boilerplate, and simplifies complex query construction, making it adaptable to dynamic SQL generation scenarios.
*/

/*
该文件定义了一个通用且可复用的 `ColumnName` 类型，旨在简化 SQL 查询字符串的构建过程。
`ColumnName` 类型表示一个列名，并提供了丰富的方法来处理各种 SQL 操作，包括比较运算、等值检查、范围查询、模式匹配、以及对 NULL 或布尔值的处理。
此外，它还支持为列生成别名，方便在复杂查询中使用，并且可以轻松获取原始列名，提供了极大的灵活性。
通过避免手动拼接查询字符串，减少了出错的风险，提升了代码的可读性与可维护性。
该工具有效减少了样板代码，且使 SQL 查询更加简洁和易于扩展，特别适用于动态生成或复杂的查询场景。
*/

// ColumnName represents a generic column name for use in SQL queries
// ColumnName 表示一个通用的列名 可用于 SQL 查询
type ColumnName[TYPE any] string

// Qs creates a SQL statement with a given operator.
// Qs 创建一个带有指定操作符的 SQL 语句。
func (columnName ColumnName[TYPE]) Qs(op string) string {
	return string(columnName) + " " + op
}

// Op creates a SQL statement with an operator and a parameter.
// Op 创建一个带有操作符和参数的 SQL 语句。
func (columnName ColumnName[TYPE]) Op(op string, x TYPE) (string, TYPE) {
	return string(columnName) + " " + op, x
}

// Eq creates a SQL statement to check if the column is equal to a given value.
// Eq 创建一个 SQL 语句来判断列是否等于给定的值。
func (columnName ColumnName[TYPE]) Eq(x TYPE) (string, TYPE) {
	return string(columnName) + "=?", x
}

// Gt creates a SQL statement to check if the column is greater than a given value.
// Gt 创建一个 SQL 语句来判断列是否大于给定的值。
func (columnName ColumnName[TYPE]) Gt(x TYPE) (string, TYPE) {
	return string(columnName) + ">?", x
}

// Lt creates a SQL statement to check if the column is less than a given value.
// Lt 创建一个 SQL 语句来判断列是否小于给定的值。
func (columnName ColumnName[TYPE]) Lt(x TYPE) (string, TYPE) {
	return string(columnName) + "<?", x
}

// Gte creates a SQL statement to check if the column is greater than or equal to a given value.
// Gte 创建一个 SQL 语句来判断列是否大于等于给定的值。
func (columnName ColumnName[TYPE]) Gte(x TYPE) (string, TYPE) {
	return string(columnName) + ">=?", x
}

// Lte creates a SQL statement to check if the column is less than or equal to a given value.
// Lte 创建一个 SQL 语句来判断列是否小于等于给定的值。
func (columnName ColumnName[TYPE]) Lte(x TYPE) (string, TYPE) {
	return string(columnName) + "<=?", x
}

// Ne creates a SQL statement to check if the column is not equal to a given value.
// Ne 创建一个 SQL 语句来判断列是否不等于给定的值。
func (columnName ColumnName[TYPE]) Ne(x TYPE) (string, TYPE) {
	return string(columnName) + "!=?", x
}

// In creates a SQL statement to check if the column's value is in a given list of values.
// In 创建一个 SQL 语句来判断列的值是否在给定的值列表中。
func (columnName ColumnName[TYPE]) In(x []TYPE) (string, []TYPE) {
	return string(columnName) + " IN(?)", x
}

// NotIn creates a SQL statement to check if the column's value is not in a given list of values.
// NotIn 创建一个 SQL 语句来判断列的值是否不在给定的值列表中。
func (columnName ColumnName[TYPE]) NotIn(x []TYPE) (string, []TYPE) {
	return string(columnName) + " NOT IN(?)", x
}

// Like creates a SQL statement to check if the column's value matches a given pattern.
// Like 创建一个 SQL 语句来判断列的值是否匹配给定的模式。
func (columnName ColumnName[TYPE]) Like(x TYPE) (string, TYPE) {
	return string(columnName) + " LIKE ?", x
}

// NotLike creates a SQL statement to check if the column's value does not match a given pattern.
// NotLike 创建一个 SQL 语句来判断列的值是否不匹配给定的模式。
func (columnName ColumnName[TYPE]) NotLike(x TYPE) (string, TYPE) {
	return string(columnName) + " NOT LIKE ?", x
}

// NotEq creates a SQL statement to check if the column is not equal to a given value.
// NotEq 创建一个 SQL 语句来判断列是否不等于给定的值。
func (columnName ColumnName[TYPE]) NotEq(x TYPE) (string, TYPE) {
	return string(columnName) + "!=?", x
}

// IsNULL creates a SQL statement to check if the column is NULL.
// IsNULL 创建一个 SQL 语句来判断列是否为 NULL。
func (columnName ColumnName[TYPE]) IsNULL() string {
	return string(columnName) + " IS NULL"
}

// IsNull creates a SQL statement to check if the column is NULL.
// IsNull 创建一个 SQL 语句来判断列是否为 NULL。
func (columnName ColumnName[TYPE]) IsNull() string {
	return string(columnName) + " IS NULL"
}

// IsNotNULL creates a SQL statement to check if the column is not NULL.
// IsNotNULL 创建一个 SQL 语句来判断列是否不为 NULL。
func (columnName ColumnName[TYPE]) IsNotNULL() string {
	return string(columnName) + " IS NOT NULL"
}

// IsNotNull creates a SQL statement to check if the column is not NULL.
// IsNotNull 创建一个 SQL 语句来判断列是否不为 NULL。
func (columnName ColumnName[TYPE]) IsNotNull() string {
	return string(columnName) + " IS NOT NULL"
}

// IsTRUE creates a SQL statement to check if the column's value is TRUE.
// IsTRUE 创建一个 SQL 语句来判断列的值是否为 TRUE。
func (columnName ColumnName[TYPE]) IsTRUE() string {
	return string(columnName) + " IS TRUE"
}

// IsTrue creates a SQL statement to check if the column's value is TRUE.
// IsTrue 创建一个 SQL 语句来判断列的值是否为 TRUE。
func (columnName ColumnName[TYPE]) IsTrue() string {
	return string(columnName) + " IS TRUE"
}

// IsFALSE creates a SQL statement to check if the column's value is FALSE.
// IsFALSE 创建一个 SQL 语句来判断列的值是否为 FALSE。
func (columnName ColumnName[TYPE]) IsFALSE() string {
	return string(columnName) + " IS FALSE"
}

// IsFalse creates a SQL statement to check if the column's value is FALSE.
// IsFalse 创建一个 SQL 语句来判断列的值是否为 FALSE。
func (columnName ColumnName[TYPE]) IsFalse() string {
	return string(columnName) + " IS FALSE"
}

// BetweenAND creates a SQL statement to check if the column's value is between two given values.
// BetweenAND 创建一个 SQL 语句来判断列的值是否介于两个给定的值之间。
func (columnName ColumnName[TYPE]) BetweenAND(arg1, arg2 TYPE) (string, TYPE, TYPE) {
	return string(columnName) + " BETWEEN ? AND ?", arg1, arg2
}

// BetweenAnd creates a SQL statement to check if the column's value is between two given values.
// BetweenAnd 创建一个 SQL 语句来判断列的值是否介于两个给定的值之间。
func (columnName ColumnName[TYPE]) BetweenAnd(arg1, arg2 TYPE) (string, TYPE, TYPE) {
	return string(columnName) + " BETWEEN ? AND ?", arg1, arg2
}

// OnEq creates a SQL statement to check if the column is equal to another column in an ON clause.
// OnEq 创建一个 SQL 语句来判断列是否在 ON 子句中等于另一个列。
func (columnName ColumnName[TYPE]) OnEq(name ColumnName[TYPE]) string {
	return string(columnName) + "=" + string(name)
}

// OnNe creates a SQL statement to check if the column is not equal to another column in an ON clause.
// OnNe 创建一个 SQL 语句来判断列是否在 ON 子句中不等于另一个列。
func (columnName ColumnName[TYPE]) OnNe(name ColumnName[TYPE]) string {
	return string(columnName) + "!=" + string(name)
}

// Name returns the raw column name.
// Name 返回原始的列名。
func (columnName ColumnName[TYPE]) Name() string {
	return string(columnName)
}

// RawName returns the raw column name.
// RawName 返回原始的列名。
func (columnName ColumnName[TYPE]) RawName() string {
	return string(columnName)
}

// AsAlias returns the column name with applied alias.
// AsAlias 返回带有别名的列名。
func (columnName ColumnName[TYPE]) AsAlias(alias string) string {
	return utils.ApplyAliasToColumn(columnName.Name(), alias)
}

// AsName generates a SQL alias for the column using another ColumnName as the alias.
// AsName 使用另一个 ColumnName 作为别名生成列的 SQL 别名。
func (columnName ColumnName[TYPE]) AsName(alias ColumnName[TYPE]) string {
	return utils.ApplyAliasToColumn(columnName.Name(), string(alias))
}
