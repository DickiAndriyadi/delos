package controller

import "delos/api/pond"

type PondController struct {
	service pond.Service
}

func NewController(s pond.Service) *PondController {
	return &PondController{service: s}
}
