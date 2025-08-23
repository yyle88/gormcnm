# `gormcnm` - 基于泛型实现的一种渐进式、类型安全的 GORM 列名工具包

## 概述

`gormcnm` 是一个前沿的 **泛型包**，旨在彻底改变您在 Go 中使用 GORM 的方式。通过充分利用 Go 泛型的强大功能，它提供了一种 **类型安全**、**高效** 和 **高度生产力** 的方式来引用模型中的数据库列。能消除硬编码列名带来的风险，增强 **重构安全性**、**可维护性**，通过减少错误提高您的开发效率。

`gormcnm` 类似于 Java 生态中的 `MyBatis Plus`，允许开发者使用像 `Example::getName` 这样的表达式动态检索列名。类似地，`gormcnm` 为 Go 带来 **类型安全** 的列引用。通过使用 `gormcnm`，开发者可以轻松执行数据库查询，具备 **编译时验证**，并避免在查询中硬编码字符串。

`gormcnm` 类似于 Python 生态系统中的 `SQLAlchemy`，开发者可以像引用 `Example.name` 这样的模型属性来动态访问列名，实现 **类型安全** 的列引用。`gormcnm` 也在 Go 中提供 **类型安全** 的动态列引用，能像 `cls.Name.Eq("abc")` 这样使用。

`gormcnm` 保证您的列引用一致且类型安全，确保代码更加简洁、稳健且易于维护。使用 `gormcnm` 后，模型定义中的任何更改都会轻松地反映在应用程序中而不破坏查询，使其成为您与数据库交互的 **安全网**。

通过 `gormcnm`，您可以充分发挥 Go 泛型的潜力，轻松实现类型安全、无重构问题且开发者友好的数据库交互。告别脆弱的硬编码字符串或运行时错误——`gormcnm` 作为您的最终安全网，确保数据库查询保持一致、可扩展，并且在应用程序发展时易于重构。

作为一个 **渐进式工具包**，`gormcnm` 无缝集成到您现有的 GORM 工作流中，允许开发者根据自己的节奏逐步采用它。您可以决定何时以及在哪些地方利用它的类型安全特性，同时继续享受 GORM 核心 API 的简单性和熟悉性，确保采用 `gormcnm` 时，学习成本最低且不干扰现有项目。

## 英文文档

[ENGLISH README](README.md)

## 安装

```bash
go get github.com/yyle88/gormcnm
```

## 快速开始

这是一个完整的可运行示例：

```go
package main

import (
	"fmt"

	"github.com/yyle88/done"
	"github.com/yyle88/gormcnm"
	"github.com/yyle88/neatjson/neatjsons"
	"github.com/yyle88/rese"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type User struct {
	Username string `gorm:"primary_key;type:varchar(100);"`
	Nickname string `gorm:"column:nickname;"`
	Age      int    `gorm:"column:age;"`
}

const (
	columnUsername = gormcnm.ColumnName[string]("username")
	columnNickname = gormcnm.ColumnName[string]("nickname")
	columnAge      = gormcnm.ColumnName[int]("age")
)

func main() {
	db := rese.P1(gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	}))
	defer rese.F0(rese.P1(db.DB()).Close)

	done.Done(db.AutoMigrate(&User{}))
	done.Done(db.Create(&User{Username: "alice", Nickname: "Alice", Age: 17}).Error)

	var user User
	done.Done(db.Where(columnUsername.Eq("alice")).First(&user).Error)
	fmt.Println(neatjsons.S(user))

	done.Done(db.Model(&user).Update(columnNickname.Kv("SuperAlice")).Error)
	done.Done(db.Where(columnUsername.Eq("alice")).First(&user).Error)
	fmt.Println(neatjsons.S(user))

	done.Done(db.Model(&user).Update(columnAge.KeAdd(1)).Error)
	done.Done(db.Where(columnUsername.Eq("alice")).First(&user).Error)
	fmt.Println(neatjsons.S(user))
}
```

## 特性

