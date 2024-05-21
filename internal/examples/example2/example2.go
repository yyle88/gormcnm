package example2

import "github.com/yyle88/gormcnm"

// Now define the fields enum vars(name, type rank)
const (
	columnName = gormcnm.ColumnName[string]("name")
	columnType = gormcnm.ColumnName[*string]("type")
	columnRank = gormcnm.ColumnName[*int]("rank")
)

// Example is a gorm model define 3 fields(name, type, rank)
type Example struct {
	Name string  `gorm:"primary_key;type:varchar(100);"`
	Type *string `gorm:"column:type;"` //通常不建议把字段定义为指针类型，但不建议不表示不能，因此要试试这种场景
	Rank *int    `gorm:"column:rank;"` //通常不建议把字段定义为指针类型，但不建议不表示不能，因此要试试这种场景
}
