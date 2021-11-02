package service

import (
	"errors"
	"log"

	"sketch/pkg/canvas"
	"sketch/storage"

	"github.com/google/uuid"
)

var (
	ErrNoCanvas         = errors.New("no valid canvas found")
	ErrInvalidRectangle = errors.New("invalid parameters provided for rectangle")
)

// sketch
type sketch struct {
	store storage.Store
	state map[string]*canvas.Canvas
}

// New returns an instance of sketch service
func New(store storage.Store) *sketch {

	// load initial state
	state, err := store.LoadState()

	if err != nil {
		log.Fatalln("can't load initial state from storage")
	}

	return &sketch{
		store: store,
		state: state,
	}
}

// NewCanvas creates and stores a new canvas with a GUID
func (s *sketch) NewCanvas() string {

	defer s.updateStore()

	id := uuid.NewString()

	// default 28x12
	// store new canvas with id
	s.state[id] = canvas.New(28, 12)

	return id
}

// getCanvas returns the canvas with the specified id from state, else returns nil
func (s sketch) getCanvas(id string) *canvas.Canvas {

	if c, ok := s.state[id]; ok {
		return c
	}

	return nil
}

// updateStore updates the underlying store with the current state of the sketch service
func (s sketch) updateStore() {

	err := s.store.SaveState(s.state)

	if err != nil {
		log.Println("failed to update store with latest state", err)
	}
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

	defer s.updateStore()

	c := s.getCanvas(canvasId)

	if c == nil {
		return "", ErrNoCanvas
	}

	var fill, outline byte

	if len(request.Fill) > 0 {
		fill = request.Fill[0]
	}

	if len(request.Outline) > 0 {
		outline = request.Outline[0]
	}

	rectangle := canvas.NewRectangle(request.PosX, request.PosY, request.Width, request.Height, fill, outline)

	if rectangle == nil {
		return "", ErrInvalidRectangle
	}

	c.Draw(*rectangle)

	return c.Print(), nil
}

// FloodFill retrieves the referenced canvas then performs the flood fill operation
func (s sketch) FloodFill(canvasId string, request FloodFillRequest) (string, error) {

	defer s.updateStore()

	c := s.getCanvas(canvasId)

	if c == nil {
		return "", ErrNoCanvas
	}

	var fill byte

	if len(request.Fill) > 0 {
		fill = request.Fill[0]
	}

	c.FloodFill(request.PosX, request.PosY, fill)

	return c.Print(), nil
}
