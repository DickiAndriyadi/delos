package repository

import "delos/api/farm"

type repository struct{}

func NewRepository() farm.Repository {
	return &repository{}
}
