package worker

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
)

type IService interface {
	FetchDNS(ctx context.Context) (result bson.M, err error)
	GetIPAddress(ctx context.Context) (res interface{}, err error)
	GetPreviousIP(ctx context.Context) (result string, err error)
	GetListDNSCloudFlare(ctx context.Context)
}
