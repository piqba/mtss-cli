package cli

import (
	"flag"
	"log"
	"os"

	"github.com/joho/godotenv"
	mtssDomain "github.com/piqba/mtss-cli/mtss/domain"
	mtssHandler "github.com/piqba/mtss-cli/mtss/handlers"
	mtssService "github.com/piqba/mtss-cli/mtss/service"
)

var (
	engine = "mongo"
	limit  = 10
)

func init() {
	flag.StringVar(&engine, "engine", engine, "--engine <mongo> Data engine to ingest on DB")
	flag.IntVar(&limit, "limit", limit, "--limit <10> Data limit to ingest on DB")
	flag.Parse()
}
func Start() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	endpointURI := os.Getenv("ENDPOINT_MTSS")
	var mtssRepository mtssDomain.MtssRepository

	switch engine {
	case "all":
		mongoDbClient := GetMongoDbClient()
		redisDbClient := GetRedisDbClient()
		mtssRepository = mtssDomain.NewMtssRepositoryAll(
			endpointURI,
			mongoDbClient,
			redisDbClient,
		)
	case "mongo":
		mongoDbClient := GetMongoDbClient()
		mtssRepository = mtssDomain.NewMtssRepositoryWithOutRedis(
			endpointURI,
			mongoDbClient,
		)
	case "redis":
		redisDbClient := GetRedisDbClient()
		mtssRepository = mtssDomain.NewMtssRepositoryWithOutMongo(
			endpointURI,
			redisDbClient,
		)
	}

	mh := mtssHandler.MtssHandler{
		Service: mtssService.NewCustomerService(mtssRepository),
	}

	mh.InsertOnDbFromAPI(engine, endpointURI, limit)

}