- **基于泛型的类型安全**：利用 Go 的泛型创建类型安全的列名，确保列引用在编译时进行验证。
- **无缝重构**：更改模型字段名或类型时，自动更新所有引用，减少错误，使重构变得轻松。
- **渐进式采用**：设计上可逐步集成，允许开发者根据需要逐步采用其特性，不会干扰现有工作流程。
- **像 `MyBatis Plus` 的列引用**：`gormcnm` 提供 `cls.Name.Eq("abc")` 的功能。
- **类似 `SQLAlchemy` 的操作**: 在 Go 中提供 **类型安全** 的动态列引用，使查询更简洁。
- **编译时验证**：通过强大的编译时验证，避免因列名错误或类型不匹配而导致的运行时错误。
- **改进的开发者体验**：享受 IDE 完整支持，包括自动补全、代码检查和针对泛型的重构工具。
- **减少人为错误**：使用定义良好的常量替代硬编码的魔法字符串，避免打字错误。
- **自文档化代码**：使用 `gormcnm` 编写的查询明确且易读，方便理解和维护代码。
- **减少模板代码**：使用 [gormcngen](https://github.com/yyle88/gormcngen) 包自动生成类型安全的列定义，节省时间和精力。

## 示例使用

假设您有一个如下定义的 GORM 模型：

```go
type Example struct {
    Name string `gorm:"primary_key;type:varchar(100);"`
    Type string `gorm:"column:type;"`
    Rank int    `gorm:"column:rank;"`
}
```

在 **Java** 中，`MyBatis Plus` 工具可以通过 `Example::getName` 获得列名，组装查询语句，查询出结果后，再使用 `result.getName()` 获取字段的值：

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

在 **Python** 中，`SQLAlchemy` 工具可以通过 `Example.name` 获得列名，组装查询语句，查询出结果后，再使用 `result.name` 获取字段的值：

```python
def test():
    result = session.query(Example).filter(Example.name == "abc").first()

    if result:
        print(result.name)
        print(result.rank)
```

但是在 Go 中目前还没有 `Example::Name` 的功能，而 example.Name 则是取的字段的值，因此在Go中目前还总是需要使用 硬编码 查询。

### 使用硬编码字段名的传统查询

通常，您会写一个查询，使用硬编码列名，如下所示：

```go
err := db.Where("name=?", "abc").First(&res).Error
```

虽然这样有效，但硬编码列名会在 **重构** 或 **类型变化** 时引入风险。稍有疏忽可能导致运行时错误或类型不匹配。

### 使用 `gormcnm` 的安全查询

使用 `gormcnm`，您可以这样定义类型安全的列：

```go
const (
    columnName = gormcnm.ColumnName[string]("name")
    columnType = gormcnm.ColumnName[string]("type")
    columnRank = gormcnm.ColumnName[int]("rank")
)
```

### 类型安全查询

这样，您就能编写 **类型安全的查询**，并在编译时验证：

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

这种方式确保列名和列类型的一致性，减少了错误的风险。

## 样例

[简单样例](internal/demos/demo1x/main.go) | [简单样例](internal/demos/demo2x/main.go)

## 优势

- **轻松重构**：列名和类型通过泛型进行强类型定义。重命名字段或更改列类型时，所有引用会自动更新。
- **编译时类型安全**：在编译时捕获类型不匹配或列引用错误，减少调试时间。
- **渐进式集成**：可以在代码库中的特定部分开始使用 `gormcnm`，随着需求的增加逐步扩展使用，确保平滑且无缝的过渡。
- **更简洁的代码**：用常量替代魔法字符串，使查询更加可读、可维护和易于理解。
- **更快的开发速度**：IDE 特性如自动补全和重构工具加快了编程速度，减少了人为错误的机会。
- **像 `MyBatis Plus` 的列引用**：`gormcnm` 使用类似于 Java 中 `MyBatis Plus` 的列引用功能。
- **类似 `SQLAlchemy` 的功能**: `gormcnm` 在 Go 中提供 **类型安全** 的动态列引用。
- **自文档化的查询**：类型安全的列引用使您的查询更加清晰和具有描述性，有助于协作和未来的维护。
- **减少调试工作**：由于所有引用都使用了定义良好的常量，定位与列相关的错误变得更容易。
- **最小化人为错误**：通过集中定义列，避免打字错误和硬编码值导致的运行时错误。
- **无缝模型同步**：模型的变化会一致地反映在您的查询中，确保您的应用始终使用最新的列定义。

## 示例使用

除了前面展示的类型安全查询，`gormcnm` 还支持复杂和灵活的操作，例如带条件的多个列更新。下面是一个全面的示例：

### 使用条件更新列

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

## 自动生成列定义

更高效的方法是，您可以使用 **[gormcngen](https://github.com/yyle88/gormcngen)** 包来 **自动生成列定义**。生成的代码可能如下所示：

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

现在，您可以像 `MyBatis Plus` 中的 `Example::getName` 一样获取列名，通过 `cls = res.Columns()` 获取列名类的对象，然后使用 `cls.Name.Eq("abc")` 进行查询。

### 使用生成的列进行查询

当列定义生成完毕，您可以在查询中使用它们：

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

### 使用 `gormcnm` 获取带条件的数据

接下来是一个使用 `gormcnm` 获取数据并应用复杂条件的示例：

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

## 使用方法

只需导入该包：

```go
import "github.com/yyle88/gormcnm"
```

## 推荐使用 `gormcnm`

1. **让重构更可靠**：放心地重命名字段或更改类型，所有列引用将保持一致。
2. **增强的安全性**：编译时验证确保没有不匹配或错误的引用。
3. **提高生产力**：类型安全的列引用和 IDE 自动补全节省开发时间，减少错误。
4. **渐进式采用**：设计上支持逐步集成，让您从小处开始，逐渐扩展使用，而不会干扰您的工作流程。
5. **简洁且易维护的代码**：查询更加清晰、可读、自解释，促进协作和未来修改。
6. **减少运行时错误**：集中定义列，减少人为错误的可能性。
7. **自动代码生成**：[gormcngen](https://github.com/yyle88/gormcngen) 包简化了设置，自动生成列定义。

## 结论

`gormcnm` 提供了一种 **简单**、**类型安全** 和 **高效** 的方式来使用 Go 中的 GORM。通过消除硬编码的列名并利用 Go 的泛型，它确保了数据库查询的安全性、一致性和易于维护。无论是构建小型应用程序还是管理复杂的系统，`gormcnm` 都能帮助您编写健壮、易于重构的代码，从而加速开发并提高工作质量。

通过采用 `gormcnm`，您可以显著提升 GORM 查询的质量，减少人为错误，并加速开发，同时保持清晰、可维护的代码库。

---

## 设计思路

[设计思路](internal/docs/CREATION_IDEAS.zh.md) && [旧版说明](internal/docs/README_OLD_DOC.zh.md)

---

## 许可

项目采用 MIT 许可证，详情请参阅 [LICENSE](LICENSE)。

---

## 贡献新代码

非常欢迎贡献代码！贡献流程：

1. 在 GitHub 上 Fork 仓库 （通过网页界面操作）。
2. 克隆Forked项目 (`git clone https://github.com/yourname/repo-name.git`)。
3. 在克隆的项目里 (`cd repo-name`)
4. 创建功能分支（`git checkout -b feature/xxx`）。
5. 添加代码 (`git add .`)。
6. 提交更改（`git commit -m "添加功能 xxx"`）。
7. 推送分支（`git push origin feature/xxx`）。
8. 发起 Pull Request （通过网页界面操作）。

请确保测试通过并更新相关文档。

---

## 贡献与支持

欢迎通过提交 pull request 或报告问题来贡献此项目。

如果你觉得这个包对你有帮助，请在 GitHub 上给个 ⭐，感谢支持！！！

**感谢你的支持！**

**祝编程愉快！** 🎉

Give me stars. Thank you!!!
