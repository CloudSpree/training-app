package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/lightstep/otel-launcher-go/launcher"
	"net/http"
	"os"
)

var lightstepToken string
var environmentName string

func init() {
	lightstepToken = os.Getenv("LIGHTSTEP_TOKEN")
	environmentName = os.Getenv("ENVIRONMENT_NAME")
}

func main() {
	ls := launcher.ConfigureOpentelemetry(
		launcher.WithServiceName("notifications"),
		launcher.WithAccessToken(lightstepToken),
		launcher.WithResourceAttributes(map[string]string{
			"environment": environmentName,
		}),
	)
	defer ls.Shutdown()

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", hello)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "hello from notifications!")
}
