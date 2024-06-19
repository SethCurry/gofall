package gofall

import (
	"testing"
	"time"
)

func Test_RateLimiter(t *testing.T) {
	t.Parallel()

	limiter := newRateLimiter(time.Second, 1)

	for i := 0; i < 5; i++ {
		ok := limiter.AddEvent()
		if i == 0 {
			if !ok {
				t.Error("did not expect to be throttled here")
			}
		} else {
			if ok {
				t.Error("expected to be throttled here")
			}
		}
	}

	time.Sleep(time.Second)

	ok := limiter.AddEvent()
	if !ok {
		t.Error("did not expect to be throttled")
	}
	ok = limiter.AddEvent()
	if ok {
		t.Error("expected to be throttled")
	}
}
