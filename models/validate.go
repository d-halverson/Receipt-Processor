package models

import (
	"fmt"
	"regexp"
)

const (
	RetailerRegex         string = "^[\\w\\s\\-&]+$"
	ShortDescriptionRegex string = "^[\\w\\s\\-]+$"
	PriceRegex            string = "^\\d+\\.\\d{2}$"
)

/*
Validates that the Receipt object r adheres to all rules outlined in
the api spec here: https://github.com/fetch-rewards/receipt-processor-challenge/blob/main/api.yml

The purchaseDate and purchaseTime fields are also validated when the Receipt struct
is being unmarshalled, but there are checks here to ensure they are initialized
in case anyone in future calls this function on an instance of Receipt that
wasn't created through json.Unmarshal
*/
func (r *Receipt) Validate() error {
	if !regexp.MustCompile(RetailerRegex).MatchString(r.Retailer) {
		return fmt.Errorf("invalid retailer format")
	}
	if r.PurchaseDate.Date.IsZero() {
		return fmt.Errorf("invalid purchase date format")
	}
	if r.PurchaseDate.Date.IsZero() {
		return fmt.Errorf("invalid purchase time format")
	}
	if !regexp.MustCompile(PriceRegex).MatchString(r.Total) {
		return fmt.Errorf("invalid total format")
	}

	if len(r.Items) == 0 {
		return fmt.Errorf("there must be at least one item in receipt")
	}
	for _, item := range r.Items {
		if !regexp.MustCompile(ShortDescriptionRegex).MatchString(item.ShortDescription) {
			return fmt.Errorf("invalid item short description format")
		}
		if !regexp.MustCompile(PriceRegex).MatchString(item.Price) {
			return fmt.Errorf("invalid item price format")
		}
	}

	return nil
}
