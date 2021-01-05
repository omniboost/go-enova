package meldeschein_test

import (
	"encoding/json"
	"log"
	"testing"
	"time"

	meldeschein "github.com/omniboost/go-avs-meldeschein"
)

func TestPostMeldeschein(t *testing.T) {
	req := client.NewPostMeldeschein()
	req.RequestBody().Identifikation.Schnittstelle = "GUESTLINE_OMNIBOOST"
	req.RequestBody().Meldeschein = meldeschein.Meldeschein{
		Buchungsnummer:         "1223",
		FirmaID:                363,
		ObjektID:               263,
		Anreise:                meldeschein.Date{time.Now()},
		Abreise:                meldeschein.Date{time.Now().AddDate(0, 0, 7)},
		KategorieID:            1,
		AnredeID:               2,
		Name:                   "Bogaert",
		Vorname:                "Leon",
		Strasse:                "Stadhuisplein",
		Hausnummer:             "3",
		Plz:                    "Terneuzen",
		Ort:                    "4563CJ",
		LandID:                 1,
		StaatsangehoerigkeitID: 1,
		WeitereAngaben:         "",
		Ausweisnr:              "",
		Kfzkennzeichen:         "",
		Geburtsdatum:           meldeschein.Date{time.Date(1983, 4, 12, 0, 0, 0, 0, time.UTC)},
	}
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {
		t.Error(err)
	}

	log.Println(string(b))
}
