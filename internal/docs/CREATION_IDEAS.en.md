# Creative Ideas

## Problem
The problem encountered is that when frequently modifying models, the `gorm` logic of `db.Where("name=?", req.Name)` can throw errors at runtime if the `name` field is deleted, renamed, or its type is modified.

A minor issue is that the word `name` is prone to typos. Although the word itself is unlikely to be misspelled, what if there are many other fields? The most reliable solution often involves copying and pasting, but this can make development less smooth.

The main issue, however, is that the project may undergo refactoring or requirement changes at any time. The goal is to "ensure models don’t change frequently," but that’s actually a luxury.

In the early stages of a project, when the requirements are still being defined, frequent changes are common. Later, when there are 20 or more models in the same project, some of which may have a `name` column, deleting a `name` field in one model becomes very cumbersome. This reduces the flexibility of business modifications. For example, due to uncertainty about whether a field is being used, models or database tables can only add fields, not delete them. If the code is full of `Where("name=?", req.Name)`, even the most skilled programmers would be cautious when modifying it. Even if they focus on solving the problem, it's a waste of energy. Our energy is limited, and only by simplifying the basics can we delve deeper into more complex problems.

## Solution
At this point, you might have thought about isolating operations related to a model into a separate file, such as creating a `repositories` directory and a `repositories/example.go` file that only operates on the `Example` model.

This approach does solve many problems. When dealing with database read and write operations in business logic (`service`), you would always call `repositories.FindExampleXxx(x,y,z)` or `repositories.UpdateExampleXxx(x,y,z)`. If there are many models, you could even categorize operations, such as `repositories.ExampleRepo(db).FindXxx(x,y,z)` or `exampleRepo.UpdateXxx(x,y,z)`.

However, a significant drawback of this approach is that reading the code requires frequent navigation between files. Additionally, some operations only need to be invoked once in a specific scenario, and there’s no need to encapsulate them. But according to the principle mentioned above, they still need to be encapsulated into functions. Furthermore, for slightly more complex queries, the number of parameters increases, leading to function calls like `exampleRepo.UpdateXxx(x,y,z,u,v,w)`. This can be confusing, especially if the arguments are not passed in the correct order. The parameter names must be meaningful, but this can lead to excessively long function names.

Also, navigating between parameters and arguments becomes cumbersome.

You might then think of using another ORM like `entgo.io/ent`, which can generate code based on model configurations, providing operations like `func (ec *ExampleCreate) SetName(s string) *ExampleCreate` and `func (eu *ExampleUpdate) SetName(s string) *ExampleUpdate`. This approach ensures static checks during compilation after modifying the model.

This is indeed a useful package, but the problem is that I don’t like using it. There’s no particular reason for this—probably because I’ve been using `gorm` for so long that I’m not accustomed to using others. New tools always require learning. I only learned the basics of `entgo.io/ent`, and I believe `gorm` is the best.

Thus, I thought, can I create something based on `gorm`? Initially, I also considered generating code, like using configuration to generate code such as:
```
func (G *DemoGorm) Where名字(v string) *DemoGorm { return G.Where名字QS("= ?", v) }
func (G *DemoGorm) Where日期(v string) *DemoGorm { return G.Where日期QS("= ?", v) }
func (G *DemoGorm) Where数值(v float64) *DemoGorm { return G.Where数值QS("= ?", v) }
```
But I quickly rejected this solution because it generated hundreds of lines of code for just defining a model, and it still felt clunky to use. Excessive unused code made my project bloated, and the IDE became sluggish. The worst part was that by using my own encapsulated methods, I might forget how to properly use `gorm` or `entgo.io/ent`, which are standard tools—something that could hurt me in job interviews.

## Result
Therefore, I gave up on generating code myself. I’m also not good at using `entgo.io/ent`, and not all projects can use it. For instance, if I join a project that doesn't use `entgo.io/ent`, I’m still stuck with `gorm`.

However, things seem to have changed with the introduction of generics in Go. I suddenly thought that I should generate less code, or even no code at all, but instead provide a third-party package that adds generic capabilities to `gorm` just by importing it.

## Expectation
For instance, if my model is:
```
type Example struct {
	Name string `gorm:"primary_key;type:varchar(100);"`
	Type string `gorm:"column:type;"`
	Rank int    `gorm:"column:rank;"`
}
```
I don’t want to do:
```
var res Example
err := db.Where("name=?", req.Name).First(&res).Error
```
Instead, I want to do:
```
var res Example
var cls = res.Columns() // This is a hypothetical operation, not actually available, but I expect something like this to be supported at the Go language level, where the class automatically contains its ORM attribute information
err := db.Where(cls.Name +"=?", req.Name).First(&res).Error // Here, cls contains the metadata (or schema) of Example
```
Ideally, even further:
```
var res Example
var cls = res.Columns()
err := db.Where(cls.Name.Eq(req.Name)).First(&res).Error // Let’s truly unleash the potential! If cls.Name.Eq(req.Name) returns ("name=?", req.Name), then it can be directly passed to db.Where()
```
Having this expectation, I decided to try and implement it.

