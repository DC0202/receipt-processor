package model

import (
	"errors"
	"fmt"
	"math"
	"receipt-processor/internal/logger"
	"receipt-processor/internal/utility"
	"strconv"

	"github.com/sirupsen/logrus"
)

type Receipt struct {
	Retailer     string `json:"retailer"`
	PurchaseDate string `json:"purchaseDate"`
	PurchaseTime string `json:"purchaseTime"`
	Items        []Item `json:"items"`
	Total        string `json:"total"`
}

type Item struct {
	ShortDescription string `json:"shortDescription"`
	Price            string `json:"price"`
}

// ReceiptDetails holds the results of processing a receipt, including points and explanation.
type ReceiptDetails struct {
	Points      int    `json:"points"`
	Explanation string `json:"explanation"`
}

// ValidateReceiptMap checks if the provided data map has all the required keys for a Receipt.
func (r Receipt) ValidateReceiptMap(dataMap map[string]interface{}) error {
	validKeys := []string{"retailer", "purchaseDate", "purchaseTime", "items", "total"}
	for key := range dataMap {
		if !contains(validKeys, key) {
			logger.Error("Invalid key in data map", logrus.Fields{
				"key":       key,
				"validKeys": validKeys,
			})
			return errors.New("incorrect receipt data")
		}
	}
	return nil
}

// Validate checks the integrity of the receipt data.
func (r Receipt) Validate() error {
	if !utility.IsValidRetailerName(r.Retailer) {
		logger.Error("Invalid retailer name", logrus.Fields{
			"retailer": r.Retailer,
		})
		return errors.New("retailer field is invalid")
	}
	if !utility.IsValidDate(r.PurchaseDate) {
		logger.Error("Invalid purchase date", logrus.Fields{
			"purchaseDate": r.PurchaseDate,
		})
		return errors.New("purchase date field is required and must be in YYYY-MM-DD format")
	}
	if !utility.IsValidTime(r.PurchaseTime) {
		logger.Error("Invalid purchase time", logrus.Fields{
			"purchaseTime": r.PurchaseTime,
		})
		return errors.New("purchase time field is required and must be in HH:MM format")
	}
	if !utility.IsValidPrice(r.Total) {
		logger.Error("Invalid total price format", logrus.Fields{
			"total": r.Total,
		})
		return errors.New("total price format is invalid, should be numeric with two decimal places")
	}
	if len(r.Items) == 0 {
		logger.Error("No items found in receipt", logrus.Fields{})
		return errors.New("at least one item is required")
	}

	total, err := strconv.ParseFloat(r.Total, 64)
	if err != nil {
		logger.Error("Error parsing total price", logrus.Fields{
			"total": r.Total,
		})
		return errors.New("total price is not a valid number")
	}

	var sum float64 = 0
	for _, item := range r.Items {
		if err := item.Validate(); err != nil {
			logger.Error("Item validation error", logrus.Fields{
				"item":  item,
				"error": err,
			})
			return fmt.Errorf("item validation error: %v", err)
		}
		price, err := strconv.ParseFloat(item.Price, 64)
		if err != nil {
			logger.Error("Error parsing item price", logrus.Fields{
				"price": item.Price,
			})
			return errors.New("item price is not a valid number")
		}
		sum += price
	}

	if math.Abs(total-sum) > 0.001 {
		logger.Error("Total does not match sum of items", logrus.Fields{
			"total": total,
			"sum":   sum,
		})
		return errors.New("total does not match the sum of item prices")
	}

	return nil
}

// Validate checks the integrity of item data.
func (i Item) Validate() error {
	if !utility.IsValidShortDescription(i.ShortDescription) {
		logger.Error("Invalid short description", logrus.Fields{
			"shortDescription": i.ShortDescription,
		})
		return errors.New("item short description is invalid")
	}
	if !utility.IsValidPrice(i.Price) {
		logger.Error("Invalid item price format", logrus.Fields{
			"price": i.Price,
		})
		return errors.New("item price format is invalid, should be numeric with two decimal places")
	}
	return nil
}

// contains checks if a string is present in a slice of strings.
func contains(slice []string, str string) bool {
	for _, v := range slice {
		if v == str {
			return true
		}
	}
	return false
}

func (r Receipt) String() string {
	return fmt.Sprintf("%s-%s-%s-%s-%v", r.Retailer, r.PurchaseDate, r.PurchaseTime, r.Total, r.Items)
}
