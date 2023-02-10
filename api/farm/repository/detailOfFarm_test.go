package repository

import (
	"fmt"
	"testing"

	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
)

func Test_farmRepository_DetailOfFarm(t *testing.T) {
	t.Run("Failed Case -  Test_farmRepository.DetailOfFarm", func(t *testing.T) {
		db, mock := initMock()
		defer func(db *gorm.DB) {
			_ = db.Close()
		}(db)

		detailMap := map[string]interface{}{
			"id": "5",
		}

		mock.ExpectExec("SELECT *").WillReturnError(fmt.Errorf("error"))

		_, err := NewRepository().DetailOfFarm(db, detailMap)
		assert.NotNil(t, err)
	})
}
