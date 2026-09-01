package flags

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/outscale/osc-sdk-go/v3/pkg/iso8601"
)

const Time = "osctime"

var Now = iso8601.Now

// TimeValue adapts time.Time for use as a flag.
type TimeValue struct {
	t *iso8601.Time
}

func NewTimeValue() *TimeValue {
	return &TimeValue{t: &iso8601.Time{}}
}

// Set time.Time value from string (rfc3339, iso8601 or duration).
func (d *TimeValue) Set(s string) error {
	s = strings.TrimSpace(s)
	// shortcuts
	switch s {
	case "bom", "beginning-of-month":
		n := Now()
		// On the first day of the month, we look at the previous month
		// The typical use case is bom to today, and it breaks on day 1 on many APIs
		if n.Day() == 1 {
			n = n.AddDate(0, -1, 0)
		}
		*d.t = iso8601.Day(n.Year(), n.Month(), 1)
		return nil
	case "t", "today":
		n := Now()
		*d.t = iso8601.Day(n.Year(), n.Month(), n.Day())
		return nil
	case "y", "yesterday":
		n := Now()
		*d.t = iso8601.Day(n.Year(), n.Month(), n.Day()-1)
		return nil
	}
	// try a iso8601 time
	v, err := iso8601.ParseString(s)
	if err == nil {
		*d.t = v
		return nil
	}
	s = strings.TrimPrefix(s, "+")
	// it might have been a duration
	dur, err := time.ParseDuration(s)
	if err == nil {
		*d.t = Now().Add(dur)
		return nil
	}
	// or a day/month/year offset
	switch {
	case strings.HasSuffix(s, "d"):
		n, err := strconv.Atoi(s[:len(s)-1])
		if err != nil {
			return fmt.Errorf("invalid day offset %q: %w", s, err)
		}
		*d.t = Now().AddDate(0, 0, n)
		return nil
	case strings.HasSuffix(s, "mo"):
		n, err := strconv.Atoi(s[:len(s)-2])
		if err != nil {
			return fmt.Errorf("invalid month offset %q: %w", s, err)
		}
		*d.t = Now().AddDate(0, n, 0)
		return nil
	case strings.HasSuffix(s, "y"):
		n, err := strconv.Atoi(s[:len(s)-1])
		if err != nil {
			return fmt.Errorf("invalid year offset %q: %w", s, err)
		}
		*d.t = Now().AddDate(n, 0, 0)
		return nil
	}
	return fmt.Errorf("invalid time format %q", s)
}

// Type name for time.Time flags.
func (d *TimeValue) Type() string {
	return "osctime"
}

func (d *TimeValue) String() string {
	if d.t == nil || d.t.IsZero() {
		return ""
	}
	return d.t.String()
}

func (d *TimeValue) Value() (iso8601.Time, bool) {
	if d.t == nil {
		return iso8601.Time{}, false
	}
	return *d.t, true
}
