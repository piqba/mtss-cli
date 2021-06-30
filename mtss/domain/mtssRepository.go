package mtss

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	DB        = "mtss-job"
	JOBS      = "jobs"
	SIGANTURE = "mtss:job"
)

type MtssRepositoryAPI struct {
	clientMgo *mongo.Client
	clientRdb *redis.Client
}

func NewMtssRepository(url string, clients ...interface{}) MtssRepositoryAPI {
	var mtssRepositoryAPI MtssRepositoryAPI
	for _, c := range clients {
		switch c.(type) {
		case *mongo.Client:
			mtssRepositoryAPI.clientMgo = c.(*mongo.Client)
		case *redis.Client:
			mtssRepositoryAPI.clientRdb = c.(*redis.Client)
		}
	}
	return mtssRepositoryAPI
}

func (m MtssRepositoryAPI) FetchAllFromAPI(url string) ([]MTSS, error) {

	var mtssArray []MTSS

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Get(url)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer func() {
		err := resp.Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	err = json.Unmarshal(body, &mtssArray)
	if err != nil {
		log.Fatal("Unmarshal", err)
		return nil, err
	}

	return mtssArray, nil
}

//CreateOne - Insert a new document in the collection.
func (m MtssRepositoryAPI) CreateOne(engine string, mtss MTSS) error {

	switch engine {
	case "redis":
		key := fmt.Sprintf("%s:%d", SIGANTURE, mtss.ID)
		value, _ := mtss.ToMAP()
		if _, err := m.clientRdb.HSet(context.TODO(), key, value).Result(); err != nil {
			log.Fatal("create: redis error: %w", err)
		}
		m.clientRdb.Expire(context.TODO(), key, 24*time.Hour)
	case "mongo":
		//Create a handle to the respective collection in the database.
		collection := m.clientMgo.Database(DB).Collection(JOBS)
		//Perform InsertOne operation & validate against the error.
		_, err := collection.InsertOne(context.TODO(), mtss)
		if err != nil {
			return err
		}
	}
	return nil
}
