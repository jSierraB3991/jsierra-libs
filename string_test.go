package eliotlibs

import (
	"testing"
	"time"
)

func TestConvertStrinToDateTime(t *testing.T) {
	testCases := []struct {
		Name       string
		Value      string
		ExpectEror error
	}{
		{
			Name:       "Ok date and time",
			Value:      "2024-09-11 13:38:18",
			ExpectEror: nil,
		},
		{
			Name:       "invalid alone date",
			Value:      "2024-09-11",
			ExpectEror: StringToDateParseError{},
		},
		{
			Name:       "invalid configuration",
			Value:      "11-09-2024 13:38:18",
			ExpectEror: StringToDateParseError{},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.Name, func(t *testing.T) {
			_, err := GetDateTimeByString(tc.Value)
			if err != tc.ExpectEror {
				t.Errorf("Error %v, got %v", tc.ExpectEror, err)
			}
		})
	}
}

func TestConvertStrinToDate(t *testing.T) {
	testCases := []struct {
		Name       string
		Value      string
		ExpectEror error
	}{
		{
			Name:       "Ok date and time",
			Value:      "2024-09-11 13:38:18",
			ExpectEror: StringToDateParseError{},
		},
		{
			Name:       "invalid alone date",
			Value:      "2024-09-11",
			ExpectEror: nil,
		},
		{
			Name:       "invalid configuration",
			Value:      "11-09-2024 13:38:18",
			ExpectEror: StringToDateParseError{},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.Name, func(t *testing.T) {
			_, err := GetDateByString(tc.Value)
			if err != tc.ExpectEror {
				t.Errorf("Error %v, got %v", tc.ExpectEror, err)
			}
		})
	}
}

func TestDateToString(t *testing.T) {
	testCases := []struct {
		Name           string
		Value          time.Time
		ExpectedOutput string
		IsShorDate     bool
	}{
		{
			Name:           "Ok time to short date",
			Value:          time.Date(2024, 12, 12, 17, 12, 23, 100, time.UTC),
			ExpectedOutput: "2024/12/12",
			IsShorDate:     true,
		},
		{
			Name:           "Ok time to long date",
			Value:          time.Date(2024, 12, 12, 17, 12, 23, 100, time.UTC),
			ExpectedOutput: "2024/12/12 17:12:23",
			IsShorDate:     false,
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.Name, func(t *testing.T) {
			var funcConvert func(date time.Time) string
			if tc.IsShorDate {
				funcConvert = GetStringToTimeShortFormat
			} else {
				funcConvert = GetStringToTimeLongFormat
			}

			valueResult := funcConvert(tc.Value)

			if valueResult != tc.ExpectedOutput {
				t.Errorf("Error expect data %s, got %s", tc.ExpectedOutput, valueResult)
			}
		})
	}

}
