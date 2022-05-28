package worker

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/thehaung/cloudflare-worker/pkg/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"io"
	"log"
	"net/http"
	"time"
)

type Service struct {
	db             *mongo.Database
	contextTimeout time.Duration
}

func NewService(db *mongo.Database, time time.Duration) IService {
	return &Service{
		db:             db,
		contextTimeout: time,
	}
}

func (s Service) GetPreviousIP(ctx context.Context) (result string, err error) {
	ctx, cancel := context.WithTimeout(ctx, s.contextTimeout)
	opts := options.FindOne().SetSort(bson.D{{"createdAt", -1}})

	findResult := bson.M{}
	defer cancel()
	if err := s.db.Collection("infos").FindOne(ctx, bson.D{}, opts).Decode(&findResult); err != nil {
		return "", err
	}

	result = fmt.Sprint(findResult["ipInfo"])

	return result, nil
}

func (s Service) GetListDNSCloudFlare(ctx context.Context) {
	//TODO implement me
	panic("implement me")
}

func (s Service) FetchDNS(ctx context.Context) (result bson.M, err error) {
	ctx, cancel := context.WithTimeout(ctx, s.contextTimeout)
	defer cancel()

	if err := s.db.Collection("infos").FindOne(ctx, bson.M{}).Decode(&result); err != nil {
		return result, err
	}

	log.Println(result)

	return result, nil
}

func (s Service) GetIPAddress(ctx context.Context) (res interface{}, err error) {
	ctx, cancel := context.WithTimeout(ctx, s.contextTimeout)
	defer cancel()

	url := config.APIGetIPUrl()
	resp, err := http.Get(url)

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(resp.Body)

	err = json.NewDecoder(resp.Body).Decode(&res)
	if err != nil {
		return "", err
	}

	return res, nil
}
