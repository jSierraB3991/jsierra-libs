package eliotlibs

import (
	"errors"
	"os"
)

func GetDataOfEnviroment(enviromentVariable string) string {
	return os.Getenv(enviromentVariable)
}

func GetDataOfEnviromentRequired(enviromentVariable string) (string, error) {
	data := GetDataOfEnviroment(enviromentVariable)
	if len(data) <= 0 {
		return "", errors.New("VARIABLE_NOT_FOUND_" + enviromentVariable)
	}
	return data, nil
}
