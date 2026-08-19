package domain

import (
	"testing"
	"time"
)

func TestDeadlineMomentIsOverdue(t *testing.T) {
	deadline := time.Date(2026, 8, 19, 12, 0, 0, 0, time.UTC)
	item := &RightsCase{Status: StatusAdjudicated, Deadline: deadline}
	if !item.IsOverdue(deadline) {
		t.Fatal("case at deadline was not treated as overdue")
	}
}
