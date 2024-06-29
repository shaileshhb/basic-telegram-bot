package telegram

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type tgBotAPI struct {
	bot       *tgbotapi.BotAPI
	channelID int64
	userID    int64
}

func NewTGBotAPI(channelID, userID int64) (*tgBotAPI, error) {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_TOKEN"))
	if err != nil {
		return nil, err
	}

	return &tgBotAPI{
		bot:       bot,
		channelID: channelID,
		userID:    userID,
	}, nil
}

func GetUpdates(bot *tgbotapi.BotAPI) error {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	// Create a new cancellable background context. Calling `cancel()` leads to the cancellation of the context
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	// `updates` is a golang channel which receives telegram updates
	updates := bot.GetUpdatesChan(u)

	// Pass cancellable context to goroutine
	go receiveUpdates(ctx, updates)

	// Tell the user the bot is online
	log.Println("Start listening for updates. Press enter to stop")

	// Wait for a newline symbol, then cancel handling updates
	bufio.NewReader(os.Stdin).ReadBytes('\n')
	cancel()

	return nil
}

func receiveUpdates(ctx context.Context, updates tgbotapi.UpdatesChannel) {
	// `for {` means the loop is infinite until we manually stop it
	for {
		select {
		// stop looping if ctx is cancelled
		case <-ctx.Done():
			return
		// receive update from channel and then handle it
		case update := <-updates:
			message := update.Message

			// this will fetch channel id
			if update.ChannelPost != nil {
				channelID := update.ChannelPost.Chat.ID
				fmt.Println("=============ChannelPost===================")
				fmt.Printf("Channel ID: %+v\n", update.ChannelPost.Chat)
				fmt.Printf("Channel ID: %d\n", channelID)
				fmt.Println("================================")
			}

			if message != nil {
				fmt.Println("=============message===================")
				fmt.Printf("Text - %+v\n", message.Text)
				fmt.Println("================================")
			}

			if update.ChannelPost.NewChatMembers != nil {
				fmt.Println("===============ChannelPost NewChatMembers=================")
				fmt.Printf("Text - %+v\n", update.ChannelPost.NewChatMembers)
				fmt.Println("================================")
			}

			if message.NewChatMembers != nil {
				fmt.Println("===============NewChatMembers=================")
				fmt.Printf("Text - %+v\n", message.NewChatMembers)
				fmt.Println("================================")
			}
		}
	}
}

func (t *tgBotAPI) CreateInviteLink() error {

	// for users who are not already in the group or where kicked out of the group
	inviteLinkConfig := tgbotapi.CreateChatInviteLinkConfig{
		ChatConfig: tgbotapi.ChatConfig{
			ChatID: t.channelID,
		},
		ExpireDate:         int(time.Now().Add(1 * time.Hour).Unix()),
		CreatesJoinRequest: false,
		MemberLimit:        1,
		Name:               "channelName",
	}

	apiResponse, err := t.bot.Request(inviteLinkConfig)
	if err != nil {
		return err
	}

	response, err := json.Marshal(apiResponse.Result)
	if err != nil {
		return err
	}

	fmt.Println("================================================================")
	fmt.Printf("response: %+v\n", string(response))
	fmt.Println("================================================================")

	return nil
}

func (t *tgBotAPI) GetChatMember() error {
	chatMembers, err := t.bot.GetChatMember(tgbotapi.GetChatMemberConfig{
		ChatConfigWithUser: tgbotapi.ChatConfigWithUser{
			ChatID: t.channelID,
			UserID: t.userID,
		},
	})

	if err != nil {
		return err
	}

	fmt.Println("================================================================")
	fmt.Printf("response: %+v\n", chatMembers)
	fmt.Println("================================================================")

	return nil
}

func (t *tgBotAPI) RemoveUser(bot *tgbotapi.BotAPI) error {

	kickMember := tgbotapi.KickChatMemberConfig{
		ChatMemberConfig: tgbotapi.ChatMemberConfig{
			ChatID: t.channelID,
			UserID: t.userID,
		},
		// RevokeMessages: true,
	}

	apiResponse, err := t.bot.Request(kickMember)
	if err != nil {
		return err
	}

	response, err := json.Marshal(apiResponse.Result)
	if err != nil {
		return err
	}

	fmt.Println("======================REMOVE USER==========================================")
	fmt.Printf("response: %+v\n", string(response))
	fmt.Println("================================================================")

	return nil
}

func (t *tgBotAPI) UnbanChatMember() error {

	unbanMember := tgbotapi.UnbanChatMemberConfig{
		ChatMemberConfig: tgbotapi.ChatMemberConfig{
			ChatID: t.channelID,
			UserID: t.userID,
		},
	}

	apiResponse, err := t.bot.Request(unbanMember)
	if err != nil {
		return err
	}

	response, err := json.Marshal(apiResponse.Result)
	if err != nil {
		return err
	}

	fmt.Println("======================Unban USER==========================================")
	fmt.Printf("response: %+v\n", string(response))
	fmt.Println("================================================================")

	return nil
}

func GetChannelMembers(bot *tgbotapi.BotAPI) error {

	return nil
}
