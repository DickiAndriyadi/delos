package service

import (
	"delos/api/farm"
	"delos/api/pond"
	"delos/config/db"
)

type service struct {
	dbManager      db.DatabaseManager
	farmRepository farm.Repository
	pondRepository pond.Repository
}

func NewService(
	dbManager db.DatabaseManager,
	farmRepository farm.Repository,
	pondRepository pond.Repository,
) farm.Service {
	return &service{
		dbManager,
		farmRepository,
		pondRepository,
	}
}
