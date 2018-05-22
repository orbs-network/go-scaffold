package logger

import (
	"gopkg.in/birkirb/loggers.v1/mappers/stdlib"
	"log"
	"os"
)

func DefaultLogger(prefix string) Interface {
	logger := log.New(os.Stderr, prefix + "\t", log.Ldate|log.Ltime)
	return stdlib.NewLogger(logger)
}