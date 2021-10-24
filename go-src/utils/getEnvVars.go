package utils

import (
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func GetEnvVar(key string) string {
	if strings.HasSuffix(os.Args[0], ".test") {
		godotenv.Load("../../../.env")
	}
	return os.Getenv(key)
}
