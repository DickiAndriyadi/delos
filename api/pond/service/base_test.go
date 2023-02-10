package service

import (
	"testing"

	"delos/api/pond"
	farmMock "delos/mocks/api/farm"
	pondMock "delos/mocks/api/pond"

	databaseMock "delos/mocks/config/db"

	"github.com/jinzhu/gorm"
)

type TestingService struct {
	Service        pond.Service
	dbManager      *databaseMock.DatabaseManager
	pondRepository *pondMock.Repository
	farmRepository *farmMock.Repository
}

var (
	database = gorm.DB{}
)

// Initialize Go Mock
func (s *TestingService) Initialize(t *testing.T) *TestingService {

	s.dbManager = &databaseMock.DatabaseManager{}
	s.dbManager.On("GetDB").Return(&database)
	s.pondRepository = &pondMock.Repository{}
	s.farmRepository = &farmMock.Repository{}
	s.Service = NewService(
		s.dbManager,
		s.pondRepository,
		s.farmRepository,
	)

	return s
}
