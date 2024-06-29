package telegram

import (
	"fmt"
	"log"
	"os"
	"time"

	gotgbot "github.com/PaulSonOfLars/gotgbot/v2"
)

type goTGBot struct {
	bot       *gotgbot.Bot
	channelID int64
	userID    int64
}

func NewGoTGBot(channelID, userID int64) *goTGBot {
	bot, err := gotgbot.NewBot(os.Getenv("TELEGRAM_TOKEN"), nil)
	if err != nil {
		log.Fatal(err)
	}

	return &goTGBot{
		bot:       bot,
		channelID: channelID,
		userID:    userID,
	}
}

func (t *goTGBot) CreateInviteLink() error {
	inviteLink, err := t.bot.CreateChatInviteLink(t.channelID, &gotgbot.CreateChatInviteLinkOpts{
		ExpireDate:  time.Now().Add(5 * time.Minute).Unix(),
		MemberLimit: 1,
	})
	if err != nil {
		return err
	}

	fmt.Println("================================================================")
	fmt.Printf("response: %+v\n", inviteLink)
	fmt.Println("================================================================")

	return nil
}

func (t *goTGBot) GetChannelMembers() error {
	chatInfo := gotgbot.ChatFullInfo{
		Id: t.channelID,
	}

	return nil
}
