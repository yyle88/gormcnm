// Package gormcnm provides ORDER BY statement building operations with flexible sorting patterns
// Auto creates OrderByBottle instances with ASC/DESC direction support and combination logic
// Supports building complex sorting clauses with multiple columns and GORM integration
//
// gormcnm 包提供 ORDER BY 语句构建操作，具有灵活的排序模式
// 自动创建 OrderByBottle 实例，支持 ASC/DESC 方向和组合逻辑
// 支持构建多列的复杂排序子句，具有 GORM 集成
package gormcnm

import "gorm.io/gorm"

// OrderByBottle represents a sort statement construction toolkit, designed with a unique naming style that reflects a materials science focus.
// OrderByBottle 代表排序语句构建器，使用了与材料学相关的命名风格。
type OrderByBottle string

// Ob concatenates the current OrderByBottle with the next one, forming a combined ordering string.
// Ob 将当前的 OrderByBottle 与下一个 OrderByBottle 连接，形成一个组合的排序字符串。
func (ob OrderByBottle) Ob(next OrderByBottle) OrderByBottle {
	return ob + " , " + next
}

// OrderByBottle concatenates the current OrderByBottle with the next one, forming a combined ordering string.
// OrderByBottle 将当前的 OrderByBottle 与下一个 OrderByBottle 连接，形成一个组合的排序字符串。
func (ob OrderByBottle) OrderByBottle(next OrderByBottle) OrderByBottle {
	return ob + " , " + next
}

// Ox converts the OrderByBottle to a string. Note that if the type is not specific, it could be ignored by GORM's logic.
// Ox 将 OrderByBottle 转换为字符串。请注意，如果类型不明确，它可能会被 GORM 的逻辑忽略。
// This is an unavoidable limitation due to GORM's handling of the Order field logic.
// 这是由于 GORM 对 Order 字段逻辑的处理所造成的无法避免的限制。
// Developers might forget to convert this to a string before passing it to GORM, so it is important to do this step.
// 开发者可能会忘记在传递给 GORM 之前将其转换为字符串，因此需要记住这一点。
// There is no elegant solution to this limitation now, but it should work fine in most cases.
// 现在没有优雅的解决方案，但在大多数情况下应该没有问题。
func (ob OrderByBottle) Ox() string {
	return string(ob)
}

// Orders converts the OrderByBottle to a string. Note that if the type is not specific, it could be ignored by GORM's logic.
// Orders 将 OrderByBottle 转换为字符串。请注意，如果类型不明确，它可能会被 GORM 的逻辑忽略。
func (ob OrderByBottle) Orders() string {
	return string(ob)
}

// Scope converts the OrderByBottle to a GORM ScopeFunction used with db.Scopes().
// It applies the ordering defined by OrderByBottle to the GORM select.
// Scope 将 OrderByBottle 转换为 GORM 的 ScopeFunction，以便被 db.Scopes() 调用。
// 它将 OrderByBottle 定义的排序规则应用于 GORM 查询。
func (ob OrderByBottle) Scope() ScopeFunction {
	return func(db *gorm.DB) *gorm.DB {
		return db.Order(string(ob))
	}
}
