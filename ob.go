package gormcnm

type ColumnOrderByAscDesc string

func (ob ColumnOrderByAscDesc) Ob(next ColumnOrderByAscDesc) ColumnOrderByAscDesc {
	return ob + " , " + next
}

// Ox 这块暂时也是没有什么办法的，就是假如不是特定的类型，就会被gorm的逻辑忽略
// 即使我在这里增加转换也不行，我觉得使用者还是会忘记的，只能说这里随缘吧
// 我自己用应该是没问题的
// 记得最后调用把这个转化为字符串再传给gorm哦
func (ob ColumnOrderByAscDesc) Ox() string {
	return string(ob)
}
