package gormcnm

import (
	"database/sql/driver"

	"github.com/pkg/errors"
	"github.com/yyle88/must"
)

// 当你在调用时报这个错时，说明你 where 条件的第一个参数不是字符串类型，而是直接使用的该项目中自定义的类型，而gorm不是能自动识别它们，因此我主动增加 panic 以提醒您出现错误
// 因为 gorm 中的 db.Where 的定义是这样的
// func (db *DB) Where(query interface{}, args ...interface{}) (tx *DB)
// 在这个项目里，你需要传的是，查询字符串 stmt 和其参数列表 args... 而不是其它类型
// 因此你需要传的是 db.Where(qx.Qx1())  db.Where(qx.Qx2())  db.Where(qx.Qx3()) 而不是 db.Where(qx)
// 这块非常抱歉，但目前尚无特别好的解决方案，因为 Where 需要的是变长的参数列表，而golang不支持变长个数的返回值，返回个数组是没用的
var valueIsNotCallable = errors.New("column.value() function is not callable")

// 就是语句 stmt 和参数列表 args 的 tuple 元组
// 因为 db.Where 恰好需要
// func (db *DB) Where(query interface{}, args ...interface{}) (tx *DB)
// 当遇到 AND 或 OR 的时候，就需要按照 AND 或 OR 的规则，合并语句和合并参数列表
// 而且 db.Select 也需要
// func (db *DB) Select(query interface{}, args ...interface{}) (tx *DB)
// 当有多个列要被选中且有过滤条件时，就要把选中的语句用逗号分隔合并，再合并参数列表
// 因此这个类就是个底层的类
// 由于 db.Where 和 db.Select 的第二个参数都是变长的参数
// 但这个类返回的 args 是个数组，很明显的，需要也返回变长参数列表，但是目前golang不支持返回变长结果
// 因此不得不作出这样的设计:
// func (qx *statementArgumentsTuple) Qx1() (string, interface{})
// func (qx *statementArgumentsTuple) Qx2() (string, interface{}, interface{})
// func (qx *statementArgumentsTuple) Qx3() (string, interface{}, interface{}, interface{})
// 你需要根据参数的数量调用 db.Where(qx.Qx2()) 这样才行
// 这也是无奈之举啊
// 当然同时为了防止你在编码时忘记，我还贴心（窝心）的增加了个 panic 的逻辑，让你在忘记是能在运行时快速定位问题
type statementArgumentsTuple struct {
	stmt string
	args []interface{}
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
