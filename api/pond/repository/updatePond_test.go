package repository

import (
	"delos/api/model/payload"
	"fmt"
	"testing"

	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
)

func Test_farmRepository_UpdatePond(t *testing.T) {
	t.Run("Failed Case -  Test_farmRepository.UpdatePond", func(t *testing.T) {
		db, mock := initMock()
		defer func(db *gorm.DB) {
			_ = db.Close()
		}(db)

		var id string
		var payload payload.AddUpdatePond

		mock.ExpectExec("UPDATE").WillReturnError(fmt.Errorf("error"))

		_, err := NewRepository().UpdatePond(db, &payload, id)
		assert.NotNil(t, err)
	})
}
