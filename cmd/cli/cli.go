package cli

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	mtssDomain "github.com/piqba/mtss-cli/mtss/domain"
	mtssHandler "github.com/piqba/mtss-cli/mtss/handlers"
	mtssService "github.com/piqba/mtss-cli/mtss/service"
)

// func init(){
// 	// 	flag.StringVar(&source, "source", source, "--source <mtss> Data source to ingest on DB")
// // 	flag.BoolVar(&batch, "batch", batch, "--batch <true> For fetch daily jobs or one by one jobs")
// // 	flag.Parse()
// }
func Start() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	endpointURI := os.Getenv("ENDPOINT_MTSS")
	var mtssRepository mtssDomain.MtssRepository
	// pass to flag
	var engine = "mongo"

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

	mh.InsertOnDbFromAPI("mongo", endpointURI, 2)

}
