package handler

import (
	"net/http"

	"sketch/service"

	"github.com/labstack/echo/v4"
)

var (
	EchoErrBadRequest = echo.NewHTTPError(http.StatusBadRequest, "invalid request body")
)

// router represents our HTTP handler
type router struct {
	svc SketchService
}

// New returns an instance of the handler
func New(svc SketchService) *router {
	return &router{svc: svc}
}

// PrintCanvas prints the canvas
func (r router) PrintCanvas(e echo.Context) error {

	id := e.Param("id")

	paint, err := r.svc.PrintCanvas(id)

	if err != nil {
		return err
	}

	return e.String(http.StatusOK, paint)
}

// DrawRectangle draws a rectangle
func (r router) DrawRectangle(e echo.Context) error {

	id := e.Param("id")

	var request service.DrawRectangleRequest

	if e.Bind(&request) != nil {
		return EchoErrBadRequest
	}

	paint, _ := r.svc.DrawRectangle(id, request)

	return e.String(http.StatusOK, paint)
}

// FloodFill performs a flood fill operation on the canvas
func (r router) FloodFill(e echo.Context) error {

	id := e.Param("id")

	var request service.FloodFillRequest

	if e.Bind(&request) != nil {
		return EchoErrBadRequest
	}

	paint, _ := r.svc.FloodFill(id, request)

	return e.String(http.StatusOK, paint)
}
