// Package fab shows your output fabulous!!
package fab

import (
	"bytes"
	"fmt"
	"math"
	"regexp"
)

//
const (
	// Start an escape sequence
	ESC = "\x1b["
	// End the escape sequence
	NND = "\x1b[0m"
)

// The Fabulous struct
type Fab struct {
	colors  []int
	count   int
	index   int
	pattern string
}

// Initializes a new Fab with an optional custom color range
func NewFab(c ...int) *Fab {
	fab := &Fab{pattern: "%s%dm%c%s"}
	if c == nil {
		fab.colors = []int{31, 32, 33, 34, 35, 36}
	} else {
		fab.colors = append(fab.colors, c...)
	}
	fab.count = len(fab.colors)
	return fab
}

// Initializes a new Fab with colors appropiate for terminals capable of
// displaying 256 colors.
func NewSuperFab() *Fab {
	pi3 := math.Pi / 3

	colors := make([]int, 0, 42)
	// https://github.com/seattlerb/minitest/blob/4373fc7ebf648e156902958466cc4678945eac29/lib/minitest/pride.rb#L78-L96
	for c, n := 0.0, 0.0; c <= 42.0; c += 1.0 {
		n = c * (1.0 / 6)
		r := (int)(3*math.Sin(n) + 3)
		g := (int)(3*math.Sin(n+2*pi3) + 3)
		b := (int)(3*math.Sin(n+4*pi3) + 3)

		colors = append(colors, 36*r+6*g+b+16)
	}
	fab := NewFab(colors...)
	fab.pattern = "%s38;5;%dm%c%s"
	return fab
}

// Paints a string with fabulous colors.
func (f *Fab) Paint(s string) string {
	if len(s) == 0 {
		return ""
	}
	var r bytes.Buffer
	for _, c := range s {
		fmt.Fprintf(&r, f.pattern, ESC, f.nextColor(), c, NND)
	}
	return r.String()
}

func (f *Fab) nextColor() int {
	color := f.colors[f.index%f.count]
	f.index++
	return color
}

// Gets a proper Fab for a provided terminal emulator.
func GetFab(term string) *Fab {
	p, _ := regexp.Compile(`^xterm|-256color$`)
	if p.MatchString(term) {
		return NewSuperFab()
	}
	return NewFab()
}
