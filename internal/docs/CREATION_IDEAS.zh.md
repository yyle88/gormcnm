# 创作背景-我想过的其它方案

## 问题
遇到的问题是，在频繁修改 models 的时候，gorm的 `db.Where("name=?", req.Name)` 的逻辑有可能会因为 `name` 被删除/改名/修改类型，而在运行时报错。

当然还有个特别小的问题是，单词`name`容易拼错，当然这个单词当然不会错，但是假如有很多其它字段呢，因此通常都是拷贝粘贴最靠谱，但这也会让开发变得不太顺滑。

当然主要的问题，还是在于项目随时可能重构，或者修改需求等，想要“确保models不常修改”，这其实是一种奢望。

特别是在项目创建初期，需求刚确定，就会很频繁的改动。
而在业务的后期，将会面临20个甚至更多的models在同一个项目里的情况，假设有三五个表都有 `name` 这个列，假设你需要删除某个模型的 `name`，就会非常麻烦，这样业务修改的灵活性将逐渐丧失，比如，因为不确定是否有用，因此模型/数据表/schema就只能增加字段，而不能删除字段。
面对代码里到处都有的 `Where("name=?", req.Name)` 再牛逼的程序员都会瑟瑟发抖/慎重对待的的吧。就算能集中精神解决这些问题，也是浪费精力的啊，我们人的精力都是有限的，只有让基础的事情简单化，这样才有可能抵达更深的深度，否则各种疥癣之疾都能把人折腾的头晕脑胀的，也就没法向更深的地方想啦。

## 方案
这个时候聪明的你想到了可以把跟某个模型相关的操作，单独写到一个文件里，比如创建 `repositories` 目录和 `repositories/example.go` 代码文件，这里面只操作 `Example` 模型。
这样做确实是能解决大部分问题，在业务代码 `service` 中涉及到读写 db 的时候，总是要调用 `repositories.FindExampleXxx(x,y,z)` 或者 `repositories.UpdateExampleXxx(x,y,z)` 的操作。
当然假如模型特别多的时候还可以按模型分类操作，比如 `repositories.ExampleRepo(db).FindXxx(x,y,z)` 或者 `exampleRepo.UpdateXxx(x,y,z)` 这样分类写操作。
但这样有个很大的弊端就是，读代码时总是要跳转的。
其次是，有些操作其实就是只需要在特定的场景下调用一次，完全没必要封装的，但按照上述的原则/约定，依然要封装函数。
还有就是，稍微复杂的查询其参数也不止两三个，面对`exampleRepo.UpdateXxx(x,y,z,u,v,w)` 这样的调用，会不会心里一惊，或者假如不小心写成`(x,y,z,w,v,u)`是不是也看不出来，毕竟函数内部逻辑用的形参和外部调用传的实参要对应位置，这本身就是个易错的事情。这就需要把参数名定的有意义些，就会让函数参数名特别长。
当然基于读代码要跳转的前提，在形参和实参间反复比对，查看传递路线也是有点费劲的。

因此这时候聪明的你想到我们可以用其它orm，比如 `go get entgo.io/ent` 能根据你的模型配置来生成代码，提供`func (ec *ExampleCreate) SetName(s string) *ExampleCreate` 和 `func (eu *ExampleUpdate) SetName(s string) *ExampleUpdate` 的操作，很明显，假如你需要改模型，改完生成代码后就能在静态检查阶段确保你的使用完全正确。
这个确实是个非常好用的包。
但问题是我不太喜欢用，没有特别的理由，估计是用 `gorm` 用的太久了不太习惯用别的，新东西总是要学习的，我学`go get entgo.io/ent`也只是学了个皮毛，而且我认为 `gorm` 是最强的。

