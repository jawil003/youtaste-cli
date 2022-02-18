package services

import "time"

type TimerService struct {
	endTime *time.Time
}

func (t *TimerService) Start(dateTilRun time.Time) {
	t.endTime = &dateTilRun
}

func (t *TimerService) IsRunning() bool {
	return t.endTime.Sub(*t.endTime).Milliseconds() > 0
}

func (t TimerService) IsActive() bool {
	return t.endTime != nil
}

func (t *TimerService) GetRemainingTime() int64 {
	if t == nil || t.endTime == nil {
		return 0
	}

	remainingTime := time.Since(*t.endTime).Milliseconds()

	if remainingTime > 0 {
		return remainingTime
	}

	return 0
}

func (t *TimerService) Clear() {
	t.endTime = nil
}
