package lib

import "os"

type JSON = map[string]interface{}

func GetEnvWithKey(key string) string {
	return os.Getenv(key)
}
