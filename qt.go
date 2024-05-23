package gormcnm

import (
	"database/sql/driver"

	"github.com/yyle88/gormcnm/internal/utils"
)

type queryArgsTuple struct {
	qc   QsCondition
	args []interface{}
}

func (qx *queryArgsTuple) safeMergeArgs(cs []*queryArgsTuple) []interface{} {
	var args []interface{}
	args = append(args, qx.args...)
	for _, c := range cs {
		args = append(args, c.args...)
	}
	return args
}

func (qx *queryArgsTuple) Value() (driver.Value, error) {
	return qx.qc.Value() //这里还是抛异常的
}

func (qx *queryArgsTuple) Qs() string {
	return string(qx.qc)
}

func (qx *queryArgsTuple) Args() []interface{} {
	return qx.args
}

// Qx0 这块暂时没有什么好的方案，我只能这样罗列下来，很期望将来能够解决这个问题
func (qx *queryArgsTuple) Qx0() string {
	utils.AssertEquals(len(qx.args), 0)
	return qx.Qs()
}

// Qx1 这块暂时没有什么好的方案，我只能这样罗列下来，很期望将来能够解决这个问题
func (qx *queryArgsTuple) Qx1() (string, interface{}) {
	utils.AssertEquals(len(qx.args), 1)
	return qx.Qs(), qx.args[0]
}

// Qx2 这块暂时没有什么好的方案，我只能这样罗列下来，很期望将来能够解决这个问题
func (qx *queryArgsTuple) Qx2() (string, interface{}, interface{}) {
	utils.AssertEquals(len(qx.args), 2)
	return qx.Qs(), qx.args[0], qx.args[1]
}

func (qx *queryArgsTuple) Qx3() (string, interface{}, interface{}, interface{}) {
	utils.AssertEquals(len(qx.args), 3)
	return qx.Qs(), qx.args[0], qx.args[1], qx.args[2]
}

func (qx *queryArgsTuple) Qx4() (string, interface{}, interface{}, interface{}, interface{}) {
	utils.AssertEquals(len(qx.args), 4)
	return qx.Qs(), qx.args[0], qx.args[1], qx.args[2], qx.args[3]
}

func (qx *queryArgsTuple) Qx5() (string, interface{}, interface{}, interface{}, interface{}, interface{}) {
	utils.AssertEquals(len(qx.args), 5)
	return qx.Qs(), qx.args[0], qx.args[1], qx.args[2], qx.args[3], qx.args[4]
}

func (qx *queryArgsTuple) Qx6() (string, interface{}, interface{}, interface{}, interface{}, interface{}, interface{}) {
	utils.AssertEquals(len(qx.args), 6)
	return qx.Qs(), qx.args[0], qx.args[1], qx.args[2], qx.args[3], qx.args[4], qx.args[5]
}

func (qx *queryArgsTuple) Qx7() (string, interface{}, interface{}, interface{}, interface{}, interface{}, interface{}, interface{}) {
	utils.AssertEquals(len(qx.args), 7)
	return qx.Qs(), qx.args[0], qx.args[1], qx.args[2], qx.args[3], qx.args[4], qx.args[5], qx.args[6]
}

func (qx *queryArgsTuple) Qx8() (string, interface{}, interface{}, interface{}, interface{}, interface{}, interface{}, interface{}, interface{}) {
	utils.AssertEquals(len(qx.args), 8)
	return qx.Qs(), qx.args[0], qx.args[1], qx.args[2], qx.args[3], qx.args[4], qx.args[5], qx.args[6], qx.args[7]
}

func (qx *queryArgsTuple) Qx9() (string, interface{}, interface{}, interface{}, interface{}, interface{}, interface{}, interface{}, interface{}, interface{}) {
	utils.AssertEquals(len(qx.args), 9)
	return qx.Qs(), qx.args[0], qx.args[1], qx.args[2], qx.args[3], qx.args[4], qx.args[5], qx.args[6], qx.args[7], qx.args[8]
}

func (qx *queryArgsTuple) Qx10() (string, interface{}, interface{}, interface{}, interface{}, interface{}, interface{}, interface{}, interface{}, interface{}, interface{}) {
	utils.AssertEquals(len(qx.args), 10)
	return qx.Qs(), qx.args[0], qx.args[1], qx.args[2], qx.args[3], qx.args[4], qx.args[5], qx.args[6], qx.args[7], qx.args[8], qx.args[9]
}

func (qx *queryArgsTuple) Qx11() (string, interface{}, interface{}, interface{}, interface{}, interface{}, interface{}, interface{}, interface{}, interface{}, interface{}, interface{}) {
	utils.AssertEquals(len(qx.args), 11)
	return qx.Qs(), qx.args[0], qx.args[1], qx.args[2], qx.args[3], qx.args[4], qx.args[5], qx.args[6], qx.args[7], qx.args[8], qx.args[9], qx.args[10]
}

func (qx *queryArgsTuple) Qx12() (string, interface{}, interface{}, interface{}, interface{}, interface{}, interface{}, interface{}, interface{}, interface{}, interface{}, interface{}, interface{}) {
	utils.AssertEquals(len(qx.args), 12)
	return qx.Qs(), qx.args[0], qx.args[1], qx.args[2], qx.args[3], qx.args[4], qx.args[5], qx.args[6], qx.args[7], qx.args[8], qx.args[9], qx.args[10], qx.args[11]
}
