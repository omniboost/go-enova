package enova_test

import (
	"log"
	"net/url"
	"os"
	"testing"

	enova "github.com/omniboost/go-enova"
)

var (
	client *enova.Client
)

func TestMain(m *testing.M) {
	var err error

	baseURLString := os.Getenv("BASE_URL")
	username := os.Getenv("USERNAME")
	password := os.Getenv("PASSWORD")
	debug := os.Getenv("DEBUG")
	var baseURL *url.URL

	if baseURLString != "" {
		baseURL, err = url.Parse(baseURLString)
		if err != nil {
			log.Fatal(err)
		}
	}

	client = enova.NewClient(nil, username, password)
	if debug != "" {
		client.SetDebug(true)
	}

	if baseURL != nil {
		client.SetBaseURL(*baseURL)
	}

	m.Run()
}
