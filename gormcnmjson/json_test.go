package gormcnmjson_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yyle88/gormcnm"
	"github.com/yyle88/gormcnm/gormcnmjson"
	"github.com/yyle88/gormcnm/internal/tests"
	"github.com/yyle88/must"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Product struct {
	Code string         `gorm:"primarykey"`
	Name string         `gorm:"type:varchar(100)"`
	Meta datatypes.JSON `gorm:"type:json"`
}

const (
	columnCode = gormcnm.ColumnName[string]("code")
	columnName = gormcnm.ColumnName[string]("name")
	columnMeta = gormcnm.ColumnName[datatypes.JSON]("meta")
)

func TestColumn_Get(t *testing.T) {
	tests.NewDBRun(t, func(db *gorm.DB) {
		must.Done(db.AutoMigrate(&Product{}))

		products := []Product{
			{Code: "P001", Name: "iPhone", Meta: datatypes.JSON([]byte(`{"brand":"Apple","price":999}`))},
			{Code: "P002", Name: "Mate60", Meta: datatypes.JSON([]byte(`{"brand":"HuaWei","price":899}`))},
			{Code: "P003", Name: "Mi14", Meta: datatypes.JSON([]byte(`{"brand":"XiaoMi","price":799}`))},
		}
		must.Done(db.Create(&products).Error)

		var results []Product
		require.NoError(t, db.Where(gormcnmjson.Raw(columnMeta).Get("brand").Eq("Apple")).Find(&results).Error)

		require.Len(t, results, 1)
		require.Equal(t, "iPhone", results[0].Name)
	})
}

func TestColumn_GetInt(t *testing.T) {
	tests.NewDBRun(t, func(db *gorm.DB) {
		must.Done(db.AutoMigrate(&Product{}))

		products := []Product{
			{Code: "P001", Name: "iPhone", Meta: datatypes.JSON([]byte(`{"brand":"Apple","price":999}`))},
			{Code: "P002", Name: "MacBook", Meta: datatypes.JSON([]byte(`{"brand":"Apple","price":1999}`))},
			{Code: "P003", Name: "Mate60", Meta: datatypes.JSON([]byte(`{"brand":"HuaWei","price":699}`))},
			{Code: "P004", Name: "Mi14", Meta: datatypes.JSON([]byte(`{"brand":"XiaoMi","price":899}`))},
		}
		must.Done(db.Create(&products).Error)

		var results []Product
		require.NoError(t, db.Where(gormcnmjson.Raw(columnMeta).GetInt("price").Gt(1000)).Find(&results).Error)

		require.Len(t, results, 1)
		require.Equal(t, "MacBook", results[0].Name)
	})
}

func TestColumn_Length(t *testing.T) {
	tests.NewDBRun(t, func(db *gorm.DB) {
		must.Done(db.AutoMigrate(&Product{}))

		products := []Product{
			{Code: "P001", Name: "iPhone", Meta: datatypes.JSON([]byte(`{"tags":["phone","5G","iOS"]}`))},
			{Code: "P002", Name: "MacBook", Meta: datatypes.JSON([]byte(`{"tags":["laptop","macOS"]}`))},
			{Code: "P003", Name: "Mate60", Meta: datatypes.JSON([]byte(`{"tags":["phone","5G"]}`))},
			{Code: "P004", Name: "Mi14", Meta: datatypes.JSON([]byte(`{"tags":["phone"]}`))},
		}
		must.Done(db.Create(&products).Error)

		var results []Product
		require.NoError(t, db.Where(gormcnmjson.Raw(columnMeta).Length("tags").Gt(2)).Find(&results).Error)

		require.Len(t, results, 1)
		require.Equal(t, "iPhone", results[0].Name)
	})
}

func TestColumn_Extract(t *testing.T) {
	tests.NewDBRun(t, func(db *gorm.DB) {
		must.Done(db.AutoMigrate(&Product{}))

		products := []Product{
			{Code: "P001", Name: "iPhone", Meta: datatypes.JSON([]byte(`{"specs":{"storage":"128GB","chip":"A17"}}`))},
			{Code: "P002", Name: "Mate60", Meta: datatypes.JSON([]byte(`{"specs":{"storage":"256GB","chip":"Kirin"}}`))},
			{Code: "P003", Name: "Mi14", Meta: datatypes.JSON([]byte(`{"specs":{"storage":"512GB","chip":"Snapdragon"}}`))},
		}
		must.Done(db.Create(&products).Error)

		var result struct {
			Name    string
			Storage string
		}
		require.NoError(t, db.Table("products").
			Select(string(columnName)+", "+string(columnMeta)+" ->> '$.specs.storage' as storage").
			Where(string(columnName)+" = ?", "iPhone").
			First(&result).Error)

		require.Equal(t, "iPhone", result.Name)
		require.Equal(t, "128GB", result.Storage)
	})
}

