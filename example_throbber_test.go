package busyIndicator_test

import (
	"time"

	"github.com/ckaznocha/busyIndicator"
)

func ExampleThrobber() {
	c := make(chan bool)

	go busyIndicator.Throbber(busyIndicator.Spin, 50*time.Millisecond, c)
	time.Sleep(1 * time.Second)
	c <- true

	go busyIndicator.Throbber(busyIndicator.Dot, 100*time.Millisecond, c)
	time.Sleep(1 * time.Second)
	c <- true

	go busyIndicator.Throbber(busyIndicator.Grow, 150*time.Millisecond, c)
	time.Sleep(10 * time.Second)
	c <- true
}

func ExampleThrobberPrefixed() {
	c := make(chan bool)

	go busyIndicator.ThrobberPrefixed("hello ", busyIndicator.Spin, 100*time.Millisecond, c)
	time.Sleep(1 * time.Second)
	c <- true
}
