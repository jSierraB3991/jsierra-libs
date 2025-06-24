package eliotlibs

import (
	"log"
	"testing"
	"time"
)

func TestValidateScheduleTask(t *testing.T) {

	timeNow := time.Now()
	testCases := []struct {
		Name       string
		Value      time.Time
		ExpectEror error
	}{
		{
			Name:       "VALID SCHEDULE TASK",
			Value:      timeNow.Add(20 * time.Second),
			ExpectEror: nil,
		},
		{
			Name:       "INVALID SCHEDULE TASK",
			Value:      timeNow.Add(-20 * time.Second),
			ExpectEror: InvalidDurationOfTask{},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.Name, func(t *testing.T) {
			err := ScheduleTask(tc.Value, func() {
				log.Printf("Log %v", tc.Name)
			})
			if err != tc.ExpectEror {
				t.Errorf("Error %v, got %v", tc.ExpectEror, err)
			}
		})
	}
	time.Sleep(21 * time.Second)
}

func TestNextBusinessDay(t *testing.T) {

	testCases := []struct {
		Name           string
		Start          time.Time
		ExpectedOutput time.Weekday
	}{
		{
			Start:          time.Date(2025, 6, 23, 0, 0, 0, 0, time.UTC),
			ExpectedOutput: time.Wednesday,
		},
		{
			Start:          time.Date(2025, 6, 24, 0, 0, 0, 0, time.UTC),
			ExpectedOutput: time.Thursday,
		},
		{
			Start:          time.Date(2025, 6, 25, 0, 0, 0, 0, time.UTC),
			ExpectedOutput: time.Friday,
		},
		{
			Start:          time.Date(2025, 6, 26, 0, 0, 0, 0, time.UTC),
			ExpectedOutput: time.Monday,
		},
		{
			Start:          time.Date(2025, 6, 27, 0, 0, 0, 0, time.UTC),
			ExpectedOutput: time.Tuesday,
		},
		{
			Start:          time.Date(2025, 6, 28, 0, 0, 0, 0, time.UTC),
			ExpectedOutput: time.Tuesday,
		},
		{
			Start:          time.Date(2025, 6, 29, 0, 0, 0, 0, time.UTC),
			ExpectedOutput: time.Tuesday,
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.Name, func(t *testing.T) {
			log.Printf("Start in day week %s", tc.Start.Weekday())
			nextDay := NextBusinessDay(tc.Start, 2)
			log.Printf("Finish in day week %s", nextDay.Weekday())
			if nextDay.Weekday() != tc.ExpectedOutput {
				t.Errorf("Error %v, got %v", tc.ExpectedOutput, nextDay.Weekday())
			}
		})
	}
	time.Sleep(21 * time.Second)
}
