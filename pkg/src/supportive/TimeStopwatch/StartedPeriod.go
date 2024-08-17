package TimeStopwatch

import (
	"github.com/loov/hrtime"
	"time"
)

//final class StartedPeriod
//{
//    private function __construct(public readonly float|int $startedAt)
//    {
//    }
//    public static function start() : self
//    {
//        return new self(\hrtime(\true));
//    }
//    public function stop() : \Qossmic\Deptrac\Supportive\Time\Period
//    {
//        return \Qossmic\Deptrac\Supportive\Time\Period::stop($this);
//    }
//}

type StartedPeriod struct {
	StartedAt time.Duration
}

func NewStartedPeriodStart() *StartedPeriod {
	return &StartedPeriod{
		StartedAt: hrtime.Now(),
	}
}

func (p *StartedPeriod) Stop() *Period {
	return NewPeriodStop(p)
}
