package configs

import (
	"os"
	"log"
)

func InitCredentials(apiKeyEnv string) (string) {
	envKey := os.Getenv(apiKeyEnv)
	if len(envKey) < 1 {
		log.Fatalf("Error: Env var %q must contain a valid MTA API key", apiKeyEnv)
	}

	return envKey
}
