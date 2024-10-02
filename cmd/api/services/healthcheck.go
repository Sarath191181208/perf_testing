package healthcheck

import (
	"net/http"
	"sarath/perf_testing/internal/json"
	"sarath/perf_testing/internal/logger"
)


type Handler struct{
  Logger logger.ApplicationLogger
}

func New(logger logger.ApplicationLogger) *Handler{
  return &Handler{
    Logger: logger,
  }
}

func (app *Handler) HandleHealthCheck(w http.ResponseWriter, r *http.Request) {
  app.Logger.Log("Got a health check request")
	data := json.Envelope{
    "status" : "OK",
  }
	json.WriteJsonToResponseWriter(data, w)
}
