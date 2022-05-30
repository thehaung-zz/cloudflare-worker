package model

type CloudFlare struct {
	Ip string `json:"ip"`

	ZoneId string `json:"zone_id"`

	Id string `json:"id"`

	Type string `json:"type"`

	Name string `json:"name"`

	Content string `json:"content"`

	TTL int8 `json:"ttl"`

	Proxied bool `json:"proxied"`
}
