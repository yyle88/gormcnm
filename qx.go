package gormcnm

// QxType 主要是传给 db.Where 的 AND OR NOT 这些关系查询的语句和参数
// 这是 db.Where 函数的定义代码
// func (db *DB) Where(query interface{}, args ...interface{}) (tx *DB)
// 需要的就是语句和参数
// 而恰好像 Eq 函数就是 返回的 string(s) + "=?", x 字符串和参数
// 但假如需要 a.Eq("xyz") 【而且】 b.Eq("uvw") 的时候
// 就需要使用这个工具类，把语句使用 (---) AND (---) 相连，把参数列表使用新数组 append 相连成为新的参数列表
// 因此这个类也是有较高使用频率的
type QxType struct {
	*stmtArgsTuple
}

func NewQx(stmt string, args ...interface{}) *QxType {
	return &QxType{
		stmtArgsTuple: newStmtArgsTuple(stmt, args),
	}
}

func Qx(stmt string, args ...interface{}) *QxType {
	return NewQx(stmt, args...)
}

func (qx *QxType) AND(cs ...*QxType) *QxType {
	var qss []QsConjunction
	var qas []*stmtArgsTuple
	for _, c := range cs {
		qss = append(qss, QsConjunction(c.stmt))
		qas = append(qas, c.stmtArgsTuple)
	}

	return &QxType{
		stmtArgsTuple: newStmtArgsTuple(string(QsConjunction(qx.stmt).AND(qss...)), qx.safeMergeArgs(qas)),
	}
}

func (qx *QxType) OR(cs ...*QxType) *QxType {
	var qss []QsConjunction
	var qas []*stmtArgsTuple
	for _, c := range cs {
		qss = append(qss, QsConjunction(c.stmt))
		qas = append(qas, c.stmtArgsTuple)
	}
	return &QxType{
		stmtArgsTuple: newStmtArgsTuple(string(QsConjunction(qx.stmt).OR(qss...)), qx.safeMergeArgs(qas)),
	}
}

func (qx *QxType) NOT() *QxType {
	return &QxType{
		stmtArgsTuple: newStmtArgsTuple(string(QsConjunction(qx.stmt).NOT()), qx.args),
	}
}

func (qx *QxType) AND1(stmt string, args ...interface{}) *QxType {
	return qx.AND(NewQx(stmt, args...))
}

func (qx *QxType) OR1(stmt string, args ...interface{}) *QxType {
	return qx.OR(NewQx(stmt, args...))
}
