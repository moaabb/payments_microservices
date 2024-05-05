package main

import (
	"context"
	"os"
	"os/signal"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/moaabb/payments_microservices/customer/config"
	"github.com/moaabb/payments_microservices/customer/db"
	"github.com/moaabb/payments_microservices/customer/db/customerdb"
	"github.com/moaabb/payments_microservices/customer/handlers"
	logging "github.com/moaabb/payments_microservices/customer/logger"
	"github.com/moaabb/payments_microservices/customer/models/domainErrors"
	"github.com/moaabb/payments_microservices/customer/observability"
	"github.com/moaabb/payments_microservices/customer/services"
	"go.opentelemetry.io/otel"
)

var cfg *config.Config
var logger *logging.ApplicationLogger

var (
	otlpEndpoint string
)

func init() {
	cfg = config.LoadConfig()
	err := logging.InitLogger(cfg.LogLevel, cfg.AppName, cfg.Environment)
	if err != nil {
		panic(err)
	}
	logger = logging.GetLogger()

	otlpEndpoint = os.Getenv("OTLP_ENDPOINT")
	if otlpEndpoint == "" {
		panic("You MUST set OTLP_ENDPOINT env variable!")
	}
}

func main() {

	ctx := context.Background()

	exp, err := observability.NewOTLPExporter(ctx, otlpEndpoint)

	if err != nil {
		logger.Fatalf("failed to initialize exporter: %s", err.Error())
	}

	// Create a new tracer provider with a batch span processor and the given exporter.
	tp := observability.NewTraceProvider(exp, cfg.AppName)

	// Handle shutdown properly so nothing leaks.
	defer func() { _ = tp.Shutdown(ctx) }()

	otel.SetTracerProvider(tp)

	cfg := config.LoadConfig()
	conn := db.ConnectToDatabase(cfg.DbUrl)
	repo := customerdb.NewCustomerRepository(conn)
	validator := domainErrors.NewValidator(validator.New())
	svc := services.NewCustomerService(repo)

	h := handlers.NewCustomerHandler(svc, validator)

	app := getRoutes(h, tp)

	go func() {
		err := app.Listen(cfg.ServerPort)
		if err != nil {
			logger.Fatalf("could not start server: %s", err)
		}
	}()

	logger.Infof("Server running and listening on port %s", cfg.ServerPort)

	s := make(chan os.Signal, 1)

	signal.Notify(s, os.Interrupt)

	sig := <-s

	logger.Infof("Signal received: %v. Shutting down the server gracefully...\n", sig)
	app.ShutdownWithTimeout(time.Second * 20)
}
