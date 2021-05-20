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
			Mblattnr:    "",
			ResID:       "11001",
			Bearbeiter:  "cardxperts",
			Abgeplant:   cardxperts.Date{Time: time.Now()},
			Ankunft:     cardxperts.Date{Time: time.Now()},
			Aufenthalte: "3",
			Reisegruppe: 0,
			Landschl: []cardxperts.Landschl{{
				Anzpers: 1,
				Lschlnr: "14",
			}},
			Gastart: cardxperts.Gastarten{
				{
					Anzpers: 1,
					Gastart: "P",
				},
			},
			Gast: cardxperts.Gasten{
				{
					Gasttyp:       "HG",
					Anrede:        "Fr.",
					Vorname:       "Marie",
					Name:          "Antoinette",
					Strasse:       "Teststrasse",
					Strasse2:      "",
					Pobox:         "",
					Nation:        "F",
					Plz:           "1234",
					Ort:           "Paris",
					Ortzusatz:     "",
					Gebdatum:      cardxperts.Date{Time: time.Now()},
					Geschlecht:    "2",
					Reisedokument: "",
					Staatsang:     "A",
					Herkunftsland: "14",
					Beruf:         "",
					Berufssparte:  "",
					Email:         "",
					Telefon:       "",
					Zusatztext:    "",
					Gaesteart:     "P",
					Gastlfdnr:     1,
				},
				// {
				// 	Staatsang: "D",
				// 	// Gebdatum:  "",
				// 	Gastlfdnr: 2,
				// },
				// {
				// 	Staatsang: "D",
				// 	// Gebdatum:  "",
				// 	Gastlfdnr: 3,
				// },
				// {
				// 	Staatsang: "D",
				// 	// Gebdatum:  "",
				// 	Gastlfdnr: 4,
				// },
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
