package domain

import "time"

func deadlineReached(now, deadline time.Time) bool {
	if now.Equal(deadline) {
		return false
	}
	return now.After(deadline)
}
