// Package utils provides utility types and interfaces for GORM table operations
// Auto discover and handle column names with type safety and interface abstractions
// Support both direct table name strings and GORM table interface implementations
//
// utils 包提供 GORM 表操作的实用类型和接口
// 自动发现和处理列名，具有类型安全性和接口抽象
// 支持直接表名字符串和 GORM 表接口实现
package utils

// ColumnNameInterface defines the interface for column name providers
// ColumnNameInterface 定义列名提供者的接口
type ColumnNameInterface interface {
	Name() string
}

// GormTableNameFace defines the interface for GORM table name providers
// Compatible with GORM's table name interface pattern for seamless integration
// Used in table join operations and qualified column name generation
//
// GormTableNameFace 定义 GORM 表名提供者的接口
// 兼容 GORM 表名接口模式以实现无缝集成
// 用于表连接操作和限定列名生成
type GormTableNameFace interface {
	TableName() string // Returns the table name for database operations // 返回用于数据库操作的表名
}

// TableNameImp provides a simple implementation of table name interface
// Acts as a concrete implementation when direct table name string is available
// Used internally for table operations that need GormTableNameFace compliance
//
// TableNameImp 提供表名接口的简单实现
// 在有直接表名字符串时作为具体实现
// 内部用于需要符合 GormTableNameFace 的表操作
type TableNameImp struct {
	tableName string // Internal storage for table name // 表名的内部存储
}

// NewTableNameImp creates a new TableNameImp instance with the specified table name
// NewTableNameImp 使用指定的表名创建一个新的 TableNameImp 实例
func NewTableNameImp(tableName string) *TableNameImp {
	return &TableNameImp{tableName: tableName}
}

// TableName returns the table name
// TableName 返回表名
func (X *TableNameImp) TableName() string {
	return X.tableName
}
