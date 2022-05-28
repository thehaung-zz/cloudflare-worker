package model

type CloudFlare struct {
	Type string `json:"type" bson:"type"`

	Name string `json:"name" bson:"name"`

	Content string `json:"content" bson:"content"`

	TTL int8 `json:"ttl" bson:"ttl"`

	Proxied bool `json:"proxied" bson:"proxied"`
}

func NewCloudFlare(typeCloudFlare, name, content string, ttl int8, proxied bool) *CloudFlare {
	return &CloudFlare{
		Type:    typeCloudFlare,
		Name:    name,
		Content: content,
		TTL:     ttl,
		Proxied: proxied,
	}
}
