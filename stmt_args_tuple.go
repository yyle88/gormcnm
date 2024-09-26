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
var erxFunctionIsNotExecutable = errors.New("column.value() function is not executable")

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
// func (qx *stmtArgsTuple) Qx1() (string, interface{})
// func (qx *stmtArgsTuple) Qx2() (string, interface{}, interface{})
// func (qx *stmtArgsTuple) Qx3() (string, interface{}, interface{}, interface{})
// 你需要根据参数的数量调用 db.Where(qx.Qx2()) 这样才行
// 这也是无奈之举啊
// 当然同时为了防止你在编码时忘记，我还贴心（窝心）的增加了个 panic 的逻辑，让你在忘记是能在运行时快速定位问题
type stmtArgsTuple struct {
	stmt string
	args []interface{}
}

func newStmtArgsTuple(stmt string, args []interface{}) *stmtArgsTuple {
	return &stmtArgsTuple{
		stmt: stmt,
		args: args,
	}
}

func (qx *stmtArgsTuple) safeMergeArgs(cs []*stmtArgsTuple) []interface{} {
	var args []interface{}
	args = append(args, qx.args...)
	for _, c := range cs {
		args = append(args, c.args...)
	}
	return args
}

// Value 这块非常重要，要避免gorm直接使用这个结构，因此要在这里panic
func (qx *stmtArgsTuple) Value() (driver.Value, error) {
	panic(erxFunctionIsNotExecutable) //当报这个错时，需要修改调用侧代码，请看这个错误码的注释
}

func (qx *stmtArgsTuple) Qs() string {
	return qx.stmt
}

func (qx *stmtArgsTuple) Args() []interface{} {
	return qx.args
}

// Qx0 这块暂时没有什么好的方案，我只能这样罗列下来，很期望将来能够解决这个问题
func (qx *stmtArgsTuple) Qx0() string {
	must.Len(qx.args, 0)
	return qx.Qs()
}

// Qx1 这块暂时没有什么好的方案，我只能这样罗列下来，很期望将来能够解决这个问题
func (qx *stmtArgsTuple) Qx1() (string, interface{}) {
	must.Len(qx.args, 1)
	return qx.Qs(), qx.args[0]
}

// Qx2 这块暂时没有什么好的方案，我只能这样罗列下来，很期望将来能够解决这个问题
func (qx *stmtArgsTuple) Qx2() (string, interface{}, interface{}) {
	must.Len(qx.args, 2)
	return qx.Qs(), qx.args[0], qx.args[1]
}

func (qx *stmtArgsTuple) Qx3() (string, interface{}, interface{}, interface{}) {
	must.Len(qx.args, 3)
	return qx.Qs(), qx.args[0], qx.args[1], qx.args[2]
}

func (qx *stmtArgsTuple) Qx4() (string, interface{}, interface{}, interface{}, interface{}) {
	must.Len(qx.args, 4)
	return qx.Qs(), qx.args[0], qx.args[1], qx.args[2], qx.args[3]
}

func (qx *stmtArgsTuple) Qx5() (string, interface{}, interface{}, interface{}, interface{}, interface{}) {
	must.Len(qx.args, 5)
	return qx.Qs(), qx.args[0], qx.args[1], qx.args[2], qx.args[3], qx.args[4]
}

func (qx *stmtArgsTuple) Qx6() (string, interface{}, interface{}, interface{}, interface{}, interface{}, interface{}) {
	must.Len(qx.args, 6)
	return qx.Qs(), qx.args[0], qx.args[1], qx.args[2], qx.args[3], qx.args[4], qx.args[5]
}

func (qx *stmtArgsTuple) Qx7() (string, interface{}, interface{}, interface{}, interface{}, interface{}, interface{}, interface{}) {
	must.Len(qx.args, 7)
	return qx.Qs(), qx.args[0], qx.args[1], qx.args[2], qx.args[3], qx.args[4], qx.args[5], qx.args[6]
}

func (qx *stmtArgsTuple) Qx8() (string, interface{}, interface{}, interface{}, interface{}, interface{}, interface{}, interface{}, interface{}) {
	must.Len(qx.args, 8)
	return qx.Qs(), qx.args[0], qx.args[1], qx.args[2], qx.args[3], qx.args[4], qx.args[5], qx.args[6], qx.args[7]
}

func (qx *stmtArgsTuple) Qx9() (string, interface{}, interface{}, interface{}, interface{}, interface{}, interface{}, interface{}, interface{}, interface{}) {
	must.Len(qx.args, 9)
	return qx.Qs(), qx.args[0], qx.args[1], qx.args[2], qx.args[3], qx.args[4], qx.args[5], qx.args[6], qx.args[7], qx.args[8]
}

func (qx *stmtArgsTuple) Qx10() (string, interface{}, interface{}, interface{}, interface{}, interface{}, interface{}, interface{}, interface{}, interface{}, interface{}) {
	must.Len(qx.args, 10)
	return qx.Qs(), qx.args[0], qx.args[1], qx.args[2], qx.args[3], qx.args[4], qx.args[5], qx.args[6], qx.args[7], qx.args[8], qx.args[9]
}

func (qx *stmtArgsTuple) Qx11() (string, interface{}, interface{}, interface{}, interface{}, interface{}, interface{}, interface{}, interface{}, interface{}, interface{}, interface{}) {
	must.Len(qx.args, 11)
	return qx.Qs(), qx.args[0], qx.args[1], qx.args[2], qx.args[3], qx.args[4], qx.args[5], qx.args[6], qx.args[7], qx.args[8], qx.args[9], qx.args[10]
}

func (qx *stmtArgsTuple) Qx12() (string, interface{}, interface{}, interface{}, interface{}, interface{}, interface{}, interface{}, interface{}, interface{}, interface{}, interface{}, interface{}) {
	must.Len(qx.args, 12)
	return qx.Qs(), qx.args[0], qx.args[1], qx.args[2], qx.args[3], qx.args[4], qx.args[5], qx.args[6], qx.args[7], qx.args[8], qx.args[9], qx.args[10], qx.args[11]
}
