package controllers

import (
	"fmt"
	"net/url"
	"strconv"

	"github.com/shaileshhb/namastebot/src/api"
)

// TelegramController
type TelegramController interface {
	WebhookPost(details map[string]interface{}) error
	HandleMessage(messageObject map[string]interface{}) error
}

type telegramController struct {
}

// NewTelegramController will create new instance of telegram controller.
func NewTelegramController() TelegramController {
	return &telegramController{}
}

// WebhookPost will receive data when some message is sent from telegram.
func (t *telegramController) WebhookPost(details map[string]interface{}) error {
	fmt.Printf("%+v\n", details)
	return nil
}

func (t *telegramController) sendMessage(messageObject map[string]interface{}, messageText string) error {

	params := url.Values{}
	chatObject := messageObject["chat"].(map[string]interface{})
	chatID := strconv.FormatFloat(chatObject["id"].(float64), 'f', 2, 64)

	params.Add("chat_id", chatID)
	params.Add("text", messageText)

	result, err := api.GetRequest("/sendMessage", params)
	if err != nil {
		return err
	}

	fmt.Printf("=== result - %+v\n", result)
	return nil
}

func (t *telegramController) HandleMessage(messageObject map[string]interface{}) error {
	messageText := messageObject["text"].(string)

	if messageText[0] == '/' {
		fmt.Println("========= command ->", messageText)
		switch messageText {
		case "/start":
			return t.sendMessage(messageObject, "Hurray!!!!!!")
		default:
			return t.sendMessage(messageObject, "This command does not exist :(")
		}
	}

	return t.sendMessage(messageObject, messageText)
}
