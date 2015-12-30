package ttime

import (
	// "github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestFreezingTime(t *testing.T) {
	// freeze time at a specific date/time (eg, test leap-year support!):
	now, err := time.Parse(time.RFC3339, "2012-02-29T00:00:00Z")
	if err != nil {
		panic("date time parse failed")
	}
	Freeze(Time(now))

	if !IsFrozen() {
		t.Error("Time should be frozen here, and was not.")
	}
	if Now().UTC() != Time(now) {
		t.Error("Time should still be set to frozen time")
	}
	t.Logf("It is now %v (frozen)", Now().UTC())
	Unfreeze()
	if Now().UTC() == Time(now) || IsFrozen() {
		t.Error("Time should no longer be frozen")
	}
	t.Logf("It is now %v (not frozen)", Now().UTC())
}

func TestAfterFreezeMode(t *testing.T) {
	// test frozen functionality.
	start := Now()
	Freeze(Now())
	<-After(10 * Millisecond)
	Unfreeze()
	elapsed := Now().Sub(start)
	t.Logf("Took %v", elapsed)
	if elapsed >= 1*Millisecond {
		t.Error("Took too long")
	}
}

func TestAfterNoFreeze(t *testing.T) {
	// test original functionality is restored
	start := Now()
	<-After(1 * Millisecond)
	elapsed := Now().Sub(start)
	t.Logf("Took %v", elapsed)
	if elapsed > 2*Millisecond {
		t.Error("Took too long")
	}
	if elapsed < 1*Millisecond {
		t.Error("Went too fast")
	}
}

func TestTickFreezeMode(t *testing.T) {
	// TODO: this leaks tickers
	// test frozen functionality.
	start := Now()
	Freeze(start)
	c := Tick(10 * Millisecond)
	<-c
	<-c
	<-c
	if !start.Add(40 * Millisecond).Equal(Now()) {
		t.Errorf("Expected Tick to advance clock, did not. %v != %v", start.Add(40*Millisecond), Now())
	}
	Unfreeze()
	elapsed := Now().Sub(start)
	t.Logf("Took %v", elapsed)
	if elapsed >= 1*Millisecond {
		t.Error("Took too long")
	}
}

func TestTickNoFreeze(t *testing.T) {
	// TODO: this leaks tickers
	// test original functionality is restored
	start := Now()
	c := Tick(1 * Millisecond)
	<-c
	<-c
	<-c
	elapsed := Now().Sub(start)
	t.Logf("Took %v", elapsed)
	if elapsed > 4*Millisecond {
		t.Error("Took too long")
	}
	if elapsed < 3*Millisecond {
		t.Error("Went too fast")
	}
}

func TestSleepNoFreeze(t *testing.T) {
	start := Now()
	Sleep(1 * Millisecond)
	elapsed := Now().Sub(start)
	if elapsed > 2*Millisecond {
		t.Error("Took too long")
	}
	if elapsed < 1*Millisecond {
		t.Error("Went too fast")
	}
}

func TestSleepFreezeMode(t *testing.T) {
	start := Now()
	Freeze(start)
	Sleep(1 * Second)
	Unfreeze()
	elapsed := Now().Sub(start)
	if elapsed > 1*Millisecond {
		t.Error("Took too long")
	}
}

/*
func TestSleepFreezeModeSleepMath(t *testing.T) {
	start := Now()
	Freeze(start)

	sleeper1 := make(chan Time)
	sleeper2 := make(chan Time)
	go func() {
		Sleep(200)
		Sleep(200)
		Sleep(200)
		sleeper1<- Now()
	}()

	go func() {
		Sleep(300)
		Sleep(300)
		sleeper2<- Now()
	}()

	sl1, sl2 := <-sleeper1,<-sleeper2
	if sl1!=sl2 {
		t.Errorf("Sleep is broken: %v, %v", sl1, sl2)
	}
}
*/
