package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/moaabb/payments_microservices/customer/config"
	"go.uber.org/zap"
)

func main() {

	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}

	cfg := config.LoadConfig()
	app := getRoutes()

	go func() {
		err = app.Listen(cfg.ServerPort)
		if err != nil {
			logger.Fatal("could not start server", zap.Error(err))
		}
	}()

	logger.Info(fmt.Sprintf("Server running and listening on port %s", cfg.ServerPort))

	s := make(chan os.Signal, 1)

	signal.Notify(s, os.Interrupt)

	sig := <-s

	logger.Info(fmt.Sprintf("Signal received: %v. Shutting down the server gracefully...\n", sig))
	app.ShutdownWithTimeout(time.Second * 20)
}
