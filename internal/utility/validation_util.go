package utility

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"receipt-processor/internal/logger"
	"regexp"
	"strconv"
	"time"

	"github.com/sirupsen/logrus"
)

// Regex patterns compiled once for efficiency.
var (
	retailerRegex         = regexp.MustCompile(`^[\w\s\-&]+$`)
	shortdescriptionRegex = regexp.MustCompile(`^[\w\s\-]+$`)
	priceRegex            = regexp.MustCompile(`^\d+\.\d{2}$`)
	dateRegex             = regexp.MustCompile(`^\d{4}-\d{2}-\d{2}$`)
	timeRegex             = regexp.MustCompile(`^(2[0-3]|[01][0-9]):([0-5][0-9])$`)
)

// IsValidRetailerName validates the retailer name against a regular expression.
func IsValidRetailerName(str string) bool {
	valid := retailerRegex.MatchString(str)
	if !valid {
		logger.Error("Invalid retailer name", logrus.Fields{
			"retailer_name": str,
		})
	}
	return valid
}

// IsValidShortDescription validates the item description format.
func IsValidShortDescription(str string) bool {
	valid := shortdescriptionRegex.MatchString(str)
	if !valid {
		logger.Error("Invalid short description", logrus.Fields{
			"short_description": str,
		})
	}
	return valid
}

// IsValidPrice validates the price format.
func IsValidPrice(str string) bool {
	// Check format using regex
	if !priceRegex.MatchString(str) {
		logger.Error("Invalid price format", logrus.Fields{
			"price": str,
		})
		return false
	}

	// Parse the price to a float to ensure it's non-negative
	price, err := strconv.ParseFloat(str, 64)
	if err != nil || price < 0 {
		logger.Error("Invalid price value", logrus.Fields{
			"price": str,
			"error": err,
		})
		return false
	}

	return true
}

// IsValidDate validates the date format.
func IsValidDate(str string) bool {
	if !dateRegex.MatchString(str) {
		logger.Error("Invalid date format", logrus.Fields{
			"date": str,
		})
		return false
	}

	// Parse the date string into a time.Time object
	parsedDate, err := time.Parse("2006-01-02", str)
	if err != nil {
		logger.Error("Date parsing error", logrus.Fields{
			"date":  str,
			"error": err,
		})
		return false
	}

	// Check if the parsed date is not in the future
	today := time.Now().Truncate(24 * time.Hour) // Truncate to remove time portion
	if parsedDate.After(today) {
		logger.Error("Date is in the future", logrus.Fields{
			"date": str,
		})
		return false
	}

	return true
}

// IsValidTime validates the time format.
func IsValidTime(str string) bool {
	valid := timeRegex.MatchString(str)
	if !valid {
		logger.Error("Invalid time format", logrus.Fields{
			"time": str,
		})
	}
	return valid
}

// ReadBody reads the full request body into a byte slice.
func ReadBody(r *http.Request) ([]byte, error) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logger.Error("Error reading body", logrus.Fields{
			"error": err,
		})
		return nil, err
	}
	defer r.Body.Close()
	return body, nil
}

// ParseJSON parses the JSON-encoded data and stores the result in the value pointed to by target.
func ParseJSON(body []byte, target any) error {
	if err := json.Unmarshal(body, target); err != nil {
		logger.Error("JSON parsing error", logrus.Fields{
			"error": err,
		})
		return err
	}
	return nil
}
