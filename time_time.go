package eliotlibs

import "time"

func GetDateTimeByString(date string) (time.Time, error) {
	layout := "2006-01-02 15:04:05"
	parsedDate, err := time.Parse(layout, date)
	if err != nil {
		layout = "2006/01/02 15:04:05"
		parsedDate, err := time.Parse(layout, date)
		if err != nil {
			return time.Time{}, StringToDateParseError{}
		}
		return parsedDate, nil
	}
	return parsedDate, nil
}

func GetDateByString(date string) (time.Time, error) {
	layout := "2006-01-02"
	parsedDate, err := time.Parse(layout, date)
	if err != nil {
		layout := "2006/01/02"
		parsedDate, err := time.Parse(layout, date)
		if err != nil {
			return time.Time{}, StringToDateParseError{}
		}
		return parsedDate, nil
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
	return GetStringNumberFor(date.Year()) + "/" + getDataTime(int(date.Month())) + "/" + getDataTime(date.Day())
}

func GetStringToTimeLongFormat(date time.Time) string {
	return GetStringNumberFor(date.Year()) + "/" + getDataTime(int(date.Month())) + "/" + getDataTime(date.Day()) + " " +
		getDataTime(date.Hour()) + ":" + getDataTime(date.Minute()) + ":" + getDataTime(date.Second())
}

func GetDateShortByNilString(date *string) (*time.Time, error) {
	if date == nil {
		return nil, nil
	}
	dateString, err := GetDateByString(*date)
	return &dateString, err
}

func GetDateLongByNilString(date *string) (*time.Time, error) {
	if date == nil {
		return nil, nil
	}
	dateString, err := GetDateTimeByString(*date)
	return &dateString, err
}

func getDataTime(datTime int) string {
	if datTime <= 9 {
		return "0" + GetStringNumberFor(datTime)
	}
	return GetStringNumberFor(datTime)
}

func NextBusinessDay(start time.Time, days int) time.Time {
	current := start
	addedDays := 0

	for addedDays < days {
		current = current.AddDate(0, 0, 1) // Añade 1 día
		if current.Weekday() != time.Saturday && current.Weekday() != time.Sunday {
			addedDays++
		}
	}

	return current
}
