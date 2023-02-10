package application

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"delos/config/db"

	"delos/api/counter/controller"
	_farmController "delos/api/farm/controller"
	_farmRepository "delos/api/farm/repository"
	_farmService "delos/api/farm/service"

	_pondController "delos/api/pond/controller"
	_pondRepository "delos/api/pond/repository"
	_pondService "delos/api/pond/service"
)

type App struct {
	E         *echo.Echo
	DBManager db.DatabaseManager
}

func (app *App) InitializeDatabase(dsn string, dbConnection string) {
	err := app.DBManager.Initialize(dsn, dbConnection)

	if err != nil {
		panic(err)
	}
}

func (app *App) Initialize() {
	app.initializeRoutes()
}

func (app *App) Start(addr string) {
	app.E.Start(addr)
}

func (app *App) initializeRoutes() {
	app.E.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAcceptEncoding},
	}))
	app.E.Use(loggerCfg())
	app.E.HTTPErrorHandler = defaultErr()

	farmRepository := _farmRepository.NewRepository()
	pondRepository := _pondRepository.NewRepository()

	farmService := _farmService.NewService(
		app.DBManager,
		farmRepository,
		pondRepository,
	)
	pondService := _pondService.NewService(
		app.DBManager,
		pondRepository,
		farmRepository,
	)

	farmController := _farmController.NewController(farmService)
	pondController := _pondController.NewController(pondService)

	app.E.POST("/v1/farms", farmController.AddNewFarm).Name = "Add New Farm"
	app.E.GET("/v1/farms", farmController.ListOfFarm).Name = "List Of Farm"
	app.E.GET("/v1/farms/:id", farmController.DetailOfFarm).Name = "Detail Of Farm"
	app.E.PUT("/v1/farms/:id", farmController.UpdateFarm).Name = "Update Farm"
	app.E.DELETE("/v1/farms/:id", farmController.DeleteFarm).Name = "Delete Farm"

	app.E.POST("/v1/ponds", pondController.AddNewPond).Name = "Add New Pond"
	app.E.GET("/v1/ponds", pondController.ListOfPond).Name = "List Of Pond"
	app.E.GET("/v1/ponds/:id", pondController.DetailOfPond).Name = "Detail Of Pond"
	app.E.PUT("/v1/ponds/:id", pondController.UpdatePond).Name = "Update Pond"
	app.E.DELETE("/v1/ponds/:id", pondController.DeletePond).Name = "Delete Pond"

	app.E.GET("/counter", controller.Counter)

}

// defaultErr is used to mask an error by echo
func defaultErr() func(err error, c echo.Context) {
	return func(err error, c echo.Context) {
		report, ok := err.(*echo.HTTPError)
		if !ok {
			report = echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		cerr := map[string]interface{}{
			"errorCode": report.Code,
			"message":   report.Message,
		}
		c.JSON(report.Code, cerr)
	}
}

// loggerCfg is used to show a log access
func loggerCfg() echo.MiddlewareFunc {
	return middleware.LoggerWithConfig(
		middleware.LoggerConfig{
			Format:           "${time_custom} | ${status}  ⇨ ${method} ${protocol} ${host}${uri} ${latency_human} \n",
			CustomTimeFormat: "15:04:05",
		})
}
