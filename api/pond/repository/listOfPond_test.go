package repository

import (
	"delos/api/pond/filter"
	"fmt"
	"testing"

	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
)

func Test_farmRepository_ListOfPond(t *testing.T) {
	t.Run("Failed Case -  Test_farmRepository.ListOfPond", func(t *testing.T) {
		db, mock := initMock()
		defer func(db *gorm.DB) {
			_ = db.Close()
		}(db)

		mock.ExpectExec("SELECT *").WillReturnError(fmt.Errorf("error"))

		var f *filter.FilterListOfPond

		_, err := NewRepository().ListOfPond(db, f)
		assert.NotNil(t, err)
	})
}
