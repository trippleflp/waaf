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

	groupMiddleware := app.Group("/groups/:id")                 // done
	groupMiddleware.Post("/", handler.GroupInfo)                // done
	groupMiddleware.Post("/deploy", handler.Deploy)             // done
	groupMiddleware.Post("/addUsers", handler.AddUsers)         // done
	groupMiddleware.Post("/removeUsers", handler.RemoveUsers)   // done
	groupMiddleware.Post("/editUserRole", handler.EditUserRole) // done
	groupMiddleware.Post("/addFunction", handler.AddFunction)   // done
	groupMiddleware.Post("/addFunctionGroups", handler.AddFunctionGroups)
	groupMiddleware.Post("/removeFunctions", handler.RemoveFunctions)
	groupMiddleware.Post("/removeFunctionGroups", handler.RemoveFunctionGroups)

	app.Post("/list", handler.ListEntitledGroups) // done
	app.Post("/create", handler.CreateGroup)      // done
	app.Post("/createInternalToken", handler.CreateToken)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	postgres.GetConnection()

	log.Fatal().Err(app.Listen(fmt.Sprintf(":%v", port)))
}
