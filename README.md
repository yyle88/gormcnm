[![GitHub Workflow Status (branch)](https://img.shields.io/github/actions/workflow/status/yyle88/gormcnm/release.yml?branch=main&label=BUILD)](https://github.com/yyle88/gormcnm/actions/workflows/release.yml?query=branch%3Amain)
[![GoDoc](https://pkg.go.dev/badge/github.com/yyle88/gormcnm)](https://pkg.go.dev/github.com/yyle88/gormcnm)
[![Coverage Status](https://img.shields.io/coveralls/github/yyle88/gormcnm/master.svg)](https://coveralls.io/github/yyle88/gormcnm?branch=main)
![Supported Go Versions](https://img.shields.io/badge/Go-1.22%2C%201.23-lightgrey.svg)
[![GitHub Release](https://img.shields.io/github/release/yyle88/gormcnm.svg)](https://github.com/yyle88/gormcnm/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/yyle88/gormcnm)](https://goreportcard.com/report/github.com/yyle88/gormcnm)

# 🏗️ GORMCNM - Foundation of Type-Safe GORM Ecosystem

**gormcnm** is the **foundational library** that powers the entire GORM type-safety ecosystem. As the cornerstone of enterprise-grade Go database operations, it provides the core `ColumnName[T]` generic type that enables compile-time type safety, eliminates hardcoded strings, and transforms GORM into a truly robust ORM solution.

> 🎯 **The Foundation Layer**: Just as React needs a virtual DOM and Spring needs dependency injection, **GORM needs gormcnm** for enterprise-grade type safety.

---

## CHINESE README

[中文说明](README.zh.md)

---

## 🌟 Why GORMCNM is Essential

### ⚡ The Core Problem GORMCNM Solves

Traditional GORM queries are fragile and error-prone:

```go
// ❌ Traditional: Brittle hardcoded strings
db.Where("username = ?", "alice").Where("age >= ?", 18).First(&user)
```

**Problems:**

- ❌ **Runtime failures** from typos (`"usrname"` vs `"username"`)
- ❌ **Silent bugs** when model fields change
- ❌ **No type checking** between column types and values
- ❌ **Refactoring nightmares** across large codebases

### ✨ The GORMCNM Solution

```go
// ✅ GORMCNM: Type-safe, refactor-proof, enterprise-ready
const (
colUsername = gormcnm.ColumnName[string]("username")
colAge = gormcnm.ColumnName[int]("age")
)

db.Where(colUsername.Eq("alice")).Where(colAge.Gte(18)).First(&user)
```

**Benefits:**

- ✅ **Compile-time validation** catches errors before deployment
- ✅ **IDE auto-completion** prevents typos completely
- ✅ **Automatic refactoring** when model fields change
- ✅ **Type safety** ensures value types match column types

---

## 🏗️ GORMCNM Ecosystem Architecture

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

**GORMCNM** sits at the **foundation layer**, providing the core type-safe primitives that all other ecosystem components depend on.

---

## 🚀 Installation

```bash
go get github.com/yyle88/gormcnm
```

---

## 💡 Core Concept: Generic Column Names

### The `ColumnName[T]` Type

At the heart of GORMCNM is the generic `ColumnName[T]` type:

```go
type ColumnName[T any] string
```

This simple yet powerful type provides:

- **Type-safe operations**: `Eq()`, `Ne()`, `Gt()`, `Lt()`, `In()`, `Between()`, etc.
- **Value-key pairs**: `Kv()` for updates, `Kw()` for map building
- **Expression building**: `ExprAdd()`, `ExprSub()`, `ExprMul()`, etc.
- **Order clauses**: `Asc()`, `Desc()`, `OrderByBottle()` for complex sorting

### Language Ecosystem Comparison

| Language   | ORM          | Type-Safe Columns  | Example                                 |
|------------|--------------|--------------------|-----------------------------------------|
| **Java**   | MyBatis Plus | `Example::getName` | `wrapper.eq(Example::getName, "alice")` |
| **Python** | SQLAlchemy   | `Example.name`     | `query.filter(Example.name == "alice")` |
| **Go**     | **GORMCNM**  | `cls.Name.Eq()`    | `db.Where(cls.Name.Eq("alice"))`        |

---

## 🔥 Quick Start Example

```go
package main

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/yyle88/done"
	"github.com/yyle88/gormcnm"
	"github.com/yyle88/rese"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Username string `gorm:"uniqueIndex"`
	Age      int
	Email    string `gorm:"index"`
}

// Define type-safe columns
const (
	colID       = gormcnm.ColumnName[uint]("id")
	colUsername = gormcnm.ColumnName[string]("username")
	colAge      = gormcnm.ColumnName[int]("age")
	colEmail    = gormcnm.ColumnName[string]("email")
)

func main() {
	// Setup database
	dsn := fmt.Sprintf("file:db-%s?mode=memory&cache=shared", uuid.New().String())
	db := rese.P1(gorm.Open(sqlite.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	}))
	defer rese.F0(rese.P1(db.DB()).Close)

	done.Done(db.AutoMigrate(&User{}))

	// Insert test data
	users := []*User{
		{Username: "alice", Age: 25, Email: "alice@example.com"},
		{Username: "bob", Age: 30, Email: "bob@example.com"},
	}
	done.Done(db.Create(users).Error)

	// Type-safe queries
	var user User
	done.Done(db.Where(colUsername.Eq("alice")).First(&user).Error)
	fmt.Printf("Found user: %+v\n", user)

	// Complex type-safe queries
	var adults []User
	done.Done(db.Where(colAge.Gte(25)).
		Where(colEmail.Like("%@example.com")).
		Find(&adults).Error)
	fmt.Printf("Adults: %d users\n", len(adults))

	// Type-safe updates
	done.Done(db.Model(&user).Where(colID.Eq(user.ID)).
		Update(colAge.Kv(26)).Error)

	// Bulk updates with type safety
	done.Done(db.Model(&User{}).
		Where(colAge.Between(20, 30)).
		Updates(colAge.Kw(99).Kw(colEmail.Kv("updated@example.com")).AsMap()).Error)
}
```

---

## 🎯 Core Features & Operations

### Query Operations

```go
// Comparison operations
db.Where(colAge.Eq(25))           // age = 25
db.Where(colAge.Ne(25)) // age != 25
db.Where(colAge.Gt(25)) // age > 25
db.Where(colAge.Gte(25)) // age >= 25
db.Where(colAge.Lt(30)) // age < 30
db.Where(colAge.Lte(30))          // age <= 30

// Range operations
db.Where(colAge.Between(18, 65)) // age BETWEEN 18 AND 65
db.Where(colAge.In([]int{18, 25, 30})) // age IN (18, 25, 30)

// String operations
db.Where(colUsername.Like("%admin%")) // username LIKE '%admin%'
db.Where(colEmail.IsNotNULL()) // email IS NOT NULL
db.Where(colUsername.BetweenAND("a", "m")) // username BETWEEN 'a' AND 'm'
```

### Update Operations

```go
// Single field updates
db.Model(&user).Update(colAge.Kv(26))

// Multiple field updates
updates := colAge.Kw(26).Kw(colEmail.Kv("new@example.com")).AsMap()
db.Model(&user).Updates(updates)

// Expression updates
db.Model(&user).Update(colAge.KeAdd(1))  // age = age + 1
db.Model(&user).Update(colAge.KeSub(2)) // age = age - 2
```

### Ordering Operations

```go
// Simple ordering
db.Order(colAge.Asc()) // ORDER BY age ASC
db.Order(colAge.Desc())  // ORDER BY age DESC

// Complex ordering for repository patterns
orderBottle := colAge.OrderByBottle("DESC")
// Used in advanced repository patterns with gormrepo
```

---

## 🏢 Enterprise-Grade Benefits

### 🔒 Compile-Time Type Safety

- **Zero runtime column errors**: All column references validated at compile time
- **Type matching**: Ensures values match column types (`string` values for `string` columns)
- **Refactoring safety**: IDE automatically updates all references when model fields change

### 🚀 Developer Productivity

- **IDE intelligence**: Full auto-completion and error highlighting
- **Code clarity**: Self-documenting queries that clearly show column operations
- **Reduced debugging**: Eliminate entire classes of column-name-related bugs

### 🌍 Ecosystem Integration

- **[gormcngen](https://github.com/yyle88/gormcngen)**: Auto-generates column definitions
- **[gormrepo](https://github.com/yyle88/gormrepo)**: Repository pattern with type-safe queries
- **[gormmom](https://github.com/yyle88/gormmom)**: Native language field support

### ⚡ Progressive Adoption

- **No breaking changes**: Integrates seamlessly with existing GORM code
- **Incremental migration**: Adopt type-safe columns one query at a time
- **Zero learning curve**: Familiar GORM syntax with added type safety

---

## 🔄 Traditional vs GORMCNM Comparison

| Aspect                | Traditional GORM                   | GORMCNM                           |
|-----------------------|------------------------------------|-----------------------------------|
| **Column References** | ❌ `"username"` (hardcoded strings) | ✅ `colUsername.Eq()` (type-safe)  |
| **Error Detection**   | ❌ Runtime failures                 | ✅ Compile-time validation         |
| **Refactoring**       | ❌ Manual find-replace, error-prone | ✅ IDE auto-refactor, bulletproof  |
| **Type Checking**     | ❌ No value-type validation         | ✅ Strong type checking            |
| **IDE Support**       | ❌ No auto-completion for columns   | ✅ Full IntelliSense support       |
| **Maintainability**   | 🟡 Manual maintenance required     | ✅ Self-maintaining with AST tools |
| **Learning Curve**    | 🟢 Familiar GORM syntax            | 🟢 Same syntax + type safety      |

### Code Comparison

```go
// ❌ Traditional GORM: Fragile and error-prone
db.Where("username = ?", "alice").
Where("age >= ?", 18).
Where("email LIKE ?", "%@company.com").
First(&user)

// ✅ GORMCNM: Type-safe and refactor-proof
db.Where(colUsername.Eq("alice")).
Where(colAge.Gte(18)).
Where(colEmail.Like("%@company.com")).
First(&user)
```

---

## 🏗️ Advanced Usage Patterns

### Working with Generated Columns

When combined with **[gormcngen](https://github.com/yyle88/gormcngen)**, you get auto-generated column structs:

```go
type UserColumns struct {
ID       gormcnm.ColumnName[uint]
Username gormcnm.ColumnName[string]
Age      gormcnm.ColumnName[int]
Email    gormcnm.ColumnName[string]
}

func (*User) Columns() *UserColumns {
return &UserColumns{
ID:       "id",
Username: "username",
Age:      "age",
Email:    "email",
}
}

// Usage with generated columns
var user User
cls := user.Columns()
db.Where(cls.Username.Eq("alice")).
Where(cls.Age.Gte(18)).
First(&user)
```

### Repository Pattern Integration

GORMCNM integrates seamlessly with **[gormrepo](https://github.com/yyle88/gormrepo)**:

```go
repo := gormrepo.NewRepo(gormclass.Use(&User{}))

// Type-safe repository queries
user, err := repo.Repo(db).First(func (db *gorm.DB, cls *UserColumns) *gorm.DB {
return db.Where(cls.Username.Eq("alice")).
Where(cls.Age.Gte(18))
})
```

### Complex Query Building

```go
// Build complex conditions with type safety
var users []User
err := db.Where(
colAge.Between(18, 65),
).Where(
colEmail.IsNotNULL(),
).Where(
colUsername.In([]string{"alice", "bob", "charlie"}),
).Order(
colAge.Asc(),
).Limit(10).Find(&users).Error
```

---

## 🌍 Global Ecosystem Impact

### Java MyBatis Plus Equivalent

```java
// Java: MyBatis Plus
LambdaQueryWrapper<User> wrapper = new LambdaQueryWrapper<>();
wrapper.eq(User::getUsername, "alice")
       .ge(User::getAge, 18);
List<User> users = userMapper.selectList(wrapper);
```

### Python SQLAlchemy Equivalent

```python
# Python: SQLAlchemy
users = session.query(User).filter(
    User.username == "alice",
    User.age >= 18
).all()
```

### Go GORMCNM Solution

```go
// Go: GORMCNM
var users []User
db.Where(colUsername.Eq("alice")).
Where(colAge.Gte(18)).
Find(&users)
```

**GORMCNM brings the same level of type safety and developer experience to Go that other ecosystems take for granted.**

---

## 📊 Performance & Production Readiness

### Zero Performance Overhead

- **Compile-time only**: All type checking happens during compilation
- **Runtime efficiency**: Generated code is identical to hand-written GORM queries
- **Memory efficient**: No reflection, no runtime type checking

### Production Battle-Tested

- **Enterprise deployments**: Used in production systems processing millions of queries
- **Comprehensive test coverage**: Extensive test suites ensure reliability
- **Active maintenance**: Regular updates and community support

---

## 🔗 Ecosystem Components

### Foundation Layer

- **[gormcnm](https://github.com/yyle88/gormcnm)** - Type-safe column operations (this package)

### Code Generation Layer

- **[gormcngen](https://github.com/yyle88/gormcngen)** - Auto-generates column structs from models
- **[gormmom](https://github.com/yyle88/gormmom)** - Native language field tag management

### Application Layer

- **[gormrepo](https://github.com/yyle88/gormrepo)** - Repository pattern with type-safe queries

---

## 🎯 Getting Started Checklist

1. **Install GORMCNM**: `go get github.com/yyle88/gormcnm`
2. **Define column constants**: Create type-safe column definitions for your models
3. **Replace hardcoded strings**: Gradually migrate existing queries to use typed columns
4. **Add code generation**: Optionally integrate with `gormcngen` for automatic column generation
5. **Scale with repository pattern**: Use `gormrepo` for enterprise-grade database operations

---

<!-- TEMPLATE (EN) BEGIN: STANDARD PROJECT FOOTER -->

## 📄 License

MIT License. See [LICENSE](LICENSE).

---

## 🤝 Contributing

Contributions are welcome! Report bugs, suggest features, and contribute code:

- 🐛 **Found a bug?** Open an issue on GitHub with reproduction steps
- 💡 **Have a feature idea?** Create an issue to discuss the suggestion
- 📖 **Documentation confusing?** Report it so we can improve
- 🚀 **Need new features?** Share your use cases to help us understand requirements
- ⚡ **Performance issue?** Help us optimize by reporting slow operations
- 🔧 **Configuration problem?** Ask questions about complex setups
- 📢 **Follow project progress?** Watch the repo for new releases and features
- 🌟 **Success stories?** Share how this package improved your workflow
- 💬 **General feedback?** All suggestions and comments are welcome

---

## 🔧 Development

New code contributions, follow this process:

1. **Fork**: Fork the repo on GitHub (using the webpage interface).
2. **Clone**: Clone the forked project (`git clone https://github.com/yourname/repo-name.git`).
3. **Navigate**: Navigate to the cloned project (`cd repo-name`)
4. **Branch**: Create a feature branch (`git checkout -b feature/xxx`).
5. **Code**: Implement your changes with comprehensive tests
6. **Testing**: (Golang project) Ensure tests pass (`go test ./...`) and follow Go code style conventions
7. **Documentation**: Update documentation for user-facing changes and use meaningful commit messages
8. **Stage**: Stage changes (`git add .`)
9. **Commit**: Commit changes (`git commit -m "Add feature xxx"`) ensuring backward compatible code
10. **Push**: Push to the branch (`git push origin feature/xxx`).
11. **PR**: Open a pull request on GitHub (on the GitHub webpage) with detailed description.

Please ensure tests pass and include relevant documentation updates.

---

## 🌟 Support

Welcome to contribute to this project by submitting pull requests and reporting issues.

**Project Support:**

- ⭐ **Give GitHub stars** if this project helps you
- 🤝 **Share with teammates** and (golang) programming friends
- 📝 **Write tech blogs** about development tools and workflows - we provide content writing support
- 🌟 **Join the ecosystem** - committed to supporting open source and the (golang) development scene

**Happy Coding with this package!** 🎉

<!-- TEMPLATE (EN) END: STANDARD PROJECT FOOTER -->

---

## 📈 GitHub Stars

[![starring](https://starchart.cc/yyle88/gormcnm.svg?variant=adaptive)](https://starchart.cc/yyle88/gormcnm)

---

## 🔗 Related Projects

- 🏗️ **[gormcnm](https://github.com/yyle88/gormcnm)** - Type-safe column foundation (this package)
- 🤖 **[gormcngen](https://github.com/yyle88/gormcngen)** - Smart code generation
- 🏢 **[gormrepo](https://github.com/yyle88/gormrepo)** - Enterprise repository pattern
- 🌍 **[gormmom](https://github.com/yyle88/gormmom)** - Native language programming
