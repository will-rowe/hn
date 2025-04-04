package server

import (
	"context"
	"log"
	"net"
	"net/http"

	"github.com/rs/cors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	report "github.com/will-rowe/hn/api/gen/go/report/v1"
	"github.com/will-rowe/hn/backend/handlers"
	"github.com/will-rowe/hn/backend/middleware"
	"github.com/will-rowe/hn/backend/reporting"
)

func Start() {
	grpcPort := ":9090"
	httpPort := ":8080"

	// gRPC setup
	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(middleware.AuthUnaryInterceptor),
	)
	svc := reporting.NewReportService()
	handler := handlers.NewReportHandler(svc)
	report.RegisterReportServiceServer(grpcServer, handler)
	reflection.Register(grpcServer)

	go func() {
		listener, err := net.Listen("tcp", grpcPort)
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
		log.Printf("gRPC server listening on %s", grpcPort)
		if err := grpcServer.Serve(listener); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	// gRPC-Gateway setup
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	gatewayMux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}

	if err := report.RegisterReportServiceHandlerFromEndpoint(ctx, gatewayMux, "localhost"+grpcPort, opts); err != nil {
		log.Fatalf("failed to register handler: %v", err)
	}

	// Add a /healthz endpoint to the default HTTP mux
	healthMux := http.NewServeMux()
	healthMux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("ok"))
	})

	// Combine health check and gRPC-Gateway mux
	rootMux := http.NewServeMux()
	rootMux.Handle("/healthz", healthMux)
	rootMux.Handle("/", middleware.AuthHTTPMiddleware(cors.Default().Handler(gatewayMux)))

	log.Printf("HTTP server listening on %s", httpPort)
	log.Fatal(http.ListenAndServe(httpPort, rootMux))
}
