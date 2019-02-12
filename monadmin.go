package main

import (
	"fmt"
	"flag"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	srvPort := flag.String("p", "1323", "server port number")
	flag.Parse()

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", hello)

	// Start server
	svrAddr := fmt.Sprintf(":%s", *srvPort)
	e.Logger.Fatal(e.Start(svrAddr))
}

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World11131!")
}
