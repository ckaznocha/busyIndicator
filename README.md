# busyIndicator

[![Build Status](http://img.shields.io/travis/ckaznocha/busyIndicator.svg?style=flat)](https://travis-ci.org/ckaznocha/busyIndicator)
[![License](http://img.shields.io/:license-mit-blue.svg)](http://ckaznocha.mit-license.org)

This is mostly a `for-fun` project as an excuse to mess with some things I don't
normally have a reason to use. There are other similar projects that may be more
mature.

--
    import "github.com/ckaznocha/busyIndicator"


## Usage

```go
var (

	//Spin is a standard spinner
	Spin = []string{"-", "\\", "|", "/"}

	//Dot is an ellipsis sequence
	Dot = []string{"", ".", "..", "..."}

	//Grow is a . that grows up to be an O
	Grow = []string{".", "ₒ", "ⱺ", "O", "ⱺ", "ₒ", "."}
)
```

#### func  Progress

```go
func Progress(
	marks ProgressMarks,
	progress chan float64,
	cancel chan bool,
)
```
Progress starts a progress bar

#### func  ProgressAnimated

```go
func ProgressAnimated(
	marks ProgressMarks,
	speed time.Duration,
	progress chan float64,
	cancel chan bool,
)
```
ProgressAnimated starts an animated progress bar

#### func  Throbber

```go
func Throbber(
	vals []string,
	speed time.Duration,
	c chan bool,
)
```
Throbber starts a throbber

#### func  ThrobberPostfixed

```go
func ThrobberPostfixed(
	postfix string,
	vals []string,
	speed time.Duration,
	c chan bool,
)
```
ThrobberPostfixed starts a throbber with a given static postfix

#### func  ThrobberPrefixed

```go
func ThrobberPrefixed(
	prefix string,
	vals []string,
	speed time.Duration,
	c chan bool,
)
```
ThrobberPrefixed starts a throbber with a given static prefix

#### func  ThrobberPrefixedAndPostfixed

```go
func ThrobberPrefixedAndPostfixed(
	prefix string,
	postfix string,
	vals []string,
	speed time.Duration,
	c chan bool,
)
```
ThrobberPrefixedAndPostfixed starts a throbber with a given static prefix and
postfix

#### type ProgressMarks

```go
type ProgressMarks struct {
	Completed  rune
	Incomplete rune
}
```

ProgressMarks holds the runes you'd like to use for your progress bar
