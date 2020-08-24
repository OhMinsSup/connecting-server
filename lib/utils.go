package lib

import "os"

func GetEnvWithKey(key string) string {
	return os.Getenv(key)
}
