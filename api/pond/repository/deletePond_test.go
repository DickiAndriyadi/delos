package repository

import (
	"fmt"
	"testing"

	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
)

func Test_farmRepository_DeletePond(t *testing.T) {
	t.Run("Failed Case -  Test_farmRepository.DeletePond", func(t *testing.T) {
		db, mock := initMock()
		defer func(db *gorm.DB) {
			_ = db.Close()
		}(db)

		deletePondMap := map[string]interface{}{
			"id": "1",
		}

		mock.ExpectExec("DELETE *").WillReturnError(fmt.Errorf("error"))

		err := NewRepository().DeletePond(db, deletePondMap)
		assert.NotNil(t, err)
	})
}
