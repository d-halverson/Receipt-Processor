package points

import (
	"encoding/json"
	"receipts/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculatePoints(t *testing.T) {
	tests := []struct {
		testName       string
		inputReceipt   string
		expectedPoints int
	}{
		{
			testName: "GithubExampleOne",
			inputReceipt: `{
								"retailer": "Target",
								"purchaseDate": "2022-01-01",
								"purchaseTime": "13:01",
								"items": [
									{
									"shortDescription": "Mountain Dew 12PK",
									"price": "6.49"
									},{
									"shortDescription": "Emils Cheese Pizza",
									"price": "12.25"
									},{
									"shortDescription": "Knorr Creamy Chicken",
									"price": "1.26"
									},{
									"shortDescription": "Doritos Nacho Cheese",
									"price": "3.35"
									},{
									"shortDescription": "   Klarbrunn 12-PK 12 FL OZ  ",
									"price": "12.00"
									}
								],
								"total": "35.35"
							}`,
			expectedPoints: 28,
		},
		{
			testName: "GithubExampleTwo",
			inputReceipt: `{
								"retailer": "M&M Corner Market",
								"purchaseDate": "2022-03-20",
								"purchaseTime": "14:33",
								"items": [
									{
									"shortDescription": "Gatorade",
									"price": "2.25"
									},{
									"shortDescription": "Gatorade",
									"price": "2.25"
									},{
									"shortDescription": "Gatorade",
									"price": "2.25"
									},{
									"shortDescription": "Gatorade",
									"price": "2.25"
									}
								],
								"total": "9.00"
							}`,
			expectedPoints: 109,
		},
		{
			testName: "OnlyRetailerPoints",
			inputReceipt: `{
								"retailer": "Madison Fresh Market", 
								"purchaseDate": "2019-09-02",
								"purchaseTime": "09:33",
								"items": [
									{
									"shortDescription": "Apple",
									"price": "3.99"
									}
								],
								"total": "3.99"
							}`,
			expectedPoints: 18,
		},
	}

	for _, test := range tests {
		t.Run(test.testName, func(t *testing.T) {
			var receipt models.Receipt
			err := json.Unmarshal([]byte(test.inputReceipt), &receipt)
			assert.NoError(t, err)
			assert.Equal(t, test.expectedPoints, CalculatePoints(&receipt))
		})
	}
}
