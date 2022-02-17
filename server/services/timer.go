package services

import "time"

type TimerService struct {
	startTime         *time.Time
	millisecondsToRun *time.Time
}

func (t *TimerService) Start(dateTilRun time.Time) {
	newTimer := time.Now()
	t.startTime = &newTimer
	t.millisecondsToRun = &dateTilRun
}

func (t *TimerService) IsRunning() bool {
	return t.millisecondsToRun.Sub(*t.startTime).Milliseconds() > 0
}

func (t TimerService) IsActive() bool {
	return t.startTime != nil && t.millisecondsToRun != nil
}

func (t *TimerService) GetRemainingTime() int64 {
	if t == nil || t.startTime == nil {
		return 0
	}

	//TODO: Fix logic to calc remaining time

	remainingTime := t.millisecondsToRun.Sub(*t.startTime).Milliseconds()

	if remainingTime > 0 {
		return remainingTime
	}

	return 0
}

func (t *TimerService) Clear() {
	t.startTime = nil
	t.millisecondsToRun = nil
}
