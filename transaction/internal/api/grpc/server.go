package grpc

import (
	"context"
	"fmt"
	"net"

	"transaction/internal/api/grpc/dto"
	"transaction/internal/helper/di"
	"transaction/internal/helper/log"

	"github.com/samber/do"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

var apiKey string

type gRPCServer struct {
	addr       string
	netPort    string
	lg         log.Logger
	grpcServer *grpc.Server
}

func NewServer(i *do.Injector) (Server, error) {
	var cfg GRPCConfig
	if err := di.GetConfigFromDI(i, &cfg); err != nil {
		return nil, fmt.Errorf("failed to get grpc config from DI: %w", err)
	}

	apiKey = cfg.GRPCAPIKey

	addr := fmt.Sprintf("%s:%d", cfg.IP, cfg.Port)

	// Define interceptors
	unaryInterceptors := []grpc.UnaryServerInterceptor{
		apiKeyInterceptor,
	}

	server := grpc.NewServer(grpc.ChainUnaryInterceptor(unaryInterceptors...))

	// Registering services
	dto.RegisterBalanceServiceServer(server, do.MustInvoke[dto.BalanceServiceServer](i))

	return &gRPCServer{
		addr:       addr,
		netPort:    cfg.NetProt,
		lg:         do.MustInvoke[log.Logger](i),
		grpcServer: server,
	}, nil
}

func (s *gRPCServer) Serve() error {
	lis, err := net.Listen(s.netPort, s.addr)
	if err != nil {
		s.lg.Fatal("Failed to listen to GRPC port: ", err)
	}

	s.lg.Info(fmt.Sprintf("Starting grpc server ... : %s", s.addr))

	return s.grpcServer.Serve(lis)
}

func (s *gRPCServer) Shutdown(ctx context.Context) error {
	s.lg.Info("Shutting down grpc server")
	s.grpcServer.GracefulStop()

	return nil
}

func apiKeyInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	// Get metadata from context
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Errorf(codes.Unauthenticated, "metadata is not provided")
	}

	// Check for auth token
	apiKeyHeader, ok := md["x-api-key"]
	if !ok || len(apiKeyHeader) == 0 {
		return nil, status.Errorf(codes.Unauthenticated, "API Key is not provided")
	}

	// Check API Key
	if apiKeyHeader[0] != apiKey {
		return nil, status.Errorf(codes.Unauthenticated, "API Key is not provided")
	}

	return handler(ctx, req)
}
