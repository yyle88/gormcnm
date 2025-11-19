// Package gormcnm provides statement and arguments tuple handling for GORM query operations
// Auto manages SQL statements with parameter binding and type conversion
// Supports driver.Valuer interface implementation for custom types in GORM WHERE clauses
//
// gormcnm 提供语句和参数元组处理，用于 GORM 查询操作
// 自动管理 SQL 语句及参数绑定和类型转换
// 支持 driver.Valuer 接口实现，用于 GORM WHERE 子句中的自定义类型
package gormcnm

import (
	"database/sql/driver"

	"github.com/pkg/errors"
	"github.com/yyle88/must"
)

// 当你在调用时报这个错时，说明你 where 条件的第一个参数不是字符串类型，而是直接使用的该项目中自定义的类型，而 gorm 不是能自动识别它们，因此我主动增加 panic 以提醒您出现错误
// 因为 gorm 中的 db.Where 的定义是这样的
// func (db *DB) Where(query interface{}, args ...interface{}) (tx *DB)
// 在这个项目里，你需要传的是，查询字符串 stmt 和其参数列表 args... 而不是其它类型
// 因此你需要传的是 db.Where(qx.Qx1())  db.Where(qx.Qx2())  db.Where(qx.Qx3()) 而不是 db.Where(qx)
// 这块非常抱歉，但目前尚无特别好的解决方案，因为 Where 需要的是变长的参数列表，而 golang 不支持变长个数的返回值，返回个数组是没用的
var valueIsNotCallable = errors.New("column.value() function is not callable")

// statementArgumentsTuple represents a tuple of SQL statement and its arguments
// Core building block for GORM query construction with proper argument handling
// Addresses Go language limitation where variadic return values are not supported
//
// 核心设计说明：
// - GORM 的 Where 和 Select 方法需要: func (db *DB) Where(query interface{}, args ...interface{})
// - 当遇到 AND/OR 时需要合并语句和参数列表
// - 由于 Go 不支持变长返回值，只能设计 Qx1(), Qx2(), Qx3()等方法
// - 包含 panic 机制防止直接传递给 GORM 导致的错误使用
//
// statementArgumentsTuple 表示 SQL 语句及其参数的元组
// GORM 查询构建的核心构建块，提供适当的参数处理
// 解决 Go 语言不支持变长返回值的限制
type statementArgumentsTuple struct {
	stmt string        // SQL statement string // SQL 语句字符串
	args []interface{} // Prepared statement arguments // 预处理语句参数
}

// newStatementArgumentsTuple creates a new statementArgumentsTuple instance with the provided statement and arguments.
// newStatementArgumentsTuple 使用提供的语句和参数创建一个新的 statementArgumentsTuple 实例。
func newStatementArgumentsTuple(stmt string, args []interface{}) *statementArgumentsTuple {
	return &statementArgumentsTuple{
		stmt: stmt,
		args: args,
	}
}

// safeCombineArguments merges the arguments from the current instance with the provided instances and returns the combined arguments list.
// safeCombineArguments 将当前实例的参数与提供的实例的参数合并，并返回合并后的参数列表。
func (qx *statementArgumentsTuple) safeCombineArguments(cs []*statementArgumentsTuple) []interface{} {
	var args []interface{}
	args = append(args, qx.args...)
	for _, c := range cs {
		args = append(args, c.args...)
	}
	return args
}

// Value panics to prevent direct use of this structure in GORM, as it is not callable in this context.
// Value 会触发 panic，以防止在 GORM 中直接使用此结构，因为在此上下文中无法调用该函数。
func (qx *statementArgumentsTuple) Value() (driver.Value, error) {
	panic(valueIsNotCallable) //当报这个错时，需要修改调用侧代码，请看这个错误码的注释
}

// Qs return the statement string of the current statementArgumentsTuple instance.
// Qs 返回当前 statementArgumentsTuple 实例的语句字符串。
func (qx *statementArgumentsTuple) Qs() string {
	return qx.stmt
}

// Args return the arguments list of the current statementArgumentsTuple instance.
// Args 返回当前 statementArgumentsTuple 实例的参数列表。
func (qx *statementArgumentsTuple) Args() []interface{} {
	return qx.args
}

// Qx0 returns the statement string when there are no arguments in the statementArgumentsTuple instance.
// Qx0 在 statementArgumentsTuple 实例没有参数时返回语句字符串。
func (qx *statementArgumentsTuple) Qx0() string {
	must.Len(qx.args, 0)
	return qx.Qs()
}

