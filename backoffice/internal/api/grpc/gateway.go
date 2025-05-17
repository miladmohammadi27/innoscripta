package grpc

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"backoffice/internal/api/grpc/dto"
	"backoffice/internal/helper/di"
	"backoffice/internal/helper/log"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/samber/do"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type gatewayServer struct {
	httpAddr   string // HTTP gateway server address
	grpcAddr   string // gRPC server address to dial
	lg         log.Logger
	httpServer *http.Server
	mux        *runtime.ServeMux
}

func NewGateway(i *do.Injector) (Gateway, error) {
	var cfg GRPCConfig
	if err := di.GetConfigFromDI(i, &cfg); err != nil {
		return nil, fmt.Errorf("failed to get gateway config from DI: %w", err)
	}

	httpAddr := fmt.Sprintf("%s:%d", cfg.GatewayIP, cfg.GatewayPort)
	grpcAddr := fmt.Sprintf("%s:%d", cfg.IP, cfg.Port)

	muxOpts := []runtime.ServeMuxOption{
		// Forward the "x-api-key" header to gRPC metadata
		runtime.WithIncomingHeaderMatcher(func(key string) (string, bool) {
			// Convert HTTP headers to gRPC metadata
			switch key {
			case "X-Api-Key", "x-api-key":
				return "x-api-key", true
			default:
				// Use default behavior for other headers
				return runtime.DefaultHeaderMatcher(key)
			}
		}),
	}

	mux := runtime.NewServeMux(muxOpts...)

	httpServer := &http.Server{
		Addr:    httpAddr,
		Handler: mux,
	}

	return &gatewayServer{
		httpAddr:   httpAddr,
		grpcAddr:   grpcAddr,
		lg:         do.MustInvoke[log.Logger](i),
		httpServer: httpServer,
		mux:        mux,
	}, nil
}

func (s *gatewayServer) ListenAndServe() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Set up a connection to the gRPC server
	conn, err := grpc.DialContext(
		ctx,
		s.grpcAddr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return fmt.Errorf("failed to dial gRPC server at %s: %w", s.grpcAddr, err)
	}
	defer conn.Close()

	// Register gRPC-Gateway handlers (add more services as needed)
	if err := dto.RegisterUserServiceHandler(ctx, s.mux, conn); err != nil {
		return fmt.Errorf("failed to register gateway handler: %w", err)
	}

	s.lg.Info(fmt.Sprintf("Starting gRPC-Gateway server on %s (proxying to gRPC at %s)", s.httpAddr, s.grpcAddr))

	if err := s.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return fmt.Errorf("failed to serve gRPC-Gateway: %w", err)
	}

	return nil
}

func (s *gatewayServer) Shutdown(ctx context.Context) error {
	s.lg.Info("Shutting down gRPC-Gateway server")
	return s.httpServer.Shutdown(ctx)
}
