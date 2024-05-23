package gormcnm

//这个文件里放些补充功能的小函数

func stmtAsAlias(stmt string, alias string) string {
	if alias != "" {
		stmt += " as " + alias
	}
	return stmt
}
