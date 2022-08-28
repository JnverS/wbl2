package main

import (
	"testing"
	"time"
)

func TestOrChannel(t *testing.T) {
	sig := func(after time.Duration) <- chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
	}()
	return c
	}

	start := time.Now()
	<-or (
		sig(2*time.Hour),
		sig(5*time.Minute),
		sig(1*time.Second),
		sig(1*time.Hour),
		sig(1*time.Minute),
	)
	since := time.Since(start)
	res := time.Second * 2
	if since > res {
		t.Errorf("%v > %v", since, res)
	}
}