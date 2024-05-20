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

但是“确保models不常修改”，这其实是一种奢望。[想要解决这种奢望的创作背景](/internal/docs/CREATION_IDEAS.md)

Demo:
[简单Demo](/internal/demos/main/main.go)
