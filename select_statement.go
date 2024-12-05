package gormcnm

import "strings"

type SxType = SelectStatement

func NewSx(stmt string, args ...interface{}) *SxType {
	return &SxType{
		statementArgumentsTuple: newStatementArgumentsTuple(stmt, args),
	}
}

// SelectStatement 就是传给 db.Select 的语句和参数
// 目前 "选中返回列" 的函数的定义是这样的
// func (db *DB) Select(query interface{}, args ...interface{}) (tx *DB)
// 在99%的场景下都是不需要传条件的
// 但在 SELECT COUNT(CASE WHEN (((name="abc") AND (type="xyz"))) THEN 1 END) as cnt FROM `examples` 这个语句里
// 很明显的 db.Select 也需要查询语句和参数 "abc" 和 "xyz"
// 而且这里条件有可能很长，而且有可能 db.Select 选中多个列的数据，就需要合并语句和合并参数
// 很明显各个列之间是逗号分隔的
// 因此定义这个类型，主要用来服务于这种场景（其实非常少用）
type SelectStatement struct {
	*statementArgumentsTuple
}

// NewSelectStatement creates a new SelectStatement with the provided query string and arguments.
// NewSelectStatement 使用提供的查询字符串和参数创建一个新的 SelectStatement 实例。
func NewSelectStatement(stmt string, args ...interface{}) *SelectStatement {
	return &SelectStatement{
		statementArgumentsTuple: newStatementArgumentsTuple(stmt, args),
	}
}

// Combine combines the current SelectStatement with other SelectStatements by merging their query strings and arguments.
// Combine 将当前的 SelectStatement 与其他 SelectStatement 合并，通过合并它们的查询字符串和参数。
func (selectStatement *SelectStatement) Combine(cs ...*SelectStatement) *SelectStatement {
	var qsVs []string
	qsVs = append(qsVs, selectStatement.Qs())
	var args []any
	args = append(args, selectStatement.Args()...)
	for _, c := range cs {
		qsVs = append(qsVs, c.Qs())
		args = append(args, c.Args()...)
	}
	var stmt = strings.Join(qsVs, ", ")      //得到的就是gorm db.Select() 的要选中的列信息，因此使用逗号分隔
	return NewSelectStatement(stmt, args...) //得到的就是 gorm db.Select() 的选中信息和附带的参数信息，比如 COUNT(CASE WHEN condition THEN 1 END) 里 condition 的参数信息
}
