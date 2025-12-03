// Package gormcnm provides GORM scope function type definitions, enabling custom GORM conditions
// Auto enables type-safe scope functions that integrate with GORM's db.Scopes() method
// Supports building reusable where modifiers and composable database operations
//
// gormcnm 提供 GORM 作用域函数类型定义，用于自定义查询条件
// 自动启用类型安全的作用域函数，与 GORM 的 db.Scopes() 方法集成
// 支持构建可重用的查询修饰符和可组合的数据库操作
package gormcnm

import (
	"gorm.io/gorm"
)

// ScopeFunction is a type alias, representing a function that modifies a GORM DB instance,
// used with db.Scopes() to use custom GORM conditions.
// See: https://github.com/go-gorm/gorm/blob/c44405a25b0fb15c20265e672b8632b8774793ca/chainable_api.go#L376
// ScopeFunction 是用于修改 GORM DB 实例的函数类型别名，
// 主要是和 db.Scopes() 配合使用，以应用自定义查询条件。
// 详见：https://github.com/go-gorm/gorm/blob/c44405a25b0fb15c20265e672b8632b8774793ca/chainable_api.go#L376
type ScopeFunction = func(db *gorm.DB) *gorm.DB
