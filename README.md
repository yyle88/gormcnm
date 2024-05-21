# gormcnm 根据 golang 定义的 models struct 字段调用 gorm 的增删改查

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

但是“确保models不常修改”，这其实是一种奢望。[想要解决这种奢望的创作背景](/internal/docs/CREATION_IDEAS.md) 这个文档只是创作背景和意图，不看也罢，接下来说明如何使用：

简单demo:

[简单demo](/internal/demos/main/main.go)

[测试case](/cname_test.go)

[测试case](/qx_test.go)

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
调用工具将生成代码 具体工具代码如何使用请直接看这里 [自动生成 gormcnm 字段定义的工具 gormcngen](https://github.com/yyle88/gormcngen)
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
{ // SELECT * FROM `examples` WHERE name="abc" AND type="xyz" AND rank>100 AND rank<200 ORDER BY `examples`.`name` LIMIT 1
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
}
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

Give me stars. Thank you!!!
