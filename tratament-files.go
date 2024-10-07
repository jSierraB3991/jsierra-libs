package eliotlibs

import (
	"os"
	"strings"
)

func IsFileExist(dir string) bool {
	if _, err := os.Stat(dir); !os.IsNotExist(err) {
		return true
	}
	return false
}

func ReadFile(templateName string) string {
	exists := IsFileExist(templateName)
	if !exists {
		return ""
	}

	data, _ := os.ReadFile(templateName)
	return string(data)
}

// remplaza info en un archivo luego lo pasa a una variable
func ReplaceTextInFile(templateName string, MapForReplace map[string]string) string {
	input := ReadFile(templateName)
	for key, value := range MapForReplace {
		input = strings.Replace(input, key, value, -1)
	}
	return input
}
