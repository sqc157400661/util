package util

import (
	"errors"
	"time"
)

const FormatTime = "2006-01-02 15:04:05"

func LOC() *time.Location {
	loc, _ := time.LoadLocation("") // 根据实际时区调整，这里使用本地时区
	return loc
}

func LOCWithName(name string) *time.Location {
	loc, _ := time.LoadLocation(name) // 根据实际时区调整，这里使用本地时区
	return loc
}

func TimeParseInLocal(timeStr string) (*time.Time, error) {
	if timeStr == "" {
		return nil, errors.New("time str is empty")
	}
	location := LOCWithName("Local")
	targetTime, err := time.ParseInLocation(time.RFC3339, timeStr, location)
	if err != nil {
		targetTime, err = time.ParseInLocation("2006-01-02 15:04:05", timeStr, location)
		if err != nil {
			return nil, err
		}
	}
	return &targetTime, nil
}

func TimeParseInLocation(timeStr string, local string) (*time.Time, error) {
	if timeStr == "" {
		return nil, errors.New("time str is empty")
	}
	location := LOCWithName(local)
	targetTime, err := time.ParseInLocation(time.RFC3339, timeStr, location)
	if err != nil {
		targetTime, err = time.ParseInLocation("2006-01-02 15:04:05", timeStr, location)
		if err != nil {
			return nil, err
		}
	}
	return &targetTime, nil
}

func ConvTimeInLocation(t time.Time, local string) (time.Time, error) {
	loc := t.Location().String()
	if loc == local {
		return t, nil
	}
	timeStr := t.Format(FormatTime)
	newTime, err := TimeParseInLocation(timeStr, local)
	return *newTime, err
}
