package utils

import (
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/joho/godotenv"
)

func GetEnvVar(key string) string {
	if strings.HasSuffix(os.Args[0], ".test") {
		godotenv.Load("../../../.env")
	} else if strings.HasSuffix(os.Args[0], "devServer") {
		_, b, _, _ := runtime.Caller(0)
		envFilePath := filepath.Join(filepath.Dir(b), "../../../.env")
		godotenv.Load(envFilePath)
	}
	return os.Getenv(key)
}
