// Package gormcnm tests validate GORM clause operations and column assignments
// Auto verifies Clause, ClauseWithTable methods and ClauseColumn operations
// Tests cover clause creation, table qualification, alias assignment, and GORM integration
//
// gormcnm 测试包验证 GORM 子句操作和列赋值
// 自动验证 Clause、ClauseWithTable 方法和 ClauseColumn 功能
// 测试涵盖子句创建、表限定、别名赋值和 GORM 集成
package gormcnm

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yyle88/neatjson/neatjsons"
)

func TestClauseColumn_Column(t *testing.T) {
	const columnName = ColumnName[string]("name")

	clause := columnName.Clause()
	column := clause.Column()
	t.Log(neatjsons.S(column))
	require.Equal(t, columnName.Name(), column.Name)
}

func TestClauseColumn_Assignment(t *testing.T) {
	const columnRank = ColumnName[int]("rank")

	clauseType := columnRank.ClauseWithTable("students")
	assignment := clauseType.Assignment(888)
	t.Log(neatjsons.S(assignment))
	require.Equal(t, assignment.Column.Name, columnRank.Name())
	require.Equal(t, 888, assignment.Value)
	require.Equal(t, "students", assignment.Column.Table)
}
