package cardxperts

import (
	"encoding/xml"

	"github.com/omniboost/go-cardxperts/omitempty"
)

type Meldedaten struct {
	XMLName    xml.Name   `xml:"gemeinde"`
	Oestat     string     `xml:"oestat,attr"`
	Version    string     `xml:"version,attr"`
	Logindaten Logindaten `xml:"Logindaten"`
	Betrieb    Betrieb    `xml:"betrieb"`
}

type Logindaten struct {
	User     string `xml:"user"`
	Mandant  string `xml:"mandant"`
	Passwort string `xml:"passwort"`
}

type Betrieb struct {
	Betriebnr  string     `xml:"betriebnr,attr"`
	Uvmnr      string     `xml:"uvmnr,attr"`
	Meldeblatt Meldeblatt `xml:"meldeblatt"`
}

type Meldeblatt struct {
	Mblattnr    string `xml:"mblattnr,attr"`
	Angeplant   Date   `xml:"angeplant,attr"`
	Ankunft     Date   `xml:"ankunft,attr,omitempty"`
	Abreise     Date   `xml:"abreise,attr"`
	Abgeplant   Date   `xml:"abgeplant,attr"`
	Bearbeiter  string `xml:"bearbeiter,attr"`
	Bemerkung   string `xml:"bemerkung,attr"`
	Aufenthalte string `xml:"aufenthalte,attr"`
	ResID       string `xml:"resid,attr"`

	Undef3 string `xml:"undef3,attr,omitempty"`
	Undef2 string `xml:"undef2,attr,omitempty"`
	Undef1 string `xml:"undef1,attr,omitempty"`
	// Reisegruppe int    `xml:"reisegruppe,attr"`
	Zahlungsart string    `xml:"zahlungsart,attr"`
	Landschl    Landschl  `xml:"landschl"`
	Gastart     Gastarten `xml:"gastart"`
	Gast        Gasten    `xml:"gast"`
}

func (mb Meldeblatt) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return omitempty.MarshalXML(mb, e, start)
}

type Landschl struct {
	Anzpers int    `xml:"anzpers,attr"`
	Lschlnr string `xml:"lschlnr,attr"`
}

type Gastarten []Gastart

type Gastart struct {
	Anzpers int    `xml:"anzpers,attr"`
	Gastart string `xml:"gastart,attr"`
}

type Gasten []Gast

type Gast struct {
	Hobby         string `xml:"hobby,attr"`
	Motiv         string `xml:"motiv,attr"`
	Zusatztext    string `xml:"zusatztext,attr"`
	Telefon       string `xml:"telefon,attr"`
	Email         string `xml:"email,attr"`
	Berufssparte  string `xml:"berufssparte,attr"`
	Beruf         string `xml:"beruf,attr"`
	Staatsang     string `xml:"staatsang,attr"`
	Reisedokument string `xml:"reisedokument,attr"`
	Geschlecht    string `xml:"geschlecht,attr"`
	Gebdatum      Date   `xml:"gebdatum,attr,omitempty"`
	Ortzusatz     string `xml:"ortzusatz,attr"`
	Ort           string `xml:"ort,attr"`
	Plz           string `xml:"plz,attr"`
	Nation        string `xml:"nation,attr"`
	Pobox         string `xml:"pobox,attr"`
	Strasse2      string `xml:"strasse2,attr"`
	Strasse       string `xml:"strasse,attr"`
	Name          string `xml:"name,attr"`
	Vorname       string `xml:"vorname,attr"`
	Titel         string `xml:"titel,attr"`
	Anrede        string `xml:"anrede,attr"`
	Gastkartenr   string `xml:"gastkartenr,attr"`
	Gasttyp       string `xml:"gasttyp,attr"`
	Gaesteart     string `xml:"gaesteart,attr"`
	Personenid    string `xml:"personenid,attr"`
	Gastlfdnr     int    `xml:"gastlfdnr,attr"`
}

func (g Gast) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return omitempty.MarshalXML(g, e, start)
}
