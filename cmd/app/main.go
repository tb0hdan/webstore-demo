package main

import (
	"net/http"
	"os"
	"strings"

	"webstore-demo/internal/server"
	"webstore-demo/internal/server/api"
	"webstore-demo/internal/store"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog"
)

func EnforceAPIJSON(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if strings.HasPrefix(c.Request().RequestURI, "/api/") &&
			c.Request().Header.Get("Content-Type") != echo.MIMEApplicationJSON {
			return c.JSON(http.StatusBadRequest, "Missing Content-Type header")
		}

		return next(c)
	}
}

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.RequestID())
	e.Use(EnforceAPIJSON)
	// Serve API documentation
	e.Static("/docs", "docs")
	// for health checks
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, c.Response().Header().Get(echo.HeaderXRequestID))
	})
	//
	storage := store.New(store.Memory)
	srv := server.NewWebStoreServer(logger, storage)

	api.RegisterHandlers(e, &srv)

	server.Run(e)
}
