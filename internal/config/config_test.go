package config

// import (
// 	"os"
// 	"testing"
// )

// func TestLoadConfig(t *testing.T) {
// 	const defaultPort = "8080"
// 	os.Setenv("PORT", "3000")
// 	defer os.Unsetenv("PORT")

// 	cfg := LoadConfig()
// 	if cfg.Port != "3000" {
// 		t.Errorf("Expected port %v, got %v", "3000", cfg.Port)
// 	}

// 	os.Unsetenv("PORT") // Test default value
// 	cfg = LoadConfig()
// 	if cfg.Port != defaultPort {
// 		t.Errorf("Expected default port %v, got %v", defaultPort, cfg.Port)
// 	}
// }
