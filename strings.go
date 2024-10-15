package eliotlibs

import (
	"crypto/rand"
	"errors"
	"math/big"
	"strconv"
	"strings"
	"time"
)

func GetStringNumberFor(number int) string {
	response := strconv.Itoa(number)
	return response
}

func GetStringUNumberFor(number uint) string {
	return GetStringNumberFor(int(number))
}

func GetStringUNumber64For(number uint64) string {
	return GetStringNumberFor(int(number))
}

type StringToDateParseError struct{}

func (StringToDateParseError) Error() string {
	return "STRING_TO_DATE_PARSE_ERROR"
}

func GetDateTimeByString(date string) (time.Time, error) {
	layout := "2006-01-02 15:04:05"
	parsedDate, err := time.Parse(layout, date)
	if err != nil {
		return time.Time{}, StringToDateParseError{}
	}
	return parsedDate, nil
}

func GetDateByString(date string) (time.Time, error) {
	layout := "2006-01-02"
	parsedDate, err := time.Parse(layout, date)
	if err != nil {
		return time.Time{}, StringToDateParseError{}
	}
	return parsedDate, nil
}

func GetStringByDateNilShort(date *time.Time) *string {
	if date == nil {
		return nil
	}
	dateString := GetStringToTimeShortFormat(*date)
	return &dateString
}

func GetStringByDateNilLong(date *time.Time) *string {
	if date == nil {
		return nil
	}
	dateString := GetStringToTimeLongFormat(*date)
	return &dateString
}

func GetStringToTimeShortFormat(date time.Time) string {
	return GetStringNumberFor(date.Year()) + "/" + GetStringNumberFor(int(date.Month())) + "/" + GetStringNumberFor(date.Day())
}

func GetStringToTimeLongFormat(date time.Time) string {
	return GetStringNumberFor(date.Year()) + "/" + GetStringNumberFor(int(date.Month())) + "/" + GetStringNumberFor(date.Day()) + " " +
		GetStringNumberFor(date.Hour()) + ":" + GetStringNumberFor(date.Minute()) + ":" + GetStringNumberFor(date.Second())
}

func GenerateRandomCode(length int, charSet string) (string, error) {
	if length <= 0 {
		return "", errors.New("length must be greater than 0")
	}

	if len(charSet) == 0 {
		return "", errors.New("character set cannot be empty")
	}

	var result strings.Builder
	charSetLength := big.NewInt(int64(len(charSet)))

	for i := 0; i < length; i++ {
		randomIndex, err := rand.Int(rand.Reader, charSetLength)
		if err != nil {
			return "", err
		}
		result.WriteByte(charSet[randomIndex.Int64()])
	}

	return result.String(), nil
}

func GetDateShortByNilString(date *string) (*time.Time, error) {
	if date == nil {
		return nil, nil
	}
	dateString, err := GetDateTimeByString(*date)
	return &dateString, err
}

func GetDateLongByNilString(date *string) (*time.Time, error) {
	if date == nil {
		return nil, nil
	}
	dateString, err := GetDateByString(*date)
	return &dateString, err
}
