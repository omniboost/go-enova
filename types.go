package meldeschein

type Identifikation struct {
	Erzeugung     Date   `xml:"erzeugung"`
	Schnittstelle string `xml:"schnittstelle"`
	Kurverwaltung int    `xml:"kurverwaltung"`
	BenutzerID    int    `xml:"benutzerid"`
	Verarbeitung  string `xml:"verarbeitung"`
}

type Fehlermeldungen struct {
	Fehler struct {
		Code         string `xml:"code"`
		Beschreibung string `xml:"beschreibung"`
		Bezug        string `xml:"bezug"`
	} `xml:"fehler"`
}
