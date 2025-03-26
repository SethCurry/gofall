package gofall

import (
	"testing"
	"time"
)

func Test_RateLimiter(t *testing.T) {
	t.Parallel()

	limiter := newRateLimiter(time.Second, 1)

	for i := 0; i < 5; i++ {
		canRun := limiter.AddEvent()
		if i == 0 {
			if !canRun {
				t.Error("did not expect to be throttled here")
			}
		} else {
			if canRun {
				t.Error("expected to be throttled here")
			}
		}
	}

	time.Sleep(time.Second)

	canRun := limiter.AddEvent()
	if !canRun {
		t.Error("did not expect to be throttled")
	}

	canRun = limiter.AddEvent()
	if canRun {
		t.Error("expected to be throttled")
	}
}
