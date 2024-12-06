package gormcnm

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yyle88/neatjson/neatjsons"
)

func TestClauseType_Column(t *testing.T) {
	const columnName = ColumnName[string]("name")

	clause := columnName.Clause()
	column := clause.Column()
	t.Log(neatjsons.S(column))
	require.Equal(t, columnName.Name(), column.Name)
}

func TestClauseType_Assignment(t *testing.T) {
	const columnRank = ColumnName[int]("rank")

	clauseType := columnRank.ClauseWithTable("students")
	assignment := clauseType.Assignment(888)
	t.Log(neatjsons.S(assignment))
	require.Equal(t, assignment.Column.Name, columnRank.Name())
	require.Equal(t, 888, assignment.Value)
	require.Equal(t, "students", assignment.Column.Table)
}
