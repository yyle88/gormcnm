package gormcnm

// ColumnValueMap 这个类型原本命名为 KeywordArguments 宽泛的表示字典，但它只被用于gorm的更新逻辑
// 因此改名为 ColumnValueMap 让意图更明显，但下面的 Kw 命名就不改啦，毕竟就是个字典，使用 Kw 很合适
type ColumnValueMap map[string]interface{}

func NewKw() ColumnValueMap {
	return make(ColumnValueMap)
}

func Kw(columnName string, value interface{}) ColumnValueMap {
	return NewKw().Kw(columnName, value)
}

// Kw 再去接收下一个Kv得到kw的map，因此这个函数命名也挺合适的，毕竟词汇量有限，而这个也没背占用
// Kw 这个命名的侧重点在于它的返回值是字典，而且和前面的 Kw 相同
// 这样有利于链式调用
func (mp ColumnValueMap) Kw(columnName string, value interface{}) ColumnValueMap {
	mp[columnName] = value
	return mp //这样就是一个链式反应的
}

// Kws 把我的类型转换为gorm能够识别的类型
// 这块也比较奇怪，就是 gorm 在 UpdateColumns 的时候只能接收有限的类型
// 比如 map[string]interface{} 或者 struct 或者 *struct
// 其他类型，就会返回 gorm.ErrInvalidData 具体代码请看gorm的源码
// 还好吧，在没有用对的时候至少是报了错 (排序那块不是特定的类型直接忽略，查询不报错，比这个还坑)
// 因此当你最终设置完以后还是要使用这个函数转换，否则执行更新就会报错
// 我建议让gorm另写个update函数，比如UpdateColumnsWithInterface(a ValuesInterface)，而我的类型也实现这个接口（比如实现Map()/ToMap()/AsMap()函数），这样就好啦
// 但现在没有的话显然还是得自己转换
// 还是老样子提供若干个不同名的函数
func (mp ColumnValueMap) Kws() map[string]interface{} {
	return mp
}

func (mp ColumnValueMap) Map() map[string]interface{} {
	return mp //换个函数名，使用者喜欢用哪个就用哪个吧
}

func (mp ColumnValueMap) AsMap() map[string]interface{} {
	return mp //换个函数名，使用者喜欢用哪个就用哪个吧
}

func (mp ColumnValueMap) ToMap() map[string]interface{} {
	return mp //换个函数名，使用者喜欢用哪个就用哪个吧
}