因此我就想，我能不能基于 `gorm` 做个东西呢。 刚开始我也是用的生成代码的路线，比如利用配置得到这样的代码：
```
func (G *DemoGorm) Where名字(v string) *DemoGorm { return G.Where名字QS("= ?", v) }
func (G *DemoGorm) Where日期(v string) *DemoGorm { return G.Where日期QS("= ?", v) }
func (G *DemoGorm) Where数值(v float64) *DemoGorm { return G.Where数值QS("= ?", v) }
```
但我很快就否了自己的这个方案，因为实在是太垃圾，就是随便定义个模型，就要生成数百行代码，而用的时候还是很蹩脚的，而且过多的无用代码也使得我的项目变得很臃肿，IDE也很卡顿，当然最可怕的是，自己封装的东西越用人就越傻，将来找工作面试啥的，只会用自己的，而不能熟练掌握 `gorm` 或者 `entgo.io/ent` 这些标准的或者经典的工具，就糟糕啦。

## 结果
因此我放弃了自己生成代码的方案，而我也不擅长使用 `entgo.io/ent` 这个工具，而且也不是所有的项目都能用 `entgo.io/ent`  的，比如当我新进一个项目，而项目本身没有使用 `entgo.io/ent` 的时候。

因此我还是得用 `gorm` 的，但是随着 golang 推出泛型以后事情似乎是有所转机，我突然觉得，我应该尽量少的生成代码，甚至是不生成代码，而是提供个三方包，只要 import 就能赋予 gorm 泛型的能力。

## 期望
假如我的模型是
```
type Example struct {
	Name string `gorm:"primary_key;type:varchar(100);"`
	Type string `gorm:"column:type;"`
	Rank int    `gorm:"column:rank;"`
}
```
我不希望这样做
```
var res Example
err := db.Where("name=?", req.Name).First(&res).Error
```
我希望能这样做
```
var res Example
var cls = res.Columns() //这是幻想出来的操作，实际没有这个东西，但我期望有这个东西，最好是能从golang语言底层就支持，让class自动蕴含cls的orm属性信息
err := db.Where(cls.Name +"=?", req.Name).First(&res).Error //注意，这里的cls值的就是 Example 的元数据(也可以说是schema或者字段定义存储类型)
```
当然更进一步的
```
var res Example
var cls = res.Columns()
err := db.Where(cls.Name.Eq(req.Name)).First(&res).Error //让我们彻底放飞自我吧。假如 cls.Name.Eq(req.Name) 就是返回两个值 ("name=?", req.Name)，则它的返回值就，恰好能被 db.Where() 接收，这样 Eq 函数本身是很容易实现的。
```
既然有这个期望
就试着实现它吧

## 结果
根据前面的期望，因此我需要以“列”为对象，让它具备泛型的效果，首先瞄准 `Eq` 这个函数去想象
```
type ColumnName[TYPE any] string

func (s ColumnName[TYPE]) Eq(x TYPE) (string, TYPE) {
	return string(s) + "=?", x
}
```
假设定义个变量
```
var name = ColumnName[string]("name") //岂不是说就能调用 name.Eq("abc") 返回 ("name=?", "abc") 两个值啦
```
而且因为泛型定义它是sting的，想使用 `name.Eq(1)` 就会在静态检查阶段/编译阶段报错啦

接下来其它的操作不也就顺理成章了嘛，"=" ">" "<" ">=" "<=" "!=" "IN" "BETWEEN...AND..." 这些操作简直是手到擒来啊。常用的 "AND" "OR" "NOT" "ORDER" 这些也都是有实现的
```
func (s ColumnName[TYPE]) Gt(x TYPE) (string, TYPE) {
	return string(s) + ">?", x
}

func (s ColumnName[TYPE]) Lt(x TYPE) (string, TYPE) {
	return string(s) + "<?", x
}
```
额好吧实在是太简单了甚至稍微有点弱智的感觉，就不一个个列举。

## 自测
假设我提供了这个包，在使用时只要定义列的变量，能直接使用啦。
```
const (
	columnName = gormcnm.ColumnName[string]("name")
	columnType = gormcnm.ColumnName[string]("type")
	columnRank = gormcnm.ColumnName[int]("rank")
)
```
就可以进行各种查询啦。
```
var res Example

err := db.Where(columnName.Eq("abc")).
    Where(columnType.Eq("xyz")).
    Where(columnRank.Gt(100)).
    Where(columnRank.Lt(200)).
    First(&res).Error

//SQL IS: SELECT * FROM `examples` WHERE name="abc" AND type="xyz" AND rank>100 AND rank<200 ORDER BY `examples`.`name` LIMIT 1
```
结果相当完美。

