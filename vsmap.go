package gormcnm

type KeywordArguments map[string]interface{}

func NewKw() KeywordArguments {
	return make(KeywordArguments)
}

func Kw(columnName string, value interface{}) KeywordArguments {
	return NewKw().Kw(columnName, value)
}

// Kw 再去接收下一个Kv得到kw的map，因此这个函数命名也挺合适的，毕竟词汇量有限，而这个也没背占用
func (mp KeywordArguments) Kw(columnName string, value interface{}) KeywordArguments {
	mp[columnName] = value
	return mp //这样就是一个链式反应的
}

// Kws 把我的类型转换为gorm能够识别的类型
// 这块也比较奇怪，就是 gorm 在 UpdateColumns 的时候只能接收有限的类型
// 比如 map[string]interface{} 或者 struct 或者 *struct
// 其他类型，就会返回 gorm.ErrInvalidData 具体代码请看gorm的源码
// 因此当你最终设置完以后还是要使用这个函数转换，否则运行时还是会报错
// 还好吧，在没有用对的时候至少是报了错，改为其他类型似乎就不是报凑那么简单啦
// 我建议让gorm另写个update函数，主要用于更新那些实现某个接口（比如cvt2map）的数据类型，这样就好啦
func (mp KeywordArguments) Kws() map[string]interface{} {
	return mp
}

func (mp KeywordArguments) Map() map[string]interface{} {
	return mp //换个函数名，使用者喜欢用哪个就用哪个吧
}

func (mp KeywordArguments) AsMap() map[string]interface{} {
	return mp //换个函数名，使用者喜欢用哪个就用哪个吧
}

func (mp KeywordArguments) ToMap() map[string]interface{} {
	return mp //换个函数名，使用者喜欢用哪个就用哪个吧
}
