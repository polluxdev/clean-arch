package main

import (
	"github.com/polluxdev/clean-arch/app/infrastructure"
)

func main() {
	infrastructure.Load()
	infrastructure.NewSQLHandler()
	infrastructure.Dispatch()
}
