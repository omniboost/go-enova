package cardxperts_test

import (
	"encoding/json"
	"log"
	"os"
	"testing"
)

func TestListGastarten(t *testing.T) {
	req := client.NewListGastarten()
	req.RequestBody().Gmdenr = os.Getenv("OESTAT")
	req.RequestBody().Mandant = os.Getenv("MANDANT")
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
