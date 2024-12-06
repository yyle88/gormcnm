package gormcnm

// ColumnValueMap is a map type used for GORM update logic to represent column-value mappings.
// ColumnValueMap 是一个用于 GORM 更新逻辑的映射类型，用于表示列与值的对应关系表。
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
// Kws 将 ColumnValueMap 转换为 GORM 的 map[string]interface{}。
func (mp ColumnValueMap) Kws() map[string]interface{} {
	return mp
}

// Map is an alias for converting ColumnValueMap to map[string]interface{}.
// Map 是将 ColumnValueMap 转换为 map[string]interface{} 的别名。
func (mp ColumnValueMap) Map() map[string]interface{} {
	return mp
}

// AsMap is an alias for converting ColumnValueMap to map[string]interface{}.
// AsMap 是将 ColumnValueMap 转换为 map[string]interface{} 的别名。
func (mp ColumnValueMap) AsMap() map[string]interface{} {
	return mp
}

// ToMap is an alias for converting ColumnValueMap to map[string]interface{}.
// ToMap 是将 ColumnValueMap 转换为 map[string]interface{} 的别名。
func (mp ColumnValueMap) ToMap() map[string]interface{} {
	return mp
}
