package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSketch_DrawRectangle(t *testing.T) {

	type check struct {
		request DrawRectangleRequest
		valid   bool
	}

	svc := New()

	id := svc.NewCanvas()

	checks := []check{
		{DrawRectangleRequest{0, 0, 4, 4, "", ""}, false},
		{DrawRectangleRequest{0, 0, 4, 4, "", "X"}, true},
	}

	for _, c := range checks {

		_, err := svc.DrawRectangle(id, c.request)

		if !c.valid {
			assert.Equal(t, ErrInvalidRectangle, err)
			continue
		}

		assert.Nil(t, err)
	}
}

func TestSketch_FloodFill(t *testing.T) {

	type check struct {
		request FloodFillRequest
		valid   bool
	}

	svc := New()

	id := svc.NewCanvas()

	checks := []check{
		{FloodFillRequest{-1, 0, ""}, true},
		{FloodFillRequest{0, -1, ""}, true},
		{FloodFillRequest{0, 0, "-"}, true},
	}

	for _, c := range checks {

		_, err := svc.FloodFill(id, c.request)

		if !c.valid {
			assert.Equal(t, ErrInvalidRectangle, err)
			continue
		}

		assert.Nil(t, err)
	}
}

func TestSketch_PrintCanvas(t *testing.T) {

	svc := New()

	id := svc.NewCanvas()

	request := DrawRectangleRequest{0, 0, 4, 4, "", "X"}

	_, err := svc.DrawRectangle(id, request)

	assert.Nil(t, err, "no error expected")

	// todo: assert paint
}

func TestSketch_NewCanvas(t *testing.T) {

	svc := New()

	id := svc.NewCanvas()

	assert.NotEmpty(t, id, "should return canvas id")
}
