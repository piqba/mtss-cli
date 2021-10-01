package cli

import (
	"log"

	"github.com/piqba/mtss-cli/pkg/logger"
	mdomain "github.com/piqba/mtss-cli/pkg/mtss/domain"
	mhandler "github.com/piqba/mtss-cli/pkg/mtss/handlers"
	mservice "github.com/piqba/mtss-cli/pkg/mtss/service"
	"github.com/piqba/mtss-cli/pkg/pgsql"
	"github.com/piqba/mtss-cli/pkg/redisdb"
	"github.com/spf13/cobra"
)

var (
	schema = `
	CREATE TABLE IF NOT EXISTS mtss_jobs
	(
		id         int PRIMARY KEY        NOT NULL,
		created_at TIMESTAMP DEFAULT now() NOT NULL,
		job  JSON
	);
	`
)
var insertCmd = &cobra.Command{
	Use:   "insert",
	Short: "Fetch data from API rest client to an specific db (redis/mongodb[unimplemented]/postgresql)",
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
		case POSTGRESQL:
			pgxDbclient, err := pgsql.PostgreSQLConnection()
			if err != nil {
				logger.LogError(err.Error())
			}
			pgxDbclient.MustExec(schema)

			mrepo := mdomain.NewMtssRepository(pgxDbclient)
			ms := mservice.NewCustomerService(mrepo)
			mhs := mhandler.NewMtssHandler(ms)
			mhs.InsertOnDbFromAPI(engine, limit)
		}
	},
}

func init() {
	insertCmd.Flags().String(flagEngine, "postgres", "select a engine for insert data")
	insertCmd.Flags().Int32(flagLimit, 10, "select a limit of jobs to fetch")

	rootCmd.AddCommand(insertCmd)

}
