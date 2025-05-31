package gormcnm

import (
	"strings"

	"github.com/yyle88/gormcnm/internal/utils"
	"gorm.io/gorm"
)

// ColumnOperationClass provides a set of common methods for handling column operations in database queries.
// ColumnOperationClass 提供一组常用的数据库列操作工具函数。
type ColumnOperationClass struct{}

// OK always returns true and provides a simple boolean condition.
// Its primary purpose is to offer a compact scope in which variables are valid only within
// the conditional block. This helps improve code clarity by limiting the scope of variables
// and making the logic easier to follow.
// OK 函数始终返回true，提供一个简单的布尔条件供控制流使用。
// 它的主要目的是提供一个小的作用域，在这个作用域内定义的变量只在条件块内有效。
// 这种做法有助于提高代码的清晰度，通过限制变量的作用域，
// 使得逻辑更加清晰和易于理解。
func (common *ColumnOperationClass) OK() bool {
	return true
}

// CreateCondition creates a new QxConjunction with the provided statement and arguments.
// CreateCondition 根据提供的语句和参数创建一个新的QxConjunction。
func (common *ColumnOperationClass) CreateCondition(stmt string, args ...interface{}) *QxConjunction {
	return NewQxConjunction(stmt, args...)
}

// NewQx creates a new QxConjunction with the provided statement and arguments.
// NewQx 根据提供的语句和参数创建一个新的QxConjunction。
func (common *ColumnOperationClass) NewQx(stmt string, args ...interface{}) *QxConjunction {
	return NewQxConjunction(stmt, args...)
}

// Qx returns a new QxConjunction with the provided statement and arguments.
// Qx 返回一个新的QxConjunction，使用提供的语句和参数。
func (common *ColumnOperationClass) Qx(stmt string, args ...interface{}) *QxConjunction {
	return NewQxConjunction(stmt, args...)
}

// CreateSelect creates a new SelectStatement with the provided statement and arguments.
// CreateSelect 根据提供的语句和参数创建一个新的SelectStatement。
func (common *ColumnOperationClass) CreateSelect(stmt string, args ...interface{}) *SelectStatement {
	return NewSelectStatement(stmt, args...)
}

// NewSx creates a new SelectStatement with the provided statement and arguments.
// NewSx 根据提供的语句和参数创建一个新的SelectStatement。
func (common *ColumnOperationClass) NewSx(stmt string, args ...interface{}) *SelectStatement {
	return NewSelectStatement(stmt, args...)
}

// Sx returns a new SelectStatement with the provided statement and arguments.
// Sx 返回一个新的SelectStatement，使用提供的语句和参数。
func (common *ColumnOperationClass) Sx(stmt string, args ...interface{}) *SelectStatement {
	return NewSelectStatement(stmt, args...)
}

// NewColumnValueMap creates a new ColumnValueMap using the NewKw function.
// NewColumnValueMap 使用NewKw函数创建一个新的ColumnValueMap。
func (common *ColumnOperationClass) NewColumnValueMap() ColumnValueMap {
	return NewKw()
}

// NewKw creates a new ColumnValueMap using the Kw function.
// NewKw 使用Kw函数创建一个新的ColumnValueMap。
func (common *ColumnOperationClass) NewKw() ColumnValueMap {
	return NewKw()
}

// CreateColumnValueMap creates a ColumnValueMap using the provided column name and value.
// CreateColumnValueMap 根据提供的列名和值创建一个ColumnValueMap。
func (common *ColumnOperationClass) CreateColumnValueMap(columnName string, value interface{}) ColumnValueMap {
	return Kw(columnName, value)
}

// Kw creates a ColumnValueMap using the provided column name and value.
// Kw 根据提供的列名和值创建一个ColumnValueMap。
func (common *ColumnOperationClass) Kw(columnName string, value interface{}) ColumnValueMap {
	return Kw(columnName, value)
}

// Where applies the provided QxConjunctions to the given gorm.DB statement.
// Where 将提供的QxConjunction应用到给定的gorm.DB语句。
func (common *ColumnOperationClass) Where(db *gorm.DB, qxs ...*QxConjunction) *gorm.DB {
	stmt := db
	for _, qx := range qxs {
		stmt = stmt.Where(qx.Qs(), qx.args...)
	}
	return stmt
}

// OrderByColumns applies the provided OrderByBottle objects to the given gorm.DB statement.
// OrderByColumns 将提供的OrderByBottle对象应用到给定的gorm.DB语句。
func (common *ColumnOperationClass) OrderByColumns(db *gorm.DB, obs ...OrderByBottle) *gorm.DB {
	stmt := db
	for _, ob := range obs {
		stmt = stmt.Order(ob.Ox())
	}
	return stmt
}

