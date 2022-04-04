package main

import (
	"context"
	"github.com/CloudSpree/training-app/pkg/span"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/lightstep/otel-launcher-go/launcher"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
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
	)
	defer ls.Shutdown()

	tracer := otel.Tracer("main")

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/api/v1/notifications", hello(tracer))

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}

// Handler
func hello(tracer trace.Tracer) func(c echo.Context) error {
	return func(c echo.Context) error {
		_, s := span.WithEnvironment(context.Background(), tracer, environmentName, "hello")
		defer s.End()

		return c.String(http.StatusOK, "hello from notifications!")
	}
}
