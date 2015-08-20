package fab

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestFabCustomColors(t *testing.T) {
	fab := NewFab(31, 32)

	assert := assert.New(t)
	assert.Equal("\x1b[31mh\x1b[0m", fab.PaintChar("h"), "they should be equal")
        assert.Equal("\x1b[32më\x1b[0m", fab.PaintChar("ë"), "they should be equal")
        assert.Equal("\x1b[31ml\x1b[0m", fab.PaintChar("l"), "they should be equal")
        assert.Equal("\x1b[32ml\x1b[0m", fab.PaintChar("l"), "they should be equal")
        assert.Equal("\x1b[31mõ\x1b[0m", fab.PaintChar("õ"), "they should be equal")
        assert.Equal("\x1b[32mh\x1b[0m\x1b[31më\x1b[0m\x1b[32ml\x1b[0m\x1b[31ml\x1b[0m\x1b[32mõ\x1b[0m", fab.Paint("hëllõ"), "they should be equal")
}

func TestFabPaintChar(t *testing.T) {
	fab := NewFab()
	assert := assert.New(t)
        assert.Equal("\x1b[31mh\x1b[0m", fab.PaintChar("h"))
        assert.Equal("\x1b[32më\x1b[0m", fab.PaintChar("ë"))
        assert.Equal("\x1b[33ml\x1b[0m", fab.PaintChar("l"))
	assert.Equal("\x1b[34ml\x1b[0m", fab.PaintChar("l"))
	assert.Equal("\x1b[35mõ\x1b[0m", fab.PaintChar("õ"))
	assert.Equal("\x1b[36m,\x1b[0m", fab.PaintChar(","))
	assert.Equal("\x1b[31m \x1b[0m", fab.PaintChar(" "))
	assert.Equal("\x1b[32mw\x1b[0m", fab.PaintChar("w"))
	assert.Equal("\x1b[33mö\x1b[0m", fab.PaintChar("ö"))
	assert.Equal("\x1b[34mr\x1b[0m", fab.PaintChar("r"))
	assert.Equal("\x1b[35ml\x1b[0m", fab.PaintChar("l"))
	assert.Equal("\x1b[36md\x1b[0m", fab.PaintChar("d"))
	assert.Equal("\x1b[31m.\x1b[0m", fab.PaintChar("."))
}

func TestSuperFabPaintChar(t *testing.T) {
	fab := NewSuperFab()
	assert := assert.New(t)

        assert.Equal("\x1b[38;5;154mh\x1b[0m", fab.PaintChar("h"))
	assert.Equal("\x1b[38;5;154më\x1b[0m", fab.PaintChar("ë"))
	assert.Equal("\x1b[38;5;148ml\x1b[0m", fab.PaintChar("l"))
	assert.Equal("\x1b[38;5;184ml\x1b[0m", fab.PaintChar("l"))
	assert.Equal("\x1b[38;5;184mõ\x1b[0m", fab.PaintChar("õ"))
	assert.Equal("\x1b[38;5;214m,\x1b[0m", fab.PaintChar(","))
	assert.Equal("\x1b[38;5;214m \x1b[0m", fab.PaintChar(" "))
	assert.Equal("\x1b[38;5;208mw\x1b[0m", fab.PaintChar("w"))
	assert.Equal("\x1b[38;5;208mö\x1b[0m", fab.PaintChar("ö"))
	assert.Equal("\x1b[38;5;203mr\x1b[0m", fab.PaintChar("r"))
	assert.Equal("\x1b[38;5;203ml\x1b[0m", fab.PaintChar("l"))
	assert.Equal("\x1b[38;5;198md\x1b[0m", fab.PaintChar("d"))
	assert.Equal("\x1b[38;5;198m.\x1b[0m", fab.PaintChar("."))
}
