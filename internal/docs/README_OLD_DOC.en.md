# `gormcnm` - Using GORM with Model Struct Fields Defined in Go for CRUD Operations

## Usage

Assume your model definition looks like this:
```go
type Example struct {
    Name string `gorm:"primary_key;type:varchar(100);"`
    Type string `gorm:"column:type;"`
    Rank int    `gorm:"column:rank;"`
}
```
It's clear that the column names for this model are `name`, `type`, and `rank`. When using GORM, you would typically query like this:
```go
err := db.Where("name=?", "abc").First(&res).Error
```
Since `name` is a hardcoded field, we expect it to remain stable in the `Example` struct, and not be modified (e.g., deleted or type-changed) as the business evolves. This forms a development habit that requires us to be cautious when defining models because they are the foundation of all logic.

However, "ensuring models are not modified" is actually an unrealistic expectation. [The background behind this idea](./CREATION_IDEAS.en.md) provides the context and motivation for solving this problem, though you can skip reading that. Let's now look at how to use `gormcnm`.

### Simple Demo:

[Simple Demo](../demos/demo1x/main.go) | [Simple Demo](../demos/demo2x/main.go)

These are the simplest examples, but they are at a demo level.

We recommend looking at the demo in this project: [Automated Column Name Generation Tool for `gormcnm` - `gormcngen`](https://github.com/yyle88/gormcngen). Please refer to the README of this project to find the demo code, which will teach you how to use the tool more effectively. This is the core and truly useful part.

### Assume your model is like this:
```go
type Example struct {
    Name string `gorm:"primary_key;type:varchar(100);"`
    Type string `gorm:"column:type;"`
    Rank int    `gorm:"column:rank;"`
}
```

The column definitions for this model are:
```go
const (
    columnName = gormcnm.ColumnName[string]("name")
    columnType = gormcnm.ColumnName[string]("type")
    columnRank = gormcnm.ColumnName[int]("rank")
)
```

Now, querying with GORM would look like this:
```go
// SELECT * FROM `examples` WHERE name="abc" AND type="xyz" AND rank>100 AND rank<200 ORDER BY `examples`.`name` LIMIT 1
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
This query avoids hardcoding column names in your code, ensuring that when refactoring, you won't forget to update the column names or set the wrong field types.

For example, if `name` is a string and `rank` is an int, using:
```go
db.Where("name=?", 1).First(&res).Error // This will cause an error because `name` is a string, and the query will definitely fail.
```
Typically, if you use:
```go
db.Where("name=?", "abc").UpdateColumns(map[string]any{"rank":"xyz"}) // This will also cause an error because `rank` is an int, not a string.
```
Since `db.Where` requires an `interface{}` type argument and `UpdateColumns` uses a `map[string]any`, this mismatch won't be caught at compile time, leading to runtime errors. This makes it harder to refactor and detect issues quickly.

Using this tool can help prevent such problems.

### Automatically Generating Column Names

If you don’t want to manually write the constant list of column names, you can use the companion automation tool to generate column names for all your models.

The tool will generate constant configuration code for column names. This tool project is called [Automated Column Name Generation package for *gormcnm* - *gormcngen*](https://github.com/yyle88/gormcngen). Just look at the README for details.

The automatically generated column name definitions look like this (to avoid duplicate fields like `id`, `created_at`, `updated_at`, and `deleted_at`, we use a custom class to encapsulate them):
```go
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

You can then use it like this:
```go
// SELECT * FROM `examples` WHERE name="abc" AND type="xyz" AND rank>100 AND rank<200 ORDER BY `examples`.`name` LIMIT 1
var res models.Example
var cls = res.Columns()
if err := db.Where(cls.Name.Eq("abc")).
    Where(cls.Type.Eq("xyz")).
    Where(cls.Rank.Gt(100)).
    Where(cls.Rank.Lt(200)).
    First(&res).Error; err != nil {
    panic(errors.WithMessage(err, "wrong")) // Just for demo
}
fmt.Println(res)
```

Other query logic is encapsulated as well, and you can gradually get familiar with how to use it in development.

## Installation

To use this project, run:
```go
go get github.com/yyle88/gormcnm
```

Then import it into your code:
```go
import "github.com/yyle88/gormcnm"
```

## Advantages

There are many advantages, some of which are discussed in the development ideas document [Development Ideas](./CREATION_IDEAS.en.md). Here are some obvious benefits:

1. **Safe Refactoring**: It ensures safe refactoring, especially when renaming, modifying, or deleting fields. You can catch errors during static analysis.
2. **Convenient Refactoring**: For example, when changing `Name` to `Username`, you can use the IDE's rename feature (e.g., Shift+F6 in GoLand) to easily rename `Name` to `Username` in both the model and the `ExampleColumns.Name` constant. This ensures you won’t affect other tables, making refactoring more efficient.
3. **Faster Coding**: Writing `db.Where(cls.Name.Eq("abc"))` is more efficient than `db.Where("name=?", "abc")`. This advantage becomes more noticeable with longer field names (thanks to IDE code suggestions that help auto-complete after typing `cls.`).
4. **Easier Code Searching**: If you want to search for changes made to `name` in the database, you can search for references to `Name` in the model and find the code location in `db.Save` or `db.Create`. Searching for `cls.Name` will help you find relevant code locations in `db.Update`, `db.UpdateColumn`, or `db.Updates`.

Give me stars. Thank you!!!
