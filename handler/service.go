package handler

import (
	"sketch/service"
)

// SketchService is an interface that wraps the SketchService functionality
type SketchService interface {
	PrintCanvas(string) (string, error)
	FloodFill(string, service.FloodFillRequest) (string, error)
	DrawRectangle(string, service.DrawRectangleRequest) (string, error)
}
