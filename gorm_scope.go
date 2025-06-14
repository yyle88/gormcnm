package gormcnm

import (
	"gorm.io/gorm"
)

// ScopeFunction is a type alias for a function that modifies a GORM DB instance,
// used with db.Scopes() to apply custom query conditions.
// See: https://github.com/go-gorm/gorm/blob/c44405a25b0fb15c20265e672b8632b8774793ca/chainable_api.go#L376
// ScopeFunction 是用于修改 GORM DB 实例的函数类型别名，
// 主要是和 db.Scopes() 配合使用，以应用自定义查询条件。
// 详见：https://github.com/go-gorm/gorm/blob/c44405a25b0fb15c20265e672b8632b8774793ca/chainable_api.go#L376
type ScopeFunction = func(db *gorm.DB) *gorm.DB
