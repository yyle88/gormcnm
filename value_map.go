// Package gormcnm provides column-value map operations for GORM update and batch operations
// Auto manages column-value mappings for Updates and UpdateColumns operations
// Supports building dynamic update maps with type-safe column assignments
//
// gormcnm 提供列值映射操作，用于 GORM 更新和批量操作
// 自动管理 Updates 和 UpdateColumns 操作的列值映射
// 支持构建动态更新映射，具备类型安全的列赋值
package gormcnm

// ColumnValueMap is a map type used for GORM update logic to represent column-value mappings.
// This type is used when calling gormrepo.Updates, gormrepo.UpdatesM, or GORM native db.Updates.
// It provides type-safe column-value assignments with chainable Kw() method.
//
// ColumnValueMap 是一个用于 GORM 更新逻辑的映射类型，用于表示列与值的对应关系表。
// 该类型在调用 gormrepo.Updates、gormrepo.UpdatesM 或 GORM 原生 db.Updates 时使用。
// 提供类型安全的列值赋值，支持链式调用 Kw() 方法。
//
// Usage with gormrepo.Updates (requires AsMap() conversion):
// 使用 gormrepo.Updates 时（需要 AsMap() 转换）：
//
//	repo.Updates(where, func(cls *AccountColumns) map[string]interface{} {
//	    return cls.
//	        Kw(cls.Nickname.Kv(newNickname)).
//	        Kw(cls.Password.Kv(newPassword)).
//	        AsMap() // Convert to map[string]interface{} // 转换为 map[string]interface{} 类型
//	})
//
// Usage with gormrepo.UpdatesM (no AsMap() needed):
// 使用 gormrepo.UpdatesM 时（无需 AsMap()）：
//
//	repo.UpdatesM(where, func(cls *AccountColumns) gormcnm.ColumnValueMap {
//	    return cls.
//	        Kw(cls.Nickname.Kv(newNickname)).
//	        Kw(cls.Password.Kv(newPassword))
//	    // No AsMap() needed! // 不需要 AsMap()！
//	})
type ColumnValueMap map[string]interface{}

// NewKw initializes a new ColumnValueMap.
// NewKw 初始化一个新的 ColumnValueMap。
func NewKw() ColumnValueMap {
	return make(ColumnValueMap)
}

// Kw creates a new ColumnValueMap and adds a key-value.
// Kw 创建一个新的 ColumnValueMap 并添加一个键值对。
func Kw(columnName string, value interface{}) ColumnValueMap {
	return NewKw().Kw(columnName, value)
}

// Kw adds a key-value pair to ColumnValueMap and supports chaining.
// Kw 向 ColumnValueMap 添加键值对并支持链式调用。
func (mp ColumnValueMap) Kw(columnName string, value interface{}) ColumnValueMap {
	mp[columnName] = value
	return mp
}

// Kws converts ColumnValueMap to map[string]interface{} for GORM.
// Required when using gormrepo.Updates, not needed when using gormrepo.UpdatesM.
// Recommend using AsMap() instead, as it has cleaner semantics and avoids naming conflicts.
//
// Kws 将 ColumnValueMap 转换为 GORM 的 map[string]interface{}。
// 使用 gormrepo.Updates 时需要调用，使用 gormrepo.UpdatesM 时不需要。
// 推荐使用 AsMap()，语义更明确且不易与其他名称冲突。
func (mp ColumnValueMap) Kws() map[string]interface{} {
	return mp
}

// Map converts ColumnValueMap to map[string]interface{} for GORM.
// Required when using gormrepo.Updates, not needed when using gormrepo.UpdatesM.
// Recommend using AsMap() instead, as it has cleaner semantics and avoids naming conflicts.
//
// Map 将 ColumnValueMap 转换为 map[string]interface{}。
// 使用 gormrepo.Updates 时需要调用，使用 gormrepo.UpdatesM 时不需要。
// 推荐使用 AsMap()，语义更明确且不易与其他名称冲突。
func (mp ColumnValueMap) Map() map[string]interface{} {
	return mp
}

// AsMap converts ColumnValueMap to map[string]interface{} for GORM.
// Required when using gormrepo.Updates, not needed when using gormrepo.UpdatesM.
// This is the recommended conversion method with clear semantics.
//
// AsMap 将 ColumnValueMap 转换为 map[string]interface{}。
// 使用 gormrepo.Updates 时需要调用，使用 gormrepo.UpdatesM 时不需要。
// 这是推荐的转换方法，语义清晰明确。
func (mp ColumnValueMap) AsMap() map[string]interface{} {
	return mp
}

// ToMap converts ColumnValueMap to map[string]interface{} for GORM.
// Required when using gormrepo.Updates, not needed when using gormrepo.UpdatesM.
// Recommend using AsMap() instead, as it has cleaner semantics and avoids naming conflicts.
//
// ToMap 将 ColumnValueMap 转换为 map[string]interface{}。
// 使用 gormrepo.Updates 时需要调用，使用 gormrepo.UpdatesM 时不需要。
// 推荐使用 AsMap()，语义更明确且不易与其他名称冲突。
func (mp ColumnValueMap) ToMap() map[string]interface{} {
	return mp
}
