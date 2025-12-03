[![GitHub Workflow Status (branch)](https://img.shields.io/github/actions/workflow/status/yyle88/gormcnm/release.yml?branch=main&label=BUILD)](https://github.com/yyle88/gormcnm/actions/workflows/release.yml?query=branch%3Amain)
[![GoDoc](https://pkg.go.dev/badge/github.com/yyle88/gormcnm)](https://pkg.go.dev/github.com/yyle88/gormcnm)
[![Coverage Status](https://img.shields.io/coveralls/github/yyle88/gormcnm/main.svg)](https://coveralls.io/github/yyle88/gormcnm?branch=main)
[![Supported Go Versions](https://img.shields.io/badge/Go-1.22%2C%201.23%2C%201.24%2C%201.25-lightgrey.svg)](https://github.com/yyle88/gormcnm)
[![GitHub Release](https://img.shields.io/github/release/yyle88/gormcnm.svg)](https://github.com/yyle88/gormcnm/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/yyle88/gormcnm)](https://goreportcard.com/report/github.com/yyle88/gormcnm)

# GORMCNM

**gormcnm** - 使用类型安全的列名和编译时验证消除 GORM 操作中的硬编码字符串。

---

## 生态系统

![GORM Type-Safe Ecosystem](assets/gormcnm-ecosystem.svg)

---

<!-- TEMPLATE (ZH) BEGIN: LANGUAGE NAVIGATION -->

## 英文文档

[ENGLISH README](README.md)
<!-- TEMPLATE (ZH) END: LANGUAGE NAVIGATION -->

---

## 语言生态系统对比

| 语言       | ORM          | 类型安全列        | 示例                                    |
|------------|--------------|-------------------|----------------------------------------|
| **Java**   | MyBatis Plus | `Example::getName` | `wrapper.eq(Example::getName, "alice")` |
| **Python** | SQLAlchemy   | `Example.name`     | `query.filter(Example.name == "alice")` |
| **Go**     | **GORMCNM**  | `cls.Name.Eq()`    | `db.Where(cls.Name.Eq("alice"))`        |

---

## 主要特性

- 🎯 **核心价值**：使用类型安全操作避免硬编码列名
- 🎯 **类型安全列操作**：泛型 `ColumnName[T]` 类型，编译时验证
- ⚡ **零运行时开销**：类型检查在编译时完成
- 🔄 **重构安全查询**：IDE 自动补全和自动重构支持
- 🌍 **丰富查询操作**：全面的比较、范围、模式和聚合操作
- 📋 **生态系统基础**：支持代码生成和仓储模式工具

---

## 安装

```bash
go get github.com/yyle88/gormcnm
```

---

## 🔥 快速开始示例

```go
package main

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/yyle88/done"
	"github.com/yyle88/gormcnm"
	"github.com/yyle88/neatjson/neatjsons"
	"github.com/yyle88/rese"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Account struct {
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
	dsn := fmt.Sprintf("file:db-%s?mode=memory&cache=shared", uuid.New().String())
	db := rese.P1(gorm.Open(sqlite.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	}))
	defer rese.F0(rese.P1(db.DB()).Close)

	//CREATE TABLE `accounts` (`username` varchar(100),`nickname` text,`age` integer,PRIMARY KEY (`username`))
	done.Done(db.AutoMigrate(&Account{}))
	//INSERT INTO `accounts` (`username`,`nickname`,`age`) VALUES ("alice","Alice",17)
	done.Done(db.Create(&Account{Username: "alice", Nickname: "Alice", Age: 17}).Error)

	//SELECT * FROM `accounts` WHERE username="alice" ORDER BY `accounts`.`username` LIMIT 1
	var account Account
	done.Done(db.Where(columnUsername.Eq("alice")).First(&account).Error)
	fmt.Println(neatjsons.S(account))

	//UPDATE `accounts` SET `nickname`="Alice-2" WHERE `username` = "alice"
	done.Done(db.Model(&account).Update(columnNickname.Kv("Alice-2")).Error)
	//SELECT * FROM `accounts` WHERE username="alice" ORDER BY `accounts`.`username` LIMIT 1
	done.Done(db.Where(columnUsername.Eq("alice")).First(&account).Error)
	fmt.Println(neatjsons.S(account))

	//UPDATE `accounts` SET `age`=18,`nickname`="Alice-3" WHERE `username` = "alice"
	done.Done(db.Model(&account).Updates(columnNickname.Kw("Alice-3").Kw(columnAge.Kv(18)).AsMap()).Error)
	//SELECT * FROM `accounts` WHERE username="alice" ORDER BY `accounts`.`username` LIMIT 1
	done.Done(db.Where(columnUsername.Eq("alice")).First(&account).Error)
	fmt.Println(neatjsons.S(account))

	//UPDATE `accounts` SET `age`=age + 1 WHERE `username` = "alice"
	done.Done(db.Model(&account).Update(columnAge.KeAdd(1)).Error)
	//SELECT * FROM `accounts` WHERE username="alice" ORDER BY `accounts`.`username` LIMIT 1
	done.Done(db.Where(columnUsername.Eq("alice")).First(&account).Error)
	fmt.Println(neatjsons.S(account))
}
```

⬆️ **源码:** [源码](internal/demos/demo1x/main.go)

---

## 🔥 高级查询示例

```go
package main

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/yyle88/gormcnm"
	"github.com/yyle88/must"
	"github.com/yyle88/rese"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Example is a gorm model define 3 fields(name, type, rank)
type Example struct {
	Name string `gorm:"primary_key;type:varchar(100);"`
	Type string `gorm:"column:type;"`
	Rank int    `gorm:"column:rank;"`
}

// Now define the fields enum vars(name, type rank)
const (
	columnName = gormcnm.ColumnName[string]("name")
	columnType = gormcnm.ColumnName[string]("type")
	columnRank = gormcnm.ColumnName[int]("rank")
)

func main() {
	//new db connection
	dsn := fmt.Sprintf("file:db-%s?mode=memory&cache=shared", uuid.New().String())
	db := rese.P1(gorm.Open(sqlite.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	}))
	defer rese.F0(rese.P1(db.DB()).Close)

	//create example data
	must.Done(db.AutoMigrate(&Example{}).Error)
	must.Done(db.Save(&Example{Name: "abc", Type: "xyz", Rank: 123}).Error)
	must.Done(db.Save(&Example{Name: "aaa", Type: "xxx", Rank: 456}).Error)

	{
		//SELECT * FROM `examples` WHERE name="abc" ORDER BY `examples`.`name` LIMIT 1
		var res Example
		must.Done(db.Where("name=?", "abc").First(&res).Error)
		fmt.Println(res)
	}
	{
		//SELECT * FROM `examples` WHERE name="abc" AND type="xyz" AND rank>100 AND rank<200 ORDER BY `examples`.`name` LIMIT 1
		var res Example
		must.Done(db.Where(columnName.Eq("abc")).
			Where(columnType.Eq("xyz")).
			Where(columnRank.Gt(100)).
			Where(columnRank.Lt(200)).
			First(&res).Error)
		fmt.Println(res)
	}
}
```

⬆️ **源码:** [源码](internal/demos/demo2x/main.go)

---

## 核心 API

`ColumnName[T]` 泛型类型提供类型安全的 SQL 操作：

```go
type ColumnName[T any] string
```

### 比较操作

```go
db.Where(columnAge.Eq(25))       // =
db.Where(columnAge.Ne(25))       // !=
db.Where(columnAge.Gt(18))       // >
db.Where(columnAge.Gte(18))      // >=
db.Where(columnAge.Lt(65))       // <
db.Where(columnAge.Lte(65))      // <=
```

### 范围和模式操作

| 方法          | SQL               | 示例                   |
|---------------|-------------------|------------------------|
| `Between(a, b)` | `BETWEEN a AND b` | `cls.Age.Between(18, 65)` |
| `In(values)`    | `IN (...)`        | `cls.ID.In([]int{1,2,3})` |
| `Like(pattern)` | `LIKE pattern`    | `cls.Name.Like("A%")`     |
| `IsNull()`      | `IS NULL`         | `cls.DeletedAt.IsNull()`  |
| `IsNotNull()`   | `IS NOT NULL`     | `cls.Email.IsNotNull()`   |

### 更新操作

| 方法      | 说明          | 示例                                                      |
|-----------|---------------|-----------------------------------------------------------|
| `Kv(value)` | 单字段更新  | `db.Model(&user).Update(cls.Age.Kv(26))`                     |
| `Kw(value)` | 构建更新映射     | `cls.Age.Kw(26).Kw(cls.Email.Kv("new@example.com")).AsMap()` |
| `KeAdd(n)`  | 表达式：加      | `db.Model(&user).Update(cls.Age.KeAdd(1))`                   |
| `KeSub(n)`  | 表达式：减 | `db.Model(&user).Update(cls.Score.KeSub(10))`                |

### 聚合与排序

| 方法          | SQL                      | 示例                            |
|---------------|--------------------------|--------------------------------|
| `Count(alias)`  | `COUNT(column) AS alias` | `db.Select(cls.ID.Count("total"))` |
| `Ob(direction)` | `ORDER BY`               | `db.Order(cls.Age.Ob("asc").Ox())` |

---

## 扩展包

本包包含用于特定数据库操作的扩展子包：

- 📦 **gormcnmjson** - 类型安全的 JSON 列操作（支持 SQLite JSON 函数）

**未来扩展**（计划中）：

| 包名          | 用途            | 状态    |
|---------------|-----------------|---------|
| gormcnmtext   | 文本搜索操作    | 计划中  |
| gormcnmdate   | 日期时间操作    | 计划中  |
| gormcnmmath   | 数学运算操作    | 计划中  |

---

## 关联项目

探索完整的 GORM 生态系统集成包：

### 核心生态

- **[gormcnm](https://github.com/yyle88/gormcnm)** - GORM 基础层，提供类型安全的列操作和查询构建器（本项目）
- **[gormcngen](https://github.com/yyle88/gormcngen)** - 使用 AST 的代码生成工具，用于类型安全的 GORM 操作
- **[gormrepo](https://github.com/yyle88/gormrepo)** - 仓储模式实现，遵循 GORM 最佳实践
- **[gormmom](https://github.com/yyle88/gormmom)** - 原生语言 GORM 标签生成引擎，支持智能列名
- **[gormzhcn](https://github.com/go-zwbc/gormzhcn)** - 完整的 GORM 中文编程接口

每个包针对 GORM 开发的不同方面，从本地化到类型安全和代码生成。

---

<!-- TEMPLATE (ZH) BEGIN: STANDARD PROJECT FOOTER -->
<!-- VERSION 2025-11-25 03:52:28.131064 +0000 UTC -->

## 📄 许可证类型

MIT 许可证 - 详见 [LICENSE](LICENSE)。

---

## 💬 联系与反馈

非常欢迎贡献代码！报告 BUG、建议功能、贡献代码：

- 🐛 **问题报告？** 在 GitHub 上提交问题并附上重现步骤
- 💡 **新颖思路？** 创建 issue 讨论
- 📖 **文档疑惑？** 报告问题，帮助我们完善文档
- 🚀 **需要功能？** 分享使用场景，帮助理解需求
- ⚡ **性能瓶颈？** 报告慢操作，协助解决性能问题
- 🔧 **配置困扰？** 询问复杂设置的相关问题
- 📢 **关注进展？** 关注仓库以获取新版本和功能
- 🌟 **成功案例？** 分享这个包如何改善工作流程
- 💬 **反馈意见？** 欢迎提出建议和意见

---

## 🔧 代码贡献

新代码贡献，请遵循此流程：

1. **Fork**：在 GitHub 上 Fork 仓库（使用网页界面）
2. **克隆**：克隆 Fork 的项目（`git clone https://github.com/yourname/repo-name.git`）
3. **导航**：进入克隆的项目（`cd repo-name`）
4. **分支**：创建功能分支（`git checkout -b feature/xxx`）
5. **编码**：实现您的更改并编写全面的测试
6. **测试**：（Golang 项目）确保测试通过（`go test ./...`）并遵循 Go 代码风格约定
7. **文档**：面向用户的更改需要更新文档
8. **暂存**：暂存更改（`git add .`）
9. **提交**：提交更改（`git commit -m "Add feature xxx"`）确保向后兼容的代码
10. **推送**：推送到分支（`git push origin feature/xxx`）
11. **PR**：在 GitHub 上打开 Merge Request（在 GitHub 网页上）并提供详细描述

请确保测试通过并包含相关的文档更新。

---

## 🌟 项目支持

非常欢迎通过提交 Merge Request 和报告问题来贡献此项目。

**项目支持：**

- ⭐ **给予星标**如果项目对您有帮助
- 🤝 **分享项目**给团队成员和（golang）编程朋友
- 📝 **撰写博客**关于开发工具和工作流程 - 我们提供写作支持
- 🌟 **加入生态** - 致力于支持开源和（golang）开发场景

**祝你用这个包编程愉快！** 🎉🎉🎉

<!-- TEMPLATE (ZH) END: STANDARD PROJECT FOOTER -->

---

## 📈 GitHub Stars

[![starring](https://starchart.cc/yyle88/gormcnm.svg?variant=adaptive)](https://starchart.cc/yyle88/gormcnm)
