package main

import (
	"os"

	"webstore-demo/internal/server"
	"webstore-demo/internal/server/api"
	"webstore-demo/internal/store"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	// Serve API documentation
	e.Static("/docs", "docs")
	//
	storage := store.New(store.StoreTypeMemory)
	srv := server.NewWebStoreServer(logger, storage)

	api.RegisterHandlers(e, &srv)

	server.Run(e)
}
