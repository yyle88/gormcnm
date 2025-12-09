// Package gormcnm tests validate core column operations including math expressions and aggregates
// Auto verifies ExprAdd, ExprSub, ExprMul, ExprDiv, ExprConcat, ExprReplace operations
// Tests cover basic expressions, aggregation functions, and column method operations
//
// gormcnm 测试包验证核心列操作，包括数学表达式和聚合函数
// 自动验证 ExprAdd、ExprSub、ExprMul、ExprDiv、ExprConcat、ExprReplace 操作
// 测试涵盖基础表达式、聚合函数和列方法功能
package gormcnm

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yyle88/gormcnm/internal/tests"
	"github.com/yyle88/neatjson/neatjsons"
	"gorm.io/gorm"
)

func TestColumnName_SafeCnm(t *testing.T) {
	type Example struct {
		Name   string `gorm:"primary_key;type:varchar(100);"`
		Create string `gorm:"column:create"`
	}

	const columnCreate = ColumnName[string]("create")

	tests.NewDBRun(t, func(db *gorm.DB) {
		require.NoError(t, db.AutoMigrate(&Example{}))
		require.NoError(t, db.Save(&Example{
			Name:   "aaa",
			Create: "abc",
		}).Error)
		require.NoError(t, db.Save(&Example{
			Name:   "xxx",
			Create: "xyz",
		}).Error)
		require.NoError(t, db.Save(&Example{
			Name:   "uuu",
			Create: "uvw",
		}).Error)

		{
			var one Example
			require.NoError(t, db.Where(columnCreate.SafeCnm(`""`).Eq("abc")).First(&one).Error)
			require.Equal(t, "aaa", one.Name)
			t.Log(neatjsons.S(one))
		}
		{
			var one Example
			require.NoError(t, db.Where(columnCreate.SafeCnm("`").Eq("xyz")).First(&one).Error)
			require.Equal(t, "xxx", one.Name)
			t.Log(neatjsons.S(one))
		}
		{
			var one Example
			require.NoError(t, db.Where(columnCreate.SafeCnm("[]").Eq("uvw")).First(&one).Error)
			require.Equal(t, "uuu", one.Name)
			t.Log(neatjsons.S(one))
		}
		{
			var one Example
			require.NoError(t, db.Where(columnCreate.SafeCnm("[-quote-]").Eq("uvw")).First(&one).Error)
			require.Equal(t, "uuu", one.Name)
			t.Log(neatjsons.S(one))
		}
	})
}

func TestColumnName_Count(t *testing.T) {
	type Example struct {
		Name string `gorm:"primary_key;type:varchar(100);"`
		Type string `gorm:"column:type;"`
	}

	const columnName = ColumnName[string]("name")

	tests.NewDBRun(t, func(db *gorm.DB) {
		require.NoError(t, db.AutoMigrate(&Example{}))
		require.NoError(t, db.Save(&Example{Name: "abc", Type: "xyz"}).Error)
		require.NoError(t, db.Save(&Example{Name: "aaa", Type: "xxx"}).Error)

		{
			var value int
			err := db.Model(&Example{}).Select(columnName.Count("cnt")).First(&value).Error
			require.NoError(t, err)
			require.Equal(t, 2, value)
		}
		{
			type resType struct {
				Cnt int64
			}
			var res resType
			err := db.Model(&Example{}).Select(columnName.CountDistinct("cnt")).First(&res).Error
			require.NoError(t, err)
			require.Equal(t, int64(2), res.Cnt)
		}
	})
}

