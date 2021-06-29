package main

import (
	"runtime"

	cli "github.com/piqba/mtss-cli/cmd/cli"
	"github.com/piqba/mtss-cli/internal"
)

func init() {
	numcpu := runtime.NumCPU()
	runtime.GOMAXPROCS(numcpu) // Try to use all available CPUs.
}

func main() {
	internal.LogInfo("Starting the applications ...")
	cli.Start()
}
