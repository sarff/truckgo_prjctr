package database

import (
	"github.com/spf13/viper"
	"testing"
)

func TestInitialize(t *testing.T) {
	t.Run("Fail to load .env file", func(t *testing.T) {
		viper.Set("env_path", "nonexistent.env")
		viper.Set("db_type", "sqlite")
		_, err := Initialize()
		if err == nil {
			t.Errorf("Expected error when loading nonexistent .env file, got nil")
		}
	})

	t.Run("Invalid database type", func(t *testing.T) {
		viper.Set("env_path", ".env")
		viper.Set("db_type", "invalid_db_type")
		_, err := Initialize()
		if err == nil {
			t.Errorf("Expected error for invalid database type, got nil")
		}
	})
}
