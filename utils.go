package gormcnm

//这个文件里放些补充功能的小函数

// 设置别名，返回类似 COUNT(*) as cnt 这样的
func stmtAsAlias(stmt string, alias string) string {
	if alias != "" {
		stmt += " as " + alias
	}
	return stmt
}

// 简单定义个接口
type nameInterface interface {
	Name() string
}

// 简单定义个接口
type tableNameFace interface {
	TableName() string
}
