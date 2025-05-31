[![GitHub Workflow Status (branch)](https://img.shields.io/github/actions/workflow/status/yyle88/gormcnm/release.yml?branch=main&label=BUILD)](https://github.com/yyle88/gormcnm/actions/workflows/release.yml?query=branch%3Amain)
[![GoDoc](https://pkg.go.dev/badge/github.com/yyle88/gormcnm)](https://pkg.go.dev/github.com/yyle88/gormcnm)
[![Coverage Status](https://img.shields.io/coveralls/github/yyle88/gormcnm/master.svg)](https://coveralls.io/github/yyle88/gormcnm?branch=main)
![Supported Go Versions](https://img.shields.io/badge/Go-1.22%2C%201.23-lightgrey.svg)
[![GitHub Release](https://img.shields.io/github/release/yyle88/gormcnm.svg)](https://github.com/yyle88/gormcnm/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/yyle88/gormcnm)](https://goreportcard.com/report/github.com/yyle88/gormcnm)

# `gormcnm` - A Progressive, Type-Safe Approach to GORM Column Names Using Generics

## Overview

`gormcnm` is a cutting-edge **generic package** designed to transform how you use GORM in Go. By leveraging the full power of Go’s generics, it offers a **type-safe**, **efficient**, and **highly productive** way to reference database columns in your models. This eliminates the risks associated with hardcoded column names, enhancing **refactoring safety**, **maintainability**, and enabling **faster development** with fewer bugs.

`gormcnm` resembles `MyBatis Plus` in the Java ecosystem, which allows developers to dynamically retrieve column names using expressions like `Example::getName`. Similarly, `gormcnm` brings **type-safe** column referencing to Go. By using `gormcnm`, developers can easily perform database queries with **compile-time validation** and avoid hardcoding strings in queries.

`gormcnm` works like `SQLAlchemy` in the Python ecosystem, where developers can reference model attributes like `Example.name` for dynamic column access, enabling **type-safe** column referencing in Go. Similarly, `gormcnm` enables **type-safe** dynamic column references in Go, like `cls.Name.Eq("abc")`.

`gormcnm` keeps your column references consistent and type-safe, ensuring that your code is cleaner, more robust, and easier to maintain. With `gormcnm`, any changes in your model definitions are effortlessly reflected across your application without breaking queries, making it the ultimate **safety net** for your database interactions.

With `gormcnm`, you unlock the full potential of Go’s generics, achieving type-safe, refactor-proof, and developer-friendly database interactions effortlessly. No more brittle hardcoded strings or runtime surprises—`gormcnm` acts as your ultimate safety net, ensuring database queries remain consistent, scalable, and easy to refactor as your application evolves.

As a **progressive package**, `gormcnm` seamlessly integrates into your existing GORM workflow, allowing developers to adopt it incrementally and at their own pace. You can decide when and where to leverage its type-safe features while continuing to benefit from the simplicity and familiarity of GORM’s core API, ensuring that adopting `gormcnm` requires minimal learning effort and no disruption to your current projects.

## CHINESE README

[中文说明](README.zh.md)

## Installation

```bash
go get github.com/yyle88/gormcnm
```

## Features

- **Generics-based Type Safety**: Harness Go's generics to create type-safe column names, ensuring that column references are validated at compile-time.
- **Seamless Refactoring**: Changing model field names or types automatically updates all references, reducing errors and making refactoring effortless.
- **Progressive Adoption**: Designed to be integrated incrementally, allowing developers to adopt its features at their own pace without disrupting existing workflows.
- **`MyBatis Plus`-like Column References**: `gormcnm` enables **type-safe** dynamic column references in Go, like `cls.Name.Eq("abc")`.
- **`SQLAlchemy`-like Access**: Like `SQLAlchemy` in Python, `gormcnm` enables **type-safe** dynamic column references in Go, making queries cleaner and less error-prone.
- **Compile-time Validation**: Prevent runtime errors caused by incorrect column names or mismatched types with robust compile-time validation.
- **Improved Developer Experience**: Enjoy full IDE support with auto-completion, linting, and refactoring tools tailored for generics.
- **Minimized Human Error**: Eliminate typos and hardcoded magic strings with well-defined constants for column names.
- **Self-documenting Code**: Queries written with `gormcnm` are explicit and readable, making it easier to understand and maintain the codebase.
- **Reduced Boilerplate**: Automatically generate type-safe column definitions with the [gormcngen](https://github.com/yyle88/gormcngen) package, saving time and effort.

## Example Usage

Suppose you have a GORM model defined as follows:

```go
type Example struct {
    Name string `gorm:"primary_key;type:varchar(100);"`
    Type string `gorm:"column:type;"`
    Rank int    `gorm:"column:rank;"`
}
```

In **Java**, the `MyBatis Plus` tool can obtain the column name through `Example::getName`, assemble the query statement, fetch the result, and then use `result.getName()` to get the value of the field:

```java
@Autowired
private ExampleMapper exampleMapper;

public void test() {
    Example result = exampleMapper.selectOne(
        new LambdaQueryWrapper<Example>().eq(Example::getName, "abc")
    );

    if (result != null) {
        System.out.println(result.getName());
        System.out.println(result.getRank());
    }
}
```

In **Python**, the `SQLAlchemy` tool can obtain the column name through `Example.name`, assemble the query statement, fetch the result, and then use `result.name` to get the value of the field:

```python
def test():
    result = session.query(Example).filter(Example.name == "abc").first()

    if result:
        print(result.name)
        print(result.rank)
```

In Go, there's no equivalent to `Example::Name`; instead, `example.Name` directly retrieves the field value, requiring hard-coded queries.

### Traditional Query with Hardcoded Field Names

Typically, you would write a query with hardcoded column names like this:

```go
err := db.Where("name=?", "abc").First(&res).Error
```

While this works, hardcoding column names introduces risks during **refactoring** or **type changes**. A simple oversight could lead to runtime errors or type mismatches.

### Safe Query with `gormcnm`

With `gormcnm`, you can define type-safe columns like this:

```go
const (
    columnName = gormcnm.ColumnName[string]("name")
    columnType = gormcnm.ColumnName[string]("type")
    columnRank = gormcnm.ColumnName[int]("rank")
)
```

### Type-Safe Queries

This allows you to write **type-safe queries** with compile-time validation:

```go
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

This approach ensures that column types and names are consistent, reducing the risk of bugs.

## Demos

[Simple demo](./internal/demos/main/main.go)

## Advantages

- **Effortless Refactoring**: Column names and types are strongly typed with generics. Renaming fields or changing column types updates all references automatically.
- **Compile-time Type Safety**: Catch issues like type mismatches or incorrect column references during compilation, reducing debugging time.
- **Progressive Integration**: Start using `gormcnm` in specific parts of your codebase and expand its use as needed, ensuring a smooth and non-disruptive transition.
- **Cleaner Code**: Replace magic strings with constants, making your queries more readable, maintainable, and understandable.
- **Faster Development**: IDE features like auto-completion and refactoring tools speed up coding and reduce the chance of human errors.
- **`MyBatis Plus`-like Syntax**: `gormcnm` uses a type-safe column referencing pattern like `MyBatis Plus` in Java, enabling dynamic column access and safe queries.
- **`SQLAlchemy`-like Access**: Similar to `SQLAlchemy` in Python, `gormcnm` provides **type-safe** dynamic column references in Go.
- **Self-documenting Queries**: Type-safe column references make your queries clearer and more descriptive, aiding collaboration and future maintenance.
- **Reduced Debugging Effort**: Locating column-related bugs is easier since all references use well-defined constants.
- **Minimized Human Error**: Avoid typos and hardcoded values that lead to runtime errors by centralizing column definitions.
- **Seamless Model Synchronization**: Model changes are consistently reflected in your queries, ensuring your application always uses up-to-date column definitions.

## Example Usage

In addition to the type-safe queries previously shown, `gormcnm` also supports complex and flexible operations such as updating multiple columns with conditions. Here's a comprehensive example:

### Updating Columns with Conditions

```go
result := db.Model(&Example{}).Where(
    Qx(columnName.Eq("aaa")).
        AND(
            Qx(columnType.Eq("xxx")),
            Qx(columnRank.Eq(123)),
        ).Qx3(),
).UpdateColumns(columnRank.Kw(100).Kw(columnType.Kv("zzz")).AsMap())
require.NoError(t, result.Error)
require.Equal(t, int64(1), result.RowsAffected)
```

## Automatically Generated Column Definitions

To make things even more efficient, you can use the **[gormcngen](https://github.com/yyle88/gormcngen)** package to **automatically generate column definitions** for your models. The generated code might look like this:

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

From now on, you can retrieve the column name like `Example::getName` in `MyBatis Plus`, obtain the column name class object with `cls = res.Columns()`, and then use `cls.Name.Eq("abc")` for the query in GORM.

### Using Auto-Generated Column Definitions in Queries

Once the columns are generated, you can use them in your queries:

```go
var res models.Example
var cls = res.Columns()
if err := db.Where(cls.Name.Eq("abc")).
    Where(cls.Type.Eq("xyz")).
    Where(cls.Rank.Gt(100)).
    Where(cls.Rank.Lt(200)).
    First(&res).Error; err != nil {
    panic(errors.WithMessage(err, "wrong"))
}
fmt.Println(res)
```

### Retrieving Data with Conditions Using `gormcnm`

Here’s an example demonstrating how to use `gormcnm` for retrieving data based on complex conditions:

```go
var res models.Example
var cls = res.Columns()
require.NoError(t, db.Where(
    Qx(cls.Name.BetweenAND("aba", "abd")).
        AND(
            Qx(cls.Type.IsNotNULL()),
        ).Qx2(),
).First(&one).Error)
require.Equal(t, "abc", res.Name)
```

## Usage

Simply import the package:

```go
import "github.com/yyle88/gormcnm"
```

## Why Use `gormcnm`?

1. **Reliable Refactoring**: Confidently rename fields or change types knowing that all column references will remain consistent.
2. **Enhanced Safety**: Compile-time validation ensures no mismatches or incorrect references.
3. **Productivity Boost**: Type-safe column references and IDE auto-completion save development time and reduce errors.
4. **Progressive Adoption**: Designed for gradual integration, letting you start small and scale as needed without disrupting your workflow.
5. **Clean and Maintainable Code**: Queries are more readable and self-explanatory, improving collaboration and future modifications.
6. **Fewer Runtime Bugs**: Centralized column definitions reduce the chances of human error.
7. **Automatic Code Generation**: The [gormcngen](https://github.com/yyle88/gormcngen) package simplifies setup, automating column definition generation.

## Conclusion

`gormcnm` brings a **simple**, **type-safe**, and **productive** approach to working with GORM in Go. By removing hardcoded column names and leveraging Go's generics, it ensures that your database queries are safe, consistent, and easy to maintain. Whether you're building small applications or managing complex systems, `gormcnm` helps you write robust, refactor-friendly code, accelerating development and improving the quality of your work.

By adopting `gormcnm`, you can significantly enhance the quality of your GORM queries, reduce human error, and accelerate development—all while maintaining a clean, maintainable codebase.

---

## Design Ideas

[CREATION_IDEAS](internal/docs/CREATION_IDEAS.en.md) && [README OLD DOC](internal/docs/README_OLD_DOC.en.md)

---

## License

MIT License. See [LICENSE](LICENSE).

---

## Contributing

Contributions are welcome! To contribute:

1. Fork the repo on GitHub (using the webpage interface).
2. Clone the forked project (`git clone https://github.com/yourname/repo-name.git`).
3. Navigate to the cloned project (`cd repo-name`)
4. Create a feature branch (`git checkout -b feature/xxx`).
5. Stage changes (`git add .`)
6. Commit changes (`git commit -m "Add feature xxx"`).
7. Push to the branch (`git push origin feature/xxx`).
8. Open a pull request on GitHub (on the GitHub webpage).

Please ensure tests pass and include relevant documentation updates.

---

## Support

Welcome to contribute to this project by submitting pull requests and reporting issues.

If you find this package valuable, give me some stars on GitHub! Thank you!!!

**Thank you for your support!**

**Happy Coding with this package!** 🎉

Give me stars. Thank you!!!

---

## GitHub Stars

[![starring](https://starchart.cc/yyle88/gormcnm.svg?variant=adaptive)](https://starchart.cc/yyle88/gormcnm)
