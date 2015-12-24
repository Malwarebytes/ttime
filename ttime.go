package ttime

// note: this is a fork from https://github.com/ssoroka/ttime
// none of this was my idea and all the credit belongs to Steven

import "time"


var (
  currentTime Time
  timeFrozen  bool
)

type Duration time.Duration
// type Location time.Location
// type Month time.Month
// type ParseError time.ParseError
// type Ticker time.Ticker
type Time time.Time
// type Timer time.Timer
// type Weekday time.Weekday

// in order to make things work I think the most important
// measure is not to "leak" any types from golang time
// so this is just a start which still needs a lot of work

// Wrap the time.Time functions
func (t Time) Add(d Duration) Time {
  return Time(time.Time(t).Add(time.Duration(d)))
}

func (t Time) Sub(u Time) Duration {
  return Duration(time.Time(t).Sub(time.Time(u)))
}

func (t Time) UTC() Time {
  return Time(time.Time(t).UTC())
}

func (t Time) Equal(u Time) bool {
  return time.Time(t).Equal(time.Time(u))
}


// existing ttime wrappers but in a none-leaky fashion
func Freeze(t Time) {
  currentTime = t
  timeFrozen = true
}

func Unfreeze() {
  timeFrozen = false
}

func IsFrozen() bool {
  return timeFrozen
}

func Now() Time {
  if timeFrozen {
    return currentTime
  } else {
    return Time(time.Now())
  }
}

func After(d Duration) <-chan Time {
  c := make(chan Time, 1)
  if timeFrozen {
    currentTime = currentTime.Add(d)
    c <- currentTime
  } else {
    c <- Time(<- time.After(time.Duration(d)))
  }
  return c
}

func Tick(d Duration) <-chan Time {
  c := make(chan Time, 1)
  go func() {
    for {
      if timeFrozen {
        currentTime = currentTime.Add(d)
        c <- currentTime
      } else {
        c <- Time(<- time.Tick(time.Duration(d)))
      }
    }
  }()
  return c
}

func Sleep(d Duration) {
  if timeFrozen && d > 0 {
    currentTime = currentTime.Add(d)
  } else {
    time.Sleep(time.Duration(d))
  }
}