// UpdateColumns updates the columns of the given gorm.DB statement with the provided ColumnValueMaps.
// UpdateColumns 使用提供的ColumnValueMap更新给定的gorm.DB语句的列。
func (common *ColumnOperationClass) UpdateColumns(db *gorm.DB, kws ...ColumnValueMap) *gorm.DB {
	mp := map[string]interface{}{}
	for _, kw := range kws {
		for k, v := range kw.AsMap() {
			mp[k] = v
		}
	}
	return db.UpdateColumns(mp)
}

// CombineColumnNames combines the names of the provided ColumnNameInterfaces into a single string.
// CombineColumnNames 将提供的ColumnNameInterface的名称组合成一个字符串。
func (common *ColumnOperationClass) CombineColumnNames(a ...utils.ColumnNameInterface) string {
	var names = make([]string, 0, len(a))
	for _, x := range a {
		names = append(names, x.Name())
	}
	return strings.Join(names, ", ")
}

// MergeNames combines the names of the provided ColumnNameInterfaces into a single string.
// MergeNames 将提供的ColumnNameInterface的名称组合成一个字符串。
func (common *ColumnOperationClass) MergeNames(a ...utils.ColumnNameInterface) string {
	return common.CombineColumnNames(a...)
}

func (common *ColumnOperationClass) CombineNamesSlices(a ...[]string) string {
	var names []string
	for _, elems := range a {
		names = append(names, elems...)
	}
	return strings.Join(names, ", ")
}

func (common *ColumnOperationClass) MergeSlices(a ...[]string) string {
	return common.CombineNamesSlices(a...)
}

// CombineStatements combines the provided SQL statements into a single string.
// CombineStatements 将提供的SQL语句组合成一个字符串。
func (common *ColumnOperationClass) CombineStatements(a ...string) string {
	return strings.Join(a, ", ")
}

// MergeStmts combines the provided SQL statements into a single string.
// MergeStmts 将提供的SQL语句组合成一个字符串。
func (common *ColumnOperationClass) MergeStmts(a ...string) string {
	return strings.Join(a, ", ")
}

// CountStmt returns a SQL statement that counts the records, applying the alias.
// CountStmt 返回一个计算记录数量的SQL语句，并应用别名。
func (common *ColumnOperationClass) CountStmt(alias string) string {
	return utils.ApplyAliasToColumn("COUNT(*)", alias)
}

// CountCaseWhenStmt returns a SQL statement that counts the records with a CASE WHEN condition, applying the alias.
// CountCaseWhenStmt 返回一个计算记录数量的SQL语句，带有CASE WHEN条件，并应用别名。
func (common *ColumnOperationClass) CountCaseWhenStmt(condition string, alias string) string {
	return utils.ApplyAliasToColumn("COUNT(CASE WHEN ("+condition+") THEN 1 END)", alias)
}

// CountCaseWhenQxSx returns a SelectStatement with a COUNT CASE WHEN condition applied, using the provided QxConjunction and alias.
// CountCaseWhenQxSx 返回一个带有COUNT CASE WHEN条件的SelectStatement，使用提供的QxConjunction和别名。
func (common *ColumnOperationClass) CountCaseWhenQxSx(qx *QxConjunction, alias string) *SelectStatement {
	return NewSelectStatement(
		utils.ApplyAliasToColumn("COUNT(CASE WHEN ("+qx.Qs()+") THEN 1 END)", alias),
		qx.Args()...,
	)
}

// CombineSelectStatements combines multiple SelectStatements into a single SelectStatement.
// CombineSelectStatements 将多个SelectStatement组合成一个SelectStatement。
func (common *ColumnOperationClass) CombineSelectStatements(cs ...SelectStatement) *SelectStatement {
	var qsVs []string
	var args []any
	for _, c := range cs {
		qsVs = append(qsVs, c.Qs())
		args = append(args, c.Args()...)
	}
	var stmt = strings.Join(qsVs, ", ")
	return NewSelectStatement(stmt, args...)
}

// CombineSxs combines multiple SelectStatements into a single SelectStatement.
// CombineSxs 将多个SelectStatement组合成一个SelectStatement。
func (common *ColumnOperationClass) CombineSxs(cs ...SelectStatement) *SelectStatement {
	return common.CombineSelectStatements(cs...)
}

// Select applies the provided SelectStatements to the given gorm.DB statement.
// Select 将提供的SelectStatement应用到给定的gorm.DB语句。
func (common *ColumnOperationClass) Select(db *gorm.DB, qxs ...*SelectStatement) *gorm.DB {
	stmt := db
	for _, qx := range qxs {
		stmt = stmt.Select(qx.Qs(), qx.args...)
	}
	return stmt
}
