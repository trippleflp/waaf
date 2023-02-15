package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"gitlab.informatik.hs-augsburg.de/flomon/waaf/services/function-group/internal/handler"
	"gitlab.informatik.hs-augsburg.de/flomon/waaf/services/function-group/internal/postgres"
	"os"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.SetGlobalLevel(zerolog.DebugLevel)

	app := fiber.New()

	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))

	groupMiddleware := app.Group("/groups")
	groupMiddleware.Post("/:id", handler.GroupInfo)         // done
	groupMiddleware.Post("/:id/addUsers", handler.AddUsers) // done
	groupMiddleware.Delete("/removeUsers", handler.RemoveUsers)
	groupMiddleware.Post("/editUserRole", handler.EditUserRole)
	groupMiddleware.Post("/addFunctions", handler.AddFunctions)
	groupMiddleware.Post("/addFunctionGroups", handler.AddFunctionGroups)
	groupMiddleware.Post("/removeFunctions", handler.RemoveFunctions)
	groupMiddleware.Post("/removeFunctionGroups", handler.RemoveFunctionGroups)

	app.Post("/list", handler.ListEntitledGroups) // done
	app.Post("/create", handler.CreateGroup)      // done
	app.Post("/createInternalToken", handler.CreateToken)
	port := os.Getenv("PORT")
	if port == "" {
		port = "10001"
	}
	postgres.GetConnection()

	log.Fatal().Err(app.Listen(fmt.Sprintf(":%v", port)))
}
