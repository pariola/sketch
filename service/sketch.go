package service

import (
	"sketch/pkg/canvas"
)

// sketch
type sketch struct {
}

// New returns an instance of sketch service
func New() *sketch {
	return &sketch{}
}

// DrawRectangle retrieves the referenced canvas then draws a new Rectangle on it
func (s sketch) DrawRectangle(canvasId string, request DrawRectangleRequest) (string, error) {

	// todo: fetch from store
	c := canvas.New(10, 12)

	rectangle := canvas.
		NewRectangle(request.PosX, request.PosY, request.Width, request.Height, request.Fill, request.Outline)

	// todo: validate rectangle

	c.Draw(*rectangle)

	return c.Print(), nil
}
