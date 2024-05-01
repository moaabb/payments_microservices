package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/moaabb/payments_microservices/customer/config"
	"github.com/moaabb/payments_microservices/customer/db"
	"github.com/moaabb/payments_microservices/customer/db/customerdb"
	"github.com/moaabb/payments_microservices/customer/handlers"
	"github.com/moaabb/payments_microservices/customer/logger"
	"github.com/moaabb/payments_microservices/customer/models/domainErrors"
	"go.uber.org/zap"
)

var log = logger.GetLogger()

func main() {
	cfg := config.LoadConfig()
	conn := db.ConnectToDatabase(cfg.DbUrl)
	repo := customerdb.NewCustomerRepository(conn)
	validator := domainErrors.NewValidator(log, validator.New())

	h := handlers.NewCustomerHandler(repo, log, validator)

	app := getRoutes(h)

	go func() {
		err := app.Listen(cfg.ServerPort)
		if err != nil {
			log.Fatal("could not start server", zap.Error(err))
		}
	}()

	log.Info(fmt.Sprintf("Server running and listening on port %s", cfg.ServerPort))

	s := make(chan os.Signal, 1)

	signal.Notify(s, os.Interrupt)

	sig := <-s

	log.Info(fmt.Sprintf("Signal received: %v. Shutting down the server gracefully...\n", sig))
	app.ShutdownWithTimeout(time.Second * 20)
}
