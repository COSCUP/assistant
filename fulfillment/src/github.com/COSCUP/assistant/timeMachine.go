package assistant

import (
	"time"
)

const startDate = "2019-08-17 00:00:00+0800"
const TIME_MACHINE_TIME_FORMAT = "2006-01-02 15:04:05Z0700"

func IsInActivity(now time.Time) bool {
	format := TIME_MACHINE_TIME_FORMAT
	startTime, _ := time.Parse(format, startDate)
	endTime := startTime.Add(2 * 24 * time.Hour)
	return now.After(startTime) && endTime.After(now)
}

func IsDayOne(now time.Time) bool {
	format := TIME_MACHINE_TIME_FORMAT
	startTime, _ := time.Parse(format, startDate)
	endTime := startTime.Add(1 * 24 * time.Hour)
	return now.After(startTime) && endTime.After(now)
}

func IsDayTwo(now time.Time) bool {
	return IsInActivity(now) && !IsDayOne(now)
}

func getUserTime(intent *DialogflowRequest) time.Time {

	// for debug
	// format := TIME_MACHINE_TIME_FORMAT
	// startTime, _ := time.Parse(format, startDate)
	// return startTime.Add(18 * time.Hour)

	return time.Now()
}

func getDay1StartTime() time.Time {
	format := TIME_MACHINE_TIME_FORMAT
	startTime, _ := time.Parse(format, startDate)
	return startTime
}

func getDay2StartTime() time.Time {
	format := TIME_MACHINE_TIME_FORMAT
	startTime, _ := time.Parse(format, startDate)
	endTime := startTime.Add(1 * 24 * time.Hour)
	return endTime
}
