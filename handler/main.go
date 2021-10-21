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

// Draw draws a rectangle
func (r router) Draw(e echo.Context) error {

	id := e.Param("id")

	var request service.DrawRequest

	if e.Bind(&request) != nil {
		return EchoErrBadRequest
	}

	paint, _ := r.svc.Draw(id, request)

	return e.String(http.StatusOK, paint)
}
