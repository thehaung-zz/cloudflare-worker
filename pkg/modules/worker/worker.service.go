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

func (s Service) GetListDNSCloudFlare(ctx context.Context) (res io.ReadCloser, err error) {
	ctx, cancel := context.WithTimeout(ctx, s.contextTimeout)
	defer cancel()
	url := fmt.Sprintf("%s/%s/%s", config.GetCloudFlareAPIUrl(), config.GetCloudFlareZoneID(), "dns_records?type=A")
	log.Println(url)
	request, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return nil, err
	}

	request.Header.Set("X-Auth-Email", config.GetCloudFlareEmail())
	request.Header.Set("X-Auth-Key", config.GetCloudFlareAPIToken())

	client := &http.Client{}
	resp, err := client.Do(request)

	return resp.Body, nil
}

func (s Service) FetchDNS(ctx context.Context) (result bson.M, err error) {
	ctx, cancel := context.WithTimeout(ctx, s.contextTimeout)
	defer cancel()

	if err := s.db.Collection("infos").FindOne(ctx, bson.M{}).Decode(&result); err != nil {
		return result, err
	}

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
