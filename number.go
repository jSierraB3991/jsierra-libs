package eliotlibs

import (
	"fmt"
	"log"
	"math"
	"strconv"

	"golang.org/x/text/message"
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

func GetStringToInt(data int) string {
	// Convierte el flotante a una cadena
	return strconv.FormatFloat(GetFloatToInt(data), 'f', -1, 64)
}

func GetStringToInt64(number int64) string {
	return strconv.Itoa(int(number))
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

func GetFloatToString(numString string) (float64, error) {
	floatValue, err := strconv.ParseFloat(numString, 32)
	if err != nil {
		fmt.Println("Error converting string to float32:", err)
		return 0, InvalidConvertStringToFloat{}
	}
	return floatValue, nil
}

func GetStringtofloat32(numString string) (float32, error) {
	floatValue, err := GetFloatToString(numString)
	if err != nil {
		return 0, err
	}

	return float32(floatValue), nil
}

func GetFormatMoneyToFloat32(price float32) string {
	return GetFormatMoneyToFloat(float64(price))
}

func GetFormatMoneyToFloat(price float64) string {
	p := message.NewPrinter(message.MatchLanguage("en"))
	return p.Sprintf("$%.2f", price)
}

func FillZeros(digits int, number int) string {
	return fmt.Sprintf("%0*d", digits, number)
}

func ConvertIntToUint(number int) uint {
	result := uint(0)
	if number > 0 {
		result = uint(number)
	}
	return result
}

func GetInt64ToUint(number uint) int64 {
	if number > math.MaxInt64 {
		log.Printf("value %v too large for int64", number)
		return 0
	}
	return int64(number)

}
