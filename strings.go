package eliotlibs

import (
	"crypto/rand"
	"errors"
	"math/big"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

func ValidateIsHaveEmojis(cadenaNoEmoji, dataError string) error {

	if ContainsEmoji(cadenaNoEmoji) {
		return ErrorBadRequest(NewContainsEmojiError(dataError))
	}
	return nil
}
func GetStringNumberFor(number int) string {
	response := strconv.Itoa(number)
	return response
}

func RemoveSpace(cadena string) string {
	return strings.ReplaceAll(cadena, " ", "")
}

func IsNumeric(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}

func GetStringUNumberFor(number uint) string {
	return GetStringNumberFor(int(number))
}

func GetStringUNumber64For(number uint64) string {
	return GetStringNumberFor(int(number))
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

func GenerateCodeByParams(params ...string) string {
	if params == nil {
		return Uppercase
	}
	code, _ := GenerateRandomCode(20, Uppercase)
	var initials strings.Builder
	for _, v := range params {
		if len(v) > 0 {
			initials.WriteRune(unicode.ToUpper(rune(v[0])))
		}
	}
	return strings.ToUpper(initials.String() + code)
}

func ContainsEmoji(s string) bool {
	var emojiRegex = regexp.MustCompile(`[\p{So}\p{Cs}]`)
	return emojiRegex.MatchString(s)
}

// Capitalizar nombre correctamente antes de guardarlo en la DB
func CapitalizeName(name string) string {
	words := strings.Fields(strings.ToLower(name))

	if len(words) == 0 {
		return ""
	}

	r := []rune(words[0])
	r[0] = unicode.ToUpper(r[0])
	words[0] = string(r)

	return strings.Join(words, " ")
}

type ImageNameTemp struct {
	UrlImageTemp string
	ImageName    string
}

func SanitizeLog(s string) string {
	s = strings.ReplaceAll(s, "\n", "")
	s = strings.ReplaceAll(s, "\r", "")
	return s
}
