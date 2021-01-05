package meldeschein_test

import (
	"encoding/json"
	"log"
	"testing"
)

func TestGetMeldescheine(t *testing.T) {
	req := client.NewGetMeldescheine()
	// req.RequestBody().Identifikation.Erzeugung = "2020-02-02"
	req.RequestBody().Identifikation.Schnittstelle = "GUESTLINE_OMNIBOOST"
	// req.RequestBody().Identifikation.Kurverwaltung = 0
	// req.RequestBody().Identifikation.Benutzerid = 1
	req.RequestBody().Identifikation.Verarbeitung = "MS-HOLEN"
	req.RequestBody().AnfrageDaten.Meldescheinnummer = "39899"
	req.RequestBody().AnfrageDaten.OrtID = 263
	// req.RequestBody().ArrivalDate = guestline.DateTime{time.Now().AddDate(0, 0, -7)}
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
