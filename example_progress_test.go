package busyIndicator_test

import (
	"testing"
	"time"

	"github.com/ckaznocha/busyIndicator"
)

func ExampleProgress(t *testing.T) {
	var (
		i float64
		c = make(chan bool)
		p = make(chan float64)
	)
	defer close(c)
	defer close(p)
	go busyIndicator.Progress(busyIndicator.ProgressMarks{'#', '~'}, p, c)

	for i < 1 {
		p <- i
		i += 0.10
		time.Sleep(1 * time.Second)
	}
	c <- true
}
func ExampleProgressAnimated() {
	var (
		i float64
		c = make(chan bool)
		p = make(chan float64)
	)
	defer close(c)
	defer close(p)
	go busyIndicator.ProgressAnimated(
		busyIndicator.ProgressMarks{'=', '-'},
		100*time.Millisecond,
		p,
		c,
	)

	for i < 1 {
		p <- i
		i += 0.10
		time.Sleep(2 * time.Second)
	}
	c <- true
}
