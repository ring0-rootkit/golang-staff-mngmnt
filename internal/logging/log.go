package logging

import (
	"fmt"
	"log"
	"os"
)

var loggers map[string]*log.Logger = make(map[string]*log.Logger, 5)

func GetFor(serviceName string) *log.Logger {
	if loggers[serviceName] == nil {
		loggers[serviceName] = log.New(os.Stdout, fmt.Sprintf("[%s service]", serviceName), log.LstdFlags)
	}
	return loggers[serviceName]
}
