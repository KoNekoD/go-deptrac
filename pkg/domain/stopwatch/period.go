package stopwatch

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

// ToSeconds - 10 seconds
func (p *Period) ToSeconds() time.Duration {
	duration := p.endedAt - p.startedAt

	return duration / 1000000000.0
}

// ToMilliseconds - 1000 milliseconds
func (p *Period) ToMilliseconds() time.Duration {
	duration := p.endedAt - p.startedAt

	return duration / 1000000.0
}

// ToFloatSeconds - 10.64 seconds
func (p *Period) ToFloatSeconds() float64 {
	return float64(p.ToMilliseconds()) / float64(1000)
}
