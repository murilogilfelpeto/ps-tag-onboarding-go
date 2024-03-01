package configuration

import (
	logger "github.com/sirupsen/logrus"
	"os"
)

func init() {
	logger.SetFormatter(&logger.JSONFormatter{})
	logger.SetOutput(os.Stdout)
	logger.SetReportCaller(true)
}