// Qx1 returns the statement string and the first argument in the statementArgumentsTuple instance.
// Qx1 返回 statementArgumentsTuple 实例的语句字符串和第一个参数。
func (qx *statementArgumentsTuple) Qx1() (string, interface{}) {
	must.Len(qx.args, 1)
	return qx.Qs(), qx.args[0]
}

// Qx2 returns the statement string and the first two arguments in the statementArgumentsTuple instance.
// Qx2 返回 statementArgumentsTuple 实例的语句字符串和前两个参数。
func (qx *statementArgumentsTuple) Qx2() (string, interface{}, interface{}) {
	must.Len(qx.args, 2)
	return qx.Qs(), qx.args[0], qx.args[1]
}

// Qx3 这块暂时没有什么好的方案，我只能这样罗列下来，很期望将来能够解决这个问题
func (qx *statementArgumentsTuple) Qx3() (string, interface{}, interface{}, interface{}) {
	must.Len(qx.args, 3)
	return qx.Qs(), qx.args[0], qx.args[1], qx.args[2]
}

// Qx4 这块暂时没有什么好的方案，我只能这样罗列下来，很期望将来能够解决这个问题
func (qx *statementArgumentsTuple) Qx4() (string, interface{}, interface{}, interface{}, interface{}) {
	must.Len(qx.args, 4)
	return qx.Qs(), qx.args[0], qx.args[1], qx.args[2], qx.args[3]
}

func (qx *statementArgumentsTuple) Qx5() (string, interface{}, interface{}, interface{}, interface{}, interface{}) {
	must.Len(qx.args, 5)
	return qx.Qs(), qx.args[0], qx.args[1], qx.args[2], qx.args[3], qx.args[4]
}

func (qx *statementArgumentsTuple) Qx6() (string, interface{}, interface{}, interface{}, interface{}, interface{}, interface{}) {
	must.Len(qx.args, 6)
	return qx.Qs(), qx.args[0], qx.args[1], qx.args[2], qx.args[3], qx.args[4], qx.args[5]
}

func (qx *statementArgumentsTuple) Qx7() (string, interface{}, interface{}, interface{}, interface{}, interface{}, interface{}, interface{}) {
	must.Len(qx.args, 7)
	return qx.Qs(), qx.args[0], qx.args[1], qx.args[2], qx.args[3], qx.args[4], qx.args[5], qx.args[6]
}

func (qx *statementArgumentsTuple) Qx8() (string, interface{}, interface{}, interface{}, interface{}, interface{}, interface{}, interface{}, interface{}) {
	must.Len(qx.args, 8)
	return qx.Qs(), qx.args[0], qx.args[1], qx.args[2], qx.args[3], qx.args[4], qx.args[5], qx.args[6], qx.args[7]
}

func (qx *statementArgumentsTuple) Qx9() (string, interface{}, interface{}, interface{}, interface{}, interface{}, interface{}, interface{}, interface{}, interface{}) {
	must.Len(qx.args, 9)
	return qx.Qs(), qx.args[0], qx.args[1], qx.args[2], qx.args[3], qx.args[4], qx.args[5], qx.args[6], qx.args[7], qx.args[8]
}

func (qx *statementArgumentsTuple) Qx10() (string, interface{}, interface{}, interface{}, interface{}, interface{}, interface{}, interface{}, interface{}, interface{}, interface{}) {
	must.Len(qx.args, 10)
	return qx.Qs(), qx.args[0], qx.args[1], qx.args[2], qx.args[3], qx.args[4], qx.args[5], qx.args[6], qx.args[7], qx.args[8], qx.args[9]
}

func (qx *statementArgumentsTuple) Qx11() (string, interface{}, interface{}, interface{}, interface{}, interface{}, interface{}, interface{}, interface{}, interface{}, interface{}, interface{}) {
	must.Len(qx.args, 11)
	return qx.Qs(), qx.args[0], qx.args[1], qx.args[2], qx.args[3], qx.args[4], qx.args[5], qx.args[6], qx.args[7], qx.args[8], qx.args[9], qx.args[10]
}

func (qx *statementArgumentsTuple) Qx12() (string, interface{}, interface{}, interface{}, interface{}, interface{}, interface{}, interface{}, interface{}, interface{}, interface{}, interface{}, interface{}) {
	must.Len(qx.args, 12)
	return qx.Qs(), qx.args[0], qx.args[1], qx.args[2], qx.args[3], qx.args[4], qx.args[5], qx.args[6], qx.args[7], qx.args[8], qx.args[9], qx.args[10], qx.args[11]
}
