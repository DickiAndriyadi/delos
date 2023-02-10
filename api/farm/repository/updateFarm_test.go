package repository

import (
	"delos/api/model/payload"
	"fmt"
	"testing"

	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
)

func Test_farmRepository_UpdateFarm(t *testing.T) {
	t.Run("Failed Case -  Test_farmRepository.UpdateFarm", func(t *testing.T) {
		db, mock := initMock()
		defer func(db *gorm.DB) {
			_ = db.Close()
		}(db)

		var id string
		var payload payload.AddUpdateFarm

		mock.ExpectExec("UPDATE").WillReturnError(fmt.Errorf("error"))

		_, err := NewRepository().UpdateFarm(db, &payload, id)
		assert.NotNil(t, err)
	})
}
