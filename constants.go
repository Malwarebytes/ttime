package ttime

import (
	"time"
)

var (
	// import a ton of constants so we can act like the time library.
	ParseDuration   = time.ParseDuration
	Date            = time.Date
	ParseInLocation = time.ParseInLocation
	FixedZone       = time.FixedZone
	LoadLocation    = time.LoadLocation
	Sunday          = time.Sunday
	Monday          = time.Monday
	Tuesday         = time.Tuesday
	Wednesday       = time.Wednesday
	Thursday        = time.Thursday
	Friday          = time.Friday
	Saturday        = time.Saturday
	ANSIC           = time.ANSIC
	UnixDate        = time.UnixDate
	RubyDate        = time.RubyDate
	RFC822          = time.RFC822
	RFC822Z         = time.RFC822Z
	RFC850          = time.RFC850
	RFC1123         = time.RFC1123
	RFC1123Z        = time.RFC1123Z
	RFC3339         = time.RFC3339
	RFC3339Nano     = time.RFC3339Nano
	Kitchen         = time.Kitchen
	Stamp           = time.Stamp
	StampMilli      = time.StampMilli
	StampMicro      = time.StampMicro
	StampNano       = time.StampNano
	// constants that I really should redefine:
	NewTimer  = time.NewTimer
	NewTicker = time.NewTicker
	Unix      = time.Unix
)

// I think they "forgot" some constants in ttime
const (
	Nanosecond  Duration = 1
	Microsecond          = 1000 * Nanosecond
	Millisecond          = 1000 * Microsecond
	Second               = 1000 * Millisecond
	Minute               = 60 * Second
	Hour                 = 60 * Minute
)
