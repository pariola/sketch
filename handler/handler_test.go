package handler

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"sketch/service"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRouter_NewCanvasPage(t *testing.T) {

	e := echo.New()

	h := New(service.New())

	req := httptest.NewRequest(http.MethodGet, "/", nil)

	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
	ctx.SetPath("/")

	err := h.NewCanvasPage(ctx)

	if assert.NoError(t, err) {
		assert.Equal(t, http.StatusTemporaryRedirect, rec.Code)
	}
}

func TestRouter_NewCanvas(t *testing.T) {
	e := echo.New()

	h := New(service.New())

	req := httptest.NewRequest(http.MethodGet, "/", nil)

	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
	ctx.SetPath("/api/canvas")

	err := h.NewCanvas(ctx)

	if assert.NoError(t, err) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.NotEmpty(t, rec.Body.String())
	}
}

func TestRouter_DrawRectangle(t *testing.T) {

	e := echo.New()

	svc := service.New()
	canvasId := svc.NewCanvas()

	h := New(svc)

	requestBody := `
{
	"pos_x": 5,
	"pos_y": 2,
	"width": 5,
	"height": 3,
	"fill": "O",
	"outline": "V"
}`

	responseBody := `VVVVV VOOOV VVVVV`

	req := httptest.
		NewRequest(http.MethodPost, "/", strings.NewReader(requestBody))

	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
	ctx.SetPath("/api/canvas/:id/draw")
	ctx.SetParamNames("id")
	ctx.SetParamValues(canvasId)

	err := h.DrawRectangle(ctx)

	if assert.NoError(t, err) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, responseBody, rec.Body.String())
	}
}

func TestRouter_FloodFill(t *testing.T) {

	e := echo.New()

	svc := service.New()
	canvasId := svc.NewCanvas()

	h := New(svc)

	requestBody := `
{
	"pos_x": 4,
	"pos_y": 3,
	"fill": "."
}`

	responseBody := `............................ ............................ ............................ ............................ ............................ ............................ ............................ ............................ ............................ ............................ ............................ ............................`

	req := httptest.
		NewRequest(http.MethodPost, "/", strings.NewReader(requestBody))

	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
	ctx.SetPath("/api/canvas/:id/floodfill")
	ctx.SetParamNames("id")
	ctx.SetParamValues(canvasId)

	err := h.FloodFill(ctx)

	if assert.NoError(t, err) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, responseBody, rec.Body.String())
	}
}

func TestRouter_PrintCanvas(t *testing.T) {

	e := echo.New()

	svc := service.New()
	canvasId := svc.NewCanvas()

	svcRequest := service.DrawRectangleRequest{
		PosX: 5, PosY: 2,
		Width: 5, Height: 3,
		Fill: "O", Outline: "X",
	}

	responseBody, err := svc.DrawRectangle(canvasId, svcRequest)

	require.NoError(t, err)

	assert.Equal(t, `XXXXX XOOOX XXXXX`, responseBody)

	h := New(svc)

	req := httptest.NewRequest(http.MethodGet, "/", nil)

	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
	ctx.SetPath("/api/canvas/:id")
	ctx.SetParamNames("id")
	ctx.SetParamValues(canvasId)

	err = h.PrintCanvas(ctx)

	if assert.NoError(t, err) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, responseBody, rec.Body.String())
	}

}
