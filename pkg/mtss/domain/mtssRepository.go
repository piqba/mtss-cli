package mtss

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/piqba/mtss-go"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	DB        = "mtss-job"
	JOBS      = "jobs"
	SIGANTURE = "mtss:job"
)

type MtssRepository struct {
	clientMgo *mongo.Client
	clientRdb *redis.Client
}

func NewMtssRepository(clients ...interface{}) MtssRepository {
	var mtssrepo MtssRepository
	for _, c := range clients {
		switch c := c.(type) {
		case *mongo.Client:
			mtssrepo.clientMgo = c
		case *redis.Client:
			mtssrepo.clientRdb = c
		}
	}
	return mtssrepo
}

func (m MtssRepository) FetchAllFromAPI(limit int32) ([]mtss.Mtss, error) {
	baseURL := mtss.URL_BASE
	skipVerify := true
	client := mtss.NewAPIClient(
		baseURL,
		skipVerify,
		nil, //custom http client, defaults to http.DefaultClient
		nil, //io.Writer (os.Stdout) to output debug messages
	)
	jobs, err := client.MtssJobs(context.TODO())
	if err != nil {
		log.Fatalf(err.Error())
	}

	return jobs[:limit], nil
}

//CreateOne - Insert a new document in the collection.
func (m MtssRepository) CreateOne(engine string, job mtss.Mtss) error {

	switch engine {
	case "redis":
		key := fmt.Sprintf("%s:%d", SIGANTURE, job.ID)
		value, _ := job.ToMAP()
		if _, err := m.clientRdb.HSet(context.TODO(), key, value).Result(); err != nil {
			log.Fatal("create: redis error: %w", err)
		}
		m.clientRdb.Expire(context.TODO(), key, 24*time.Hour)
	case "mongo":
		//Create a handle to the respective collection in the database.
		collection := m.clientMgo.Database(DB).Collection(JOBS)
		//Perform InsertOne operation & validate against the error.
		_, err := collection.InsertOne(context.TODO(), job)
		if err != nil {
			return err
		}
	}
	return nil
}
