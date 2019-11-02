package fbposts

import (
	. "time"
)

type CheckInterval struct {
	Min int
	Max int
}

func IsPostedToday(postTime Time) bool {
	return Now().Day() == postTime.Day()
}

func IsTimeToCheck(currentTime Time, checkTime Time, interval CheckInterval) bool {
	timeLeft := currentTime.Sub(checkTime).Minutes()
	return timeLeft >= float64(interval.Min) && timeLeft <= float64(interval.Max)
}

func StartTicking(callback func(Time), intervalMin int) {
	for now := range Tick(Duration(intervalMin) * Minute) {
		callback(now)
	}
}
