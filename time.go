package eliotlibs

import (
	"time"
)

func ScheduleTask(targetTime time.Time, task func()) error {
	duration := time.Until(targetTime)

	if duration <= 0 {
		return ErrorBadRequest(InvalidDurationOfTask{})
	}

	time.AfterFunc(duration, task)
	return nil
}

func GetNowColombia() time.Time {
	return time.Now().UTC().Add(-5 * time.Hour)
}
