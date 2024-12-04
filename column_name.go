package gormcnm

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SafeCnm returns a safe column name by enclosing it in backticks.
// SafeCnm 返回一个安全的列名，将其用反引号括起来。
// If the column name conflicts with a SQL keyword (e.g., "create"), enclosing it in backticks ensures proper execution.
// 如果列名与SQL关键字（例如"create"）冲突，使用反引号将其括起来，确保正确执行。
// This function is useful when using db.Select("`type`").Find(&one) as an example.
// 该函数在使用 db.Select("`type`").Find(&one) 等查询时非常有用。
func (s ColumnName[TYPE]) SafeCnm(quote string) ColumnName[TYPE] {
	switch len(quote) {
	case 0: // If no quote is provided, we simply add spaces around the column name.
		// 如果没有提供引号，我们将列名周围加上空格作为默认处理。
		return ColumnName[TYPE](" " + string(s) + " ")
	case 1: // If the quote is a single character like "`", `"`
		// 如果引号是一个单字符（例如 "`" 或 `"`）
		return ColumnName[TYPE](quote + string(s) + quote)
	case 2: // If the quote is two characters like `""`, "``", "[]"
		// 如果引号是两个字符（例如 `""`、 "``"、 "[]")
		return ColumnName[TYPE](quote[:1] + string(s) + quote[1:])
	default: // Not recommended
		// 这里不推荐这样做
		return ColumnName[TYPE](quote[:1] + string(s) + quote[len(quote)-1:])
	}
}

// ExprAdd creates a GORM expression to add a value to the column.
// ExprAdd: 创建一个GORM表达式，将一个值加到列中。
func (s ColumnName[TYPE]) ExprAdd(v TYPE) clause.Expr {
	return gorm.Expr(string(s)+" + ?", v)
}

// ExprSub creates a GORM expression to subtract a value from the column.
// ExprSub: 创建一个GORM表达式，从列中减去一个值。
func (s ColumnName[TYPE]) ExprSub(v TYPE) clause.Expr {
	return gorm.Expr(string(s)+" - ?", v)
}

// Ob creates an order-by clause for the column with the specified direction (ASC or DESC).
// Ob: 创建一个带有指定方向（ASC或DESC）的ORDER BY子句。
func (s ColumnName[TYPE]) Ob(direction string) OrderByBottle {
	return OrderByBottle(string(s) + " " + direction)
}

// OrderByDirection creates an order-by clause for the column with the given direction (ASC or DESC).
// OrderByDirection: 创建一个带有给定方向（ASC或DESC）的ORDER BY子句。
func (s ColumnName[TYPE]) OrderByDirection(direction string) OrderByBottle {
	return OrderByBottle(string(s) + " " + direction)
}

// Qc creates a condition for the column using the provided operator (e.g., '=', '>', etc.).
// Qc: 使用提供的运算符（例如 '=', '>', 等）为列创建条件。
func (s ColumnName[TYPE]) Qc(op string) QsConjunction {
	return QsConjunction(string(s) + " " + op)
}

// ColumnCondition creates a condition for the column using the provided operator (e.g., '=', '>', etc.).
// ColumnCondition: 使用提供的运算符（例如 '=', '>', 等）为列创建条件。
func (s ColumnName[TYPE]) ColumnCondition(op string) QsConjunction {
	return QsConjunction(string(s) + " " + op)
}

// Qx creates a condition with an operator and a value for the column, useful for building complex queries.
// Qx: 创建带有运算符和值的条件，用于构建复杂的查询。
func (s ColumnName[TYPE]) Qx(op string, x TYPE) *QxType {
	stmt := string(s.Qc(op))
	args := []interface{}{x}
	return NewQx(stmt, args...)
}

// ColumnConditionWithValue creates a condition with an operator and a value for the column, useful for building complex queries.
// ColumnConditionWithValue: 创建带有运算符和值的列条件，用于构建复杂的查询。
func (s ColumnName[TYPE]) ColumnConditionWithValue(op string, x TYPE) *QxType {
	stmt := string(s.ColumnCondition(op))
	args := []interface{}{x}
	return NewQx(stmt, args...)
}

// Kw creates a map with a single key-value.
// Kw: 创建一个包含单个键值对的map。
func (s ColumnName[TYPE]) Kw(x TYPE) ColumnValueMap {
	return ColumnValueMap{string(s): x}
}

// CreateColumnValueMap creates a map with a single key-value.
// CreateColumnValueMap: 创建一个包含单个键值对的map。
func (s ColumnName[TYPE]) CreateColumnValueMap(x TYPE) ColumnValueMap {
	return ColumnValueMap{string(s): x}
}

// Kv returns a key-value, works with GORM's Update function.
// Kv: 返回键值对，适用于GORM的Update函数。
func (s ColumnName[TYPE]) Kv(x TYPE) (string, TYPE) {
	return string(s), x
}

// ColumnKeyAndValue returns a key-value, useful for GORM's Update function.
// ColumnKeyAndValue: 返回键值对，适用于GORM的Update函数。
func (s ColumnName[TYPE]) ColumnKeyAndValue(x TYPE) (string, TYPE) {
	return string(s), x
}

// KeExp extends Kv by returning a key-expression, useful in GORM's Update function with expressions.
// KeExp: 扩展Kv，返回键表达式对，适用于GORM的Update函数，支持表达式。
func (s ColumnName[TYPE]) KeExp(x clause.Expr) (string, clause.Expr) {
	return string(s), x
}

// KeAdd is used for updates where a value is added to the field (e.g., incrementing a value).
// KeAdd: 用于更新字段时，将一个值加到字段上（例如递增一个值）。
func (s ColumnName[TYPE]) KeAdd(x TYPE) (string, clause.Expr) {
	return s.KeExp(s.ExprAdd(x))
}

// KeSub is used for updates where a value is subtracted from the field (e.g., decrementing a value).
// KeSub: 用于更新字段时，将一个值从字段中减去（例如递减一个值）。
func (s ColumnName[TYPE]) KeSub(x TYPE) (string, clause.Expr) {
	return s.KeExp(s.ExprSub(x))
}

// Count creates a COUNT query for the column, excluding NULL values.
// Count: 创建一个COUNT查询，只统计非NULL值的列。
func (s ColumnName[TYPE]) Count(alias string) string {
	return applyAliasToColumn("COUNT("+string(s)+")", alias)
}

// CountDistinct creates a COUNT DISTINCT query for the given column, skipping NULL values in the count.
// CountDistinct: 创建一个COUNT DISTINCT查询，用于给定列，跳过NULL值。
func (s ColumnName[TYPE]) CountDistinct(alias string) string {
	return applyAliasToColumn("COUNT(DISTINCT("+string(s)+"))", alias)
}
