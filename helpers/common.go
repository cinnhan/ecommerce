package helpers

import (
	"fmt"
	"github.com/joho/godotenv"
)

func init() {
	// load environment variables
	if err := LoadEnv(); err != nil {
		fmt.Printf("Error loading .env file: %v\n", err)
		return
	}
}

func LoadEnv() error {
	return godotenv.Load()
}
