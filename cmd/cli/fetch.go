package main

import (
	"log"

	"github.com/spf13/cobra"

	mdomain "github.com/piqba/mtss-cli/pkg/mtss/domain"
	mhandler "github.com/piqba/mtss-cli/pkg/mtss/handlers"
	mservice "github.com/piqba/mtss-cli/pkg/mtss/service"
)

var fetchCmd = &cobra.Command{
	Use:   "fetch",
	Short: "Fetch data from API rest client",
	Run: func(cmd *cobra.Command, args []string) {
		limit, err := cmd.Flags().GetInt32(flagLimit)
		if err != nil {
			log.Fatalf(err.Error())
		}
		mrepo := mdomain.NewMtssRepository()
		ms := mservice.NewCustomerService(mrepo)
		mhs := mhandler.NewMtssHandler(ms)
		mhs.FetchAllFromAPI(limit)
	},
}

func init() {
	fetchCmd.Flags().Int32(flagLimit, 10, "select a limit of jobs to fetch")
	rootCmd.AddCommand(fetchCmd)

}
