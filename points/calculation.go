package points

import "receipts/models"

func CalculatePoints(receipt *models.Receipt) int {
	points := 0

	for _, rule := range GetReceiptRules() {
		points = points + rule(receipt)
	}

	return points
}
