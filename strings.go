package eliotlibs

import (
	"crypto/rand"
	"errors"
	"math/big"
	"regexp"
	"strconv"
	"strings"
)

func GetStringNumberFor(number int) string {
	response := strconv.Itoa(number)
	return response
}

func RemoveSpace(cadena string) string {
	return strings.ReplaceAll(cadena, " ", "")
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

func ContainsEmoji(s string) bool {
	var emojiRegex = regexp.MustCompile(`[\p{So}\p{Cs}]`)
	return emojiRegex.MatchString(s)
}

type ImageNameTemp struct {
	UrlImageTemp string
	ImageName    string
}
