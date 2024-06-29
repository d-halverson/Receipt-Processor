package points

import (
	"receipts/models"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

/*
Only worried about testing valid receipt cases in this file,
invalid cases are handled in validate_test.go, and by the
time points are calculated we can assume we have a valid receipt.
*/

func TestRetailerRule(t *testing.T) {
	tests := []struct {
		testName       string
		inputRetailer  string
		expectedPoints int
	}{
		{
			testName:       "BasicRetailer",
			inputRetailer:  "Target",
			expectedPoints: 6,
		},
		{
			testName:       "NumbersAndLetters",
			inputRetailer:  "Target2",
			expectedPoints: 7,
		},
		{
			testName:       "InnerWhitespace",
			inputRetailer:  "Target 2",
			expectedPoints: 7,
		},
		{
			testName:       "OutterWhitespace",
			inputRetailer:  "   Target   ",
			expectedPoints: 6,
		},
	}

	for _, test := range tests {
		t.Run(test.testName, func(t *testing.T) {
			assert.Equal(t, test.expectedPoints, RetailerRule(&models.Receipt{Retailer: test.inputRetailer}))
		})
	}
}

func TestTotalRoundRule(t *testing.T) {
	tests := []struct {
		testName       string
		inputTotal     string
		expectedPoints int
	}{
		{
			testName:       "RoundDollar",
			inputTotal:     "1.00",
			expectedPoints: 50,
		},
		{
			testName:       "Cents",
			inputTotal:     "1.25",
			expectedPoints: 0,
		},
		{
			testName:       "CentsOneZero",
			inputTotal:     "1.20",
			expectedPoints: 0,
		},
	}

	for _, test := range tests {
		t.Run(test.testName, func(t *testing.T) {
			assert.Equal(t, test.expectedPoints, TotalRoundRule(&models.Receipt{Total: test.inputTotal}))
		})
	}
}

func TestTotalMultipleRule(t *testing.T) {
	tests := []struct {
		testName       string
		inputTotal     string
		expectedPoints int
	}{
		{
			testName:       "RoundDollar",
			inputTotal:     "1.00",
			expectedPoints: 25,
		},
		{
			testName:       "25Cents",
			inputTotal:     "1.25",
			expectedPoints: 25,
		},
		{
			testName:       "20Cents",
			inputTotal:     "1.20",
			expectedPoints: 0,
		},
		{
			testName:       "50Cents",
			inputTotal:     "1.50",
			expectedPoints: 25,
		},
		{
			testName:       "75Cents",
			inputTotal:     "1.75",
			expectedPoints: 25,
		},
		{
			testName:       "75Cents",
			inputTotal:     "1.75",
			expectedPoints: 25,
		},
		{
			testName:       "Free",
			inputTotal:     "0.00",
			expectedPoints: 0,
		},
		{
			testName:       "LessThan25Cents",
			inputTotal:     "0.10",
			expectedPoints: 0,
		},
	}

	for _, test := range tests {
		t.Run(test.testName, func(t *testing.T) {
			assert.Equal(t, test.expectedPoints, TotalMultipleRule(&models.Receipt{Total: test.inputTotal}))
		})
	}
}

func TestNumItemsRule(t *testing.T) {
	tests := []struct {
		testName       string
		inputNumItems  int
		expectedPoints int
	}{
		{
			testName:       "OneItem",
			inputNumItems:  1,
			expectedPoints: 0,
		},
		{
			testName:       "TwoItems",
			inputNumItems:  2,
			expectedPoints: 5,
		},
		{
			testName:       "ThreeItems",
			inputNumItems:  3,
			expectedPoints: 5,
		},
		{
			testName:       "11Items",
			inputNumItems:  11,
			expectedPoints: 25,
		},
		{
			testName:       "100Items",
			inputNumItems:  100,
			expectedPoints: 250,
		},
	}

	for _, test := range tests {
		t.Run(test.testName, func(t *testing.T) {
			items := []models.Item{}
			for i := 0; i < test.inputNumItems; i++ {
				items = append(items, models.Item{})
			}

			assert.Equal(t, test.expectedPoints, NumItemsRule(&models.Receipt{Items: items}))
		})
	}
}

func TestItemDescriptionRule(t *testing.T) {
	tests := []struct {
		testName       string
		inputItems     []models.Item
		expectedPoints int
	}{
		{
			testName:       "NoApplicableItems",
			inputItems:     []models.Item{{ShortDescription: "Hello", Price: "1.25"}},
			expectedPoints: 0,
		},
		{
			testName:       "OneApplicableItemWhitespace",
			inputItems:     []models.Item{{ShortDescription: " This string length is multiple of three      ", Price: "10.00"}},
			expectedPoints: 2,
		},
		{
			testName:       "OneApplicableItem",
			inputItems:     []models.Item{{ShortDescription: "Hello", Price: "1.25"}, {ShortDescription: "Hey", Price: "1.00"}},
			expectedPoints: 1,
		},
		{
			testName: "ManyApplicableItem",
			inputItems: []models.Item{
				{ShortDescription: "This string length is multiple of three", Price: "0.10"},
				{ShortDescription: "Hello", Price: "1.25"}, // N/A
				{ShortDescription: "Hey", Price: "1.00"},
				{ShortDescription: "Hey", Price: "555.12"},
			},
			expectedPoints: 114,
		},
	}

	for _, test := range tests {
		t.Run(test.testName, func(t *testing.T) {
			assert.Equal(t, test.expectedPoints, ItemDescriptionRule(&models.Receipt{Items: test.inputItems}))
		})
	}
}

func TestPurchaseDayRule(t *testing.T) {
	tests := []struct {
		testName          string
		inputPurchaseDate string
		expectedPoints    int
	}{
		{
			testName:          "OddDay",
			inputPurchaseDate: "2022-01-01",
			expectedPoints:    6,
		},
		{
			testName:          "EvenDay",
			inputPurchaseDate: "2022-01-02",
			expectedPoints:    0,
		},
		{
			testName:          "EvenMonthOddDay",
			inputPurchaseDate: "1999-10-13",
			expectedPoints:    6,
		},
		{
			testName:          "EvenMonthEvenDay",
			inputPurchaseDate: "1999-10-14",
			expectedPoints:    0,
		},
	}

	for _, test := range tests {
		t.Run(test.testName, func(t *testing.T) {
			parsedDate, err := time.Parse(models.DateLayout, test.inputPurchaseDate)
			assert.NoError(t, err)
			assert.Equal(t, test.expectedPoints, PurchaseDayRule(&models.Receipt{PurchaseDate: models.PurchaseDate{Date: parsedDate}}))
		})
	}
}

func TestPurchaseTimeRule(t *testing.T) {
	tests := []struct {
		testName          string
		inputPurchaseTime string
		expectedPoints    int
	}{
		{
			testName:          "Before2",
			inputPurchaseTime: "10:10",
			expectedPoints:    0,
		},
		{
			testName:          "At2",
			inputPurchaseTime: "14:00",
			expectedPoints:    0,
		},
		{
			testName:          "In2Hour",
			inputPurchaseTime: "14:30",
			expectedPoints:    10,
		},
		{
			testName:          "At3",
			inputPurchaseTime: "15:00",
			expectedPoints:    10,
		},
		{
			testName:          "In3Hour",
			inputPurchaseTime: "15:30",
			expectedPoints:    10,
		},
		{
			testName:          "At4",
			inputPurchaseTime: "16:00",
			expectedPoints:    0,
		},
		{
			testName:          "In4Hour",
			inputPurchaseTime: "16:50",
			expectedPoints:    0,
		},
		{
			testName:          "After4",
			inputPurchaseTime: "20:10",
			expectedPoints:    0,
		},
	}

	for _, test := range tests {
		t.Run(test.testName, func(t *testing.T) {
			parsedTime, err := time.Parse(models.TimeLayout, test.inputPurchaseTime)
			assert.NoError(t, err)
			assert.Equal(t, test.expectedPoints, PurchaseTimeRule(&models.Receipt{PurchaseTime: models.PurchaseTime{Time: parsedTime}}))
		})
	}
}
