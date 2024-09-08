package gormcnm

import "gorm.io/gorm/clause"

func (s ColumnName[TYPE]) Clause() *ClauseType[TYPE] {
	return &ClauseType[TYPE]{
		cnm: s,
	}
}

func (s ColumnName[TYPE]) ClauseWithTableName(tableName string) *ClauseType[TYPE] {
	return &ClauseType[TYPE]{
		cnm: s,
		tab: tableName,
	}
}

type ClauseType[TYPE any] struct {
	cnm ColumnName[TYPE]
	tab string
}

func (X *ClauseType[TYPE]) WithTableName(tab string) *ClauseType[TYPE] {
	X.tab = tab
	return X
}

func (X *ClauseType[TYPE]) Column() clause.Column {
	return clause.Column{
		Table: X.tab,
		Name:  X.cnm.Name(),
		Alias: "",    //TODO 将来有需求时再添加
		Raw:   false, //非raw的会被继续加工为合理的语句，比如增加标名，以及增加转义符号等，因此这里使用false就行
	}
}

func (X *ClauseType[TYPE]) Assignment(x TYPE) clause.Assignment {
	return clause.Assignment{
		Column: X.Column(),
		Value:  x,
	}
}
