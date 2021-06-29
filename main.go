package main

import (
	"runtime"

	cli "github.com/piqba/mtss-cli/cmd/cli"
)

func init() {
	numcpu := runtime.NumCPU()
	runtime.GOMAXPROCS(numcpu) // Try to use all available CPUs.
}

func main() {
	cli.Start()
}
