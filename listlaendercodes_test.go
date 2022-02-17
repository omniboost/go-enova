package enova_test

import (
	"encoding/json"
	"log"
	"os"
	"testing"
)

func TestListlaendercodes(t *testing.T) {
	req := client.NewListlaendercodes()
	// req.RequestBody().Identifikation.Erzeugung = "2020-02-02"
	req.RequestBody().Mandant = os.Getenv("MANDANT")
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
