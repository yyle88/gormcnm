package tests_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yyle88/gormcnm/internal/tests"
	"gorm.io/gorm"
)

func TestNewDBRun(t *testing.T) {
	tests.NewDBRun(t, func(db *gorm.DB) {
		var result int
		require.NoError(t, db.Raw("SELECT 1").Scan(&result).Error)
		require.Equal(t, 1, result)
	})
}