## Result
To implement the previous expectation, I need to treat "columns" as objects, enabling their generic behavior. First, I focus on imagining the `Eq` function:
```
type ColumnName[TYPE any] string

func (s ColumnName[TYPE]) Eq(x TYPE) (string, TYPE) {
	return string(s) + "=?", x
}
```
Suppose I define a variable:
```
var name = ColumnName[string]("name") // Now I can call name.Eq("abc") to get ("name=?", "abc")
```
Because it’s generically defined as a string, using `name.Eq(1)` would result in a compile-time error.

Next, other operations like `=`, `>`, `<`, `>=`, `<=`, `!=`, `IN`, `BETWEEN...AND...` follow naturally:
```
func (s ColumnName[TYPE]) Gt(x TYPE) (string, TYPE) {
	return string(s) + ">?", x
}

func (s ColumnName[TYPE]) Lt(x TYPE) (string, TYPE) {
	return string(s) + "<?", x
}
```
These operations are simple enough to be quickly implemented, so I won’t list them all here.

## Self-test
Let’s say I provide this package. When used, you can directly define columns as variables and start using them:
```
const (
	columnName = gormcnm.ColumnName[string]("name")
	columnType = gormcnm.ColumnName[string]("type")
	columnRank = gormcnm.ColumnName[int]("rank")
)
```
Now you can query like this:
```
var res Example

err := db.Where(columnName.Eq("abc")).
    Where(columnType.Eq("xyz")).
    Where(columnRank.Gt(100)).
    Where(columnRank.Lt(200)).
    First(&res).Error

// SQL IS: SELECT * FROM `examples` WHERE name="abc" AND type="xyz" AND rank>100 AND rank<200 ORDER BY `examples`.`name` LIMIT 1
```
The result is perfect.

## Advantages
I believe this tool has several advantages:
1. As seen in the self-test code above, you only need to `go get github.com/yyle88/gormcnm`, then define constants/variables to enjoy the convenience this package brings. You can define one or two columns or more if needed.
2. It doesn’t change `gorm`’s usage or logic, so it doesn’t affect the readability of your code.
3. Since it follows `gorm`'s query logic, it doesn’t hinder you from learning this powerful tool. This is crucial—tools should not make developers lazy, and using non-standard tools can harm interview performance.
4. It helps you become more proficient in using `gorm`.
5. This tool is lightweight. You can choose to use or not use it as you wish. If you haven’t defined the `rank` field constant, you can simply use the original `Where("rank>?", 100)`. This flexibility is very valuable.
6. Since you can choose whether to use this tool, introducing it into an existing `gorm` project is minimal cost and doesn’t break anything.

## Potential of `gormcnm`

To use this tool, you need to run `go get github.com/yyle88/gormcnm` and `import "github.com/yyle88/gormcnm"`, then define the column constants:

```go
const (
    columnName = gormcnm.ColumnName[string]("name")
    columnType = gormcnm.ColumnName[string]("type")
    columnRank = gormcnm.ColumnName[int]("rank")
)
```

However, it's obvious that by analyzing the `models` definitions, we can determine which columns exist. In other words, these field definitions **could** potentially be automatically generated through `generate` tools. (This has already been implemented.)

Moreover, it's clear that these constants don't have a namespace concept.

For example, what if in another table the `rank` field is no longer an `int`, but instead a `string` with values like "S", "A", "B", "C"? If we define a constant again, it would lead to confusion.

Therefore, you should do the following:

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

Then, you can use it like this:

```go
var res models.Example
var cls = res.Columns()
err := db.Where(cls.Name.Eq("abc")).
    Where(cls.Type.Eq("xyz")).
    Where(cls.Rank.Gt(100)).
    Where(cls.Rank.Lt(200)).
    First(&res).Error
```

As you can see, the code generated by my tool is exactly like this.

```bash
go get github.com/yyle88/gormcngen
```

For usage instructions of this generation tool, please refer to the corresponding project's documentation: [Automatic gormcnm field definition generator pkg gormcngen](https://github.com/yyle88/gormcngen).

## Some others

This project is actually a tribute to `gorm`. However, `gormcnm` might not sound very elegant. If you don't like the name, you can check out [gormrepo](https://github.com/yyle88/gormrepo).

I occasionally think that the name `gormcls` might be better, especially since when using a Python ORM in the past, I often used `a.cls` for operations.

However, it’s obvious that this project deals with operations related to column names. Therefore, shortening `gorm column name` to `gormcnm` is just perfect—it’s both loud and elegant, truly fitting for the task.

## Conclusion

This creative background and usage method introduction is now complete.

Give me stars. Thank you!!!
