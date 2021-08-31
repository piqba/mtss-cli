package cli

import (
	"log"

	mdomain "github.com/piqba/mtss-cli/pkg/mtss/domain"
	mhandler "github.com/piqba/mtss-cli/pkg/mtss/handlers"
	mservice "github.com/piqba/mtss-cli/pkg/mtss/service"
	"github.com/piqba/mtss-cli/pkg/redisdb"
	"github.com/spf13/cobra"
)

var insertCmd = &cobra.Command{
	Use:   "insert",
	Short: "Insert Data from API rest client",
	Run: func(cmd *cobra.Command, args []string) {
		engine, err := cmd.Flags().GetString(flagEngine)
		if err != nil {
			log.Fatalf(err.Error())
		}
		limit, err := cmd.Flags().GetInt32(flagLimit)
		if err != nil {
			log.Fatalf(err.Error())
		}

		switch engine {
		case REDIS:
			redisDbClient := redisdb.GetRedisDbClient()
			mrepo := mdomain.NewMtssRepository(redisDbClient)
			ms := mservice.NewCustomerService(mrepo)
			mhs := mhandler.NewMtssHandler(ms)
			mhs.InsertOnDbFromAPI(engine, limit)
		}
	},
}

func init() {
	insertCmd.Flags().String(flagEngine, "redis", "select a engine for insert data")
	insertCmd.Flags().Int32(flagLimit, 10, "select a limit of jobs to fetch")

	rootCmd.AddCommand(insertCmd)

}
