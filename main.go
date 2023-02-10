package main

import (
	"log"
	"os"

	"delos/application"
	"delos/config/db"

	"github.com/joho/godotenv"
	"github.com/labstack/echo"
)

func main() {
	app := &application.App{
		E:         echo.New(),
		DBManager: db.NewDatabaseManager(),
	}

	_ = godotenv.Load(".env")

	if envCheck() {
		user := os.Getenv("DB_USERNAME")
		port := os.Getenv("DB_PORT")
		addr := os.Getenv("DB_ADDR")
		pswd := os.Getenv("DB_PASSWORD")
		dbname := os.Getenv("DB_NAME")

		// assigning a main connection databases variable with db env structure
		dsn := user + ":" + pswd + "@tcp(" + addr + ":" + port + ")/" + dbname + "?parseTime=true"
		app.InitializeDatabase(
			dsn,
			os.Getenv("DB_CONNECTION"),
		)

		defer app.DBManager.GetDB().Close()
	} else {
		log.Println("error env")
		os.Exit(0)
	}

	app.Initialize()
	app.Start(":" + os.Getenv("PORT"))

}

func envCheck() bool {
	s := true
	env := []string{
		// "ENVIRONMENT",
		// "IP_ADDR",
		"PORT",

		"DB_ADDR",
		"DB_PORT",
		"DB_USERNAME",
		// "DB_PASSWORD",
		"DB_NAME",
		"DB_CONNECTION",
	}

	for _, e := range env {
		if os.Getenv(e) == "" {
			s = false
		}
	}
	return s
}
