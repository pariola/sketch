package handler

import (
	"sketch/service"
)

// SketchService is an interface that wraps the SketchService functionality
type SketchService interface {
	FloodFill(string, service.FloodFillRequest) (string, error)
	DrawRectangle(string, service.DrawRectangleRequest) (string, error)
}
