package main

import (
	"fmt"
	"function-uploader/internal/handler"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.SetGlobalLevel(zerolog.DebugLevel)

	app := fiber.New()

	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))

	app.Post("/:functionGroup/:functionName", handler.UploadHandler)
	port := os.Getenv("PORT")
	if port == "" {
		port = "10004"
	}
	log.Fatal().Err(app.Listen(fmt.Sprintf(":%v", port)))
}
