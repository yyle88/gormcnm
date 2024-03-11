package gormcnm

import (
	"database/sql/driver"

	"github.com/yyle88/gormcnm/utilsgormcnm"
)

type QxType struct {
	qc   QsCondition
	args []interface{}
}

func NewQx(qs string, args ...interface{}) *QxType {
	return &QxType{
		qc:   QsCondition(qs),
		args: args,
	}
}

func Qx(qs string, args ...interface{}) *QxType {
	return NewQx(qs, args...)
}

func (qx *QxType) AND(cs ...*QxType) *QxType {
	var qss []QsCondition
	for _, c := range cs {
		qss = append(qss, c.qc)
	}
	return &QxType{
		qc:   qx.qc.AND(qss...),
		args: qx.safeMergeArgs(cs),
	}
}

func (qx *QxType) OR(cs ...*QxType) *QxType {
	var qss []QsCondition
	for _, c := range cs {
		qss = append(qss, c.qc)
	}
	return &QxType{
		qc:   qx.qc.OR(qss...),
		args: qx.safeMergeArgs(cs),
	}
}

func (qx *QxType) safeMergeArgs(cs []*QxType) []interface{} {
	var args []interface{}
	args = append(args, qx.args...)
	for _, c := range cs {
		args = append(args, c.args...)
	}
	return args
}

func (qx *QxType) NOT() *QxType {
	return &QxType{
		qc:   qx.qc.NOT(),
		args: qx.args,
	}
}

func (qx *QxType) AND1(qs string, args ...interface{}) *QxType {
	return qx.AND(NewQx(qs, args...))
}

func (qx *QxType) OR1(qs string, args ...interface{}) *QxType {
	return qx.OR(NewQx(qs, args...))
}

func (qx *QxType) Value() (driver.Value, error) {
	return qx.qc.Value() //这里还是抛异常的
}

func (qx *QxType) Qs() string {
	return string(qx.qc)
}

func (qx *QxType) Args() []interface{} {
	return qx.args
}

// Qx0 这块暂时没有什么好的方案，我只能这样罗列下来，很期望将来能够解决这个问题
func (qx *QxType) Qx0() string {
	utilsgormcnm.AssertEquals(len(qx.args), 0)
	return qx.Qs()
}

// Qx1 这块暂时没有什么好的方案，我只能这样罗列下来，很期望将来能够解决这个问题
func (qx *QxType) Qx1() (string, interface{}) {
	utilsgormcnm.AssertEquals(len(qx.args), 1)
	return qx.Qs(), qx.args[0]
}

// Qx2 这块暂时没有什么好的方案，我只能这样罗列下来，很期望将来能够解决这个问题
func (qx *QxType) Qx2() (string, interface{}, interface{}) {
	utilsgormcnm.AssertEquals(len(qx.args), 2)
	return qx.Qs(), qx.args[0], qx.args[1]
}

func (qx *QxType) Qx3() (string, interface{}, interface{}, interface{}) {
	utilsgormcnm.AssertEquals(len(qx.args), 3)
	return qx.Qs(), qx.args[0], qx.args[1], qx.args[2]
}

func (qx *QxType) Qx4() (string, interface{}, interface{}, interface{}, interface{}) {
	utilsgormcnm.AssertEquals(len(qx.args), 4)
	return qx.Qs(), qx.args[0], qx.args[1], qx.args[2], qx.args[3]
}

func (qx *QxType) Qx5() (string, interface{}, interface{}, interface{}, interface{}, interface{}) {
	utilsgormcnm.AssertEquals(len(qx.args), 5)
	return qx.Qs(), qx.args[0], qx.args[1], qx.args[2], qx.args[3], qx.args[4]
}

func (qx *QxType) Qx6() (string, interface{}, interface{}, interface{}, interface{}, interface{}, interface{}) {
	utilsgormcnm.AssertEquals(len(qx.args), 6)
	return qx.Qs(), qx.args[0], qx.args[1], qx.args[2], qx.args[3], qx.args[4], qx.args[5]
}

func (qx *QxType) Qx7() (string, interface{}, interface{}, interface{}, interface{}, interface{}, interface{}, interface{}) {
	utilsgormcnm.AssertEquals(len(qx.args), 7)
	return qx.Qs(), qx.args[0], qx.args[1], qx.args[2], qx.args[3], qx.args[4], qx.args[5], qx.args[6]
}

func (qx *QxType) Qx8() (string, interface{}, interface{}, interface{}, interface{}, interface{}, interface{}, interface{}, interface{}) {
	utilsgormcnm.AssertEquals(len(qx.args), 8)
	return qx.Qs(), qx.args[0], qx.args[1], qx.args[2], qx.args[3], qx.args[4], qx.args[5], qx.args[6], qx.args[7]
}

func (qx *QxType) Qx9() (string, interface{}, interface{}, interface{}, interface{}, interface{}, interface{}, interface{}, interface{}, interface{}) {
	utilsgormcnm.AssertEquals(len(qx.args), 9)
	return qx.Qs(), qx.args[0], qx.args[1], qx.args[2], qx.args[3], qx.args[4], qx.args[5], qx.args[6], qx.args[7], qx.args[8]
}

func (qx *QxType) Qx10() (string, interface{}, interface{}, interface{}, interface{}, interface{}, interface{}, interface{}, interface{}, interface{}, interface{}) {
	utilsgormcnm.AssertEquals(len(qx.args), 10)
	return qx.Qs(), qx.args[0], qx.args[1], qx.args[2], qx.args[3], qx.args[4], qx.args[5], qx.args[6], qx.args[7], qx.args[8], qx.args[9]
}

func (qx *QxType) Qx11() (string, interface{}, interface{}, interface{}, interface{}, interface{}, interface{}, interface{}, interface{}, interface{}, interface{}, interface{}) {
	utilsgormcnm.AssertEquals(len(qx.args), 11)
	return qx.Qs(), qx.args[0], qx.args[1], qx.args[2], qx.args[3], qx.args[4], qx.args[5], qx.args[6], qx.args[7], qx.args[8], qx.args[9], qx.args[10]
}

func (qx *QxType) Qx12() (string, interface{}, interface{}, interface{}, interface{}, interface{}, interface{}, interface{}, interface{}, interface{}, interface{}, interface{}, interface{}) {
	utilsgormcnm.AssertEquals(len(qx.args), 12)
	return qx.Qs(), qx.args[0], qx.args[1], qx.args[2], qx.args[3], qx.args[4], qx.args[5], qx.args[6], qx.args[7], qx.args[8], qx.args[9], qx.args[10], qx.args[11]
}
