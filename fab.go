package fab

import (
	"fmt"
	"math"
	"regexp"
)

const (
	ESC = "\x1b["
	NND = "\x1b[0m"
)

type Fab struct {
	colors []int
	count int
	index int
	pattern string
}

func NewFab(c ...int) *Fab {
	fab := &Fab{pattern: "%s%dm%s%s"}
	if c == nil {
		fab.colors = []int{31, 32, 33, 34, 35, 36}
	} else {
		fab.colors = append(fab.colors, c...)
	}
	fab.count = len(fab.colors)
	return fab
}

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
	fab.pattern = "%s38;5;%dm%s%s"
	return fab
}

func (f *Fab) Paint(s string) string {
	if len(s) == 0 {
		return ""
	}
	p, _ := regexp.Compile(`.`)
	m := p.FindAllString(s, -1)
	var r string
	for _, c := range m {
		r = fmt.Sprintf("%s%s", r, f.PaintChar(c))
	}
	return r
}

func (f *Fab) PaintChar(c string) string {
	return fmt.Sprintf(f.pattern, ESC, f.NextColor(), c, NND)
}

func (f *Fab) NextColor() int {
	color := f.colors[f.index%f.count]
	f.index++
	return color
}

func GetFab(term string) *Fab {
	p, _ := regexp.Compile(`^xterm|-256color$`)
	if p.MatchString(term) {
		return NewSuperFab()
	}
	return NewFab()
}