func TestColumn_Type(t *testing.T) {
	tests.NewDBRun(t, func(db *gorm.DB) {
		must.Done(db.AutoMigrate(&Product{}))

		products := []Product{
			{Code: "P001", Name: "iPhone", Meta: datatypes.JSON([]byte(`{"key":"value"}`))},
			{Code: "P002", Name: "Mate60", Meta: datatypes.JSON([]byte(`{"key":"value"}`))},
			{Code: "P003", Name: "Mi14", Meta: datatypes.JSON([]byte(`{"key":"value"}`))},
		}
		must.Done(db.Create(&products).Error)

		var result struct {
			JSONType string
		}
		require.NoError(t, db.Table("products").
			Select("JSON_TYPE("+string(columnMeta)+") as json_type").
			Limit(1).
			Scan(&result).Error)

		require.Equal(t, "object", result.JSONType)
	})
}

func TestColumn_Valid(t *testing.T) {
	tests.NewDBRun(t, func(db *gorm.DB) {
		must.Done(db.AutoMigrate(&Product{}))

		products := []Product{
			{Code: "P001", Name: "iPhone", Meta: datatypes.JSON([]byte(`{"valid":true}`))},
			{Code: "P002", Name: "Mate60", Meta: datatypes.JSON([]byte(`{"valid":true}`))},
			{Code: "P003", Name: "Mi14", Meta: datatypes.JSON([]byte(`{"valid":true}`))},
		}
		must.Done(db.Create(&products).Error)

		var result struct {
			IsValid int
		}
		require.NoError(t, db.Table("products").
			Select("JSON_VALID("+string(columnMeta)+") as is_valid").
			Limit(1).
			Scan(&result).Error)

		require.Equal(t, 1, result.IsValid)
	})
}

func TestColumn_Update(t *testing.T) {
	tests.NewDBRun(t, func(db *gorm.DB) {
		must.Done(db.AutoMigrate(&Product{}))

		products := []Product{
			{Code: "P001", Name: "iPhone", Meta: datatypes.JSON([]byte(`{"brand":"Apple","price":999}`))},
			{Code: "P002", Name: "Mate60", Meta: datatypes.JSON([]byte(`{"brand":"HuaWei","price":899}`))},
			{Code: "P003", Name: "Mi14", Meta: datatypes.JSON([]byte(`{"brand":"XiaoMi","price":799}`))},
		}
		must.Done(db.Create(&products).Error)

		require.NoError(t, db.Model(&Product{}).
			Where(columnCode.Eq("P001")).
			Update(string(columnMeta), gorm.Expr(gormcnmjson.Raw(columnMeta).Set("price", 1099).Name())).Error)

		var updated Product
		require.NoError(t, db.Where(columnCode.Eq("P001")).First(&updated).Error)

		var meta map[string]interface{}
		require.NoError(t, json.Unmarshal(updated.Meta, &meta))
		require.Equal(t, "1099", meta["price"])
	})
}

func TestColumn_New(t *testing.T) {
	type ProductWithString struct {
		Code string `gorm:"primarykey"`
		Name string `gorm:"type:varchar(100)"`
		Meta string `gorm:"type:json"`
	}

	const stringMeta = gormcnm.ColumnName[string]("meta")

	tests.NewDBRun(t, func(db *gorm.DB) {
		must.Done(db.AutoMigrate(&ProductWithString{}))

		products := []ProductWithString{
			{Code: "P001", Name: "iPhone", Meta: `{"brand":"Apple","price":999}`},
			{Code: "P002", Name: "Mate60", Meta: `{"brand":"HuaWei","price":899}`},
			{Code: "P003", Name: "Mi14", Meta: `{"brand":"XiaoMi","price":799}`},
		}
		must.Done(db.Create(&products).Error)

		var results []ProductWithString
		require.NoError(t, db.Where(gormcnmjson.New(stringMeta).Get("brand").Eq("Apple")).Find(&results).Error)

		require.Len(t, results, 1)
		require.Equal(t, "iPhone", results[0].Name)
	})
}
