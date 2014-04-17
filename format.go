package friendlypastdate

import (
	"fmt"
	"time"
)

func Format(now time.Time, comparison time.Time) string {
	if now.Sub(comparison) < time.Minute {
		return "just now"
	} else if now.Sub(comparison) < (time.Minute * 2) {
		return "a minute ago"
	} else if now.Sub(comparison) < time.Hour {
		return fmt.Sprintf("%.0f minutes ago", now.Sub(comparison).Minutes())
	} else if now.Sub(comparison) < (time.Hour * 2) {
		return "an hour ago"
	} else if now.Sub(comparison) < (time.Hour * 24) {
		return fmt.Sprintf("%.0f hours ago", now.Sub(comparison).Hours())
	} else {
		layout := "Mon, 2 Jan 2006"
		return comparison.Format(layout)
	}
}
