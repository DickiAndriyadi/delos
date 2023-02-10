package controller

import (
	"delos/api/counter/controller"
	"delos/api/model/response"
	"delos/api/pond/filter"
	"net/http"

	"github.com/labstack/echo"
)

func (ctrl *PondController) ListOfPond(c echo.Context) error {
	controller.CountEndpointHits(c)

	var (
		result response.DataResponse
	)

	filter := getFilter(c)

	data, err := ctrl.service.ListOfPond(&filter)
	if err != nil {
		return echo.NewHTTPError(
			err.ErrCode,
			err.Message.Error(),
		)
	}

	result.Data = map[string]interface{}{
		"items": data,
	}

	return c.JSON(http.StatusOK, result)
}

func getFilter(c echo.Context) filter.FilterListOfPond {
	return filter.FilterListOfPond{
		FarmID: c.QueryParam("farm_id"),
	}
}
