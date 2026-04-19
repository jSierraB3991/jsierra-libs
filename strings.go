package eliotlibs

import (
	"crypto/rand"
	"errors"
	"math/big"
	"regexp"
	"strconv"
	"strings"
	"unicode"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func ValidateIsHaveEmojis(cadenaNoEmoji, dataError string) error {

	if ContainsEmoji(cadenaNoEmoji) {
		return NewContainsEmojiError(dataError)
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

type StringToDateParseError struct{}

func (StringToDateParseError) Error() string {
	return STRING_TO_DATE_PARSE_ERROR
}

type ContainsEmojiError struct {
	data string
}

func NewContainsEmojiError(data string) ContainsEmojiError {
	return ContainsEmojiError{
		data: data,
	}
}

func (ce ContainsEmojiError) Error() string {
	return strings.ToUpper(ce.data) + "_CONTAINS_EMOJI_ERROR"
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
	name = strings.TrimSpace(name)         // Elimina espacios en blanco extra
	caser := cases.Title(language.English) // Capitaliza adecuadamente
	return caser.String(name)
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
