package eliotlibs

import "time"

func BoolP(data bool) *bool {
	return &data
}

func UIntP(data uint) *uint {
	return &data
}

func UInt64P(data uint64) *uint64 {
	return &data
}

func TimeP(data time.Time) *time.Time {
	return &data
}

func StringP(data string) *string {
	return &data
}

func IntP(data int) *int {
	return &data
}

func Int64P(data int64) *int64 {
	return &data
}

func floatP(data float64) *float64 {
	return &data
}

func float32P(data float32) *float32 {
	return &data
}
