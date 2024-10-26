package scheduler

import "time"

type Trigger interface {
	Trigger(input Input) error
}

type Input struct {
	Timeout time.Duration
}
