package controller

import "delos/api/farm"

type FarmController struct {
	service farm.Service
}

func NewController(s farm.Service) *FarmController {
	return &FarmController{service: s}
}
