package cli

import (
	"flag"
	"os"

	"github.com/joho/godotenv"
	"github.com/piqba/mtss-cli/internal"
	mtssDomain "github.com/piqba/mtss-cli/mtss/domain"
	mtssHandler "github.com/piqba/mtss-cli/mtss/handlers"
	mtssService "github.com/piqba/mtss-cli/mtss/service"
)

var (
	engine     = "mongo"
	limit      = 10
	ErrLoadEnv = internal.NewError("dotenv: Error loading .env file")
)

func init() {
	flag.StringVar(&engine, "engine", engine, "--engine <mongo> Data engine to ingest on DB")
	flag.IntVar(&limit, "limit", limit, "--limit <10> Data limit to ingest on DB")
	flag.Parse()
}
func Start() {
	err := godotenv.Load()
	if err != nil {
		internal.LogError(ErrLoadEnv.Error())
	}
	endpointURI := os.Getenv("ENDPOINT_MTSS")
	var mtssRepository mtssDomain.MtssRepository

	switch engine {
	case "all":
		mongoDbClient := GetMongoDbClient()
		redisDbClient := GetRedisDbClient()
		mtssRepository = mtssDomain.NewMtssRepository(
			endpointURI,
			mongoDbClient,
			redisDbClient,
		)
	case "mongo":
		mongoDbClient := GetMongoDbClient()
		mtssRepository = mtssDomain.NewMtssRepository(
			endpointURI,
			mongoDbClient,
		)
	case "redis":
		redisDbClient := GetRedisDbClient()
		mtssRepository = mtssDomain.NewMtssRepository(
			endpointURI,
			redisDbClient,
		)
	}

	mh := mtssHandler.MtssHandler{
		Service: mtssService.NewCustomerService(mtssRepository),
	}

	mh.InsertOnDbFromAPI(engine, endpointURI, limit)

}
