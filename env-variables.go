package eliotlibs

import (
	"errors"
	"os"
)

func GetDataOfEnviroment(enviromentVariable string) string {
	return os.Getenv(enviromentVariable)
}
func GetDataOfEnviromentByDefault(enviromentVariable string, defaultValue string) string {
	data := GetDataOfEnviroment(enviromentVariable)
	if len(data) <= 0 {
		return defaultValue
	}
	return data
}

func GetDataOfEnviromentRequired(enviromentVariable string) (string, error) {
	data := GetDataOfEnviroment(enviromentVariable)
	if len(data) <= 0 {
		return "", errors.New("VARIABLE_NOT_FOUND_" + enviromentVariable)
	}
	return data, nil
}
