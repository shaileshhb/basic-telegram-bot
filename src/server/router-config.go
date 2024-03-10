package server

import (
	"github.com/shaileshhb/namastebot/src/controllers"
	"github.com/shaileshhb/namastebot/src/routes"
)

func (ser *Server) CreateRouterInstance() {
	ser.InitializeRouter()

	telegramcon := controllers.NewTelegramController()
	telegramroute := routes.NewTelegramRouter(telegramcon)

	ser.RegisterRoutes([]Controller{telegramroute})
}
