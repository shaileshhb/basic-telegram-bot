package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/gofiber/fiber/v2/log"
	"github.com/joho/godotenv"
	"github.com/shaileshhb/namastebot/src/server"
	"golang.ngrok.com/ngrok"
	"golang.ngrok.com/ngrok/config"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	var wg sync.WaitGroup

	ser := server.NewServer("Basic telegram bot", &wg)
	ser.CreateRouterInstance()

	err = run(ser, context.Background())
	if err != nil {
		panic(err)
	}
	// log.Error(ser.App.Listen(":8080"))

	// Stop Server On System Call or Interrupt.
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt, syscall.SIGINT)
	<-ch
	stopApp(ser)
}

func stopApp(ser *server.Server) {
	// app.Stop()
	ser.WG.Wait()
	fmt.Println("After wait")
	os.Exit(0)
}

func run(ser *server.Server, ctx context.Context) error {
	log.Info("inside run function")
	listener, err := ngrok.Listen(ctx,
		config.HTTPEndpoint(),
		ngrok.WithAuthtokenFromEnv(),
	)
	if err != nil {
		return err
	}

	log.Info("App URL ", listener.URL())
	log.Error(ser.App.Listen(":8080"))
	return nil
}
