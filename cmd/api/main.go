package main

import (
	"runtime"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joho/godotenv"
	"github.com/piqba/mtss-cli/pkg/constants"
	"github.com/piqba/mtss-cli/pkg/logger"
	mdomain "github.com/piqba/mtss-cli/pkg/mtss/domain"
	mhandler "github.com/piqba/mtss-cli/pkg/mtss/handlers"
	mservice "github.com/piqba/mtss-cli/pkg/mtss/service"
	"github.com/piqba/mtss-cli/pkg/pgsql"
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
	app := fiber.New()
	app.Use(cors.New())
	app.Use(recover.New())
	pgxDbclient, err := pgsql.PostgreSQLConnection()
	if err != nil {
		logger.LogError(err.Error())
	}
	pgxDbclient.MustExec(constants.SchemaMTTS)
	mrepo := mdomain.NewMtssRepository(pgxDbclient)
	ms := mservice.NewCustomerService(mrepo)
	h := mhandler.New(ms)
	app.Get("/mtss/jobs", h.GetMtssJobs)

	logger.LogError(app.Listen(":3000").Error())
}
