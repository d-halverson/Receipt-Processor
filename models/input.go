package models

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"
)

// This file contains structs that represent receipt input a request body.

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

/*
Returns dollars before period in a price string.
If format is invalid, returns -1.
*/
func GetDollars(price string) int {
	// Split the price string into dollars and cents
	parts := strings.Split(price, ".")
	if len(parts) != 2 || parts[0] == "" || parts[1] == "" {
		return -1
	}

	dollars, err := strconv.Atoi(parts[0])
	if err != nil {
		return -1
	}

	return dollars
}

/*
Returns cents after period in a price string.
Assumes valid format, ex: "1.25"
*/
func GetCents(price string) int {
	// Split the price string into dollars and cents
	parts := strings.Split(price, ".")
	if len(parts) != 2 || parts[0] == "" || parts[1] == "" {
		return -1
	}

	cents, err := strconv.Atoi(parts[1])
	if err != nil {
		return -1
	}

	return cents
}

func (d *PurchaseDate) UnmarshalJSON(data []byte) error {
	var rawDate string
	err := json.Unmarshal(data, &rawDate)
	if err != nil {
		return fmt.Errorf("purchaseDate field must be a string")
	}

	parsedDate, err := time.Parse(DateLayout, rawDate)
	if err != nil {
		return fmt.Errorf("purchaseDate must be in format YYYY-MM-DD")
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
		return fmt.Errorf("purchaseTime field must be a string")
	}

	parsedTime, err := time.Parse(TimeLayout, rawTime)
	if err != nil {
		return fmt.Errorf("purchaseTime must be in format HH:MM")
	}

	t.Time = parsedTime
	return nil
}

func (p PurchaseTime) String() string {
	return fmt.Sprint(p.Time.Format(TimeLayout))
}
