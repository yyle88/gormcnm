package gormcnm

import "gorm.io/gorm/clause"

func (s ColumnName[TYPE]) Clause() *ClauseType[TYPE] {
	return &ClauseType[TYPE]{
		cnm: s,
	}
}

func (s ColumnName[TYPE]) ClauseWithTable(tableName string) *ClauseType[TYPE] {
	return &ClauseType[TYPE]{
		cnm:   s,
		table: tableName,
	}
}

type ClauseType[TYPE any] struct {
	cnm   ColumnName[TYPE]
	table string
	alias string
	raw   bool
}

func (X *ClauseType[TYPE]) WithTable(tableName string) *ClauseType[TYPE] {
	X.table = tableName
	return X
}

func (X *ClauseType[TYPE]) WithAlias(alias string) *ClauseType[TYPE] {
	X.alias = alias
	return X
}

func (X *ClauseType[TYPE]) WithRaw(raw bool) *ClauseType[TYPE] {
	X.raw = raw
	return X
}

func (X *ClauseType[TYPE]) Column() clause.Column {
	return clause.Column{
		Table: X.table,
		Name:  X.cnm.Name(),
		Alias: X.alias,
		Raw:   X.raw, //非raw的会被继续加工为合理的语句，比如增加表名，增加转义符号等，因此这里推荐使用false(默认值)
	}
}

func (X *ClauseType[TYPE]) Assignment(value TYPE) clause.Assignment {
	return clause.Assignment{Column: X.Column(), Value: value}
}
