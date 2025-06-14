package gormcnm

import "gorm.io/gorm"

type QxType = QxConjunction

func NewQx(stmt string, args ...interface{}) *QxType {
	return &QxType{
		statementArgumentsTuple: newStatementArgumentsTuple(stmt, args),
	}
}

// QxConjunction is used for constructing relational query statements (AND, OR, NOT) and their arguments for db.Where.
// QxConjunction 用于构造关系查询语句（AND、OR、NOT）以及其对应的参数，以供 db.Where 使用。
// Example: When combining conditions like a.Eq("xyz") and b.Eq("uvw"), this class concatenates statements with (---) AND (---) and merges arguments into a new list.
// 示例：当需要组合条件如 a.Eq("xyz") 和 b.Eq("uvw") 时，该工具类将语句用 (---) AND (---) 连接，并将参数列表合并为新的列表。
type QxConjunction struct {
	*statementArgumentsTuple
}

// NewQxConjunction creates a new instance of QxConjunction with the provided statement and arguments.
// NewQxConjunction 使用提供的语句和参数创建一个新的 QxConjunction 实例。
func NewQxConjunction(stmt string, args ...interface{}) *QxConjunction {
	return &QxConjunction{
		statementArgumentsTuple: newStatementArgumentsTuple(stmt, args),
	}
}

// Qx is shorthand for creating a new QxConjunction instance.
// Qx 是创建 QxConjunction 实例的简写形式。
func Qx(stmt string, args ...interface{}) *QxConjunction {
	return NewQxConjunction(stmt, args...)
}

// AND combines the current QxConjunction instance with multiple QxConjunction instances using "AND".
// AND 使用 "AND" 将当前 QxConjunction 实例与多个 QxConjunction 实例组合在一起。
func (qx *QxConjunction) AND(cs ...*QxConjunction) *QxConjunction {
	var qss []QsConjunction
	var qas []*statementArgumentsTuple
	for _, c := range cs {
		qss = append(qss, QsConjunction(c.stmt))
		qas = append(qas, c.statementArgumentsTuple)
	}

	return &QxConjunction{
		statementArgumentsTuple: newStatementArgumentsTuple(string(QsConjunction(qx.stmt).AND(qss...)), qx.safeCombineArguments(qas)),
	}
}

// OR combines the current QxConjunction instance with multiple QxConjunction instances using "OR".
// OR 使用 "OR" 将当前 QxConjunction 实例与多个 QxConjunction 实例组合在一起。
func (qx *QxConjunction) OR(cs ...*QxConjunction) *QxConjunction {
	var qss []QsConjunction
	var qas []*statementArgumentsTuple
	for _, c := range cs {
		qss = append(qss, QsConjunction(c.stmt))
		qas = append(qas, c.statementArgumentsTuple)
	}
	return &QxConjunction{
		statementArgumentsTuple: newStatementArgumentsTuple(string(QsConjunction(qx.stmt).OR(qss...)), qx.safeCombineArguments(qas)),
	}
}

// NOT negates the current QxConjunction instance by wrapping the statement with "NOT".
// NOT 通过在语句外包裹 "NOT" 来对当前 QxConjunction 实例进行逻辑取反。
func (qx *QxConjunction) NOT() *QxConjunction {
	return &QxConjunction{
		statementArgumentsTuple: newStatementArgumentsTuple(string(QsConjunction(qx.stmt).NOT()), qx.args),
	}
}

// AND1 creates a new QxConjunction instance with the given statement and arguments, then combines it with the current instance using "AND".
// AND1 使用给定的语句和参数创建一个新的 QxConjunction 实例，然后使用 "AND" 将其与当前实例组合。
func (qx *QxConjunction) AND1(stmt string, args ...interface{}) *QxConjunction {
	return qx.AND(NewQxConjunction(stmt, args...))
}

// OR1 creates a new QxConjunction instance with the given statement and arguments, then combines it with the current instance using "OR".
// OR1 使用给定的语句和参数创建一个新的 QxConjunction 实例，然后使用 "OR" 将其与当前实例组合。
func (qx *QxConjunction) OR1(stmt string, args ...interface{}) *QxConjunction {
	return qx.OR(NewQxConjunction(stmt, args...))
}

// Scope converts the QxConjunction to a GORM ScopeFunction used with db.Scopes().
// It applies the query conditions defined by QxConjunction to the GORM select.
// Scope 将 QxConjunction 转换为 GORM 的 ScopeFunction，以便于被 db.Scopes() 调用。
// 它将 QxConjunction 定义的查询条件应用于 GORM 查询。
func (qx *QxConjunction) Scope() ScopeFunction {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where(qx.Qs(), qx.args...)
	}
}