## 优势
我认为这个工具它有诸多的优势： 
1. 就像上面的自测代码那样，你只需要 `go get github.com/yyle88/gormcnm` 再定义 const/vars 就能享受到这个包带来的便利，只定义一两个列也是可以的。
2. 基本能够看到，这套基于`gorm`的东西，没有改变 `gorm` 的使用方法和逻辑，不仅让你的代码具备可读性。
3. 由于沿用 `gorm` 查询逻辑，它也不会影响你学习 `gorm` 这个非常牛逼的东西。这个我认为非常重要，就是工具不能让开发者变傻，假如让开发者变傻不擅长使用通用的东西，就不好啦，就有可能让开发者在面试时吃亏，我自己也是开发者我当然不希望有这样的负面作用。
4. 使用这个东西，还会让你更擅长使用 `gorm`。
5. 这个工具非常轻量级，就像上面的自测的代码，在代码里可以随时选择使用/不使用这个工具。
假如你没有定义 rank 字段的列常量，你就直接用原来的 Where("rank>?", 100) 就行。这个有啥优势呢，这个优势很大很大，它就能让我的工具不需要做得很完整，毕竟大家都是很忙的，有啥没实现的功能，就直接用原来的 `gorm` 解决就行，这样我也就没必要把全部东西都做出来啦（毕竟免费做东西造个轮子就行，没必要造车的）。
6. 由于可以随时用或者不用这个工具，把这个工具引入到现用使用 `gorm` 项目里的代价就很小，而且不会破坏任何东西。

## 潜力
使用这个工具你需要 `go get github.com/yyle88/gormcnm` 和 `import "github.com/yyle88/gormcnm"` 再定义列的常量:
```
const (
	columnName = gormcnm.ColumnName[string]("name")
	columnType = gormcnm.ColumnName[string]("type")
	columnRank = gormcnm.ColumnName[int]("rank")
)
```
但是很明显的，根据 models 的定义能通过分析知道有哪些列，也就是说这些字段定义都是【有可能】通过 `generate` 而手段自动生成的。【实际已经实现啦】

而且，很明显的，这几个常量没有 namespace 的概念。
假设在其它表里 `rank` 这个字段不再是 `int` 而是 `string` "S" "A" "B" "C" 呢，再定义常量，岂不是容易混淆啦。

因此你要这样做
```
type ExampleColumns struct {
	Name gormcnm.ColumnName[string]
	Type gormcnm.ColumnName[string]
	Rank gormcnm.ColumnName[int]
}

func (*Example) Columns() *ExampleColumns {
	return &ExampleColumns{
		Name: "name",
		Type: "type",
		Rank: "rank",
	}
}
```
因此你要这样用
```
var res models.Example
var cls = res.Columns()
err := db.Where(cls.Name.Eq("abc")).
    Where(cls.Type.Eq("xyz")).
    Where(cls.Rank.Gt(100)).
    Where(cls.Rank.Lt(200)).
    First(&res).Error
```
因此很明显的，我写的生成代码的工具自动生成出来的代码，就是这样的。
```
go get github.com/yyle88/gormcngen
```
至于这个生成工具的用法就请移步到对应的项目里查看他的文档吧。[自动生成 gormcnm 字段定义的工具 gormcngen](https://github.com/yyle88/gormcngen)

## 其它
这个项目其实是致敬 `gorm` 的，但 `gormcnm` 或许有些不雅观，假如觉得不适也可以看这里 [gormcls](https://github.com/yyle88/gormcls)
我偶尔也会认为这个名字 `gormcls` 其实更好些，特别是以前用python某个orm的时候就是使用 `a.cls` 操作的。
但很显然这个项目解决的是“跟列名字段名有关的操作”，因此 `gorm column name` 缩写为 `gormcnm` 简直不能再合适啦，响亮而文雅，齐得隆东强。

## 完结
这个创作背景和使用方法介绍，完。
