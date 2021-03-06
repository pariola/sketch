package canvas

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_inBounds(t *testing.T) {

	type boundTest struct {
		x, y   int
		result bool
	}

	// check a 2x2 matrix
	tests := []boundTest{
		{0, 0, true},
		{0, 1, true},
		{0, 2, false},
		{0, 3, false},
		{1, 0, true},
		{1, 1, true},
		{1, 2, false},
		{1, 3, false},
		{2, 0, false},
		{2, 1, false},
		{2, 2, false},
		{2, 3, false},
		{3, 0, false},
		{3, 1, false},
		{3, 2, false},
		{3, 3, false},
	}

	for _, b := range tests {
		result := inBounds(b.x, b.y, 2, 2)
		assert.Equalf(t, b.result, result, "cell [%d, %d]", b.x, b.y)
	}

}

func Test_inXBoundary(t *testing.T) {

	type boundary struct {
		x, y   int
		result bool
	}

	r := NewRectangle(0, 0, 3, 3, '.', empty)

	xBoundaries := []boundary{
		{0, 0, true},
		{0, 1, false},
		{0, 2, true},
		{0, 3, false},
		{1, 0, true},
		{1, 1, false},
		{1, 2, true},
		{1, 3, false},
		{2, 0, true},
		{2, 1, false},
		{2, 2, true},
		{2, 3, false},
		{3, 0, false},
		{3, 1, false},
		{3, 2, false},
		{3, 3, false},
	}

	for _, b := range xBoundaries {
		result := inXBoundary(b.x, b.y, *r)
		assert.Equalf(t, b.result, result, "cell [%d, %d]", b.x, b.y)
	}
}

func Test_inYBoundary(t *testing.T) {

	type boundary struct {
		x, y   int
		result bool
	}

	r := NewRectangle(0, 0, 3, 3, '.', empty)

	yBoundaries := []boundary{
		{0, 0, true},
		{0, 1, true},
		{0, 2, true},
		{0, 3, false},
		{1, 0, false},
		{1, 1, false},
		{1, 2, false},
		{1, 3, false},
		{2, 0, true},
		{2, 1, true},
		{2, 2, true},
		{2, 3, false},
		{3, 0, false},
		{3, 1, false},
		{3, 2, false},
		{3, 3, false},
	}

	for _, b := range yBoundaries {
		result := inYBoundary(b.x, b.y, *r)
		assert.Equalf(t, b.result, result, "cell [%d, %d]", b.x, b.y)
	}
}
