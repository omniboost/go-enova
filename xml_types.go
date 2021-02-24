package meldeschein

import (
	"encoding/json"
	"encoding/xml"
	"time"

	"github.com/pkg/errors"
)

type Date struct {
	time.Time
}

func (d Date) MarshalSchema() string {
	return d.Time.Format("2006-01-02")
}

func (d Date) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(d.Time.Format("2006-01-02"), start)
}

func (d *Date) UnmarshalXML(dec *xml.Decoder, start xml.StartElement) error {
	var value string
	err := dec.DecodeElement(&value, &start)
	if err != nil {
		return err
	}

	if value == "" {
		return nil
	}

	// first try standard date
	layout := time.RFC3339
	d.Time, err = time.Parse(layout, value)
	if err == nil {
		return nil
	}

	layout = "2006-01-02"
	d.Time, err = time.Parse(layout, value)
	if err == nil {
		return errors.WithStack(err)
	}

	return nil
}

func (d Date) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.Format("2006-01-02"))
}

func (d *Date) UnmarshalJSON(text []byte) (err error) {
	var value string
	err = json.Unmarshal(text, &value)
	if err != nil {
		return err
	}

	if value == "" {
		return nil
	}

	// 28-1-2008
	layout := "2006-01-02"
	time, err := time.Parse(layout, value)
	d.Time = time
	return err
}

type DateTime struct {
	time.Time
}
