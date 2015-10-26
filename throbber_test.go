package busyIndicator

import (
	"bufio"
	"bytes"
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestRingFromStringSlice(t *testing.T) {
	var (
		tests = []struct {
			vals []string
			max  int
		}{
			{Spin, 1},
			{Dot, 3},
			{Grow, 1},
		}
	)
	for _, test := range tests {
		r, m := ringFromStringSlice(test.vals)

		if m != test.max {
			t.Errorf("Expected max string %d, got %d", test.max, m)
		}

		if r.Len() != len(test.vals) {
			t.Errorf(
				"Expected ring length %d, got %d",
				r.Len(),
				len(test.vals),
			)
		}

		for _, v := range test.vals {
			if r.Value != v {
				t.Errorf("Expected next value %s, got %s", v, r.Value)
			}
			r = r.Next()
		}
	}
}

func TestReplaceLineAfter(t *testing.T) {
	var (
		wantString = func(s string) []byte {
			return []byte(fmt.Sprintf(
				"%s%s%s",
				bytes.Repeat(backSpace, len(s)),
				clearFromCursor,
				s,
			))
		}
		tests = []struct {
			b   *bytes.Buffer
			in  []byte
			out []byte
		}{
			{new(bytes.Buffer), []byte("hello\n"), wantString("hello\n")},
			{new(bytes.Buffer), []byte("foo\n"), wantString("foo\n")},
			{new(bytes.Buffer), []byte("bar\n"), wantString("bar\n")},
		}
	)
	for _, test := range tests {
		bufRead := bufio.NewReader(test.b)
		bufWrite := bufio.NewWriter(test.b)
		replaceLineAfter(bufWrite, test.in, 1*time.Microsecond)
		out, err := bufRead.ReadBytes('\n')
		if err != nil {
			t.Error(err)
		}
		if !reflect.DeepEqual(test.out, out) {
			t.Errorf("Expected %s, got %s", test.out, out)
		}
	}
}
