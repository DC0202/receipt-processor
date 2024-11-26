package hash

import (
	"crypto/sha1"
	"fmt"
	"receipt-processor/internal/logger"

	"github.com/sirupsen/logrus"
)

// GenerateHash computes a SHA-1 hash for the provided data string.
func GenerateHash(data string) string {
	h := sha1.New()
	h.Write([]byte(data))
	bs := h.Sum(nil)
	hash := fmt.Sprintf("%x", bs)
	logger.Info("Generated hash", logrus.Fields{
		"input_data_length": len(data),
		"generated_hash":    hash,
	})
	return hash
}
