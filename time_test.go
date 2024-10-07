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
