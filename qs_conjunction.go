package gormcnm

import (
	"database/sql/driver"
	"fmt"
	"strings"
)

// QsConjunction means gorm query statement conjunction. example: OR AND NOT
// 就是表示 "或" "且" "非" 的连接词
// 在语法学中，conjunction（连词）是一种词类，用来连接词、短语、语句
type QsConjunction string

func (qc QsConjunction) AND(qcs ...QsConjunction) QsConjunction {
	var qss = make([]string, 0, 1+len(qcs)) //新空间-确保线程安全
	qss = append(qss, "("+qc.Qs()+")")
	for _, c := range qcs {
		qss = append(qss, "("+c.Qs()+")") //把能加括号的地方都加上以免出问题
	}
	return "(" + QsConjunction(strings.Join(qss, " AND ")) + ")"
}

func (qc QsConjunction) OR(qcs ...QsConjunction) QsConjunction {
	var qss = make([]string, 0, 1+len(qcs)) //新空间-确保线程安全
	qss = append(qss, "("+qc.Qs()+")")
	for _, c := range qcs {
		qss = append(qss, "("+c.Qs()+")") //把能加括号的地方都加上以免出问题
	}
	return "(" + QsConjunction(strings.Join(qss, " OR ")) + ")"
}

func (qc QsConjunction) NOT() QsConjunction {
	return QsConjunction(fmt.Sprintf("NOT(%s)", qc))
}

// Value 这块非常重要，要避免gorm直接使用这个结构，因此要在这里panic
func (qc QsConjunction) Value() (driver.Value, error) {
	panic(erxFunctionIsNotExecutable) //当报这个错时，需要修改调用侧代码，请看这个错误码的注释
}

// Qs 查询语句
// 因为我发现 db.Where(xxx("name = ? AND type = ?")) 这个是不行的
// 基本的结论是：
// 当你定义了一个类型为 type xxx string 的自定义类型，并尝试将其用作查询条件时，GORM 会将整个自定义类型作为一个值来处理，而不是将其展开为 SQL 语句中的占位符。
// 当你执行 db.Where(xxx("name = ? AND type = ?")) 时，GORM 会将 xxx("name = ? AND type = ?") 视为一个单独的值，再将其传递给查询条件，而不会将其展开为具体的 SQL 语句。
// 因此:
// 假如你不把结果再转换为string，就会在where条件中触发value异常
// 这也正是我在前面，实现一个错误的value函数的原因
func (qc QsConjunction) Qs() string {
	return string(qc)
}

func (qc QsConjunction) Qx() *QxType {
	var args = make([]interface{}, 0) //就这样吧，避免忘记
	return NewQx(string(qc), args...)
}
