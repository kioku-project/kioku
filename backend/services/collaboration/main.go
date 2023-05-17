package main

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/kioku-project/kioku/pkg/helper"
	"github.com/kioku-project/kioku/services/collaboration/handler"
	pb "github.com/kioku-project/kioku/services/collaboration/proto"
	"github.com/kioku-project/kioku/store"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"

	"go-micro.dev/v4"
	"go-micro.dev/v4/logger"
	"go-micro.dev/v4/server"

	_ "github.com/go-micro/plugins/v4/registry/kubernetes"
	"github.com/go-micro/plugins/v4/wrapper/trace/opentelemetry"

	grpcc "github.com/go-micro/plugins/v4/client/grpc"
	grpcs "github.com/go-micro/plugins/v4/server/grpc"
)

var (
	service        = "collaboration"
	version        = "latest"
	serviceAddress = fmt.Sprintf("%s%s", os.Getenv("HOSTNAME"), ":8080")
	jaegerUrl = os.Getenv("JAEGER_ADDRESS")
)

func main() {

	// Initialize the database connection
	dbStore, err := store.NewCollaborationStore()
	if err != nil {
		logger.Fatal("Failed to initialize database:", err)
	}

	logger.Info("Trying to listen on: ", serviceAddress)

	opts := []micro.Option{
		micro.Name(service),
		micro.Version(version),
		micro.Address(serviceAddress),
	}
	// Create service
	srv := micro.NewService(
		micro.Server(grpcs.NewServer(server.Address(serviceAddress), server.Wait(nil))),
		micro.Client(grpcc.NewClient()),
	)

	// Initiliaze tracing
	if jaegerUrl != "" {
		tp, err := helper.NewTracerProvider(service, srv.Server().Options().Id, version, jaegerUrl)
		if err != nil {
			logger.Fatal(err)
		}
		defer func() {
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
			defer cancel()
			if err := tp.Shutdown(ctx); err != nil {
				logger.Fatal(err)
			}
		}()
		otel.SetTracerProvider(tp)
		otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))
		traceOpts := []opentelemetry.Option{
			opentelemetry.WithHandleFilter(func(ctx context.Context, r server.Request) bool {
				if e := r.Endpoint(); strings.HasPrefix(e, "Health.") {
					return true
				}
				return false
			}),
		}
		opts = append(opts, micro.WrapHandler(opentelemetry.NewHandlerWrapper(traceOpts...)))
	}
	srv.Init(opts...)

	// Create a new instance of the service handler with the initialized database connection
	svc := handler.New(dbStore)

	// Register handler
	if err := pb.RegisterCollaborationHandler(srv.Server(), svc); err != nil {
		logger.Fatal(err)
	}
	if err := pb.RegisterHealthHandler(srv.Server(), new(handler.Health)); err != nil {
		logger.Fatal(err)
	}

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
