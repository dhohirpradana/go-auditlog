package main

import (
	"fmt"
	"os"
)

func ValidateEnvVars() error {
	requiredVars := []string{
		"ES_URL",
		"ES_USER",
		"ES_PASS",
		"ES_INDEX",
	}

	for _, key := range requiredVars {
		value, exists := os.LookupEnv(key)
		if !exists || value == "" {
			return fmt.Errorf("environment variable %s is not set or is empty", key)
		}
	}
	return nil
}
