package fab

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestFabCustomColors(t *testing.T) {
	fab := NewFab(31, 32)
	assert := assert.New(t)
        assert.Equal("\x1b[31mh\x1b[0m\x1b[32më\x1b[0m\x1b[31ml\x1b[0m\x1b[32ml\x1b[0m\x1b[31mõ\x1b[0m", fab.Paint("hëllõ"), "they should be equal")
}

func TestFabPaintChar(t *testing.T) {
	fab := NewFab()
	assert := assert.New(t)
        assert.Equal("\x1b[31mh\x1b[0m\x1b[32më\x1b[0m\x1b[33ml\x1b[0m\x1b[34ml\x1b[0m\x1b[35mõ\x1b[0m", fab.Paint("hëllõ"), "they should be equal")
}

func TestSuperFabPaintChar(t *testing.T) {
	fab := NewSuperFab()
	assert := assert.New(t)
	assert.Equal("\x1b[38;5;154mh\x1b[0m\x1b[38;5;154më\x1b[0m\x1b[38;5;148ml\x1b[0m\x1b[38;5;184ml\x1b[0m\x1b[38;5;184mõ\x1b[0m", fab.Paint("hëllõ"), "they should be equal")
}
