package service

import (
	"delos/api/farm"
	"delos/api/pond"
	"delos/config/db"
)

type service struct {
	dbManager      db.DatabaseManager
	pondRepository pond.Repository
	farmRepository farm.Repository
}

func NewService(
	dbManager db.DatabaseManager,
	pondRepository pond.Repository,
	farmRepository farm.Repository,
) pond.Service {
	return &service{
		dbManager,
		pondRepository,
		farmRepository,
	}
}
