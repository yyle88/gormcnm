package gormcnm

type OrderByBottle string //我还是更喜欢用 Bottle 这个单词，以维持名称在材料学方面的重心平衡

func (ob OrderByBottle) Ob(next OrderByBottle) OrderByBottle {
	return ob + " , " + next
}

// Ox 这块暂时也是没有什么办法的，就是假如不是特定的类型，就会被gorm的逻辑忽略
// 即使我在这里增加转换也不行，我觉得使用者还是会忘记的，只能说这里随缘吧，因为除非使用有限的几种，任何包一层的管理器，都不能被gorm的Order逻辑识别
// 目前暂无优雅的解决方案
// 我自己用应该是没问题的
// 记得最后调用把这个转化为字符串再传给gorm的Order函数
func (ob OrderByBottle) Ox() string {
	return string(ob)
}
