package gormcnm

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

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

func (s ColumnName[TYPE]) NotEq(x TYPE) (string, TYPE) {
	return string(s) + "!=?", x
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

// Name Get column name. Column Name is abbreviated as cnm. return raw string.
// 因为偶尔需要 db.Select("name").Find(&one) 这样的操作，使用 string(cls.Name) 也行，但使用次数多了发现还是得封装个 Name 方法更方便
func (s ColumnName[TYPE]) Name() string {
	return string(s)
}

// AsAlias Return a raw string: "column_name as alias"
// 因为查询时偶尔要取别名，以适配结果，比如当自定义 resType 收集返回结果时，其字段就很有可能 和 models 里的字段是不同名的
func (s ColumnName[TYPE]) AsAlias(alias string) string {
	return stmtAsAlias(s.Name(), alias)
}

// SafeCnm returns a safe column name by enclosing it in backticks. Example: column name "type" -> "`type`" is safe.
// 比如你的列名就叫 create 的时候，就会和 SQL 的 create 冲突，就得使用引号把它引起来，否则不能执行。
// Use it when using db.Select("`type`").Find(&one) as an example.
// 就是当列名和数据库SQL关键字冲突时，需要用特殊手段使其不冲突，在gorm里就是添加反引号把字段引起来。
// 这样范型设计，代码就会变得很简单，比如当需要使用 Type 字段的时候，就可以使用 cls.Type.SafeCnm("“").Eq("value") 就能解决问题啦，能够完美贴合已有的所有逻辑。
// 至于自动化识别关键字的操作，我懒得做，因为实际使用场景也是很少的(Less is more)，当然主要是最初设计的时候忽略了这个情况，假如遇事不决都加引号也会比较繁琐。
// 注意：这个反引号 (`) 通常是用于 MySQL 数据库系统中，而在 Postgres 中应该使用双引号 (") 来引用列名。
func (s ColumnName[TYPE]) SafeCnm(quote string) ColumnName[TYPE] {
	switch len(quote) {
	case 0: //认为这种情况应该是误用或者是默认情况，就简单的两边增加空格就行
		return ColumnName[TYPE](" " + string(s) + " ")
	case 1: //比如参数是 quote := "\"" 或者 quote := "`" 的时候
		return ColumnName[TYPE](quote + string(s) + quote)
	case 2: //比如参数是 quote := "\"\"" 或者 quote := "``" 或者 "[]" 的时候
		return ColumnName[TYPE](quote[:1] + string(s) + quote[1:])
	default: //这里就不报错啦，而是取两端的字符，当gorm执行报错时，使用者自然能够找到原因
		return ColumnName[TYPE](quote[:1] + string(s) + quote[len(quote)-1:])
	}
}

func (s ColumnName[TYPE]) ExprAdd(v TYPE) clause.Expr {
	return gorm.Expr(string(s)+" + ?", v)
}

func (s ColumnName[TYPE]) ExprSub(v TYPE) clause.Expr {
	return gorm.Expr(string(s)+" - ?", v)
}

func (s ColumnName[TYPE]) Ob(direction string) ColumnOrderByAscDesc {
	return ColumnOrderByAscDesc(string(s) + " " + direction)
}

func (s ColumnName[TYPE]) Qc(op string) QsConjunction {
	return QsConjunction(string(s) + " " + op)
}

func (s ColumnName[TYPE]) Qx(op string, x TYPE) *QxType {
	stmt := string(s.Qc(op))
	args := []interface{}{x}
	return NewQx(stmt, args...)
}

// Kw 得到只有1个元素的kw的map，这样能继续去增加元素
func (s ColumnName[TYPE]) Kw(x TYPE) ColumnValueMap {
	return ColumnValueMap{string(s): x}
}

// Kv 只是简单返回个 k,v 的结果，因为用的是泛型，因此能避免类型错误，而这个 k,v 的结果恰巧可以传给gorm的Update函数(完美)。Example: db.Where(k.Eq("a")).Update(k.Kv("b")).Error (非常完美)。
func (s ColumnName[TYPE]) Kv(x TYPE) (string, TYPE) {
	return string(s), x
}

// KeExp 是在Kv的基础上新增的，返回个 k,expression 的结果，能传给gorm的Update函数。这个函数预计使用率不高，因此就不实现Kw的对应的KwExp啦(KwExp命名含义不明确)，因为使用率会更低些(主要是不知道该起啥名)。
func (s ColumnName[TYPE]) KeExp(x clause.Expr) (string, clause.Expr) {
	return string(s), x
}

// KeAdd 在db.Update的时候，通常字段自增或者加某个值更新的情况多些(因为这个函数的出现，我把Ke函数名称改为KeExp，因为Ke略短不太适合用于这种使用频率过低的函数，而KExpr不太美观)。
func (s ColumnName[TYPE]) KeAdd(x TYPE) (string, clause.Expr) {
	return s.KeExp(s.ExprAdd(x))
}

// KeSub 在db.Update的时候，让某个字段减去某个值，返回的还是 k,expression 的结果
func (s ColumnName[TYPE]) KeSub(x TYPE) (string, clause.Expr) {
	return s.KeExp(s.ExprSub(x))
}

func (s ColumnName[TYPE]) IFNULLFnStmt(sfn string, dfv string, alias string) string {
	if dfv == "" {
		dfv = "0"
	}
	//IFNULL 是 MySQL 特定的函数，在其他数据库系统中可能不支持
	stmt := "IFNULL(" + sfn + "(" + string(s) + "), " + dfv + ")"
	//when alias is not none return "stmt as alias"
	return stmtAsAlias(stmt, alias)
}

func (s ColumnName[TYPE]) CoalesceStmt(sfn string, dfv string, alias string) string {
	if dfv == "" {
		dfv = "0"
	}
	//COALESCE 是 SQL 标准中的函数，在大多数数据库系统中都支持
	stmt := "COALESCE(" + sfn + "(" + string(s) + "), " + dfv + ")"
	return stmtAsAlias(stmt, alias)
}

func (s ColumnName[TYPE]) CoalesceSumStmt(alias string) string {
	return s.CoalesceStmt("SUM", "0", alias)
}

func (s ColumnName[TYPE]) CoalesceMaxStmt(alias string) string {
	return s.CoalesceStmt("MAX", "0", alias)
}

func (s ColumnName[TYPE]) CoalesceMinStmt(alias string) string {
	return s.CoalesceStmt("MIN", "0", alias)
}

func (s ColumnName[TYPE]) CoalesceAvgStmt(alias string) string {
	return s.CoalesceStmt("AVG", "0", alias)
}

// Count COUNT(column_name) will only count rows where the given column is NOT NULL value
// 这个和表级别的 count(*) 还是有区别的
func (s ColumnName[TYPE]) Count(alias string) string {
	stmt := "COUNT(" + string(s) + ")"
	return stmtAsAlias(stmt, alias)
}

// CountDistinct COUNT(DISTINCT column_name) 在遇到 NULL 值时,会自动跳过,不将 NULL 计入统计结果
func (s ColumnName[TYPE]) CountDistinct(alias string) string {
	stmt := "COUNT(DISTINCT(" + string(s) + "))"
	return stmtAsAlias(stmt, alias)
}
