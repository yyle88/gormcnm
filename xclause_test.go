package gormcnm

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yyle88/neatjson/neatjsons"
)

func TestClauseType_Column(t *testing.T) {
	const name = ColumnName[string]("name")

	clause := name.Clause()
	column := clause.Column()
	t.Log(neatjsons.S(column))
	require.Equal(t, name.Name(), column.Name)
}

func TestClauseType_Assignment(t *testing.T) {
	const rank = ColumnName[int]("rank")

	clauseType := rank.ClauseWithTable("students")
	assignment := clauseType.Assignment(888)
	t.Log(neatjsons.S(assignment))
	require.Equal(t, assignment.Column.Name, rank.Name())
	require.Equal(t, 888, assignment.Value)
	require.Equal(t, "students", assignment.Column.Table)
}
