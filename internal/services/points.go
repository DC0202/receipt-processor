package services

import (
	"math"
	"receipt-processor/internal/logger"
	"receipt-processor/internal/model"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

func CalculatePoints(receipt model.Receipt) (int, string) {
	points := 0
	explanation := "Breakdown:\n"

	// Alphanumeric characters in the retailer name
	retailerChars := regexp.MustCompile(`[a-zA-Z0-9]+`).FindAllString(receipt.Retailer, -1)
	numChars := 0
	for _, match := range retailerChars {
		numChars += len(match) // Count characters in each match
	}
	points += numChars
	explanation += strconv.Itoa(numChars) + " points - retailer name has " + strconv.Itoa(numChars) + " alphanumeric characters\n"
	logger.Info("Calculated points for retailer name", logrus.Fields{
		"retailer": receipt.Retailer,
		"points":   numChars,
	})

	// Round dollar amount
	total, err := strconv.ParseFloat(receipt.Total, 64)
	if err != nil {
		logger.Error("Error parsing total amount", logrus.Fields{
			"total": receipt.Total,
			"error": err,
		})
	} else {
		if total == math.Floor(total) {
			points += 50
			explanation += "50 points - total is a round dollar amount\n"
		}
		if math.Mod(total*100, 25) == 0 {
			points += 25
			explanation += "25 points - total is a multiple of 0.25\n"
		}
		logger.Info("Added points for total", logrus.Fields{
			"round_total_points":      50,
			"multiple_of_0.25_points": 25,
			"total":                   total,
		})
	}

	// Points for every two items
	itemPairs := len(receipt.Items) / 2
	points += itemPairs * 5
	explanation += strconv.Itoa(itemPairs*5) + " points - " + strconv.Itoa(len(receipt.Items)) + " items (2 pairs @ 5 points each)\n"
	logger.Info("Added points for item pairs", logrus.Fields{
		"item_count": len(receipt.Items),
		"points":     itemPairs * 5,
	})

	// Points based on item descriptions
	for _, item := range receipt.Items {
		trimmedDescription := strings.TrimSpace(item.ShortDescription)
		if len(trimmedDescription)%3 == 0 {
			itemPrice, err := strconv.ParseFloat(item.Price, 64)
			if err != nil {
				logger.Error("Error parsing item price", logrus.Fields{
					"price": item.Price,
					"error": err,
				})
			} else {
				itemPoints := int(math.Ceil(itemPrice * 0.2))
				points += itemPoints
				explanation += strconv.Itoa(itemPoints) + " Points - \"" + trimmedDescription + "\" is " +
					strconv.Itoa(len(trimmedDescription)) + " characters (a multiple of 3)\n" +
					"          item price of " + item.Price + " * 0.2 = " + strconv.FormatFloat(itemPrice*0.2, 'f', 2, 64) + ", rounded up is " + strconv.Itoa(itemPoints) + " points\n"
				logger.Info("Added points for item description", logrus.Fields{
					"item_description": trimmedDescription,
					"item_price":       itemPrice,
					"points":           itemPoints,
				})
			}
		}
	}

	// Points for odd purchase date
	date, err := time.Parse("2006-01-02", receipt.PurchaseDate)
	if err != nil {
		logger.Error("Error parsing purchase date", logrus.Fields{
			"purchase_date": receipt.PurchaseDate,
			"error":         err,
		})
	} else if date.Day()%2 != 0 {
		points += 6
		explanation += "6 points - purchase day is odd\n"
		logger.Info("Added points for odd purchase day", logrus.Fields{
			"purchase_date": receipt.PurchaseDate,
			"points":        6,
		})
	}

	// Points for time between 2:00 PM and 4:00 PM
	purchaseTime, err := time.Parse("15:04", receipt.PurchaseTime)
	if err != nil {
		logger.Error("Error parsing purchase time", logrus.Fields{
			"purchase_time": receipt.PurchaseTime,
			"error":         err,
		})
	} else {
		hour := purchaseTime.Hour()
		if hour >= 14 && hour < 16 {
			points += 10
			explanation += "10 points - time of purchase is between 2:00pm and 4:00pm\n"
			logger.Info("Added points for time of purchase", logrus.Fields{
				"purchase_time": receipt.PurchaseTime,
				"points":        10,
			})
		}
	}

	explanation += "  + ---------\n  = " + strconv.Itoa(points) + " points"
	logger.Info("Total points calculated", logrus.Fields{
		"total_points": points,
		"explanation":  explanation,
	})
	return points, explanation
}
