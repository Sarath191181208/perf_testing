package api

import (
	"sarath/perf_testing/internal/data"
	"sarath/perf_testing/internal/logger"
)



type Application struct{
  Logger logger.ApplicationLogger
  Db *data.Models
}

