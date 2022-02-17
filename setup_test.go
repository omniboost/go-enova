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
	dbName := os.Getenv("DB_NAME")
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	guid := os.Getenv("EXTERNAL_SYSTEM_GUID")
	tableName := os.Getenv("TABLE_NAME")
	schemaName := os.Getenv("SCHEMA_NAME")
	debug := os.Getenv("DEBUG")
	var baseURL *url.URL

	if baseURLString != "" {
		baseURL, err = url.Parse(baseURLString)
		if err != nil {
			log.Fatal(err)
		}
	}

	client = enova.NewClient(nil)
	if debug != "" {
		client.SetDebug(true)
	}

	if baseURL != nil {
		client.SetBaseURL(*baseURL)
	}

	client.SetDBName(dbName)
	client.SetDBUsername(dbUsername)
	client.SetDBPassword(dbPassword)
	client.SetExternalSystemGUID(guid)
	client.SetTableName(tableName)
	client.SetSchemaName(schemaName)

	m.Run()
}
