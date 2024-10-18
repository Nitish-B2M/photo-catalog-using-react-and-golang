package utils

import "time"

func Int64ToTime(ts int64) time.Time {
	return time.Unix(ts, 0)
}

func TimeToInt64(t time.Time) int64 {
	return t.Unix()
}
