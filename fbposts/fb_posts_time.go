package fbposts

import . "time"

type CheckInterval struct {
	Min float64
	Max float64
}

func IsTimeToCheck(currentTime Time, checkTime Time, interval CheckInterval) bool {
	timeLeft := currentTime.Sub(checkTime).Minutes()
	return timeLeft >= interval.Min && timeLeft <= interval.Max
}

func StartTicking(callback func(Time), intervalMin int) {
	for now := range Tick(Duration(intervalMin) * Minute) {
		callback(now)
	}
}
