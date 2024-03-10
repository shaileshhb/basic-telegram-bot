package server

import (
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/gofiber/fiber/v2"
)

// Controller is implemented by the controllers.
type Controller interface {
	RegisterRoutes(router fiber.Router)
}

// Server Struct For Start the equisplit service.
type Server struct {
	Name            string
	App             *fiber.App
	Router          fiber.Router
	WG              *sync.WaitGroup
	TelegramBaseUrl string
}

// auth security.Authentication,
// NewServer will create new instance for the server.
func NewServer(name string, wg *sync.WaitGroup) *Server {
	return &Server{
		Name:            name,
		WG:              wg,
		TelegramBaseUrl: fmt.Sprintf("%s%s", os.Getenv("TELEGRAM_BASEURL"), os.Getenv("TELEGRAM_TOKEN")),
	}
}

// InitializeRouter Register the route.
func (ser *Server) InitializeRouter() {
	app := fiber.New(fiber.Config{
		AppName: ser.Name,
	})

	app.Get("/", func(c *fiber.Ctx) error {
		log.Println("get test")
		return c.Status(200).JSON(fiber.Map{
			"message": "Get hello world!!!",
		})
	})

	apiV1 := app.Group("")

	ser.App = app
	ser.Router = apiV1
}

// RegisterRoutes will register the specified routes in controllers.
func (ser *Server) RegisterRoutes(controllers []Controller) {
	for _, controller := range controllers {
		controller.RegisterRoutes(ser.Router)
	}
}
