package gormcnm

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
