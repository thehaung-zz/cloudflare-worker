package worker

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"io"
)

type IService interface {
	FetchDNS(ctx context.Context) (result bson.M, err error)
	GetIPAddress(ctx context.Context) (res interface{}, err error)
	GetPreviousIP(ctx context.Context) (result string, err error)
	GetListDNSCloudFlare(ctx context.Context) (res io.ReadCloser, err error)
	UpdateAllIPPublic(ctx context.Context, ip string) (res interface{}, err error)
}
