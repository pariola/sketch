package canvas

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewRectangle(t *testing.T) {

	r := NewRectangle(0, 0, 5, 5, "X", "@")

	// coordinates
	assert.Equal(t, 0, r.posX, "invalid position x-axis")
	assert.Equal(t, 0, r.posY, "invalid position y-axis")

	// size
	assert.Equal(t, 5, r.width, "invalid width size")
	assert.Equal(t, 5, r.height, "invalid height size")

	// fill & outline
	assert.Equal(t, "X", r.fillChar, "invalid fill character")
	assert.Equal(t, "@", r.outlineChar, "invalid outline character")
}
