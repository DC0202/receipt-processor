package hash

import (
	"testing"
)

func TestGenerateHash(t *testing.T) {
	data := "test data"
	hash1 := GenerateHash(data)
	if hash1 == "" {
		t.Errorf("GenerateHash() should not return an empty string")
	}

	// Ensure consistent hash generation
	hash2 := GenerateHash(data)
	if hash1 != hash2 {
		t.Errorf("GenerateHash() should return consistent hash values for the same input")
	}
}
