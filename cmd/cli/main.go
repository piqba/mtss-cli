package main

import (
	"runtime"

	"github.com/joho/godotenv"
	"github.com/piqba/mtss-cli/pkg/logger"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		logger.LogError(err.Error())
	}
	numcpu := runtime.NumCPU()
	runtime.GOMAXPROCS(numcpu) // Try to use all available CPUs.
}

func main() {
	Execute()
}
