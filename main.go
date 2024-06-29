package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/shaileshhb/namastebot/src/server"
	"github.com/shaileshhb/namastebot/src/telegram"
)

var channelID int64
var userID int64

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	id, err := strconv.Atoi(os.Getenv("CHANNEL_ID"))
	if err != nil {
		panic(err)
	}

	channelID = int64(id)

	id, err = strconv.Atoi(os.Getenv("USER_ID"))
	if err != nil {
		panic(err)
	}

	userID = int64(id)

	// ========================================================================
	// bot, err := telegram.ConfigureTGBotAPI()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// // bot.Debug = true

	// err = telegram.CreateInviteLinkTGBOTAPI(bot)
	// if err != nil {
	// 	panic(err)
	// }

	// err = telegram.GetChatMemberTGBOTAPI(bot)
	// if err != nil {
	// 	panic(err)
	// }

	// err = telegram.GetUpdates(bot)
	// if err != nil {
	// 	panic(err)
	// }
	// ========================================================================

	testGoTgBot()

	// ========================================================================
	// var wg sync.WaitGroup

	// ser := server.NewServer("Basic telegram bot", &wg)
	// ser.CreateRouterInstance()

	// err = run(ser, context.Background())
	// if err != nil {
	// 	panic(err)
	// }
	// // log.Error(ser.App.Listen(":8080"))

	// // Stop Server On System Call or Interrupt.
	// ch := make(chan os.Signal, 1)
	// signal.Notify(ch, os.Interrupt, syscall.SIGINT)
	// <-ch
	// stopApp(ser)
}

func testGoTgBot() {
	bot := telegram.NewGoTGBot(channelID, userID)

	err := bot.CreateInviteLink()
	if err != nil {
		panic(err)
	}

}

func stopApp(ser *server.Server) {
	// app.Stop()
	ser.WG.Wait()
	fmt.Println("After wait")
	os.Exit(0)
}

// func run(ser *server.Server, ctx context.Context) error {
// 	listener, err := ngrok.Listen(ctx,
// 		config.HTTPEndpoint(),
// 		ngrok.WithAuthtokenFromEnv(),
// 	)
// 	if err != nil {
// 		return err
// 	}

// 	log.Info("App URL ", listener.URL())
// 	log.Error(ser.App.Listen(":8080"))
// 	return nil
// }
