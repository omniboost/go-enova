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
	Uvmnr      string     `xml:"uvmnr,attr,omitempty"`
	Meldeblatt Meldeblatt `xml:"meldeblatt"`
}

type Meldeblatt struct {
	Mblattnr    string `xml:"mblattnr,attr,omitempty"`
	ResID       string `xml:"resid,attr"`
	Angeplant   Date   `xml:"angeplant,attr,omitempty"`
	Ankunft     Date   `xml:"ankunft,attr,omitempty"`
	Abreise     Date   `xml:"abreise,attr"`
	Bearbeiter  string `xml:"bearbeiter,attr"`
	Bemerkung   string `xml:"bemerkung,attr"`
	Aufenthalte string `xml:"aufenthalte,attr,omitempty"`
	Abgeplant   Date   `xml:"abgeplant,attr"`

	Undef3      string     `xml:"undef3,attr,omitempty"`
	Undef2      string     `xml:"undef2,attr,omitempty"`
	Undef1      string     `xml:"undef1,attr,omitempty"`
	Reisegruppe int        `xml:"reisegruppe,attr"`
	Zahlungsart string     `xml:"zahlungsart,attr,omitempty"`
	Landschl    []Landschl `xml:"landschl"`
	Gastart     Gastarten  `xml:"gastart"`
	Gast        Gasten     `xml:"gast"`
}

func (mb Meldeblatt) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return omitempty.MarshalXML(mb, e, start)
}

type Landschl struct {
	Lschlnr string `xml:"lschlnr,attr"`
	Anzpers int    `xml:"anzpers,attr"`
}

type Gastarten []Gastart

type Gastart struct {
	Gastart string `xml:"gastart,attr"`
	Anzpers int    `xml:"anzpers,attr"`
}

type Gasten []Gast

type Gast struct {
	Gastlfdnr     int    `xml:"gastlfdnr,attr"`
	Gasttyp       string `xml:"gasttyp,attr"`
	Anrede        string `xml:"anrede,attr"`
	Vorname       string `xml:"vorname,attr"`
	Name          string `xml:"name,attr"`
	Strasse       string `xml:"strasse,attr"`
	Strasse2      string `xml:"strasse2,attr"`
	Pobox         string `xml:"pobox,attr"`
	Nation        string `xml:"nation,attr"`
	Plz           string `xml:"plz,attr"`
	Ort           string `xml:"ort,attr"`
	Ortzusatz     string `xml:"ortzusatz,attr"`
	Gebdatum      Date   `xml:"gebdatum,attr,omitempty"`
	Geschlecht    string `xml:"geschlecht,attr"`
	Reisedokument string `xml:"reisedokument,attr"`
	Staatsang     string `xml:"staatsang,attr"`
	Herkunftsland string `xml:"herkunftsland,attr"`
	Beruf         string `xml:"beruf,attr"`
	Berufssparte  string `xml:"berufssparte,attr"`
	Email         string `xml:"email,attr"`
	Telefon       string `xml:"telefon,attr"`
	Zusatztext    string `xml:"zusatztext,attr"`
	Gaesteart     string `xml:"gaesteart,attr"`

	// Hobby         string `xml:"hobby,attr,omitempty"`
	// Motiv         string `xml:"motiv,attr,omitempty"`
	// // Titel         string `xml:"titel,attr"`
	GastkarteNr string `xml:"gastkartenr,attr"`
	PersonenID  string `xml:"personenid,attr"`
}

func (g Gast) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return omitempty.MarshalXML(g, e, start)
}
