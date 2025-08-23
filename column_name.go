package gormcnm

import (
	"github.com/yyle88/gormcnm/internal/utils"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SafeCnm returns a safe column name by enclosing it in backticks.
// SafeCnm 返回一个安全的列名，将其用反引号括起来。
// If the column name conflicts with a SQL keyword (e.g., "create"), enclosing it in backticks ensures proper execution.
// 如果列名与SQL关键字（例如"create"）冲突，使用反引号将其括起来，确保正确执行。
// This function is useful when using db.Select("`type`").Find(&one) as an example.
// 该函数在使用 db.Select("`type`").Find(&one) 等查询时非常有用。
func (columnName ColumnName[TYPE]) SafeCnm(quote string) ColumnName[TYPE] {
	switch len(quote) {
	case 0: // If no quote is provided, we simply add spaces around the column name.
		// 如果没有提供引号，我们将列名周围加上空格作为默认处理。
		return ColumnName[TYPE](" " + string(columnName) + " ")
	case 1: // If the quote is a single character like "`", `"`
		// 如果引号是一个单字符（例如 "`" 或 `"`）
		return ColumnName[TYPE](quote + string(columnName) + quote)
	case 2: // If the quote is two characters like `""`, "``", "[]"
		// 如果引号是两个字符（例如 `""`、 "``"、 "[]")
		return ColumnName[TYPE](quote[:1] + string(columnName) + quote[1:])
	default: // Not recommended but not panic, use default quote
		// 不推荐这样做，但也不抛异常，给个默认的结果
		return ColumnName[TYPE](`"` + string(columnName) + `"`)
	}
}

// ExprAdd creates a GORM expression to add a value to the column.
// ExprAdd: 创建一个GORM表达式，将一个值加到列中。
func (columnName ColumnName[TYPE]) ExprAdd(v TYPE) clause.Expr {
	return gorm.Expr(string(columnName)+" + ?", v)
}

// ExprSub creates a GORM expression to subtract a value from the column.
// ExprSub: 创建一个GORM表达式，从列中减去一个值。
func (columnName ColumnName[TYPE]) ExprSub(v TYPE) clause.Expr {
	return gorm.Expr(string(columnName)+" - ?", v)
}

// ExprMul creates a GORM expression to multiply the column by a value.
// ExprMul: 创建一个GORM表达式，将列乘以一个值。
func (columnName ColumnName[TYPE]) ExprMul(v TYPE) clause.Expr {
	return gorm.Expr(string(columnName)+" * ?", v)
}

// ExprDiv creates a GORM expression to divide the column by a value.
// ExprDiv: 创建一个GORM表达式，将列除以一个值。
func (columnName ColumnName[TYPE]) ExprDiv(v TYPE) clause.Expr {
	return gorm.Expr(string(columnName)+" / ?", v)
}

// ExprConcat creates a GORM expression to concatenate a string to the column.
// ExprConcat: 创建一个GORM表达式，将字符串连接到列。
func (columnName ColumnName[TYPE]) ExprConcat(v TYPE) clause.Expr {
	return gorm.Expr("CONCAT("+string(columnName)+", ?)", v)
}

// ExprReplace creates a GORM expression to replace text in the column.
// ExprReplace: 创建一个GORM表达式，替换列中的文本。
func (columnName ColumnName[TYPE]) ExprReplace(oldValue, newValue TYPE) clause.Expr {
	return gorm.Expr("REPLACE("+string(columnName)+", ?, ?)", oldValue, newValue)
}

// Ob creates an order-by clause for the column with the specified direction (ASC or DESC).
// Ob: 创建一个带有指定方向（ASC或DESC）的ORDER BY子句。
func (columnName ColumnName[TYPE]) Ob(direction string) OrderByBottle {
	return OrderByBottle(string(columnName) + " " + direction)
}

// OrderByBottle creates an order-by clause for the column with the given direction (ASC or DESC).
// OrderByBottle: 创建一个带有给定方向（ASC或DESC）的ORDER BY子句。
func (columnName ColumnName[TYPE]) OrderByBottle(direction string) OrderByBottle {
	return OrderByBottle(string(columnName) + " " + direction)
}

// Qc creates a condition for the column using the provided operator (e.g., '=', '>', etc.).
// Qc: 使用提供的运算符（例如 '=', '>', 等）为列创建条件。
func (columnName ColumnName[TYPE]) Qc(op string) QsConjunction {
	return QsConjunction(string(columnName) + " " + op)
}

// ColumnCondition creates a condition for the column using the provided operator (e.g., '=', '>', etc.).
// ColumnCondition: 使用提供的运算符（例如 '=', '>', 等）为列创建条件。
func (columnName ColumnName[TYPE]) ColumnCondition(op string) QsConjunction {
	return QsConjunction(string(columnName) + " " + op)
}

// Qx creates a condition with an operator and a value for the column, useful for building complex queries.
// Qx: 创建带有运算符和值的条件，用于构建复杂的查询。
func (columnName ColumnName[TYPE]) Qx(op string, x TYPE) *QxConjunction {
	stmt := string(columnName.Qc(op))
	args := []interface{}{x}
	return NewQxConjunction(stmt, args...)
}

// ColumnConditionWithValue creates a condition with an operator and a value for the column, useful for building complex queries.
// ColumnConditionWithValue: 创建带有运算符和值的列条件，用于构建复杂的查询。
func (columnName ColumnName[TYPE]) ColumnConditionWithValue(op string, x TYPE) *QxConjunction {
	stmt := string(columnName.ColumnCondition(op))
	args := []interface{}{x}
	return NewQxConjunction(stmt, args...)
}

// Kw creates a map with a single key-value.
// Kw: 创建一个包含单个键值对的map。
func (columnName ColumnName[TYPE]) Kw(x TYPE) ColumnValueMap {
	return ColumnValueMap{string(columnName): x}
}

// CreateColumnValueMap creates a map with a single key-value.
// CreateColumnValueMap: 创建一个包含单个键值对的map。
func (columnName ColumnName[TYPE]) CreateColumnValueMap(x TYPE) ColumnValueMap {
	return ColumnValueMap{string(columnName): x}
}

// Kv returns a key-value, works with GORM's Update function.
// Kv: 返回键值对，适用于GORM的Update函数。
func (columnName ColumnName[TYPE]) Kv(x TYPE) (string, TYPE) {
	return string(columnName), x
}

// ColumnKeyAndValue returns a key-value, useful for GORM's Update function.
// ColumnKeyAndValue: 返回键值对，适用于GORM的Update函数。
func (columnName ColumnName[TYPE]) ColumnKeyAndValue(x TYPE) (string, TYPE) {
	return string(columnName), x
}

// KeExp extends Kv by returning a key-expression, useful in GORM's Update function with expressions.
// KeExp: 扩展Kv，返回键表达式对，适用于GORM的Update函数，支持表达式。
func (columnName ColumnName[TYPE]) KeExp(x clause.Expr) (string, clause.Expr) {
	return string(columnName), x
}

// KeAdd is used for updates where a value is added to the field (e.g., incrementing a value).
// Returns (columnName, gorm.Expr) tuple for use in UpdateColumns operations.
//
// With Gorm:
//
//	db.UpdateColumns(map[string]interface{}{"price": gorm.Expr("price + ?", 10)})
//
// With KeAdd:
//
//	db.UpdateColumns(map[string]interface{}{cls.Price.KeAdd(10)})
//
// Both generate: "UPDATE ... SET price = price + 10"
//
// KeAdd: 用于更新字段时，将一个值加到字段上（例如递增一个值）。
// 返回 (列名, gorm.Expr) 元组，用于 UpdateColumns 操作。
//
// 传统写法：
//
//	db.UpdateColumns(map[string]interface{}{"price": gorm.Expr("price + ?", 10)})
//
// 使用 KeAdd：
//
//	db.UpdateColumns(map[string]interface{}{cls.Price.KeAdd(10)})
//
// 都生成："UPDATE ... SET price = price + 10"
func (columnName ColumnName[TYPE]) KeAdd(x TYPE) (string, clause.Expr) {
	return columnName.KeExp(columnName.ExprAdd(x))
}

// KeSub is used for updates where a value is subtracted from the field (e.g., decrementing a value).
// Returns (columnName, gorm.Expr) tuple for use in UpdateColumns operations.
//
// With Gorm:
//
//	db.UpdateColumns(map[string]interface{}{"stock": gorm.Expr("stock - ?", 1)})
//
// With KeSub:
//
//	db.UpdateColumns(map[string]interface{}{cls.Stock.KeSub(1)})
//
// Both generate: "UPDATE ... SET stock = stock - 1"
//
// KeSub: 用于更新字段时，将一个值从字段中减去（例如递减一个值）。
// 返回 (列名, gorm.Expr) 元组，用于 UpdateColumns 操作。
//
// 传统写法：
//
//	db.UpdateColumns(map[string]interface{}{"stock": gorm.Expr("stock - ?", 1)})
//
// 使用 KeSub：
//
//	db.UpdateColumns(map[string]interface{}{cls.Stock.KeSub(1)})
//
// 都生成："UPDATE ... SET stock = stock - 1"
func (columnName ColumnName[TYPE]) KeSub(x TYPE) (string, clause.Expr) {
	return columnName.KeExp(columnName.ExprSub(x))
}

// KeMul is used for updates where a field is multiplied by a value (e.g., scaling a value).
// KeMul: 用于更新字段时，将字段乘以一个值（例如缩放一个值）。
func (columnName ColumnName[TYPE]) KeMul(x TYPE) (string, clause.Expr) {
	return columnName.KeExp(columnName.ExprMul(x))
}

// KeDiv is used for updates where a field is divided by a value (e.g., splitting a value).
// KeDiv: 用于更新字段时，将字段除以一个值（例如分割一个值）。
func (columnName ColumnName[TYPE]) KeDiv(x TYPE) (string, clause.Expr) {
	return columnName.KeExp(columnName.ExprDiv(x))
}

// KeConcat is used for updates where a string is concatenated to the field (e.g., appending text).
// KeConcat: 用于更新字段时，将字符串连接到字段（例如追加文本）。
func (columnName ColumnName[TYPE]) KeConcat(x TYPE) (string, clause.Expr) {
	return columnName.KeExp(columnName.ExprConcat(x))
}

// KeReplace is used for updates where text in the field is replaced (e.g., updating patterns).
// KeReplace: 用于更新字段时，替换字段中的文本（例如更新模式）。
func (columnName ColumnName[TYPE]) KeReplace(oldValue, newValue TYPE) (string, clause.Expr) {
	return columnName.KeExp(columnName.ExprReplace(oldValue, newValue))
}

// Count creates a COUNT query for the column, excluding NULL values.
// Count: 创建一个COUNT查询，只统计非NULL值的列。
func (columnName ColumnName[TYPE]) Count(alias string) string {
	return utils.ApplyAliasToColumn("COUNT("+string(columnName)+")", alias)
}

// CountDistinct creates a COUNT DISTINCT query for the given column, skipping NULL values in the count.
// CountDistinct: 创建一个COUNT DISTINCT查询，用于给定列，跳过NULL值。
func (columnName ColumnName[TYPE]) CountDistinct(alias string) string {
	return utils.ApplyAliasToColumn("COUNT(DISTINCT("+string(columnName)+"))", alias)
}
