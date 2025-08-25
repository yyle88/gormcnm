[![GitHub Workflow Status (branch)](https://img.shields.io/github/actions/workflow/status/yyle88/gormcnm/release.yml?branch=main&label=BUILD)](https://github.com/yyle88/gormcnm/actions/workflows/release.yml?query=branch%3Amain)
[![GoDoc](https://pkg.go.dev/badge/github.com/yyle88/gormcnm)](https://pkg.go.dev/github.com/yyle88/gormcnm)
[![Coverage Status](https://img.shields.io/coveralls/github/yyle88/gormcnm/master.svg)](https://coveralls.io/github/yyle88/gormcnm?branch=main)
![Supported Go Versions](https://img.shields.io/badge/Go-1.22%2C%201.23-lightgrey.svg)
[![GitHub Release](https://img.shields.io/github/release/yyle88/gormcnm.svg)](https://github.com/yyle88/gormcnm/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/yyle88/gormcnm)](https://goreportcard.com/report/github.com/yyle88/gormcnm)

# 🏗️ GORMCNM - GORM 类型安全列名基础层

**gormcnm** 是整个 **GORM 生态系统的基础层**，提供完全类型安全的列名操作，彻底消除数据库操作中的硬编码字符串。

> 🎯 **零运行时错误：在编译时捕获所有列名和类型错误**

---

## 英文文档

[ENGLISH README](README.md)

---

## 🎯 核心理念

### ✨ 类型安全的列名操作
- **泛型列名定义**：`ColumnName[T]` 确保类型安全
- **编译时验证**：消除所有硬编码字符串错误
- **IDE 智能支持**：完整的代码补全和类型检查

### 🔧 完整的 SQL 操作符
- **比较操作**：`Eq()`、`Ne()`、`Gt()`、`Gte()`、`Lt()`、`Lte()`
- **范围操作**：`In()`、`NotIn()`、`Between()`、`NotBetween()`
- **模式匹配**：`Like()`、`NotLike()`、`ILike()`
- **空值检查**：`IsNull()`、`IsNotNull()`

### 📊 数学表达式构建
- **算术运算**：`ExprAdd()`、`ExprSub()`、`ExprMul()`、`ExprDiv()`
- **聚合函数**：`Sum()`、`Count()`、`Avg()`、`Max()`、`Min()`
- **条件表达式**：`CASE WHEN` 构建、`COALESCE` 支持
- **排序支持**：`OrderBy()`、`OrderByBottle()` 方法

---

## 🏗️ 生态系统架构

```
┌─────────────────────────────────────────────────────────────────────┐
│                    GORM Type-Safe Ecosystem                         │
├─────────────────────────────────────────────────────────────────────┤
│                                                                     │
│  ┌─────────────┐    ┌─────────────┐    ┌─────────────┐              │
│  │  gormzhcn   │    │  gormmom    │    │  gormrepo   │              │
│  │ Chinese API │───▶│ Native Lang │───▶│  Package    │─────┐        │
│  │  Localize   │    │  Smart Tags │    │  Pattern    │     │        │
│  └─────────────┘    └─────────────┘    └─────────────┘     │        │
│         │                   │                              │        │
│         │                   ▼                              ▼        │
│         │            ┌─────────────┐              ┌─────────────┐   │
│         │            │ gormcngen   │              │Application  │   │
│         │            │Code Generate│─────────────▶│Custom Code  │   │
│         │            │AST Operation│              │             │   │
│         │            └─────────────┘              └─────────────┘   │
│         │                   │                              ▲        │
│         │                   ▼                              │        │
│         └────────────▶┌─────────────┐◄─────────────────────┘        │
│                       │   GORMCNM   │                               │
│                       │ FOUNDATION  │                               │
│                       │ Type-Safe   │                               │
│                       │ Core Logic  │                               │
│                       └─────────────┘                               │
│                              │                                      │
│                              ▼                                      │
│                       ┌─────────────┐                               │
│                       │    GORM     │                               │
│                       │  Database   │                               │
│                       └─────────────┘                               │
│                                                                     │
└─────────────────────────────────────────────────────────────────────┘
```

**GORMCNM** 作为**基础层**，为整个生态系统提供类型安全的核心逻辑。

---

## 📦 生态系统价值

### 🔹 作为基础层的核心价值
**gormcnm** 是整个生态系统的**基石**，为上层组件提供：
- 类型安全的列名定义和操作
- 完整的 SQL 表达式构建能力
- 编译时错误检测机制

### 🔹 与上层组件的协作
- **[gormcngen](https://github.com/yyle88/gormcngen)** 依赖 gormcnm 生成类型安全的列结构
- **[gormrepo](https://github.com/yyle88/gormrepo)** 使用 gormcnm 实现仓储模式
- **[gormmom](https://github.com/yyle88/gormmom)** 基于 gormcnm 提供原生语言支持
- **[gormzhcn](https://github.com/go-zwbc/gormzhcn)** 利用 gormcnm 实现中文编程接口

---

## 🚀 快速开始

### 安装

```bash
go get github.com/yyle88/gormcnm
```

### 基础使用

#### 1. 定义列名类型

```go
package main

import (
    "github.com/yyle88/gormcnm"
)

// 定义用户表的列名
type UserColumns struct {
    ID       gormcnm.ColumnName[uint]   `json:"id"`
    Username gormcnm.ColumnName[string] `json:"username"`
    Email    gormcnm.ColumnName[string] `json:"email"`
    Age      gormcnm.ColumnName[int]    `json:"age"`
    IsActive gormcnm.ColumnName[bool]   `json:"is_active"`
}

// 实例化列名
func GetUserColumns() *UserColumns {
    return &UserColumns{
        ID:       "id",
        Username: "username", 
        Email:    "email",
        Age:      "age",
        IsActive: "is_active",
    }
}
```

#### 2. 类型安全的查询条件

```go
func queryUsers(db *gorm.DB) {
    cols := GetUserColumns()
    
    // 类型安全的条件构建
    var users []User
    err := db.Where(cols.Username.Eq("alice")).          // username = 'alice'
           Where(cols.Age.Gte(18)).                      // AND age >= 18
           Where(cols.IsActive.Eq(true)).                // AND is_active = true
           Where(cols.Email.Like("%@gmail.com")).        // AND email LIKE '%@gmail.com'
           Find(&users).Error
    
    if err != nil {
        log.Fatal(err)
    }
}
```

#### 3. 复杂条件和表达式

```go
func complexQueries(db *gorm.DB) {
    cols := GetUserColumns()
    
    // 范围查询
    db.Where(cols.Age.Between(18, 65)).Find(&users)
    
    // IN 查询  
    db.Where(cols.Username.In([]string{"alice", "bob", "carol"})).Find(&users)
    
    // 空值检查
    db.Where(cols.Email.IsNotNull()).Find(&users)
    
    // 数学表达式
    db.Select(cols.Age.ExprAdd(10).As("future_age")).Find(&users)
    
    // 聚合查询
    var avgAge float64
    db.Model(&User{}).Select(cols.Age.Avg()).Scan(&avgAge)
    
    // 排序
    db.Order(cols.Username.OrderBy("ASC")).
       Order(cols.Age.OrderBy("DESC")).
       Find(&users)
}
```

#### 4. 更新操作

```go
func updateUsers(db *gorm.DB) {
    cols := GetUserColumns()
    
    // 类型安全的更新
    err := db.Model(&User{}).
            Where(cols.Username.Eq("alice")).
            Update(cols.Age.ColName(), 30).Error  // 使用 ColName() 获取列名
    
    // 批量更新
    updates := map[string]interface{}{
        cols.Age.ColName():      25,
        cols.IsActive.ColName(): true,
    }
    err = db.Model(&User{}).
            Where(cols.Age.Lt(18)).
            Updates(updates).Error
}
```

---

## 🔧 核心 API 参考

### 基础比较操作

| 方法 | SQL 等价 | 描述 | 示例 |
|------|---------|------|------|
| `Eq(value)` | `= value` | 等于 | `cols.Name.Eq("alice")` |
| `Ne(value)` | `<> value` | 不等于 | `cols.Age.Ne(0)` |
| `Gt(value)` | `> value` | 大于 | `cols.Age.Gt(18)` |
| `Gte(value)` | `>= value` | 大于等于 | `cols.Age.Gte(18)` |
| `Lt(value)` | `< value` | 小于 | `cols.Age.Lt(65)` |
| `Lte(value)` | `<= value` | 小于等于 | `cols.Age.Lte(65)` |

### 范围和集合操作

| 方法 | SQL 等价 | 描述 | 示例 |
|------|---------|------|------|
| `In(values)` | `IN (values)` | 在集合中 | `cols.ID.In([]int{1,2,3})` |
| `NotIn(values)` | `NOT IN (values)` | 不在集合中 | `cols.Status.NotIn([]string{"deleted"})` |
| `Between(a, b)` | `BETWEEN a AND b` | 范围查询 | `cols.Age.Between(18, 65)` |
| `NotBetween(a, b)` | `NOT BETWEEN a AND b` | 不在范围内 | `cols.Score.NotBetween(0, 60)` |

### 模式匹配操作

| 方法 | SQL 等价 | 描述 | 示例 |
|------|---------|------|------|
| `Like(pattern)` | `LIKE pattern` | 模式匹配 | `cols.Name.Like("A%")` |
| `NotLike(pattern)` | `NOT LIKE pattern` | 不匹配模式 | `cols.Email.NotLike("%spam%")` |
| `ILike(pattern)` | `ILIKE pattern` | 大小写不敏感匹配 | `cols.Name.ILike("alice")` |

### 空值操作

| 方法 | SQL 等价 | 描述 | 示例 |
|------|---------|------|------|
| `IsNull()` | `IS NULL` | 为空 | `cols.DeletedAt.IsNull()` |
| `IsNotNull()` | `IS NOT NULL` | 不为空 | `cols.Email.IsNotNull()` |

### 数学表达式

| 方法 | SQL 等价 | 描述 | 示例 |
|------|---------|------|------|
| `ExprAdd(n)` | `column + n` | 加法 | `cols.Age.ExprAdd(1)` |
| `ExprSub(n)` | `column - n` | 减法 | `cols.Score.ExprSub(10)` |
| `ExprMul(n)` | `column * n` | 乘法 | `cols.Price.ExprMul(1.1)` |
| `ExprDiv(n)` | `column / n` | 除法 | `cols.Total.ExprDiv(100)` |

### 聚合函数

| 方法 | SQL 等价 | 描述 | 示例 |
|------|---------|------|------|
| `Sum()` | `SUM(column)` | 求和 | `cols.Amount.Sum()` |
| `Count()` | `COUNT(column)` | 计数 | `cols.ID.Count()` |
| `Avg()` | `AVG(column)` | 平均值 | `cols.Score.Avg()` |
| `Max()` | `MAX(column)` | 最大值 | `cols.Age.Max()` |
| `Min()` | `MIN(column)` | 最小值 | `cols.Price.Min()` |

---

## 💡 最佳实践

### 🎯 列名定义模式

```go
// ✅ 推荐：使用描述性的结构体
type ProductColumns struct {
    ID          gormcnm.ColumnName[uint]      `json:"id"`
    Name        gormcnm.ColumnName[string]   `json:"name"`
    Price       gormcnm.ColumnName[decimal.Decimal] `json:"price"`
    CategoryID  gormcnm.ColumnName[uint]      `json:"category_id"`
    CreatedAt   gormcnm.ColumnName[time.Time] `json:"created_at"`
    UpdatedAt   gormcnm.ColumnName[time.Time] `json:"updated_at"`
}

// ✅ 工厂函数模式
func NewProductColumns() *ProductColumns {
    return &ProductColumns{
        ID:         "id",
        Name:       "name", 
        Price:      "price",
        CategoryID: "category_id",
        CreatedAt:  "created_at", 
        UpdatedAt:  "updated_at",
    }
}
```

### 🔧 复杂查询构建

```go
func findProducts(db *gorm.DB, filters ProductFilters) ([]Product, error) {
    cols := NewProductColumns()
    query := db.Model(&Product{})
    
    // 动态条件构建
    if filters.MinPrice > 0 {
        query = query.Where(cols.Price.Gte(filters.MinPrice))
    }
    
    if filters.MaxPrice > 0 {
        query = query.Where(cols.Price.Lte(filters.MaxPrice))
    }
    
    if len(filters.Categories) > 0 {
        query = query.Where(cols.CategoryID.In(filters.Categories))
    }
    
    if filters.NamePattern != "" {
        query = query.Where(cols.Name.Like("%" + filters.NamePattern + "%"))
    }
    
    // 排序
    query = query.Order(cols.CreatedAt.OrderBy("DESC"))
    
    var products []Product
    err := query.Find(&products).Error
    return products, err
}
```

### 📊 聚合查询示例

```go
func getStatistics(db *gorm.DB) (*ProductStats, error) {
    cols := NewProductColumns()
    
    type Result struct {
        TotalProducts int             `json:"total_products"`
        AvgPrice      decimal.Decimal `json:"avg_price"`
        MaxPrice      decimal.Decimal `json:"max_price"`
        MinPrice      decimal.Decimal `json:"min_price"`
    }
    
    var result Result
    err := db.Model(&Product{}).
            Select(
                cols.ID.Count().As("total_products"),
                cols.Price.Avg().As("avg_price"),
                cols.Price.Max().As("max_price"),
                cols.Price.Min().As("min_price"),
            ).
            Where(cols.CreatedAt.Gte(time.Now().AddDate(0, -1, 0))). // 最近一个月
            Scan(&result).Error
            
    return &ProductStats{
        TotalProducts: result.TotalProducts,
        AvgPrice:     result.AvgPrice,
        MaxPrice:     result.MaxPrice,
        MinPrice:     result.MinPrice,
    }, err
}
```

---

## 🌟 核心优势

### ✨ 编译时安全
- **类型检查**：编译器确保类型匹配
- **IDE 支持**：完整的智能提示和重构
- **重构友好**：字段重命名自动更新所有引用

### ⚡ 性能优化
- **零反射**：纯静态类型定义
- **内联优化**：编译器优化表达式构建
- **缓存友好**：预定义的列名常量

### 🎯 开发体验
- **清晰的 API**：直观的方法命名
- **链式调用**：支持流畅的查询构建
- **错误减少**：消除硬编码字符串错误

---

## 📝 完整示例

查看 [examples](internal/examples) 目录获取完整使用示例。

---

<!-- TEMPLATE (ZH) BEGIN: STANDARD PROJECT FOOTER -->

## 📄 许可证类型

MIT 许可证。详见 [LICENSE](LICENSE)。

---

## 🤝 项目贡献

非常欢迎贡献代码！报告 BUG、建议功能、贡献代码：

- 🐛 **发现问题？** 在 GitHub 上提交问题并附上重现步骤
- 💡 **功能建议？** 创建 issue 讨论您的想法
- 📖 **文档疑惑？** 报告问题，帮助我们改进文档
- 🚀 **需要功能？** 分享使用场景，帮助理解需求
- ⚡ **性能瓶颈？** 报告慢操作，帮助我们优化性能
- 🔧 **配置困扰？** 询问复杂设置的相关问题
- 📢 **关注进展？** 关注仓库以获取新版本和功能
- 🌟 **成功案例？** 分享这个包如何改善工作流程
- 💬 **意见反馈？** 欢迎所有建议和宝贵意见

---

## 🔧 代码贡献

新代码贡献，请遵循此流程：

1. **Fork**：在 GitHub 上 Fork 仓库（使用网页界面）
2. **克隆**：克隆 Fork 的项目（`git clone https://github.com/yourname/repo-name.git`）
3. **导航**：进入克隆的项目（`cd repo-name`）
4. **分支**：创建功能分支（`git checkout -b feature/xxx`）
5. **编码**：实现您的更改并编写全面的测试
6. **测试**：（Golang 项目）确保测试通过（`go test ./...`）并遵循 Go 代码风格约定
7. **文档**：为面向用户的更改更新文档，并使用有意义的提交消息
8. **暂存**：暂存更改（`git add .`）
9. **提交**：提交更改（`git commit -m "Add feature xxx"`）确保向后兼容的代码
10. **推送**：推送到分支（`git push origin feature/xxx`）
11. **PR**：在 GitHub 上打开 Pull Request（在 GitHub 网页上）并提供详细描述

请确保测试通过并包含相关的文档更新。

---

## 🌟 项目支持

非常欢迎通过提交 Pull Request 和报告问题来为此项目做出贡献。

**项目支持：**

- ⭐ **给予星标**如果项目对您有帮助
- 🤝 **分享项目**给团队成员和（golang）编程朋友
- 📝 **撰写博客**关于开发工具和工作流程 - 我们提供写作支持
- 🌟 **加入生态** - 致力于支持开源和（golang）开发场景

**使用这个包快乐编程！** 🎉

<!-- TEMPLATE (ZH) END: STANDARD PROJECT FOOTER -->

---

## 📈 GitHub Stars

[![starring](https://starchart.cc/yyle88/gormcnm.svg?variant=adaptive)](https://starchart.cc/yyle88/gormcnm)

---

## 🔗 相关项目

- 🏗️ **[gormcnm](https://github.com/yyle88/gormcnm)** - 类型安全列基础包
- 🤖 **[gormcngen](https://github.com/yyle88/gormcngen)** - 智能代码生成
- 🏢 **[gormrepo](https://github.com/yyle88/gormrepo)** - 企业仓储模式
- 🌍 **[gormmom](https://github.com/yyle88/gormmom)** - 原生语言编程支持
