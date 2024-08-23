package time_stopwatch

import (
	"github.com/loov/hrtime"
	"time"
)

type Period struct {
	startedAt time.Duration
	endedAt   time.Duration
}

func NewPeriod(startedAt time.Duration, endedAt time.Duration) *Period {
	return &Period{
		startedAt: startedAt,
		endedAt:   endedAt,
	}
}

func NewPeriodStop(startedPeriod *StartedPeriod) *Period {
	return NewPeriod(
		startedPeriod.StartedAt,
		hrtime.Now(),
	)
}

func (p *Period) ToSeconds() time.Duration {
	duration := p.endedAt - p.startedAt

	return duration / 1000000000.0
}
