package main

import (
	"flag"
	"sync"

	"github.com/labstack/echo"
)

var (
	sm       sync.Map
	flagPort = flag.String("port", ":8080", "define the port")
)

func main() {
	flag.Parse()

	e := echo.New()
	e.HideBanner = true

	e.GET("/", list)

	e.POST("/", create)

	e.DELETE("/:key", delete)

	e.Logger.Fatal(e.Start(*flagPort))
}
