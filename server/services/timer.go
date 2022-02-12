package services

import "time"

type TimerService struct {
	startTime         *time.Time
	millisecondsToRun int64
}

func (t *TimerService) Start(millisecondsToRun int64) {
	newTimer := time.Now()
	t.startTime = &newTimer
	t.millisecondsToRun = millisecondsToRun
}

func (t *TimerService) IsRunning() bool {
	return time.Since(*t.startTime).Milliseconds() <= t.millisecondsToRun
}

func (t TimerService) IsActive() bool {
	return t.startTime != nil && t.millisecondsToRun > 0
}

func (t *TimerService) GetRemainingTime() int64 {
	if t == nil || t.startTime == nil {
		return 0
	}

	return t.startTime.Add(time.Millisecond * time.Duration(t.millisecondsToRun)).Sub(time.Now()).Milliseconds()
}

func (t *TimerService) Clear() {
	t.startTime = nil
	t.millisecondsToRun = 0
}
