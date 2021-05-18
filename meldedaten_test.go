package cardxperts_test

import (
	"encoding/json"
	"log"
	"os"
	"testing"
	"time"

	"github.com/omniboost/go-cardxperts"
)

func TestMeldedaten(t *testing.T) {
	req := client.NewMeldedatenRequest()
	// req.RequestBody().Identifikation.Erzeugung = "2020-02-02"
	req.RequestBody().Meldedaten.Oestat = os.Getenv("OESTAT")
	req.RequestBody().Meldedaten.Logindaten = cardxperts.Logindaten{
		User:     client.Username(),
		Mandant:  os.Getenv("MANDANT"),
		Passwort: client.Password(),
	}
	req.RequestBody().Meldedaten.Betrieb = cardxperts.Betrieb{
		Betriebnr: os.Getenv("BETRIEBNR"),
		Meldeblatt: cardxperts.Meldeblatt{
			// Mblattnr:  "11001",
			ResID:     "",
			Abgeplant: cardxperts.Date{Time: time.Now()},
			Angeplant: cardxperts.Date{Time: time.Now()},
			Landschl: cardxperts.Landschl{
				Anzpers: 4,
				Lschlnr: "D",
			},
			Gastart: cardxperts.Gastarten{
				{
					Anzpers: 2,
					Gastart: "E",
				},
				{
					Anzpers: 2,
					Gastart: "K",
				},
			},
			Gast: cardxperts.Gasten{
				{
					Staatsang: "D",
					// Gebdatum:  "",
					Gastlfdnr: 1,
				},
				{
					Staatsang: "D",
					// Gebdatum:  "",
					Gastlfdnr: 2,
				},
				{
					Staatsang: "D",
					// Gebdatum:  "",
					Gastlfdnr: 3,
				},
				{
					Staatsang: "D",
					// Gebdatum:  "",
					Gastlfdnr: 4,
				},
			},
		},
	}
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
