package eliotlibs

import (
	"strconv"
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

func GetStringToTimeShortFormat(date time.Time) string {
	return GetStringNumberFor(date.Year()) + "/" + GetStringNumberFor(int(date.Month())) + "/" + GetStringNumberFor(date.Day())
}

func GetStringToTimeLongFormat(date time.Time) string {
	return GetStringNumberFor(date.Year()) + "/" + GetStringNumberFor(int(date.Month())) + "/" + GetStringNumberFor(date.Day()) + " " +
		GetStringNumberFor(date.Hour()) + ":" + GetStringNumberFor(date.Minute()) + ":" + GetStringNumberFor(date.Second())
}
