package models

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateReceipt(t *testing.T) {
	tests := []struct {
		testName            string
		inputReceipt        string // json format of receipt
		expectJsonError     bool
		expectValidateError bool
	}{
		{
			testName: "simple-receipt.json",
			inputReceipt: `{
					"retailer": "Target",
					"purchaseDate": "2022-01-02",
					"purchaseTime": "13:13",
					"total": "1.25",
					"items": [
						{"shortDescription": "Pepsi - 12-oz", "price": "1.25"}
					]
				}`,
			expectJsonError:     false,
			expectValidateError: false,
		},
		{
			testName: "morning-receipt.json",
			inputReceipt: `{
					"retailer": "Walgreens",
					"purchaseDate": "2022-01-02",
					"purchaseTime": "08:13",
					"total": "2.65",
					"items": [
						{"shortDescription": "Pepsi - 12-oz", "price": "1.25"},
						{"shortDescription": "Dasani", "price": "1.40"}
					]
				}`,
			expectJsonError:     false,
			expectValidateError: false,
		},
		{
			testName: "&InRetailerIsValid",
			inputReceipt: `{
					"retailer": "M&M Corner Market",
					"purchaseDate": "2022-01-02",
					"purchaseTime": "08:13",
					"total": "2.65",
					"items": [
						{"shortDescription": "Pepsi - 12-oz", "price": "1.25"},
						{"shortDescription": "Dasani", "price": "1.40"}
					]
				}`,
			expectJsonError:     false,
			expectValidateError: false,
		},
		{
			testName:            "empty",
			inputReceipt:        ``,
			expectJsonError:     true,
			expectValidateError: true,
		},
		{
			testName: "EmptyRetailer",
			inputReceipt: `{
					"retailer": "", 
					"purchaseDate": "2022-01-02",
					"purchaseTime": "08:13",
					"total": "2.65",
					"items": [
						{"shortDescription": "Pepsi - 12-oz", "price": "1.25"},
						{"shortDescription": "Dasani", "price": "1.40"}
					]
				}`,
			expectJsonError:     false,
			expectValidateError: true,
		},
		{
			testName: "InvalidRetailer", // ! is invalid special character
			inputReceipt: `{
					"retailer": "M&M Corner Market!", 
					"purchaseDate": "2022-01-02",
					"purchaseTime": "08:13",
					"total": "2.65",
					"items": [
						{"shortDescription": "Pepsi - 12-oz", "price": "1.25"},
						{"shortDescription": "Dasani", "price": "1.40"}
					]
				}`,
			expectJsonError:     false,
			expectValidateError: true,
		},
		{
			testName: "EmptyDate",
			inputReceipt: `{
					"retailer": "Madison Fresh Market", 
					"purchaseDate": "",
					"purchaseTime": "08:13",
					"total": "2.65",
					"items": [
						{"shortDescription": "Pepsi - 12-oz", "price": "1.25"},
						{"shortDescription": "Dasani", "price": "1.40"}
					]
				}`,
			expectJsonError:     true,
			expectValidateError: true,
		},
		{
			testName: "InvalidDateYear",
			inputReceipt: `{
					"retailer": "Madison Fresh Market", 
					"purchaseDate": "22-01-02",
					"purchaseTime": "08:13",
					"total": "2.65",
					"items": [
						{"shortDescription": "Pepsi - 12-oz", "price": "1.25"},
						{"shortDescription": "Dasani", "price": "1.40"}
					]
				}`,
			expectJsonError:     true,
			expectValidateError: true,
		},
		{
			testName: "InvalidDateMonth",
			inputReceipt: `{
					"retailer": "Madison Fresh Market", 
					"purchaseDate": "2022-1-02",
					"purchaseTime": "08:13",
					"total": "2.65",
					"items": [
						{"shortDescription": "Pepsi - 12-oz", "price": "1.25"},
						{"shortDescription": "Dasani", "price": "1.40"}
					]
				}`,
			expectJsonError:     true,
			expectValidateError: true,
		},
		{
			testName: "InvalidDateDay",
			inputReceipt: `{
					"retailer": "Madison Fresh Market", 
					"purchaseDate": "2022-01-1",
					"purchaseTime": "08:13",
					"total": "2.65",
					"items": [
						{"shortDescription": "Pepsi - 12-oz", "price": "1.25"},
						{"shortDescription": "Dasani", "price": "1.40"}
					]
				}`,
			expectJsonError:     true,
			expectValidateError: true,
		},
		{
			testName: "InvalidDateUnderscore",
			inputReceipt: `{
					"retailer": "Madison Fresh Market", 
					"purchaseDate": "2022_01_02",
					"purchaseTime": "08:13",
					"total": "2.65",
					"items": [
						{"shortDescription": "Pepsi - 12-oz", "price": "1.25"},
						{"shortDescription": "Dasani", "price": "1.40"}
					]
				}`,
			expectJsonError:     true,
			expectValidateError: true,
		},
		{
			testName: "EmptyTime",
			inputReceipt: `{
					"retailer": "Madison Fresh Market", 
					"purchaseDate": "2022-01-02",
					"purchaseTime": "",
					"total": "2.65",
					"items": [
						{"shortDescription": "Pepsi - 12-oz", "price": "1.25"},
						{"shortDescription": "Dasani", "price": "1.40"}
					]
				}`,
			expectJsonError:     true,
			expectValidateError: true,
		},
		{
			testName: "InvalidTimeHour",
			inputReceipt: `{
					"retailer": "Madison Fresh Market", 
					"purchaseDate": "2022-01-02",
					"purchaseTime": "100:13",
					"total": "2.65",
					"items": [
						{"shortDescription": "Pepsi - 12-oz", "price": "1.25"},
						{"shortDescription": "Dasani", "price": "1.40"}
					]
				}`,
			expectJsonError:     true,
			expectValidateError: true,
		},
		{
			testName: "InvalidTimeMin",
			inputReceipt: `{
					"retailer": "Madison Fresh Market", 
					"purchaseDate": "2022-01-02",
					"purchaseTime": "10:113",
					"total": "2.65",
					"items": [
						{"shortDescription": "Pepsi - 12-oz", "price": "1.25"},
						{"shortDescription": "Dasani", "price": "1.40"}
					]
				}`,
			expectJsonError:     true,
			expectValidateError: true,
		},
		{
			testName: "InvalidTimeUnderscore",
			inputReceipt: `{
					"retailer": "Madison Fresh Market", 
					"purchaseDate": "2022-01-02",
					"purchaseTime": "10_11",
					"total": "2.65",
					"items": [
						{"shortDescription": "Pepsi - 12-oz", "price": "1.25"},
						{"shortDescription": "Dasani", "price": "1.40"}
					]
				}`,
			expectJsonError:     true,
			expectValidateError: true,
		},
		{
			testName: "EmptyTotal",
			inputReceipt: `{
					"retailer": "Madison Fresh Market", 
					"purchaseDate": "2022-01-02",
					"purchaseTime": "10:11",
					"total": "",
					"items": [
						{"shortDescription": "Pepsi - 12-oz", "price": "1.25"},
						{"shortDescription": "Dasani", "price": "1.40"}
					]
				}`,
			expectJsonError:     false,
			expectValidateError: true,
		},
		{
			testName: "InvalidTotalCents",
			inputReceipt: `{
					"retailer": "Madison Fresh Market", 
					"purchaseDate": "2022-01-02",
					"purchaseTime": "10:11",
					"total": "2.655",
					"items": [
						{"shortDescription": "Pepsi - 12-oz", "price": "1.25"},
						{"shortDescription": "Dasani", "price": "1.40"}
					]
				}`,
			expectJsonError:     false,
			expectValidateError: true,
		},
		{
			testName: "InvalidTotalCentsTooShort",
			inputReceipt: `{
					"retailer": "Madison Fresh Market", 
					"purchaseDate": "2022-01-02",
					"purchaseTime": "10:11",
					"total": "2.5",
					"items": [
						{"shortDescription": "Pepsi - 12-oz", "price": "1.25"},
						{"shortDescription": "Dasani", "price": "1.40"}
					]
				}`,
			expectJsonError:     false,
			expectValidateError: true,
		},
		{
			testName: "InvalidTotalTwoPeriod",
			inputReceipt: `{
					"retailer": "Madison Fresh Market", 
					"purchaseDate": "2022-01-02",
					"purchaseTime": "10:11",
					"total": "2.65.65",
					"items": [
						{"shortDescription": "Pepsi - 12-oz", "price": "1.25"},
						{"shortDescription": "Dasani", "price": "1.40"}
					]
				}`,
			expectJsonError:     false,
			expectValidateError: true,
		},
		{
			testName: "InvalidTotalLetters",
			inputReceipt: `{
					"retailer": "Madison Fresh Market", 
					"purchaseDate": "2022-01-02",
					"purchaseTime": "10:11",
					"total": "a2.65",
					"items": [
						{"shortDescription": "Pepsi - 12-oz", "price": "1.25"},
						{"shortDescription": "Dasani", "price": "1.40"}
					]
				}`,
			expectJsonError:     false,
			expectValidateError: true,
		},
		{
			testName: "InvalidTotalNoPeriod",
			inputReceipt: `{
					"retailer": "Madison Fresh Market", 
					"purchaseDate": "2022-01-02",
					"purchaseTime": "10:11",
					"total": "265",
					"items": [
						{"shortDescription": "Pepsi - 12-oz", "price": "1.25"},
						{"shortDescription": "Dasani", "price": "1.40"}
					]
				}`,
			expectJsonError:     false,
			expectValidateError: true,
		},
		{
			testName: "EmptyItems",
			inputReceipt: `{
					"retailer": "Madison Fresh Market", 
					"purchaseDate": "2022-01-02",
					"purchaseTime": "10:11",
					"total": "265",
					"items": []
				}`,
			expectJsonError:     false,
			expectValidateError: true,
		},
		{
			testName: "EmptyItemDescription",
			inputReceipt: `{
					"retailer": "Madison Fresh Market", 
					"purchaseDate": "2022-01-02",
					"purchaseTime": "10:11",
					"total": "265",
					"items": [
						{"shortDescription": "", "price": "1.25"},
						{"shortDescription": "Dasani", "price": "1.40"}
					]
				}`,
			expectJsonError:     false,
			expectValidateError: true,
		},
		{
			testName: "InvalidItemDescription", // ! is invalid special character
			inputReceipt: `{
					"retailer": "Madison Fresh Market", 
					"purchaseDate": "2022-01-02",
					"purchaseTime": "10:11",
					"total": "265",
					"items": [
						{"shortDescription": "Pepsi - 12-oz!", "price": "1.25"},
						{"shortDescription": "Dasani", "price": "1.40"}
					]
				}`,
			expectJsonError:     false,
			expectValidateError: true,
		},
		{
			testName: "InvalidItemDescription", // & is invalid special character for shortDescription
			inputReceipt: `{
					"retailer": "Madison Fresh Market", 
					"purchaseDate": "2022-01-02",
					"purchaseTime": "10:11",
					"total": "265",
					"items": [
						{"shortDescription": "Pepsi &", "price": "1.25"},
						{"shortDescription": "Dasani", "price": "1.40"}
					]
				}`,
			expectJsonError:     false,
			expectValidateError: true,
		},
		{
			testName: "EmptyItemPrice",
			inputReceipt: `{
					"retailer": "Madison Fresh Market", 
					"purchaseDate": "2022-01-02",
					"purchaseTime": "10:11",
					"total": "265",
					"items": [
						{"shortDescription": "Pepsi - 12-oz", "price": "1.25"},
						{"shortDescription": "Dasani", "price": ""}
					]
				}`,
			expectJsonError:     false,
			expectValidateError: true,
		},
		{
			testName: "InvalidItemPrice",
			inputReceipt: `{
					"retailer": "Madison Fresh Market", 
					"purchaseDate": "2022-01-02",
					"purchaseTime": "10:11",
					"total": "265",
					"items": [
						{"shortDescription": "Pepsi - 12-oz", "price": "1.25"},
						{"shortDescription": "Dasani", "price": "$1.00"}
					]
				}`,
			expectJsonError:     false,
			expectValidateError: true,
		},
		{
			testName: "MissingRetailer",
			inputReceipt: `{
					"purchaseDate": "2022-01-02",
					"purchaseTime": "10:11",
					"total": "265",
					"items": [
						{"shortDescription": "Pepsi - 12-oz", "price": "1.25"},
						{"shortDescription": "Dasani", "price": "1.00"}
					]
				}`,
			expectJsonError:     false,
			expectValidateError: true,
		},
		{
			testName: "MissingPurchaseDate",
			inputReceipt: `{
					"retailer": "Madison Fresh Market", 
					"purchaseTime": "10:11",
					"total": "265",
					"items": [
						{"shortDescription": "Pepsi - 12-oz", "price": "1.25"},
						{"shortDescription": "Dasani", "price": "1.00"}
					]
				}`,
			expectJsonError:     false,
			expectValidateError: true,
		},
		{
			testName: "MissingPurchaseTime",
			inputReceipt: `{
					"retailer": "Madison Fresh Market", 
					"purchaseDate": "2022-01-02",
					"total": "265",
					"items": [
						{"shortDescription": "Pepsi - 12-oz", "price": "1.25"},
						{"shortDescription": "Dasani", "price": "1.00"}
					]
				}`,
			expectJsonError:     false,
			expectValidateError: true,
		},
		{
			testName: "MissingTotal",
			inputReceipt: `{
					"retailer": "Madison Fresh Market", 
					"purchaseDate": "2022-01-02",
					"purchaseTime": "10:11",
					"items": [
						{"shortDescription": "Pepsi - 12-oz", "price": "1.25"},
						{"shortDescription": "Dasani", "price": "1.00"}
					]
				}`,
			expectJsonError:     false,
			expectValidateError: true,
		},
		{
			testName: "MissingItems",
			inputReceipt: `{
					"retailer": "Madison Fresh Market", 
					"purchaseDate": "2022-01-02",
					"purchaseTime": "10:11",
					"total": "265"
				}`,
			expectJsonError:     false,
			expectValidateError: true,
		},
		{
			testName: "MissingShortDescription",
			inputReceipt: `{
					"retailer": "Madison Fresh Market", 
					"purchaseDate": "2022-01-02",
					"purchaseTime": "10:11",
					"total": "265",
					"items": [
						{"price": "1.25"},
						{"shortDescription": "Dasani", "price": "1.00"}
					]
				}`,
			expectJsonError:     false,
			expectValidateError: true,
		},
		{
			testName: "MissingPrice",
			inputReceipt: `{
					"retailer": "Madison Fresh Market", 
					"purchaseDate": "2022-01-02",
					"purchaseTime": "10:11",
					"total": "265",
					"items": [
						{"shortDescription": "Pepsi - 12-oz"},
						{"shortDescription": "Dasani", "price": "1.00"}
					]
				}`,
			expectJsonError:     false,
			expectValidateError: true,
		},
	}

	for _, test := range tests {
		t.Run(test.testName, func(t *testing.T) {
			var parsedInputReceipt Receipt
			if test.expectJsonError {
				assert.Error(t, json.Unmarshal([]byte(test.inputReceipt), &parsedInputReceipt))
			} else {
				assert.NoError(t, json.Unmarshal([]byte(test.inputReceipt), &parsedInputReceipt))
			}
			if test.expectValidateError {
				assert.Error(t, parsedInputReceipt.Validate())
			} else {
				assert.NoError(t, parsedInputReceipt.Validate())
			}
		})
	}
}