func TestColumnName_ExprOperations(t *testing.T) {
	type Example struct {
		ID     uint    `gorm:"primary_key"`
		Price  float64 `gorm:"column:price"`
		Amount int     `gorm:"column:amount"`
		Title  string  `gorm:"column:title"`
	}

	const (
		columnPrice  = ColumnName[float64]("price")
		columnAmount = ColumnName[int]("amount")
	)

	tests.NewDBRun(t, func(db *gorm.DB) {
		require.NoError(t, db.AutoMigrate(&Example{}))
		require.NoError(t, db.Create(&Example{
			ID:     1,
			Price:  100.0,
			Amount: 10,
			Title:  "Product",
		}).Error)

		// Test ExprAdd and ExprMul with UpdateColumns
		{
			updateMap := map[string]interface{}{}
			key1, expr1 := columnPrice.KeAdd(10.0) // price = price + 10
			key2, expr2 := columnAmount.KeMul(2)   // quantity = quantity * 2
			updateMap[key1] = expr1
			updateMap[key2] = expr2

			result := db.Model(&Example{}).Where("id = ?", 1).UpdateColumns(updateMap)
			require.NoError(t, result.Error)
			require.Equal(t, int64(1), result.RowsAffected)

			var updated Example
			require.NoError(t, db.First(&updated, 1).Error)
			require.Equal(t, 110.0, updated.Price) // 100 + 10
			require.Equal(t, 20, updated.Amount)   // 10 * 2
			t.Log("updated result:", neatjsons.S(updated))
		}

		// Test ExprSub and ExprDiv
		{
			updateMap := map[string]interface{}{}
			key1, expr1 := columnPrice.KeSub(10.0) // price = price - 10
			key2, expr2 := columnAmount.KeDiv(4)   // quantity = quantity / 4
			updateMap[key1] = expr1
			updateMap[key2] = expr2

			result := db.Model(&Example{}).Where("id = ?", 1).UpdateColumns(updateMap)
			require.NoError(t, result.Error)
			require.Equal(t, int64(1), result.RowsAffected)

			var updated Example
			require.NoError(t, db.First(&updated, 1).Error)
			require.Equal(t, 100.0, updated.Price) // 110 - 10
			require.Equal(t, 5, updated.Amount)    // 20 / 4
			t.Log("updated result:", neatjsons.S(updated))
		}
	})
}

func TestColumnName_StringExprOperations(t *testing.T) {
	type Example struct {
		ID    uint   `gorm:"primary_key"`
		Title string `gorm:"column:title"`
		Email string `gorm:"column:email"`
	}

	const (
		columnTitle = ColumnName[string]("title")
		columnEmail = ColumnName[string]("email")
	)

	tests.NewDBRun(t, func(db *gorm.DB) {
		require.NoError(t, db.AutoMigrate(&Example{}))
		require.NoError(t, db.Create(&Example{
			ID:    1,
			Title: "Product",
			Email: "user@company.com",
		}).Error)

		// Test ExprConcat
		{
			updateMap := map[string]interface{}{}
			key1, expr1 := columnTitle.KeConcat(" [Hot]") // title = CONCAT(title, ' [Hot]')
			updateMap[key1] = expr1

			result := db.Model(&Example{}).Where("id = ?", 1).UpdateColumns(updateMap)
			require.NoError(t, result.Error)
			require.Equal(t, int64(1), result.RowsAffected)

			var updated Example
			require.NoError(t, db.First(&updated, 1).Error)
			require.Equal(t, "Product [Hot]", updated.Title)
			t.Log("updated result:", neatjsons.S(updated))
		}

		// Test ExprReplace
		{
			updateMap := map[string]interface{}{}
			key1, expr1 := columnEmail.KeReplace("@company.com", "@newcompany.com") // email = REPLACE(email, '@company.com', '@newcompany.com')
			updateMap[key1] = expr1

			result := db.Model(&Example{}).Where("id = ?", 1).UpdateColumns(updateMap)
			require.NoError(t, result.Error)
			require.Equal(t, int64(1), result.RowsAffected)

			var updated Example
			require.NoError(t, db.First(&updated, 1).Error)
			require.Equal(t, "user@newcompany.com", updated.Email)
			t.Log("updated result:", neatjsons.S(updated))
		}
	})
}
