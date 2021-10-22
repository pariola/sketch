package main

import (
	"log"

	"sketch/handler"
	"sketch/service"

	"github.com/labstack/echo/v4"
)

func main() {

	// setup handler
	h := handler.New(
		service.New(),
	)

	// register routes
	e := echo.New()

	e.File("/*", "static/index.html")

	e.GET("/api/canvas", h.NewCanvas)
	e.GET("/api/canvas/:id", h.PrintCanvas)
	e.GET("/api/canvas/:id/ws", h.PrintCanvasWS)
	e.POST("/api/canvas/:id/draw", h.DrawRectangle)
	e.POST("/api/canvas/:id/floodfill", h.FloodFill)

	// start server
	log.Fatalln(e.Start(":5000"))
}
