package gormcnm

// COALESCE 是 SQL 标准中的函数，在大多数数据库系统中都支持
func (s ColumnName[TYPE]) COALESCE() *coalesceQs {
	return newCoalesceQs("COALESCE", string(s))
}

// IFNULLFn 调用 IFNULL 函数，因需要与 IsNULL 这个更常用的函数区分，因此把名称改为这样的
// IFNULL 是 MySQL 特定的函数，在其他数据库系统中可能不支持
func (s ColumnName[TYPE]) IFNULLFn() *coalesceQs {
	return newCoalesceQs("IFNULL", string(s))
}

type coalesceQs struct {
	ident string //就是关键字 "IFNULL"/"COALESCE"
	cname string //这就是列名
}

func newCoalesceQs(ident string, cname string) *coalesceQs {
	return &coalesceQs{
		ident: ident,
		cname: cname,
	}
}

func (qs *coalesceQs) Stmt(sfn string, dfv string, alias string) string {
	if dfv == "" {
		dfv = "0"
	}
	// COALESCE 是 SQL 标准中的函数，在大多数数据库系统中都支持
	// IFNULL 是 MySQL 特定的函数，在其他数据库系统中可能不支持
	stmt := qs.ident + "(" + sfn + "(" + string(qs.cname) + "), " + dfv + ")"
	return stmtAsAlias(stmt, alias) // when alias is not none return "stmt as alias"
}

func (qs *coalesceQs) SumStmt(alias string) string {
	return qs.Stmt("SUM", "0", alias)
}

func (qs *coalesceQs) MaxStmt(alias string) string {
	return qs.Stmt("MAX", "0", alias)
}

func (qs *coalesceQs) MinStmt(alias string) string {
	return qs.Stmt("MIN", "0", alias)
}

func (qs *coalesceQs) AvgStmt(alias string) string {
	return qs.Stmt("AVG", "0", alias)
}
