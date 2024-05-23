package gormcnm

import "strings"

type SxType struct {
	*queryArgsTuple
}

func NewSx(qs string, args ...interface{}) *SxType {
	return &SxType{
		&queryArgsTuple{
			qc:   QsCondition(qs),
			args: args,
		},
	}
}

func Sx(qs string, args ...interface{}) *SxType {
	return NewSx(qs, args...)
}

func (sx *SxType) Combine(cs ...*SxType) *SxType {
	var qsVs []string
	qsVs = append(qsVs, sx.Qs())
	var args []any
	args = append(args, sx.Args()...)
	for _, c := range cs {
		qsVs = append(qsVs, c.Qs())
		args = append(args, c.Args()...)
	}
	var stmt = strings.Join(qsVs, ", ") //得到的就是gorm db.Select() 的要选中的列信息，因此使用逗号分隔
	return NewSx(stmt, args...)         //得到的就是 gorm db.Select() 的选中信息和附带的参数信息，比如 COUNT(CASE WHEN condition THEN 1 END) 里 condition 的参数信息
}
