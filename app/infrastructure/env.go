package infrastructure

import (
	"github.com/joho/godotenv"
	"github.com/polluxdev/clean-arch/app/usecases"
)

func Load(logger usecases.Logger) {
	err := godotenv.Load()
	if err != nil {
		logger.LogError("%s", err)
	}
}
