# gormcnm 根据 golang 定义的 models struct 字段调用 gorm 的增删改查

## 使用
假设你的 model 定义是这样的:
```
type Example struct {
	Name string `gorm:"primary_key;type:varchar(100);"`
	Type string `gorm:"column:type;"`
	Rank int    `gorm:"column:rank;"`
}
```
很明显的你会得到它们的列名是 `name` `type` `rank`，在使用gorm时，通常是这样查询的
```
err := db.Where("name=?", "abc").First(&res).Error
```
由于 `name` 是个硬编码的字段，因此我们期望的时在 `Example` 里这个字段是稳定的，随着业务变化也不会被修改的（比如删除或者换类型）
由此就会让我们形成开发习惯，即定义models的时候需要慎重，因为它是一切逻辑的基石。

但是“确保models不常修改”，这其实是一种奢望。[想要解决这种奢望的创作背景](./CREATION_IDEAS.zh.md) 这个文档只是创作背景和意图，不看也罢，接下来说明如何使用：

简单demo:

[简单demo](../demos/demo1x/main.go) | [简单demo](../demos/demo2x/main.go)

这些都是最简单的，但实际上也只是demo级别的

推荐你看这个项目里的demo: [自动生成 gormcnm 字段定义的工具 gormcngen](https://github.com/yyle88/gormcngen) 请在这个项目的README中找到demo代码，将会教会你如何更好的使用。
这个才是核心的，确实非常好用的。

首先假设你的模型是
```
type Example struct {
	Name string `gorm:"primary_key;type:varchar(100);"`
	Type string `gorm:"column:type;"`
	Rank int    `gorm:"column:rank;"`
}
```
这个模型的各个列字段定义就是这样的
```
const (
	columnName = gormcnm.ColumnName[string]("name")
	columnType = gormcnm.ColumnName[string]("type")
	columnRank = gormcnm.ColumnName[int]("rank")
)
```
这时候使用 gorm 查询就是这样的：
```
//SELECT * FROM `examples` WHERE name="abc" AND type="xyz" AND rank>100 AND rank<200 ORDER BY `examples`.`name` LIMIT 1
var res Example
if err := db.Where(columnName.Eq("abc")).
    Where(columnType.Eq("xyz")).
    Where(columnRank.Gt(100)).
    Where(columnRank.Lt(200)).
    First(&res).Error; err != nil {
    panic(errors.WithMessage(err, "wrong"))
}
fmt.Println(res)
```
这样查询就能避免代码中出现大量的列名字符串，确保在重构代码的时候不会漏改，或者字段类型设置错误。

比如 name 是 string 而 rank 是 int 的，这时候使用
```
db.Where("name=?", 1).First(&res).Error //就会报错，这是因为 name 是 string 的，查询结果肯定是不正确的
```
通常的假如使用
```
db.Where("name=?","abc).UpdateColumns(map[string]any{"rank":"xyz"}) //也会报错，因为 rank 是 int 而非 string 类型
```
由于查询时 db.Where 需要的是 `interface{}` 类型参数，而更新的时候传的是 `map[string]any` 类型参数
就使得在静态检查/编译阶段，都不能即时发现错误，只在业务运行时报错，这就不利于快速重构和发现问题。

而使用我的工具就能避免这个问题。

假如不想手写各个类名的常数列表，还可以使用配套的自动化生成工具，自动化生成自定义的所有 models 的列名。

调用工具将生成常量配置的代码。这是工具项目 [自动生成 gormcnm 字段定义的工具包 gormcngen](https://github.com/yyle88/gormcngen) 只看README就行。

自动化生成的列名定义是这样的（这里为了防止各个模型都有相同的字段，比如 id created_at updated_at deleted_at，就使用了个自定义的类将它们装起来）：
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
接下来使用
```
// SELECT * FROM `examples` WHERE name="abc" AND type="xyz" AND rank>100 AND rank<200 ORDER BY `examples`.`name` LIMIT 1
var res models.Example
var cls = res.Columns()
if err := db.Where(cls.Name.Eq("abc")).
    Where(cls.Type.Eq("xyz")).
    Where(cls.Rank.Gt(100)).
    Where(cls.Rank.Lt(200)).
    First(&res).Error; err != nil {
    panic(errors.WithMessage(err, "wrong")) //只是demo这样写写
}
fmt.Println(res)
```
其它查询逻辑都有封装，具体在开发中可慢慢体会如何使用。

## 安装
使用本项目
```
go get github.com/yyle88/gormcnm
```

```
import "github.com/yyle88/gormcnm"
```

## 优势
有诸多优势，有的写在了开发思路文档里 [开发思路](./CREATION_IDEAS.zh.md)
这里补充些显而易见的优势
1. 能够让你的代码重构变得安全，特别是重命名字段/修改字段类型/删除字段时，它能让你能在静态检查阶段发现问题。
2. 能够让你的重构变得便捷，比如当想修改 `Name` -> `Username` 的时候，就直接通过IDE的重命名(比如使用GOLAND的shift+F6快捷键)，把模型里的 `Name` 改为 `Username` 再把 `ExampleColumns . Name` 重命名，把 `Name: "name"` 修改为 `Username: "username"` 就行，这就能保证你不会改到别的表，非常便捷。
3. 能够提高你的编码速度，就比如 `db.Where(cls.Name.Eq("abc"))` 就会比 `db.Where("name=?", "abc")` 的编码效率高，而当字段名比较长的时候这种优势会很明显（毕竟有IDE代码自动提示，只要输`cls.`就能自动提示后面的）。
4. 能够让你的代码搜索更方便，比如想搜索db中 `name` 改变的原因，通过搜索模型中 `Name` 的引用，就能在 `db.Save` 或 `db.Create` 里面找到赋值的代码位置，再搜索 `cls.Name` 的引用，就能找到`db.Update` `db.UpdateColumn` `db.Updates` 里跟设置 `Name` 相关的代码位置。

Give me stars. Thank you!!!
