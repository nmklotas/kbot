package fbposts

import (
	"math"
	"strings"
	"time"
)

type CheckInterval struct {
	Min int
	Max int
}

func ParseToLocalTime(fbTime string) (time.Time, error) {
	timeInRfc3339 := strings.Replace(fbTime, "+0000", "", 1) + "Z"
	parsedTime, err := time.ParseInLocation(time.RFC3339, timeInRfc3339, time.Now().Location())
	if err != nil {
		return time.Time{}, err
	}

	result := parsedTime.In(time.Now().Location())
	return result, nil
}

func IsPostedToday(postTime time.Time) bool {
	return time.Now().Day() == postTime.Day()
}

func IsTimeToCheck(currentTime time.Time, checkTime time.Time, interval CheckInterval) bool {
	timeLeft := math.Abs(currentTime.Sub(checkTime).Minutes())
	return timeLeft >= float64(interval.Min) && timeLeft <= float64(interval.Max)
}

func StartTicking(callback func(time.Time), intervalMin int) {
	for now := range time.Tick(time.Duration(intervalMin) * time.Minute) {
		callback(now)
	}
}
