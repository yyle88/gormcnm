// Package gormcnmjson provides type-safe JSON column operations for GORM
// Supports SQLite JSON functions with compile-time type checking
// Works with both string and []byte JSON column types
//
// gormcnmjson 为 GORM 提供类型安全的 JSON 列操作
// 支持 SQLite JSON 函数并提供编译时类型检查
// 同时支持 string 和 []byte 类型的 JSON 列
package gormcnmjson

import (
	"fmt"

	"github.com/yyle88/gormcnm"
)

// Column represents a JSON column with type-safe SQL operations
// Provides methods to generate JSON-specific SQL expressions
//
// Column 表示一个 JSON 列，提供类型安全的 SQL 操作
// 提供生成 JSON 特定 SQL 表达式的方法
type Column struct {
	name string // Column name in database // 数据库中的列名
}

// New creates a Column from a ColumnName with generic type support
// Accepts both string and []byte column types with versatile options
//
// New 从 ColumnName 创建 Column，支持泛型类型
// 接受 string 和 []byte 列类型，提供多种选择
func New[T ~string | ~[]byte](columnName gormcnm.ColumnName[T]) Column {
	return Column{name: columnName.Name()}
}

// Raw creates a Column from a []byte-based ColumnName
// Dedicated constructor for datatypes.JSON and related []byte types
//
// Raw 从基于 []byte 的 ColumnName 创建 Column
// 专门用于 datatypes.JSON 和相关的 []byte 类型
func Raw[T ~[]byte](columnName gormcnm.ColumnName[T]) Column {
	return Column{name: columnName.Name()}
}

// Name returns the underlying column name as a string
// Use this when you need the raw column name in SQL expressions
//
// Name 返回底层列名字符串
// 当需要在 SQL 表达式中使用原始列名时使用此方法
func (co Column) Name() string {
	return co.name
}

// Get extracts a JSON value as text using the ->> operation
// Returns a type-safe string ColumnName for chaining conditions
//
// Get 使用 ->> 操作将 JSON 值提取为文本
// 返回类型安全的字符串 ColumnName 用于链式条件
func (co Column) Get(path string) gormcnm.ColumnName[string] {
	return gormcnm.ColumnName[string](
		fmt.Sprintf("%s ->> '$.%s'", co.name, path),
	)
}

// Extract extracts a JSON sub-object using the -> operation
// Returns a Column for additional nested operations
//
// Extract 使用 -> 操作提取 JSON 子对象
// 返回 Column 用于额外的嵌套操作
func (co Column) Extract(path string) Column {
	return Column{
		name: fmt.Sprintf("%s -> '$.%s'", co.name, path),
	}
}

// GetInt extracts a JSON value as an integer with type casting
// Returns a type-safe int ColumnName for numeric comparisons
//
// GetInt 将 JSON 值提取为整数并进行类型转换
// 返回类型安全的 int ColumnName 用于数值比较
func (co Column) GetInt(path string) gormcnm.ColumnName[int] {
	return gormcnm.ColumnName[int](
		fmt.Sprintf("CAST(%s ->> '$.%s' AS INTEGER)", co.name, path),
	)
}

// Length returns the length of a JSON text/object using JSON_ARRAY_LENGTH
// If path is empty, measures the root JSON; otherwise measures the nested path
//
// Length 使用 JSON_ARRAY_LENGTH 返回 JSON 数组的长度
// 如果 path 为空则测量根 JSON，否则测量嵌套路径
func (co Column) Length(path string) gormcnm.ColumnName[int] {
	if path == "" {
		return gormcnm.ColumnName[int](
			fmt.Sprintf("JSON_ARRAY_LENGTH(%s)", co.name),
		)
	}
	return gormcnm.ColumnName[int](
		fmt.Sprintf("JSON_ARRAY_LENGTH(%s, '$.%s')", co.name, path),
	)
}

// Type returns the JSON type of a value using JSON_TYPE function
// If path is empty, checks root JSON type; otherwise checks the nested path
//
// Type 使用 JSON_TYPE 函数返回 JSON 值的类型
// 如果 path 为空则检查根 JSON 类型，否则检查嵌套路径
func (co Column) Type(path string) gormcnm.ColumnName[string] {
	if path == "" {
		return gormcnm.ColumnName[string](
			fmt.Sprintf("JSON_TYPE(%s)", co.name),
		)
	}
	return gormcnm.ColumnName[string](
		fmt.Sprintf("JSON_TYPE(%s, '$.%s')", co.name, path),
	)
}

// Valid checks if the JSON text is well-formed using JSON_VALID
// Returns 1 for valid JSON, 0 for invalid JSON
//
// Valid 使用 JSON_VALID 检查 JSON 文本是否格式正确
// 返回 1 表示有效的 JSON，返回 0 表示无效的 JSON
func (co Column) Valid() gormcnm.ColumnName[int] {
	return gormcnm.ColumnName[int](
		fmt.Sprintf("JSON_VALID(%s)", co.name),
	)
}

// Set updates a JSON value at the specified path using JSON_SET
// Returns a Column with the modified expression for use in UPDATE statements
//
// Set 使用 JSON_SET 在指定路径更新 JSON 值
// 返回包含修改表达式的 Column 用于 UPDATE 语句
func (co Column) Set(path string, value interface{}) Column {
	return Column{
		name: fmt.Sprintf("JSON_SET(%s, '$.%s', '%v')", co.name, path, value),
	}
}

// Remove deletes a value at the specified path using JSON_REMOVE
// Returns a Column with the removal expression for use in UPDATE statements
//
// Remove 使用 JSON_REMOVE 删除指定路径的值
// 返回包含删除表达式的 Column 用于 UPDATE 语句
func (co Column) Remove(path string) Column {
	return Column{
		name: fmt.Sprintf("JSON_REMOVE(%s, '$.%s')", co.name, path),
	}
}

// AsAlias creates a column alias for use in SELECT statements
// Returns the column expression with the specified alias name
//
// AsAlias 为 SELECT 语句创建列别名
// 返回带有指定别名的列表达式
func (co Column) AsAlias(alias string) string {
	return co.name + " as " + alias
}
