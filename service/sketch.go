package service

import (
	"errors"

	"sketch/pkg/canvas"

	"github.com/google/uuid"
)

var (
	ErrNoCanvas         = errors.New("no valid canvas found")
	ErrInvalidRectangle = errors.New("invalid parameters provided for rectangle")
)

// sketch
type sketch struct {
	store map[string]*canvas.Canvas
}

// New returns an instance of sketch service
func New() *sketch {
	return &sketch{
		store: make(map[string]*canvas.Canvas, 0),
	}
}

// NewCanvas creates and stores a new canvas with a GUID
func (s *sketch) NewCanvas() string {

	id := uuid.NewString()

	// default 28x12
	// store new canvas with id
	s.store[id] = canvas.New(28, 12)

	return id
}

// getCanvas returns the canvas with the specified id from store, else returns nil
func (s sketch) getCanvas(id string) *canvas.Canvas {

	if c, ok := s.store[id]; ok {
		return c
	}

	return nil
}

// PrintCanvas returns the string representation of the referenced canvas
func (s sketch) PrintCanvas(canvasId string) (string, error) {

	c := s.getCanvas(canvasId)

	if c == nil {
		return "", ErrNoCanvas
	}

	return c.Print(), nil
}

// DrawRectangle retrieves the referenced canvas then draws a new Rectangle on it
func (s sketch) DrawRectangle(canvasId string, request DrawRectangleRequest) (string, error) {

	c := s.getCanvas(canvasId)

	if c == nil {
		return "", ErrNoCanvas
	}

	rectangle := canvas.
		NewRectangle(request.PosX, request.PosY, request.Width, request.Height, request.Fill, request.Outline)

	if rectangle == nil {
		return "", ErrInvalidRectangle
	}

	c.Draw(*rectangle)

	return c.Print(), nil
}

// FloodFill retrieves the referenced canvas then performs the flood fill operation
func (s sketch) FloodFill(canvasId string, request FloodFillRequest) (string, error) {

	c := s.getCanvas(canvasId)

	if c == nil {
		return "", ErrNoCanvas
	}

	c.FloodFill(request.PosX, request.PosY, request.Fill)

	return c.Print(), nil
}
