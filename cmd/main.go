package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"sarath/perf_testing/cmd/api"
	"sarath/perf_testing/internal/data"
	"sarath/perf_testing/internal/logger"

	_ "github.com/lib/pq"
)

func main() {
	PORT := os.Getenv("PORT")
	db_dsn := os.Getenv("DB_DSN")

  // Creating a Logger
	app_logger := &logger.SysoutLogger{
		Logger: log.New(os.Stdout, "", 0),
	}

  // Connecting to db 
	db_conn, err := OpenDB(db_dsn, app_logger)
	if err != nil {
		app_logger.Log(fmt.Sprint("Can't establish a db conn because of", err.Error()))
		return
	}
	app_logger.Log("Connected to DB")
	defer db_conn.Close()

  // Creating a application struct
	app := api.Application{
		Logger: app_logger,
		Db:     data.New(db_conn),
	}

  // Creating a server
	server := http.Server{
		Addr:    fmt.Sprint(":", PORT),
		Handler: app.Routes(),
	}

	app.Logger.Log(fmt.Sprint("The server is running on port:", PORT))

	err = server.ListenAndServe()
	if err != nil {
		app.Logger.Log(fmt.Sprintf("The server failed because of : %v", err))
	}

	app.Logger.Log("The application is shutdown now.")
}

func OpenDB(db_dsn string, logger logger.ApplicationLogger) (*sql.DB, error) {
	db, err := sql.Open("postgres", db_dsn)
	if err != nil {
		logger.Log(err.Error())
		return nil, err
	}
	return db, nil
}
