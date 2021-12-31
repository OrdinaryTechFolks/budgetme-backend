package env

import (
	"fmt"
	"os"
	"strconv"
)

func GetString(key string, defaultValue string) string {
	value, exists := os.LookupEnv(key)

	if !exists || len(value) == 0 {
		return defaultValue
	}

	return value
}

func GetInt(key string, defaultValue int) int {
	value, exists := os.LookupEnv(key)

	if !exists || len(value) == 0 {
		return defaultValue
	}

	intValue, err := strconv.Atoi(value)
	if err != nil {
		panic(fmt.Sprintf("Env with key %s is not a valid number", key))
	}

	return intValue
}
