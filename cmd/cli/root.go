package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

const (
	flagLimit  = "limit"
	flagEngine = "engine"
	REDIS      = "redis"
	POSTGRESQL = "postgres"
	MONGODB    = "mongodb"
	ALL        = "all"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "mtssctl",
	Short: "A brief description of your application",
	Long:  `this is Mtss ctl`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		//TODO change this with tomatto-logger
		fmt.Println(err)
		os.Exit(1)
	}

}
func init() {

}
