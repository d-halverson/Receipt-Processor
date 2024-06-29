package points

import (
	"math"
	"receipts/models"
	"strconv"
	"strings"
	"unicode"
)

// Given a receipt, return number of points gained from this rule
type ReceiptRule func(*models.Receipt) int

// Returns a list of all current receipt point rules
func GetReceiptRules() []ReceiptRule {
	return []ReceiptRule{RetailerRule, TotalRoundRule, TotalMultipleRule, NumItemsRule, ItemDescriptionRule, PurchaseDayRule, PurchaseTimeRule}
}

// Returns 1 point for each alphanumeric character in the retailer
func RetailerRule(receipt *models.Receipt) int {
	count := 0
	for _, char := range receipt.Retailer {
		if unicode.IsLetter(char) || unicode.IsNumber(char) {
			count++
		}
	}
	return count
}

// Returns 50 points if the receipt total is a round dollar amount with no cents.
func TotalRoundRule(receipt *models.Receipt) int {
	if models.GetCents(receipt.Total) == 0 {
		return 50
	}
	return 0
}

// Returns 25 points if the receipt total is a multiple of 0.25
func TotalMultipleRule(receipt *models.Receipt) int {
	total, err := strconv.ParseFloat(receipt.Total, 64)
	if err != nil {
		return 0
	}
	if total >= 0.25 && math.Mod(total, 0.25) == 0 {
		return 25
	}
	return 0
}

// Returns 5 points for every two items on the receipt
func NumItemsRule(receipt *models.Receipt) int {
	return (len(receipt.Items) / 2) * 5
}

/*
For every item with shortDescription with trimmed length being
a multiple of three, the price of the item is multiplied by 0.2
and rounded up to the nearest integer. This value is points gained
from each item. This rule sums up points from this calculation on
all items for the receipt.
*/
func ItemDescriptionRule(receipt *models.Receipt) int {
	points := 0

	for _, item := range receipt.Items {
		if len(strings.TrimSpace(item.ShortDescription))%3 == 0 {
			price, err := strconv.ParseFloat(item.Price, 64)
			if err != nil {
				continue
			}
			points = points + int(math.Ceil(price*0.2))
		}
	}

	return points
}

// Returns 6 points if the purchaseDate is on an odd day.
func PurchaseDayRule(receipt *models.Receipt) int {
	if receipt.PurchaseDate.Date.Day()%2 != 0 {
		return 6
	}
	return 0
}

/*
Returns 10 points if the purchaseTime of receipt is in between
2:00PM and 4:00PM exclusive.
*/
func PurchaseTimeRule(receipt *models.Receipt) int {
	hour := receipt.PurchaseTime.Time.Hour()
	minutes := receipt.PurchaseTime.Time.Minute()

	if hour == 15 {
		return 10
	}
	if hour == 14 && minutes > 0 {
		return 10
	}

	return 0
}
