package controller

import (
	"delos/api/counter/controller"
	"delos/api/model/payload"
	"delos/api/model/response"
	"net/http"

	"github.com/labstack/echo"
)

func (ctrl *PondController) AddNewPond(c echo.Context) error {
	controller.CountEndpointHits(c)

	var (
		result response.DataResponse
	)

	payload := new(payload.AddUpdatePond)

	if err := c.Bind(payload); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	payload.Req = "Add"

	if err := payload.Validate(); err != nil {
		return echo.NewHTTPError(
			http.StatusBadRequest,
			err,
		)
	}

	data, err := ctrl.service.AddNewPond(*payload)
	if err != nil {
		return echo.NewHTTPError(
			err.ErrCode,
			err.Message.Error(),
		)
	}

	result.Data = data

	return c.JSON(http.StatusCreated, result)
}
