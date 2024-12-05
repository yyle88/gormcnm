package gormcnm

import (
	"database/sql/driver"
	"fmt"
	"strings"
)

type QsType = QsConjunction

func NewQs(stmt string) QsType {
	return QsType(stmt)
}

// QsConjunction means gorm query statement conjunction. example: OR AND NOT
// QsConjunction 表示 GORM 查询语句中的连接词，例如 OR、AND、NOT
// 就是表示 "或"、"且"、"非" 的连接词
// 在语法学中，conjunction（连词）是一种词类，用来连接词、短语、语句
type QsConjunction string

func NewQsConjunction(stmt string) QsConjunction {
	return QsConjunction(stmt)
}

// AND constructs a conjunction using "AND" between QsConjunction instances.
// AND 使用 "AND" 在多个 QsConjunction 实例之间构造连接语句。
func (qsConjunction QsConjunction) AND(qcs ...QsConjunction) QsConjunction {
	var qss = make([]string, 0, 1+len(qcs)) // New slice to ensure thread safety // 新建一个切片以确保线程安全
	qss = append(qss, "("+qsConjunction.Qs()+")")
	for _, c := range qcs {
		qss = append(qss, "("+c.Qs()+")") // Add parentheses around each component to avoid logic issues // 在每个组件周围加括号以避免逻辑问题
	}
	return "(" + QsConjunction(strings.Join(qss, " AND ")) + ")"
}

// OR constructs a conjunction using "OR" between QsConjunction instances.
// OR 使用 "OR" 在多个 QsConjunction 实例之间构造连接语句。
func (qsConjunction QsConjunction) OR(qcs ...QsConjunction) QsConjunction {
	var qss = make([]string, 0, 1+len(qcs)) // New slice to ensure thread safety // 新建一个切片以确保线程安全
	qss = append(qss, "("+qsConjunction.Qs()+")")
	for _, c := range qcs {
		qss = append(qss, "("+c.Qs()+")") // Add parentheses around each component to avoid logic issues // 在每个组件周围加括号以避免逻辑问题
	}
	return "(" + QsConjunction(strings.Join(qss, " OR ")) + ")"
}

// NOT negates the QsConjunction instance by wrapping it with "NOT".
// NOT 通过添加 "NOT" 来对 QsConjunction 实例进行逻辑取反。
func (qsConjunction QsConjunction) NOT() QsConjunction {
	return QsConjunction(fmt.Sprintf("NOT(%s)", qsConjunction))
}

// Value prevents GORM from directly using QsConjunction by causing a panic.
// Value 阻止 GORM 直接使用 QsConjunction，通过触发 panic 实现。
// 当你定义了一个类型为 type xxx string 的自定义类型，并尝试将其用作查询条件时，GORM 会将整个自定义类型作为一个值来处理，而不是将其展开为 SQL 语句中的占位符。
// 因此也就是说 db.Where(xxx("name = ? AND type = ?")) 这个是不行的
// 基本的结论是：
// 当你执行 db.Where(xxx("name = ? AND type = ?")) 时，GORM 会将 xxx("name = ? AND type = ?") 视为一个单独的值，再将其传递给查询条件，而不会将其展开为具体的 SQL 语句。
// 这时，给 type xxx string 增加 Value 函数，而且里面报panic，就能避免用错
// 这就是这里给出 Value { panic } 的原因
// 这样，假如你不把结果转换为string，而是直接往where条件里传递，就会在where条件中触发value异常
func (qsConjunction QsConjunction) Value() (driver.Value, error) {
	panic(valueIsNotCallable) // If this error occurs, the caller must adjust their code. See error code comments.
	// 如果报这个错误，调用侧需要修改代码。具体可参考错误码的注释。
}

// Qs converts the QsConjunction instance into a string representation.
// Qs 将 QsConjunction 实例转换为字符串表示。
func (qsConjunction QsConjunction) Qs() string {
	return string(qsConjunction)
}

// Qx converts the QsConjunction instance into a QxConjunction object with arguments.
// Qx 将 QsConjunction 实例转换为带参数的 QxConjunction 对象。
func (qsConjunction QsConjunction) Qx() *QxConjunction {
	var args = make([]interface{}, 0) // Means no args // 意味着没有参数
	return NewQxConjunction(string(qsConjunction), args...)
}
