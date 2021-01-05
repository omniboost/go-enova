package meldeschein_test

import (
	"encoding/json"
	"log"
	"testing"
)

func TestGetConfigurationList(t *testing.T) {
	req := client.NewGetConfigurationList()
	// req.RequestBody().Identifikation.Erzeugung = "2020-02-02"
	req.RequestBody().Identifikation.Schnittstelle = "GUESTLINE_OMNIBOOST"
	// req.RequestBody().Identifikation.Kurverwaltung = 0
	// req.RequestBody().Identifikation.Benutzerid = 1
	// req.RequestBody().Identifikation.Verarbeitung = "CONFIG-L-LAND"
	// req.RequestBody().Identifikation.Verarbeitung = "CONFIG-L-ANREDE"
	req.RequestBody().Identifikation.Verarbeitung = "CONFIG-L-KAT"
	// req.RequestBody().ArrivalDate = guestline.DateTime{time.Now().AddDate(0, 0, -7)}
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
