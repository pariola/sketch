package main

import (
	"log"

	"sketch/handler"
	"sketch/service"
	"sketch/storage"

	"github.com/labstack/echo/v4"
)

func main() {

	store := storage.New("sketch.gob")

	// setup handler
	h := handler.New(
		service.New(store),
	)

	// register routes
	e := echo.New()

	// client view
	e.GET("/", h.NewCanvasPage)
	e.File("/:id", "static/index.html")

	// backend
	e.GET("/api/canvas", h.NewCanvas)
	e.GET("/api/canvas/:id", h.PrintCanvas)
	e.GET("/api/canvas/:id/ws", h.PrintCanvasWS)
	e.POST("/api/canvas/:id/draw", h.DrawRectangle)
	e.POST("/api/canvas/:id/floodfill", h.FloodFill)

	// start server
	log.Fatalln(e.Start(":5000"))
}
