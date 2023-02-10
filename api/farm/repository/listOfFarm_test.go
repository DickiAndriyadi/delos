package repository

import (
	"fmt"
	"testing"

	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
)

func Test_farmRepository_ListOfFarm(t *testing.T) {
	t.Run("Failed Case -  Test_farmRepository.ListOfFarm", func(t *testing.T) {
		db, mock := initMock()
		defer func(db *gorm.DB) {
			_ = db.Close()
		}(db)

		mock.ExpectExec("SELECT *").WillReturnError(fmt.Errorf("error"))

		_, err := NewRepository().ListOfFarm(db)
		assert.NotNil(t, err)
	})
}
