package repository

import "delos/api/pond"

type repository struct{}

func NewRepository() pond.Repository {
	return &repository{}
}
