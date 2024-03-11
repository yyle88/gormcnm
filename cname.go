package gormcnm

type ColumnName[TYPE any] string

func (s ColumnName[TYPE]) Qs(op string) string {
	return string(s) + " " + op
}

func (s ColumnName[TYPE]) Op(op string, x TYPE) (string, TYPE) {
	return string(s) + " " + op, x
}

func (s ColumnName[TYPE]) Eq(x TYPE) (string, TYPE) {
	return string(s) + "=?", x
}

func (s ColumnName[TYPE]) Gt(x TYPE) (string, TYPE) {
	return string(s) + ">?", x
}

func (s ColumnName[TYPE]) Lt(x TYPE) (string, TYPE) {
	return string(s) + "<?", x
}

func (s ColumnName[TYPE]) Gte(x TYPE) (string, TYPE) {
	return string(s) + ">=?", x
}

func (s ColumnName[TYPE]) Lte(x TYPE) (string, TYPE) {
	return string(s) + "<=?", x
}

func (s ColumnName[TYPE]) Ne(x TYPE) (string, TYPE) {
	return string(s) + "!=?", x
}

func (s ColumnName[TYPE]) In(x []TYPE) (string, []TYPE) {
	return string(s) + " IN(?)", x
}

func (s ColumnName[TYPE]) NotIn(x []TYPE) (string, []TYPE) {
	return string(s) + " NOT IN(?)", x
}

func (s ColumnName[TYPE]) Like(x TYPE) (string, TYPE) {
	return string(s) + " LIKE ?", x
}

func (s ColumnName[TYPE]) NotLike(x TYPE) (string, TYPE) {
	return string(s) + " NOT LIKE ?", x
}

func (s ColumnName[TYPE]) IsNULL() string {
	return string(s) + " IS NULL"
}

func (s ColumnName[TYPE]) IsNotNULL() string {
	return string(s) + " IS NOT NULL"
}

func (s ColumnName[TYPE]) IsTRUE() string {
	return string(s) + " IS TRUE"
}

func (s ColumnName[TYPE]) IsFALSE() string {
	return string(s) + " IS FALSE"
}

func (s ColumnName[TYPE]) BetweenAND(arg1, arg2 TYPE) (string, TYPE, TYPE) {
	return string(s) + " BETWEEN ? AND ?", arg1, arg2
}

func (s ColumnName[TYPE]) Ob(direction string) ColumnOrderByAscDesc {
	return ColumnOrderByAscDesc("`" + string(s) + "`" + " " + direction)
}

func (s ColumnName[TYPE]) Qc(op string) QsCondition {
	return QsCondition(string(s) + " " + op)
}

func (s ColumnName[TYPE]) Qx(op string, x TYPE) *QxType {
	return &QxType{
		qc:   s.Qc(op),
		args: []interface{}{x},
	}
}

// Kw 得到只有1个元素的kw的map，这样能继续去增加元素
func (s ColumnName[TYPE]) Kw(x TYPE) KeywordArguments {
	return KeywordArguments{string(s): x}
}

// Kv 只是简单返回个k+v的结果，因为用的是泛型，因此能避免类型错误
func (s ColumnName[TYPE]) Kv(x TYPE) (string, TYPE) {
	return string(s), x
}
