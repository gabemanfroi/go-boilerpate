package api

import (
	"fmt"
	"github.com/gabemanfroi/go-boilerplate/application"
	"github.com/gabemanfroi/go-boilerplate/infra/IoC"
	"github.com/gabemanfroi/go-boilerplate/infra/core"
	"github.com/gabemanfroi/go-boilerplate/infra/db"
	"github.com/gabemanfroi/go-boilerplate/internal/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"log"
)

func init() {

	log.Printf("Setting Up Your Server...")

	utils.Check(godotenv.Load(), "Error while trying to read .env file...")

	core.LoadConfig()
	IoC.InitContainer()
	db.Migrate()
	log.Printf("Server Setup complete...")
}

func StartServer() {

	app := fiber.New()

	application.RegisterRoutes(app)

	log.Println(fmt.Sprintf("Starting Server on port %s", core.AppConfig.AppPort))
	log.Fatal(app.Listen(fmt.Sprintf(":%v", core.AppConfig.AppPort)))

}
