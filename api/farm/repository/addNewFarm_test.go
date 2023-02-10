package repository

import (
	"delos/api/model"
	"fmt"
	"testing"

	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
)

func Test_farmRepository_AddNewFarm(t *testing.T) {
	t.Run("Failed Case -  Test_farmRepository.AddNewFarm", func(t *testing.T) {
		db, mock := initMock()
		defer func(db *gorm.DB) {
			_ = db.Close()
		}(db)

		mock.ExpectQuery("INSERT INTO").WillReturnError(fmt.Errorf("error"))

		_, err := NewRepository().AddNewFarm(db, &model.Farm{})
		assert.NotNil(t, err)
	})
}
