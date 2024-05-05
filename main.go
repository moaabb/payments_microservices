package main

import (
	"context"
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
	"github.com/moaabb/payments_microservices/customer/observability"
	"github.com/moaabb/payments_microservices/customer/services"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
)

var log = logger.GetLogger()

var (
	tracer       trace.Tracer
	otlpEndpoint string
)

func init() {
	otlpEndpoint = os.Getenv("OTLP_ENDPOINT")
	if otlpEndpoint == "" {
		log.Fatal("You MUST set OTLP_ENDPOINT env variable!")
	}
}

func main() {
	ctx := context.Background()

	// For testing to print out traces to the console
	// exp, err := newConsoleExporter()
	exp, err := observability.NewOTLPExporter(ctx, otlpEndpoint)

	if err != nil {
		log.Fatal("failed to initialize exporter", zap.Error(err))
	}

	// Create a new tracer provider with a batch span processor and the given exporter.
	tp := observability.NewTraceProvider(exp)

	// Handle shutdown properly so nothing leaks.
	defer func() { _ = tp.Shutdown(ctx) }()

	otel.SetTracerProvider(tp)

	// Finally, set the tracer that can be used for this package.
	tracer = tp.Tracer("customer_svc")

	cfg := config.LoadConfig()
	conn := db.ConnectToDatabase(cfg.DbUrl)
	repo := customerdb.NewCustomerRepository(conn, tracer)
	validator := domainErrors.NewValidator(log, validator.New())
	svc := services.NewCustomerService(repo, log, tracer)

	h := handlers.NewCustomerHandler(svc, log, validator, tracer)

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
