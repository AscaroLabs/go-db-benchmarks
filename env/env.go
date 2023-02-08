package env

import (
	"os"
	"strconv"
)

// All variables for project
var (
	Mode = GetterInt("MODE", 0)

	DbHost     = Getter("DB_HOST", "")
	DbPort     = Getter("DB_PORT", "")
	DbUser     = Getter("DB_USER", "")
	DbName     = Getter("DB_NAME", "")
	DbPassword = Getter("DB_PASSWORD", "")

	Workers = GetterInt("WORKERS", 1)

	TntUser     = Getter("TDB_USER", "")
	TntPassword = Getter("TDB_PASSWORD", "")
	TntHost     = Getter("TDB_HOST", "")
	TntPort     = Getter("TDB_PORT", "")
)

// Getter -
func Getter(key, defaultValue string) string {
	env, ok := os.LookupEnv(key)
	if ok {
		return env
	}
	return defaultValue
}

// GetterInt -
func GetterInt(key string, defaultValue int) int {
	env, ok := os.LookupEnv(key)
	if ok {
		res, err := strconv.ParseInt(env, 10, 32)
		if err == nil {
			return int(res)
		}
	}
	return defaultValue
}

func GetterBool(key string, defaultValue bool) bool {
	env, ok := os.LookupEnv(key)
	if ok {
		res, err := strconv.ParseBool(env)
		if err == nil {
			return res
		}
	}
	return defaultValue
}
