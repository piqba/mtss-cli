package mtss

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/jmoiron/sqlx"
	"github.com/piqba/mtss-cli/pkg/logger"
	"github.com/piqba/mtss-go"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	DB         = "mtss-job"
	JOBS       = "jobs"
	SIGANTURE  = "mtss:job"
	REDIS      = "redis"
	POSTGRESQL = "postgres"
	MONGODB    = "mongodb"
)

type MtssRepository struct {
	clientMgo *mongo.Client
	clientRdb *redis.Client
	clientPgx *sqlx.DB
}

func NewMtssRepository(clients ...interface{}) MtssRepository {
	var mtssrepo MtssRepository
	for _, c := range clients {
		switch c := c.(type) {
		case *mongo.Client:
			mtssrepo.clientMgo = c
		case *redis.Client:
			mtssrepo.clientRdb = c
		case *sqlx.DB:
			mtssrepo.clientPgx = c
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
	if limit > 0 {

		return jobs[:limit], nil
	}
	return jobs, nil

}

//CreateOne - Insert a new document in the collection.
func (m MtssRepository) CreateOne(engine string, job mtss.Mtss) error {

	switch engine {
	case REDIS:
		key := fmt.Sprintf("%s:%d", SIGANTURE, job.ID)
		value, _ := job.ToMAP()
		if _, err := m.clientRdb.HSet(context.TODO(), key, value).Result(); err != nil {
			log.Fatal("create: redis error: %w", err)
		}
		m.clientRdb.Expire(context.TODO(), key, 24*time.Hour)
	case MONGODB:
		//Create a handle to the respective collection in the database.
		collection := m.clientMgo.Database(DB).Collection(JOBS)
		//Perform InsertOne operation & validate against the error.
		_, err := collection.InsertOne(context.TODO(), job)
		if err != nil {
			return err
		}
	case POSTGRESQL:
		tx := m.clientPgx.MustBegin()
		query := "INSERT INTO mtss_jobs (id, created_at, job) VALUES ($1, $2, $3) ON CONFLICT (id) DO UPDATE SET job = $3"
		jobToJSON, err := json.Marshal(job)
		if err != nil {
			logger.LogError(err.Error())
		}
		tx.MustExec(
			query,
			job.ID,
			time.Now(),
			string(jobToJSON),
		)
		tx.Commit()
	}
	return nil
}

func (m MtssRepository) SendDataToRedisStream(rdb *redis.Client, key string, value map[string]interface{}) error {

	err := rdb.XAdd(context.TODO(), &redis.XAddArgs{
		Stream: key,
		Values: value,
	}).Err()
	if err != nil {
		return err
	}
	return nil
}

func (m MtssRepository) GetMtssJobs(limit, offset int) ([]mtss.Mtss, error) {
	var jobs []mtss.Mtss
	query := ""
	if limit == 0 && offset == 0 {

		query = "select job from mtss_jobs"
	} else {
		query = fmt.Sprintf("select job from mtss_jobs limit %d offset %d;", limit, offset)
	}
	rows, err := m.clientPgx.Queryx(query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var data []byte
		var job mtss.Mtss
		err = rows.Scan(&data)
		json.Unmarshal(data, &job)
		if err != nil {
			return nil, err
		}
		jobs = append(jobs, job)
	}
	return jobs, nil
}
