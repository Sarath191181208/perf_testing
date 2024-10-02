package api

import (
	"net/http"
	healthcheck "sarath/perf_testing/cmd/api/services"
	"sarath/perf_testing/cmd/api/services/users"

	"github.com/gorilla/mux"
)

func (app *Application) Routes() *mux.Router{
	mux := mux.NewRouter()

  health_check_handler := healthcheck.New(app.Logger)
  users_handler := users.New(app.Logger, app.Db)

  mux.HandleFunc("/health", health_check_handler.HandleHealthCheck)
  mux.HandleFunc("/register", users_handler.RegisterUsers).Methods(http.MethodPost)
  mux.HandleFunc("/get/{id}", users_handler.FindUser)

  return mux
}
