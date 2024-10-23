package eliotlibs

import (
	"log"
	"strconv"
)

func GetNumberForString(number string) int {
	response, err := strconv.Atoi(number)
	if err != nil {
		log.Println(number)
		log.Println("GetNumberForString")
		return 0
	}
	return response
}

func GetUNumberForString(number string) uint {
	response, err := strconv.Atoi(number)
	if err != nil {
		log.Println(number)
		log.Println("GetNumberForString")
		return 0
	}
	return uint(response)
}

func GetFloatStringToUInt(data uint) string {
	// Convierte el flotante a una cadena
	return strconv.FormatFloat(GetFloatToUInt(data), 'f', -1, 64)
}

func GetFloatStringToInt(data int) string {
	// Convierte el flotante a una cadena
	return strconv.FormatFloat(GetFloatToInt(data), 'f', -1, 64)
}

func GetFloatToUInt(data uint) float64 {
	return float64(data)
}

func GetFloatToInt(data int) float64 {
	return float64(data)
}
