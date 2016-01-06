package ttime

// note: this is a fork from https://github.com/ssoroka/ttime
// none of this was my idea and all the credit belongs to Steven

import (
	"sync"
	"time"
)

var lock sync.RWMutex

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

func (t Time) After(u Time) bool {
	return time.Time(t).After(time.Time(u))
}

func (t Time) Before(u Time) bool {
	return time.Time(t).Before(time.Time(u))
}

func (t Time) Format(layout string) string {
	return time.Time(t).Format(layout)
}

// existing ttime wrappers but in a none-leaky fashion
func Freeze(t Time) {
	lock.Lock()
	defer lock.Unlock()
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
		lock.RLock()
		defer lock.RUnlock()
		return currentTime
	} else {
		return Time(time.Now())
	}
}

func After(d Duration) <-chan Time {
	c := make(chan Time, 1)
	if timeFrozen {
		lock.Lock()
		currentTime = currentTime.Add(d)
		lock.Unlock()
		c <- currentTime
	} else {
		c <- Time(<-time.After(time.Duration(d)))
	}
	return c
}

func Tick(d Duration) <-chan Time {
	c := make(chan Time, 1)
	go func() {
		for {
			if timeFrozen {
				lock.Lock()
				currentTime = currentTime.Add(d)
				lock.Unlock()
				c <- currentTime
			} else {
				c <- Time(<-time.Tick(time.Duration(d)))
			}
		}
	}()
	return c
}

func Sleep(d Duration) {
	// thinking more about this...
	// Add() does not appear to be the right strategy in the concurrent case.
	// I have goroutines that all Sleep (and therefore Add to the clock). As a
	// consequence I get distorted times and can not use timestamps in testing.
	// For me this is bad since this is exactly what I want to do. Therefore
	// I am going to work on a ttime version that behaves more realistic but
	// with an "infinitely" fast clock.
	// Now, enough with the wining! Fix it!
	if timeFrozen {
		if d > 0 {
			// for now I might get away with a mutex
			lock.Lock()
			currentTime = currentTime.Add(d)
			lock.Unlock()
		}
	} else {
		time.Sleep(time.Duration(d))
	}
}

func Parse(layout, value string) (Time, error) {
	t, err := time.Parse(layout, value)
	return Time(t), err
}
