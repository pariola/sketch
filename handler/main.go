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

// NewCanvas creates a new canvas
func (r router) NewCanvas(e echo.Context) error {
	id := r.svc.NewCanvas()
	return e.String(http.StatusOK, id)
}

// PrintCanvas prints the canvas
func (r router) PrintCanvas(e echo.Context) error {

	id := e.Param("id")

	paint, err := r.svc.PrintCanvas(id)

	switch err {
	case nil:
	case service.ErrNoCanvas:
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	default:
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
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

	paint, err := r.svc.DrawRectangle(id, request)

	switch err {
	case nil:
	case service.ErrNoCanvas:
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	case service.ErrInvalidRectangle:
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	default:
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return e.String(http.StatusOK, paint)
}

// FloodFill performs a flood fill operation on the canvas
func (r router) FloodFill(e echo.Context) error {

	id := e.Param("id")

	var request service.FloodFillRequest

	if e.Bind(&request) != nil {
		return EchoErrBadRequest
	}

	paint, err := r.svc.FloodFill(id, request)

	switch err {
	case nil:
	case service.ErrNoCanvas:
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	default:
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return e.String(http.StatusOK, paint)
}
