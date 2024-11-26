package utility

import (
	"testing"
)

func TestGenerateID(t *testing.T) {
	id1 := GenerateID()
	id2 := GenerateID()
	if id1 == "" || id2 == "" {
		t.Error("Generated IDs should not be empty")
	}
	if id1 == id2 {
		t.Error("Generated IDs should be unique")
	}
}
