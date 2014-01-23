package time

import (
	"time"
)

type Timer struct {
	IsRun    bool
	Interval int
}

func NewTimer(interval int) Timer {
	var timer Timer
	timer.Interval = interval
	timer.IsRun = false
	return timer
}

func (this *Timer) SetInterval(interval int) {
	this.Interval = interval
}

func (this *Timer) Start(inner_function func()) {
	if this.IsRun {
		return
	}
	this.IsRun = true
	ticker := time.NewTicker(time.Millisecond * time.Duration(this.Interval))
	for _ = range ticker.C {
		if !this.IsRun {
			break
		}
		inner_function()
	}
}

func (this *Timer) Stop() {
	this.IsRun = false
}
