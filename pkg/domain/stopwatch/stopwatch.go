package stopwatch

import (
	"github.com/KoNekoD/go-deptrac/pkg/domain/apperrors"
)

type Stopwatch struct {
	periods map[string]*StartedPeriod
}

func NewStopwatch() *Stopwatch {
	return &Stopwatch{
		periods: make(map[string]*StartedPeriod),
	}
}

func (s *Stopwatch) Start(event string) *apperrors.StopwatchException {
	err := s.assertPeriodNotStarted(event)
	if err != nil {
		return err
	}

	s.periods[event] = NewStartedPeriodStart()

	return nil
}

func (s *Stopwatch) Stop(event string) (*Period, error) {
	err := s.AssertPeriodStarted(event)
	if err != nil {
		return nil, err
	}

	period := s.periods[event].Stop()

	delete(s.periods, event)

	return period, nil
}

func (s *Stopwatch) assertPeriodNotStarted(event string) *apperrors.StopwatchException {
	if _, ok := s.periods[event]; ok {
		return apperrors.NewStopwatchExceptionPeriodAlreadyStarted(event)
	}

	return nil
}

func (s *Stopwatch) AssertPeriodStarted(event string) *apperrors.StopwatchException {
	if _, ok := s.periods[event]; !ok {
		return apperrors.NewStopwatchExceptionPeriodNotStarted(event)
	}

	return nil
}
