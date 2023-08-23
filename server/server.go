package server

import (
	"fmt"
	"goscrape/config"
	"os"

	db "goscrape/database"
	router "goscrape/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

const keyServerAddr = "serverAddr"

func InitializeServer() error {
	err := config.LoadEnv()
	if err != nil {
		return err
	}

	err = db.InitDb()
	if err != nil {
		return err
	}

	defer db.DisconnectDb()

	fmt.Println("STARTING SERVER")

	server := fiber.New()
	server.Use(recover.New())
	server.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path} ${latency}\n",
	}))

	router.InitRouter(server)

	port := os.Getenv("server_port")
	server.Listen(":" + port)

	return nil
}
