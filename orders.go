package gormcnm

import "gorm.io/gorm"

// OrderByBottle represents a sort statement builder, designed with a unique naming style that reflects a materials science focus.
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

// Ox converts the OrderByBottle to a string. Note that if the type is not specific, it may be ignored by GORM's logic.
// Ox 将 OrderByBottle 转换为字符串。请注意，如果类型不明确，它可能会被 GORM 的逻辑忽略。
// This is an unavoidable limitation due to GORM's handling of the Order field logic.
// 这是由于 GORM 对 Order 字段逻辑的处理所造成的无法避免的限制。
// Users may forget to convert this to a string before passing it to GORM, so it is important to remember this step.
// 用户可能会忘记在传递给 GORM 之前将其转换为字符串，因此需要记住这一点。
// There is currently no elegant solution to this limitation, but it should work fine for personal use.
// 目前没有优雅的解决方案，但对于个人使用来说应该没有问题。
func (ob OrderByBottle) Ox() string {
	return string(ob)
}

// Orders converts the OrderByBottle to a string. Note that if the type is not specific, it may be ignored by GORM's logic.
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
