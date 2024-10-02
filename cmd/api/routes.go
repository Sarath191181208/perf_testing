package api

import (
	"net/http"
	healthcheck "sarath/perf_testing/cmd/api/services"
	"sarath/perf_testing/cmd/api/services/users"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func (app *Application) Routes() *mux.Router{
	mux := mux.NewRouter()

  health_check_handler := healthcheck.New(app.Logger)
  users_handler := users.New(app.Logger, app.Db)

  mux.HandleFunc("/health", health_check_handler.HandleHealthCheck)
  mux.HandleFunc("/register", users_handler.RegisterUsers).Methods(http.MethodPost)
  mux.HandleFunc("/get/{id}", users_handler.FindUser)
  mux.Handle("/metrics", promhttp.Handler())

  return mux
}
