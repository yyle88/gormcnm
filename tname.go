package gormcnm

func (s ColumnName[TYPE]) TC(tab tableNameFace) *ColumnInTableOperationClass[TYPE] {
	return &ColumnInTableOperationClass[TYPE]{
		tab: tab,
		cnm: s,
	}
}

func (s ColumnName[TYPE]) TN(tableName string) *ColumnInTableOperationClass[TYPE] {
	return &ColumnInTableOperationClass[TYPE]{
		tab: &classImplementsTableName{tableName: tableName}, //把名称转换为接口要不然同样的代码得写两遍
		cnm: s,
	}
}

// ColumnInTableOperationClass 就是表名和列名的二元组
// 我们这个包是以列名为中心设计的，但依然有部分操作需要用到表名，特别是在join的时候，因此还是增加了这个类型的
type ColumnInTableOperationClass[TYPE any] struct {
	tab tableNameFace
	cnm ColumnName[TYPE]
}

// Eq Return a raw string: "table_a.column_a = table_b.column_b"
// 就是在 join 的时候，让不同表的两个列划等号，能省去些许代码，而且能确保两列的类型是相同的
func (tc *ColumnInTableOperationClass[TYPE]) Eq(xc *ColumnInTableOperationClass[TYPE]) string {
	return tc.Name() + " = " + xc.Name()
}

// Sp 就是当有除了相等以外的其它情况时，就用这个函数，但感觉这种情况时很少见的
// Sp 基本表示 "ship" 的缩写，以我这种贫穷的英语，也只能这样起名吧
func (tc *ColumnInTableOperationClass[TYPE]) Sp(op string, xc *ColumnInTableOperationClass[TYPE]) string {
	return tc.Name() + " " + op + " " + xc.Name()
}

// Name Return a raw string: "table.column_name"
// 当你需要使用join查询时，就需要指定表名，比如需要 Select("users.id, orders.amount") 这样的语句，就需要给列名指定表名
// 这个方法还叫 Name 的原因是，让这个新类型也实现 Name 接口，这样在使用时或许能有些许方便
func (tc *ColumnInTableOperationClass[TYPE]) Name() string {
	return tc.tab.TableName() + "." + tc.cnm.Name()
}

// ColumnName 就是返回合并后的列名，这个虽然文雅，但是命名太长其实也不太方便
func (tc *ColumnInTableOperationClass[TYPE]) ColumnName() ColumnName[TYPE] {
	return ColumnName[TYPE](tc.tab.TableName() + "." + string(tc.cnm))
}

// Cnm 这个虽然是个很粗鄙的命名，但是确实比较简单，而且还很生僻不容易和其它标志符重名，因此还是相当不错的
func (tc *ColumnInTableOperationClass[TYPE]) Cnm() ColumnName[TYPE] {
	return ColumnName[TYPE](tc.tab.TableName() + "." + string(tc.cnm))
}

// Ob 还是排序相关的，增加个语法糖，避免老在外面使用 Cnm 函数，这样不太优雅
func (tc *ColumnInTableOperationClass[TYPE]) Ob(direction string) OrderByBottle {
	return tc.Cnm().Ob(direction)
}

// AsAlias Return a raw string: "table.column_name as alias"
// 当你需要使用join查询时，就需要指定表名，比如需要 Select("users.id as user_id, orders.amount as order_amount") 这样的语句
func (tc *ColumnInTableOperationClass[TYPE]) AsAlias(alias string) string {
	return stmtAsAlias(tc.Name(), alias)
}

// AsName Return a raw string: "table.column_name as alias"
// 当你需要使用join查询时，就需要指定表名，比如需要 Select("users.id as user_id, orders.amount as order_amount") 这样的语句
// 但是假如当 user_id 也是另一个表的 column_name 的时候，你就可以直接传这个类型，而不是 raw string
// 当然这种使用场景是很少的
func (tc *ColumnInTableOperationClass[TYPE]) AsName(newColumnName nameInterface) string {
	return stmtAsAlias(tc.Name(), newColumnName.Name())
}
