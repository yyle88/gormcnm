package gormcnmstub

import (
	"github.com/yyle88/gormcnm"
	"github.com/yyle88/gormcnm/internal/utils"
	"gorm.io/gorm"
)

func OK() bool {
	return stub.OK()
}
func CreateCondition(stmt string, args ...interface{}) *gormcnm.QxConjunction {
	return stub.CreateCondition(stmt, args...)
}
func NewQx(stmt string, args ...interface{}) *gormcnm.QxConjunction {
	return stub.NewQx(stmt, args...)
}
func Qx(stmt string, args ...interface{}) *gormcnm.QxConjunction {
	return stub.Qx(stmt, args...)
}
func CreateSelect(stmt string, args ...interface{}) *gormcnm.SelectStatement {
	return stub.CreateSelect(stmt, args...)
}
func NewSx(stmt string, args ...interface{}) *gormcnm.SelectStatement {
	return stub.NewSx(stmt, args...)
}
func Sx(stmt string, args ...interface{}) *gormcnm.SelectStatement {
	return stub.Sx(stmt, args...)
}
func NewColumnValueMap() gormcnm.ColumnValueMap {
	return stub.NewColumnValueMap()
}
func NewKw() gormcnm.ColumnValueMap {
	return stub.NewKw()
}
func CreateColumnValueMap(columnName string, value interface{}) gormcnm.ColumnValueMap {
	return stub.CreateColumnValueMap(columnName, value)
}
func Kw(columnName string, value interface{}) gormcnm.ColumnValueMap {
	return stub.Kw(columnName, value)
}
func Where(db *gorm.DB, qxs ...*gormcnm.QxConjunction) *gorm.DB {
	return stub.Where(db, qxs...)
}
func OrderByColumns(db *gorm.DB, obs ...gormcnm.OrderByBottle) *gorm.DB {
	return stub.OrderByColumns(db, obs...)
}
func UpdateColumns(db *gorm.DB, kws ...gormcnm.ColumnValueMap) *gorm.DB {
	return stub.UpdateColumns(db, kws...)
}
func CombineColumnNames(a ...utils.ColumnNameInterface) string {
	return stub.CombineColumnNames(a...)
}
func MergeNames(a ...utils.ColumnNameInterface) string {
	return stub.MergeNames(a...)
}
func CombineNamesSlices(a ...[]string) string {
	return stub.CombineNamesSlices(a...)
}
func MergeSlices(a ...[]string) string {
	return stub.MergeSlices(a...)
}
func CombineStatements(a ...string) string {
	return stub.CombineStatements(a...)
}
func MergeStmts(a ...string) string {
	return stub.MergeStmts(a...)
}
func CountStmt(alias string) string {
	return stub.CountStmt(alias)
}
func CountCaseWhenStmt(condition string, alias string) string {
	return stub.CountCaseWhenStmt(condition, alias)
}
func CountCaseWhenQxSx(qx *gormcnm.QxConjunction, alias string) *gormcnm.SelectStatement {
	return stub.CountCaseWhenQxSx(qx, alias)
}
func CombineSelectStatements(cs ...gormcnm.SelectStatement) *gormcnm.SelectStatement {
	return stub.CombineSelectStatements(cs...)
}
func CombineSxs(cs ...gormcnm.SelectStatement) *gormcnm.SelectStatement {
	return stub.CombineSxs(cs...)
}
func Select(db *gorm.DB, qxs ...*gormcnm.SelectStatement) *gorm.DB {
	return stub.Select(db, qxs...)
}
func LEFTJOIN(tableName string) *gormcnm.TableJoin {
	return stub.LEFTJOIN(tableName)
}
func RIGHTJOIN(tableName string) *gormcnm.TableJoin {
	return stub.RIGHTJOIN(tableName)
}
func INNERJOIN(tableName string) *gormcnm.TableJoin {
	return stub.INNERJOIN(tableName)
}
func CROSSJOIN(tableName string) *gormcnm.TableJoin {
	return stub.CROSSJOIN(tableName)
}
