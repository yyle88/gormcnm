package gormcnm

type QxType struct {
	*queryArgsTuple
}

func NewQx(qs string, args ...interface{}) *QxType {
	return &QxType{
		&queryArgsTuple{
			qc:   QsCondition(qs),
			args: args,
		},
	}
}

func Qx(qs string, args ...interface{}) *QxType {
	return NewQx(qs, args...)
}

func (qx *QxType) AND(cs ...*QxType) *QxType {
	var qss []QsCondition
	var qas []*queryArgsTuple
	for _, c := range cs {
		qss = append(qss, c.qc)
		qas = append(qas, c.queryArgsTuple)
	}

	return &QxType{
		&queryArgsTuple{
			qc:   qx.qc.AND(qss...),
			args: qx.safeMergeArgs(qas),
		},
	}
}

func (qx *QxType) OR(cs ...*QxType) *QxType {
	var qss []QsCondition
	var qas []*queryArgsTuple
	for _, c := range cs {
		qss = append(qss, c.qc)
		qas = append(qas, c.queryArgsTuple)
	}
	return &QxType{
		&queryArgsTuple{
			qc:   qx.qc.OR(qss...),
			args: qx.safeMergeArgs(qas),
		},
	}
}

func (qx *QxType) NOT() *QxType {
	return &QxType{
		&queryArgsTuple{
			qc:   qx.qc.NOT(),
			args: qx.args,
		},
	}
}

func (qx *QxType) AND1(qs string, args ...interface{}) *QxType {
	return qx.AND(NewQx(qs, args...))
}

func (qx *QxType) OR1(qs string, args ...interface{}) *QxType {
	return qx.OR(NewQx(qs, args...))
}
