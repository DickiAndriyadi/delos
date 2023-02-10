package service

import (
	"testing"

	"delos/api/farm"

	farmMock "delos/mocks/api/farm"
	pondMock "delos/mocks/api/pond"
	databaseMock "delos/mocks/config/db"

	"github.com/jinzhu/gorm"
)

type TestingService struct {
	Service        farm.Service
	dbManager      *databaseMock.DatabaseManager
	farmRepository *farmMock.Repository
	pondRepository *pondMock.Repository
}

var (
	database = gorm.DB{}
)

// Initialize Go Mock
func (s *TestingService) Initialize(t *testing.T) *TestingService {

	s.dbManager = &databaseMock.DatabaseManager{}
	s.dbManager.On("GetDB").Return(&database)
	s.farmRepository = &farmMock.Repository{}
	s.pondRepository = &pondMock.Repository{}
	s.Service = NewService(
		s.dbManager,
		s.farmRepository,
		s.pondRepository,
	)

	return s
}
