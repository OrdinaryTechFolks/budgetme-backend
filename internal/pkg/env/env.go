package env

import (
	"fmt"
	"os"
	"strconv"
)

func Get(key string, defaultValue string) string {
	value, exists := os.LookupEnv(key)

	if !exists || len(value) == 0 {
		return defaultValue
	}

	return value
}

func GetString(key string, defaultValue string) string {
	return Get(key, defaultValue)
}

func GetInt(key string, defaultValue string) int {
	value := Get(key, defaultValue)

	intValue, err := strconv.Atoi(value)
	if err != nil {
		panic(fmt.Sprintf("Env with key %s is not a valid number", key))
	}

	return intValue
}
