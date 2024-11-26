package utility

import (
	"testing"
)

// Test IsValidRetailerName
func TestIsValidRetailerName(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"M&M Corner Market", true},
		{"Retailer_123", true},
		{"Invalid@Name", false}, // Special character '@' is invalid
		{"", false},             // Empty string is invalid
	}

	for _, test := range tests {
		result := IsValidRetailerName(test.input)
		if result != test.expected {
			t.Errorf("IsValidRetailerName(%q) = %v; want %v", test.input, result, test.expected)
		}
	}
}

// Test IsValidShortDescription
func TestIsValidShortDescription(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"Gatorade", true},
		{"Energy Drink 250ml", true},
		{"Invalid@Description", false}, // Special character '@' is invalid
		{"", false},                    // Empty string is invalid
	}

	for _, test := range tests {
		result := IsValidShortDescription(test.input)
		if result != test.expected {
			t.Errorf("IsValidShortDescription(%q) = %v; want %v", test.input, result, test.expected)
		}
	}
}

// Test IsValidPrice
func TestIsValidPrice(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"9.99", true},
		{"100.00", true},
		{"invalid", false}, // Non-numeric string is invalid
		{"10.1", false},    // Incorrect decimal places
		{"", false},        // Empty string is invalid
	}

	for _, test := range tests {
		result := IsValidPrice(test.input)
		if result != test.expected {
			t.Errorf("IsValidPrice(%q) = %v; want %v", test.input, result, test.expected)
		}
	}
}

// Test IsValidDate
func TestIsValidDate(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"2022-03-20", true},
		{"20-03-2022", false}, // Incorrect format
		{"2022/03/20", false}, // Incorrect separator
		{"", false},           // Empty string is invalid
	}

	for _, test := range tests {
		result := IsValidDate(test.input)
		if result != test.expected {
			t.Errorf("IsValidDate(%q) = %v; want %v", test.input, result, test.expected)
		}
	}
}

// Test IsValidTime
func TestIsValidTime(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"14:33", true},
		{"2:33 PM", false}, // Incorrect format
		{"14:60", false},   // Invalid minutes
		{"", false},        // Empty string is invalid
	}

	for _, test := range tests {
		result := IsValidTime(test.input)
		if result != test.expected {
			t.Errorf("IsValidTime(%q) = %v; want %v", test.input, result, test.expected)
		}
	}
}
