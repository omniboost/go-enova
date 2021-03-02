package cardxperts_test

import (
	"encoding/json"
	"log"
	"testing"
)

func TestListGastarten(t *testing.T) {
	req := client.NewListGastarten()
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
