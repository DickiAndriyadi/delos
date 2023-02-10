package repository

import (
	"fmt"
	"testing"

	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
)

func Test_farmRepository_DeleteFarm(t *testing.T) {
	t.Run("Failed Case -  Test_farmRepository.DeleteFarm", func(t *testing.T) {
		db, mock := initMock()
		defer func(db *gorm.DB) {
			_ = db.Close()
		}(db)

		var id string

		mock.ExpectExec("DELETE *").WillReturnError(fmt.Errorf("error"))

		err := NewRepository().DeleteFarm(db, id)
		assert.NotNil(t, err)
	})
}
