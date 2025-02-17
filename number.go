package eliotlibs

import (
	"fmt"
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

func GetStrinToFloat(numFloat float32) string {
	return GetStrinToFloat64(float64(numFloat))
}

func GetStrinToFloat64(numFloat float64) string {
	return strconv.FormatFloat(numFloat, 'f', 2, 32)
}

type InvalidConvertStringToFloat struct{}

func (InvalidConvertStringToFloat) Error() string {
	return "InvalidConvertStringToFloat"
}

func GetStringtofloat32(numString string) (float32, error) {
	floatValue, err := strconv.ParseFloat(numString, 32)
	if err != nil {
		fmt.Println("Error converting string to float32:", err)
		return float32(0), InvalidConvertStringToFloat{}
	}

	// Cast the float64 result to float32
	return float32(floatValue), nil
}
