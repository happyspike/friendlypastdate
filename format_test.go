package friendlypastdate

import (
	"testing"
	"time"
)

func TestWithinLastMinuteGivesJustNow(t *testing.T) {
	comparison := time.Date(2000, 1, 1, 0, 0, 0, 0, time.Local)
	now := comparison.Add(time.Second * 59)

	if Format(now, comparison) != "just now" {
		t.Fail()
	}
}

func Test1MinuteGivesAMinuteAgo(t *testing.T) {
	comparison := time.Date(2000, 1, 1, 0, 0, 0, 0, time.Local)
	now := comparison.Add(time.Minute)

	if Format(now, comparison) != "a minute ago" {
		t.Fail()
	}
}

func TestJustMoreThan1MinuteGivesAMinuteAgo(t *testing.T) {
	comparison := time.Date(2000, 1, 1, 0, 0, 0, 0, time.Local)
	now := comparison.Add(time.Minute + time.Second)

	if Format(now, comparison) != "a minute ago" {
		t.Fail()
	}
}

func TestWithinLastHourGivesXMinutesAgo(t *testing.T) {
	comparison := time.Date(2000, 1, 1, 0, 0, 0, 0, time.Local)
	now := comparison.Add(time.Minute * 59)

	if Format(now, comparison) != "59 minutes ago" {
		t.Fail()
	}
}

func Test1HourGivesAnHourAgo(t *testing.T) {
	comparison := time.Date(2000, 1, 1, 0, 0, 0, 0, time.Local)
	now := comparison.Add(time.Hour)

	if Format(now, comparison) != "an hour ago" {
		t.Fail()
	}
}

func TestJustMoreThan1HourGivesAnHourAgo(t *testing.T) {
	comparison := time.Date(2000, 1, 1, 0, 0, 0, 0, time.Local)
	now := comparison.Add(time.Hour + time.Second)

	if Format(now, comparison) != "an hour ago" {
		t.Fail()
	}
}

func TestWithinLastDayGivesXHoursAgo(t *testing.T) {
	comparison := time.Date(2000, 1, 1, 0, 0, 0, 0, time.Local)
	now := comparison.Add(time.Hour * 23)

	if Format(now, comparison) != "23 hours ago" {
		t.Fail()
	}
}

func TestMoreThanADayGivesTheDateAndTime(t *testing.T) {
	comparison := time.Date(2014, 1, 1, 0, 0, 0, 0, time.Local)
	now := comparison.Add(time.Hour*26 + time.Second)

	if Format(now, comparison) != "Wed, 1 Jan 2014" {
		t.Log(Format(now, comparison))
		t.Fail()
	}
}
