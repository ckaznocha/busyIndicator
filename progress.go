package busyIndicator

import (
	"bufio"
	"bytes"
	"fmt"
	"math"
	"os"
	"time"

	"github.com/nsf/termbox-go"
)

//ProgressMarks holds the runes you'd like to use for your progress bar
type ProgressMarks struct {
	Completed  rune
	Incomplete rune
}

//Progress starts a progress bar
func Progress(
	marks ProgressMarks,
	progress chan float64,
	cancel chan bool,
) {
	doProgress(false, 0, progress, cancel, marks)
}

//ProgressAnimated starts an animated progress bar
func ProgressAnimated(
	marks ProgressMarks,
	speed time.Duration,
	progress chan float64,
	cancel chan bool,
) {
	doProgress(true, speed, progress, cancel, marks)
}

func doProgress(
	animated bool,
	speed time.Duration,
	progress chan float64,
	cancel chan bool,
	marks ProgressMarks,
) {
	width, err := getWidth()
	if err != nil {
		return
	}
	var (
		buffer = bufio.NewWriter(os.Stdout)
		max    = int(math.Ceil(width - 9))
		last   int
	)
	for {
		select {
		case <-cancel:
			return
		case p := <-progress:
			completed := int(math.Ceil(float64(max) * p))
			if animated != true {
				last = completed
			}
			for last <= completed && completed <= 100 {
				done := bytes.Repeat(
					[]byte(string(marks.Completed)),
					int(last),
				)
				toDo := bytes.Repeat(
					[]byte(string(marks.Incomplete)),
					int(max-last),
				)
				replaceLineAfter(
					buffer,
					[]byte(fmt.Sprintf("%6.2f%%|%s%s|", p*100, done, toDo)),
					speed,
				)
				last++
			}
			last = completed
		}
	}
}

func getWidth() (float64, error) {
	var (
		err   error
		width int
	)
	err = termbox.Init()
	if err != nil {
		return 0.0, err
	}
	defer termbox.Close()
	width, _ = termbox.Size()
	return float64(width), nil
}
