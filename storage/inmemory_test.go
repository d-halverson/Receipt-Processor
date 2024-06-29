package storage

import (
	"encoding/json"
	"receipts/models"
	"strconv"
	"sync"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestGetReceipt(t *testing.T) {
	receiptStorage := NewReceiptStorage()

	// Should be nil when not set yet
	id := uuid.New()
	assert.Nil(t, receiptStorage.GetReceipt(id))

	rawReceipt := `{
					"retailer": "Walgreens",
					"purchaseDate": "2022-01-02",
					"purchaseTime": "08:13",
					"total": "2.65",
					"items": [
						{"shortDescription": "Pepsi - 12-oz", "price": "1.25"},
						{"shortDescription": "Dasani", "price": "1.40"}
					]
				}`
	var receipt models.Receipt
	assert.NoError(t, json.Unmarshal([]byte(rawReceipt), &receipt))
	receiptStorage.SetReceipt(id, &receipt)

	// Should be found now that it is set
	assert.Equal(t, &receipt, receiptStorage.GetReceipt(id))
}

// Testing that getting many in sequence does not cause deadlock
func TestGetReceiptMany(t *testing.T) {
	receiptStorage := NewReceiptStorage()
	id := uuid.New()

	rawReceipt := `{
					"retailer": "Walgreens",
					"purchaseDate": "2022-01-02",
					"purchaseTime": "08:13",
					"total": "2.65",
					"items": [
						{"shortDescription": "Pepsi - 12-oz", "price": "1.25"},
						{"shortDescription": "Dasani", "price": "1.40"}
					]
				}`
	var receipt models.Receipt
	assert.NoError(t, json.Unmarshal([]byte(rawReceipt), &receipt))
	receiptStorage.SetReceipt(id, &receipt)

	for i := 0; i < 1000; i++ {
		assert.Equal(t, &receipt, receiptStorage.GetReceipt(id))
	}
}

// Testing that setting twice updates the stored receipt
func TestSetReceipt(t *testing.T) {
	receiptStorage := NewReceiptStorage()

	id := uuid.New()

	rawReceipt := `{
					"retailer": "Walgreens",
					"purchaseDate": "2022-01-02",
					"purchaseTime": "08:13",
					"total": "2.65",
					"items": [
						{"shortDescription": "Pepsi - 12-oz", "price": "1.25"},
						{"shortDescription": "Dasani", "price": "1.40"}
					]
				}`
	var receipt models.Receipt
	assert.NoError(t, json.Unmarshal([]byte(rawReceipt), &receipt))
	receiptStorage.SetReceipt(id, &receipt)
	assert.Equal(t, &receipt, receiptStorage.GetReceipt(id))

	// Update receipt
	receipt.Retailer = "Madison Fresh Market"
	receiptStorage.SetReceipt(id, &receipt)
	assert.Equal(t, &receipt, receiptStorage.GetReceipt(id))
}

// Testing that setting many updates and getting in sequence does not cause deadlock
func TestSetReceiptMany(t *testing.T) {
	receiptStorage := NewReceiptStorage()
	id := uuid.New()

	rawReceipt := `{
					"retailer": "Walgreens",
					"purchaseDate": "2022-01-02",
					"purchaseTime": "08:13",
					"total": "2.65",
					"items": [
						{"shortDescription": "Pepsi - 12-oz", "price": "1.25"},
						{"shortDescription": "Dasani", "price": "1.40"}
					]
				}`
	var receipt models.Receipt
	assert.NoError(t, json.Unmarshal([]byte(rawReceipt), &receipt))
	receiptStorage.SetReceipt(id, &receipt)

	for i := 0; i < 1000; i++ {
		receipt.Retailer = strconv.Itoa(i)
		receiptStorage.SetReceipt(id, &receipt)
		assert.Equal(t, &receipt, receiptStorage.GetReceipt(id))
	}
}

/*
Testing many concurrent sets and gets to ensure no deadlock is possible

This will fail if deadlock occurs due to timeout.
*/
func TestSetGetConcurrent(t *testing.T) {
	receiptStorage := NewReceiptStorage()
	id := uuid.New()

	rawReceipt := `{
					"retailer": "Walgreens",
					"purchaseDate": "2022-01-02",
					"purchaseTime": "08:13",
					"total": "2.65",
					"items": [
						{"shortDescription": "Pepsi - 12-oz", "price": "1.25"},
						{"shortDescription": "Dasani", "price": "1.40"}
					]
				}`
	var receipt models.Receipt
	assert.NoError(t, json.Unmarshal([]byte(rawReceipt), &receipt))

	var wg sync.WaitGroup
	wg.Add(2)

	// Starting many sets
	go func() {
		defer wg.Done()
		for i := 0; i < 10000; i++ {
			receipt.Retailer = strconv.Itoa(i)
			receiptStorage.SetReceipt(id, &receipt)
		}
	}()

	// Starting many gets
	go func() {
		defer wg.Done()
		for i := 0; i < 10000; i++ {
			receiptStorage.GetReceipt(id)
		}
	}()

	wg.Wait()
}
