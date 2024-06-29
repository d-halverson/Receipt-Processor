package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"receipts/models"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestProcessReceipt(t *testing.T) {
	router := CreateRouter()
	tests := []struct {
		testName       string
		inputReceipt   string
		expectedStatus int
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
			expectedStatus: http.StatusOK,
		},
		{
			testName: "InvalidJson", // missing closing bracket
			inputReceipt: `{
					"retailer": "Target",
					"purchaseDate": "2022-01-02",
					"purchaseTime": "13:13",
					"total": "1.25",
					"items": [
						{"shortDescription": "Pepsi - 12-oz", "price": "1.25"}
					]
				`,
			expectedStatus: http.StatusBadRequest,
		},
		{
			testName: "InvalidDate",
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
			expectedStatus: http.StatusBadRequest,
		},
		{
			testName: "InvalidTotal",
			inputReceipt: `{
					"retailer": "Madison Fresh Market", 
					"purchaseDate": "2022-01-10",
					"purchaseTime": "08:13",
					"total": "2.652",
					"items": [
						{"shortDescription": "Pepsi - 12-oz", "price": "1.25"},
						{"shortDescription": "Dasani", "price": "1.40"}
					]
				}`,
			expectedStatus: http.StatusBadRequest,
		},
	}

	for _, test := range tests {
		t.Run(test.testName, func(t *testing.T) {
			req, err := http.NewRequest("POST", "/receipts/process", bytes.NewBuffer([]byte(test.inputReceipt)))
			assert.NoError(t, err)

			responseRecorder := httptest.NewRecorder()
			router.ServeHTTP(responseRecorder, req)

			assert.Equal(t, test.expectedStatus, responseRecorder.Code)

			// In 200 case, make sure response is in correct format
			if test.expectedStatus == http.StatusOK {
				var response models.Id
				err := json.NewDecoder(responseRecorder.Body).Decode(&response)
				assert.NoError(t, err)
				_, err = uuid.Parse(response.Id)
				assert.NoError(t, err)
			}
		})
	}
}

/*
Only worried about testing getting receipt when it is present or not.
The calculation of points is tested directly on CalculatePoints() function.
*/
func TestGetPoints(t *testing.T) {
	router := CreateRouter()
	var responseRecorder *httptest.ResponseRecorder

	// Test GetPoints on invalid uuid format
	t.Run("InvalidIdFormat", func(t *testing.T) {
		responseRecorder = httptest.NewRecorder()
		req, err := http.NewRequest("GET", "/receipts/1234/points", nil)
		assert.NoError(t, err)
		router.ServeHTTP(responseRecorder, req)
		assert.Equal(t, http.StatusNotFound, responseRecorder.Code)
	})

	// Test GetPoints of non existent receipt
	t.Run("NonExistentReceipt", func(t *testing.T) {
		responseRecorder = httptest.NewRecorder()
		req, err := http.NewRequest("GET", "/receipts/"+uuid.New().String()+"/points", nil)
		assert.NoError(t, err)
		router.ServeHTTP(responseRecorder, req)
		assert.Equal(t, http.StatusNotFound, responseRecorder.Code)
	})

	// Invalid process of receipt, then test GetPoints of non existent receipt again
	t.Run("NonExistentReceiptAfterProcess", func(t *testing.T) {
		// Invalid process
		invalidDate := `{
							"retailer": "Madison Fresh Market", 
							"purchaseDate": "2022-01-1",
							"purchaseTime": "08:13",
							"total": "2.65",
							"items": [
								{"shortDescription": "Pepsi - 12-oz", "price": "1.25"},
								{"shortDescription": "Dasani", "price": "1.40"}
							]
						}`
		responseRecorder = httptest.NewRecorder()
		req, err := http.NewRequest("POST", "/receipts/process", bytes.NewBuffer([]byte(invalidDate)))
		assert.NoError(t, err)
		router.ServeHTTP(responseRecorder, req)

		// Call points
		responseRecorder = httptest.NewRecorder()
		req, err = http.NewRequest("GET", "/receipts/"+uuid.New().String()+"/points", nil)
		assert.NoError(t, err)
		router.ServeHTTP(responseRecorder, req)
		assert.Equal(t, http.StatusNotFound, responseRecorder.Code)
	})

	t.Run("ExistingReceipt", func(t *testing.T) {
		// Valid process of receipt
		ValidReceipt := `{
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
						}`
		responseRecorder = httptest.NewRecorder()
		req, err := http.NewRequest("POST", "/receipts/process", bytes.NewBuffer([]byte(ValidReceipt)))
		assert.NoError(t, err)
		router.ServeHTTP(responseRecorder, req)
		var responseId models.Id
		err = json.NewDecoder(responseRecorder.Body).Decode(&responseId)
		assert.NoError(t, err)
		_, err = uuid.Parse(responseId.Id)
		assert.NoError(t, err)

		// Call points
		responseRecorder = httptest.NewRecorder()
		req, err = http.NewRequest("GET", "/receipts/"+responseId.Id+"/points", nil)
		assert.NoError(t, err)
		router.ServeHTTP(responseRecorder, req)
		assert.Equal(t, http.StatusOK, responseRecorder.Code)
		var responsePoints models.Points
		err = json.NewDecoder(responseRecorder.Body).Decode(&responsePoints)
		assert.NoError(t, err)
		assert.Equal(t, 28, responsePoints.Points)
	})
}
