package config

import (
	"fmt"
	"os"
	"strconv"
)

func GetEnvironment() string {
	return os.Getenv("ENV")
}

func GetTimeDeadlineAPI() int {
	res, err := strconv.Atoi(os.Getenv("TIME_DEADLINE_API"))
	if err != nil {
		panic(err)
	}
	return res
}

func GetServerPort() string {
	res, err := strconv.Atoi(os.Getenv("SERVER_PORT"))
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf(":%d", res)
}

func GetAPIUrl() string {
	return os.Getenv("API_URL")
}

func APIGetIPUrl() string {
	return os.Getenv("API_GET_IP_URL")
}

func GetCloudFlareAPIUrl() string {
	return os.Getenv("CLOUD_FLARE_BASE_URL")
}

func GetCloudFlareEmail() string {
	return os.Getenv("CLOUD_FLARE_EMAIL")
}

func GetCloudFlareAPIToken() string {
	return os.Getenv("CLOUD_FLARE_API_TOKEN")
}

func GetCloudFlareZoneID() string {
	return os.Getenv("CLOUD_FLARE_ZONE_ID")
}
