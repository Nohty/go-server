package main

import (
	"fmt"
	"log"

	"github.com/Nohty/api/config"
	"github.com/Nohty/api/database"
	"github.com/Nohty/api/router"
	"github.com/Nohty/api/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New(fiber.Config{
		ErrorHandler: utils.ErrorHandler,
	})
	app.Use(cors.New())

	database.ConnectDB()

	router.SetupRoutes(app)
	log.Fatal(app.Listen(fmt.Sprintf(":%s", config.PORT)))
}
