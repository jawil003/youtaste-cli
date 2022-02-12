package services

import "time"

type TimerService struct {
	startTime         time.Time
	millisecondsToRun int64
}

func (t *TimerService) Start(millisecondsToRun int64) {
	t.startTime = time.Now()
	t.millisecondsToRun = millisecondsToRun
}

func (t *TimerService) IsRunning() bool {
	return time.Since(t.startTime).Milliseconds() <= t.millisecondsToRun
}

func (t *TimerService) GetRemainingTime() int64 {
	return time.Since(t.startTime).Milliseconds()
}
