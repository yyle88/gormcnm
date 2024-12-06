package gormcnm

import (
	"strings"

	"github.com/yyle88/gormcnm/internal/utils"
	"gorm.io/gorm"
)

// ColumnOperationClass 让外部能够继承它，这样就能继承操作函数，让查询更便捷
// 自动生成的代码，不仅要包含各个列的基本信息，能进行简单查询，还要能够有便捷的操作函数，以便于复杂的查询
type ColumnOperationClass struct{}

func (c *ColumnOperationClass) OK() bool {
	return true //这个函数有奇效，让你把变量的创建放在if{}代码块里
}

func (c *ColumnOperationClass) NewQx(stmt string, args ...interface{}) *QxConjunction {
	return NewQxConjunction(stmt, args...)
}

func (c *ColumnOperationClass) Qx(stmt string, args ...interface{}) *QxConjunction {
	return NewQxConjunction(stmt, args...)
}

func (c *ColumnOperationClass) NewSx(stmt string, args ...interface{}) *SelectStatement {
	return NewSelectStatement(stmt, args...)
}

func (c *ColumnOperationClass) Sx(stmt string, args ...interface{}) *SelectStatement {
	return NewSelectStatement(stmt, args...)
}

func (c *ColumnOperationClass) NewKw() ColumnValueMap {
	return NewKw()
}

func (c *ColumnOperationClass) Kw(columnName string, value interface{}) ColumnValueMap {
	return Kw(columnName, value)
}

// Where 设置查询条件
// 很明显这样做会破坏gorm链式操作的【链式调用语句】，但这样也是可行的，也能简化些代码
func (c *ColumnOperationClass) Where(db *gorm.DB, qxs ...*QxConjunction) *gorm.DB {
	stmt := db
	for _, qx := range qxs {
		stmt = stmt.Where(qx.Qs(), qx.args...)
	}
	return stmt
}

// OrderByColumns 设置排序方向
// 很明显这样做会破坏gorm链式操作的【链式调用语句】，但这样也是可行的，也能简化些代码
func (c *ColumnOperationClass) OrderByColumns(db *gorm.DB, obs ...OrderByBottle) *gorm.DB {
	stmt := db
	for _, ob := range obs {
		stmt = stmt.Order(ob.Ox())
	}
	return stmt
}

// UpdateColumns 根据字典更新数据
// 很明显这样做会破坏gorm链式操作的【链式调用语句】，但这样也是可行的，也能简化些代码
func (c *ColumnOperationClass) UpdateColumns(db *gorm.DB, kws ...ColumnValueMap) *gorm.DB {
	mp := map[string]interface{}{}
	for _, kw := range kws {
		for k, v := range kw.AsMap() {
			mp[k] = v
		}
	}
	return db.UpdateColumns(mp)
}

// MergeNames join column names with comma ", ". return a string. Use it when using db.Select() / db.Group(). thank you!
// 因为恰好有这个函数 func (s ColumnName[TYPE]) Name() string 因此这个函数可以直接接受自定义的列名，再以逗号拼接，相当于是简化逻辑
func (c *ColumnOperationClass) MergeNames(a ...utils.ColumnNameInterface) string {
	var names = make([]string, 0, len(a))
	for _, x := range a {
		names = append(names, x.Name())
	}
	return strings.Join(names, ", ")
}

// MergeStmts join some SQL statements with comma ", ". return a string. Use it when using db.Select() / db.Group(). thank you!
// 理论上函数名叫 Merge 就行，但假如别人定义的 model 里也有 Merge 呢，就有可能冲突（也可能不冲突，毕竟列名通常建议使用名词），因此把函数名写长些，避免发生冲突。
func (c *ColumnOperationClass) MergeStmts(a ...string) string {
	return strings.Join(a, ", ")
}

// CountStmt 统计表中行的数量 count(*) 和 count(1) 两者是等价的，这里使用 count(*) 因为这个更加常用
func (c *ColumnOperationClass) CountStmt(alias string) string {
	return utils.ApplyAliasToColumn("COUNT(*)", alias)
}

// CountCaseWhenStmt COUNT(CASE WHEN condition THEN 1 END): 根据条件统计符合条件的行数
// 这个比较不常用
func (c *ColumnOperationClass) CountCaseWhenStmt(condition string, alias string) string {
	return utils.ApplyAliasToColumn("COUNT(CASE WHEN ("+condition+") THEN 1 END)", alias)
}

func (c *ColumnOperationClass) CountCaseWhenQxSx(qx *QxConjunction, alias string) *SelectStatement {
	return NewSelectStatement(
		utils.ApplyAliasToColumn("COUNT(CASE WHEN ("+qx.Qs()+") THEN 1 END)", alias),
		qx.Args()...,
	)
}

func (c *ColumnOperationClass) CombineSxs(cs ...SelectStatement) *SelectStatement {
	var qsVs []string
	var args []any
	for _, c := range cs {
		qsVs = append(qsVs, c.Qs())
		args = append(args, c.Args()...)
	}
	var stmt = strings.Join(qsVs, ", ")
	return NewSelectStatement(stmt, args...)
}

// Select 选择列和筛选数据
// 和前面 Where OrderByColumns UpdateColumns 类似的，这也会破坏gorm的【链式调用语句】，这里阐释 sx 这个类型的定义目的和使用场景
func (c *ColumnOperationClass) Select(db *gorm.DB, qxs ...*SelectStatement) *gorm.DB {
	stmt := db
	for _, qx := range qxs {
		stmt = stmt.Select(qx.Qs(), qx.args...)
	}
	return stmt
}
