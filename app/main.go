package main

import (
	"github.com/polluxdev/clean-arch/app/infrastructure"
)

func main() {
	logger := infrastructure.NewLogger()

	infrastructure.Load(logger)

	sqlHandler, err := infrastructure.NewSQLHandler()
	if err != nil {
		logger.LogError("%s", err)
	}

	infrastructure.Dispatch(logger, sqlHandler)
}
