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

	e.POST("/:id/draw", h.DrawRectangle)

	// start server
	log.Fatalln(e.Start(":5000"))
}
