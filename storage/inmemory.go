package storage

import (
	"receipts/models"
	"sync"

	"github.com/google/uuid"
)

type ReceiptStorage struct {
	*sync.RWMutex
	idToReceipt map[uuid.UUID]*models.Receipt
}

func NewReceiptStorage() *ReceiptStorage {
	return &ReceiptStorage{
		RWMutex:     &sync.RWMutex{},
		idToReceipt: make(map[uuid.UUID]*models.Receipt),
	}
}

/*
If receipt exists, returns the receipt, otherwise returns nil.

Waits for read lock.
*/
func (rs *ReceiptStorage) GetReceipt(id uuid.UUID) *models.Receipt {
	rs.RLock()
	defer rs.RUnlock()
	return rs.idToReceipt[id]
}

/*
Saves the id to receipt mapping after waiting for the read / write lock.
*/
func (rs *ReceiptStorage) SetReceipt(id uuid.UUID, receipt *models.Receipt) {
	rs.Lock()
	defer rs.Unlock()
	rs.idToReceipt[id] = receipt
}
