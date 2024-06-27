package models

import (
	"encoding/json"
	"fmt"
	"time"
)

const (
	DateLayout string = "2006-01-02"
	TimeLayout string = "15:04"
)

/*
Using wrapper struct for purchaseDate to support automatic calling
of logic to parse into time.Time when Receipt struct is unmarshalled
*/
type PurchaseDate struct {
	Date time.Time
}

/*
Using wrapper struct for purchaseTime to support automatic calling
of logic to parse into time.Time when Receipt struct is unmarshalled
*/
type PurchaseTime struct {
	Time time.Time
}

type Receipt struct {
	Retailer     string       `json:"retailer"`
	PurchaseDate PurchaseDate `json:"purchaseDate"`
	PurchaseTime PurchaseTime `json:"purchaseTime"`
	Items        []Item       `json:"items"`
	Total        string       `json:"total"`
}

type Item struct {
	ShortDescription string `json:"shortDescription"`
	Price            string `json:"price"`
}

func (d *PurchaseDate) UnmarshalJSON(data []byte) error {
	var rawDate string
	err := json.Unmarshal(data, &rawDate)
	if err != nil {
		return err
	}

	parsedDate, err := time.Parse(DateLayout, rawDate)
	if err != nil {
		return err
	}

	d.Date = parsedDate
	return nil
}

func (d PurchaseDate) String() string {
	return fmt.Sprint(d.Date.Format(DateLayout))
}

func (t *PurchaseTime) UnmarshalJSON(data []byte) error {
	var rawTime string
	err := json.Unmarshal(data, &rawTime)
	if err != nil {
		return err
	}

	parsedTime, err := time.Parse(TimeLayout, rawTime)
	if err != nil {
		return err
	}

	t.Time = parsedTime
	return nil
}

func (p PurchaseTime) String() string {
	return fmt.Sprint(p.Time.Format(TimeLayout))
}
