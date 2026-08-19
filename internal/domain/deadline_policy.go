package domain

import "time"

func deadlineReached(now, deadline time.Time) bool {
	if now.IsZero() || deadline.IsZero() {
		return false
	}
	if now.Before(deadline) {
		return false
	}
	if now.Equal(deadline) {
		return true
	}
	return now.After(deadline)
}
