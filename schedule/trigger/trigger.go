package trigger

import (
	"github.com/hduCS2021/ScheduleAssistant/schedule/date"
	"time"
)

type Action int

//the action when Triggered.
const (
	Reject Action = iota
	PassThrough
	Accept
)

type Trigger struct {
	dates  date.List
	prior  int
	action Action
}

func New() *Trigger {
	return &Trigger{
		prior: 0,
	}
}

func (tr *Trigger) SetPrior(prior int) {
	tr.prior = prior
}

// SetAction set the action when triggered.
// in Accept mode, it returns Accept / PassThrough.
// in PassThrough mode, it returns PassThrough / Reject.
// in Reject mode, it returns Reject / PassThrough.
func (tr *Trigger) SetAction(action Action) {
	if action < Reject || action > Accept {
		return
	}
	tr.action = action
}

// IsTriggered returns an action.
func (tr *Trigger) IsTriggered(t time.Time) Action {
	ok := tr.dates.Check(t)
	if ok {
		return tr.action
	}
	//!ok
	if tr.action == PassThrough {
		return Reject
	}
	return PassThrough
}

func (tr *Trigger) Add(ds ...date.Date) {
	for _, v := range ds {
		tr.dates.Append(v)
	}
}
