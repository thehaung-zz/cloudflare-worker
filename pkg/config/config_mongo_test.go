package config

import (
	"fmt"
	"log"
	"os"
	"testing"
)

const succeed = "\u2713"
const failed = "\u2717"

func TestConfig(t *testing.T) {
	err := os.Setenv("MONGO_SERVER_HOST", ":1234")
	if err != nil {
		log.Fatal(err)
		return
	}
	err = os.Setenv("MONGO_USER", "john")
	if err != nil {
		log.Fatal(err)
		return
	}
	err = os.Setenv("MONGO_PASSWORD", "doe")
	if err != nil {
		log.Fatal(err)
		return
	}
	err = os.Setenv("MONGO_HOST", "localhost")
	if err != nil {
		log.Fatal(err)
		return
	}
	err = os.Setenv("MONGO_PORT", "27017")
	if err != nil {
		log.Fatal(err)
		return
	}

	config := NewConfig()

	mongoURI := fmt.Sprintf("mongodb+srv://%s:%s@%s",
		"john",
		"doe",
		"localhost",
	)

	if config.MongoURI() != mongoURI {
		t.Fatalf("%s there is an problem in mongo connection uri generator. %s != %s", failed, config.MongoURI(), mongoURI)
	}
	t.Logf("%s Testing MongoDB connection uri generator is successful", succeed)
}
