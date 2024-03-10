package routes

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/shaileshhb/namastebot/src/controllers"
)

// telegramRouter
type telegramRouter struct {
	con controllers.TelegramController
}

// NewTelegramRouter will create new instance of telegramRouter.
func NewTelegramRouter(con controllers.TelegramController) *telegramRouter {
	return &telegramRouter{
		con: con,
	}
}

// RegisterRoutes will register routes for user-group router.
func (t *telegramRouter) RegisterRoutes(router fiber.Router) {
	// router.Post("/", t.webhookPost)
	router.Post("/", t.handleMessage)

	log.Info("Telegram routes registered")
}

// webhookPost will receive data when some message is sent from telegram.
func (t telegramRouter) webhookPost(c *fiber.Ctx) error {
	details := make(map[string]interface{})

	err := c.BodyParser(&details)
	if err != nil {
		log.Error(err)
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	err = t.con.WebhookPost(details)
	if err != nil {
		log.Error(err)
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(http.StatusCreated).JSON(details)
}

func (t *telegramRouter) handleMessage(c *fiber.Ctx) error {
	var requestObj map[string]interface{}

	err := c.BodyParser(&requestObj)
	if err != nil {
		log.Error(err)
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	messageObj := requestObj["message"].(map[string]interface{})

	err = t.con.HandleMessage(messageObj)
	if err != nil {
		log.Error(err)
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(nil)
}
