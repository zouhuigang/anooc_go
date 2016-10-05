package model

import (
	"github.com/polaris1119/logger"
	"golang.org/x/net/context"
	"os"
)

func GetLogger(ctx context.Context) *logger.Logger {
	if ctx == nil {
		return logger.New(os.Stdout)
	}

	_logger, ok := ctx.Value("logger").(*logger.Logger)
	if ok {
		return _logger
	}

	return logger.New(os.Stdout)
}
