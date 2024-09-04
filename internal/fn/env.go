package fn

import (
	"os"
	"strconv"
	"strings"
)

func GetEnvStr(key, defaultValue string) string {
	v, ok := os.LookupEnv(key)
	if !ok {
		return defaultValue
	}
	value := strings.TrimSpace(v)
	return value
}

func GetEnvBool(key string, defaultValue bool) bool {
	value := strings.TrimSpace(os.Getenv(key))
	if value == "" {
		return defaultValue
	}
	if value == "true" || value == "1" || value == "on" || value == "yes" {
		return true
	}
	return false
}

func GetEnvInt(key string, defaultValue int) int {
	value := strings.TrimSpace(os.Getenv(key))
	if value == "" {
		return defaultValue
	}

	i, err := strconv.Atoi(value)
	if err != nil {
		return defaultValue
	}
	return i
}
