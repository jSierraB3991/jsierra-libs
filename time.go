package eliotlibs

import (
	"time"
)

type InvalidDurationOfTask struct{}

func (InvalidDurationOfTask) Error() string {
	return "INVALID_DURATION_OF_TASK"
}

func ScheduleTask(targetTime time.Time, task func()) error {
	duration := time.Until(targetTime)

	if duration <= 0 {
		return InvalidDurationOfTask{}
	}

	time.AfterFunc(duration, task)
	return nil
}
