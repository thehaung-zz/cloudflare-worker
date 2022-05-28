package model

type Info struct {
	IpInfo    string `json:"ipInfo" bson:"ipInfo"`
	CreatedAt string `json:"createdAt" bson:"createdAt"`
}

func NewInfo(ipInfo string, createdAt string) *Info {
	return &Info{
		IpInfo:    ipInfo,
		CreatedAt: createdAt,
	}
}
