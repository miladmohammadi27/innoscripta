package grpc

type GRPCConfig struct {
	// gRPC server
	IP            string `env:"GRPC_NETWORK_IP"             envDefault:"localhost"`
	Port          uint16 `env:"GRPC_NETWORK_PORT"           envDefault:"9000"`
	NetProt       string `env:"GRPC_NETWORK_PROTOCOL"       envDefault:"tcp"`
	KeyFile       string `env:"GRPC_NETWORK_KEY"            envDefault:""`
	CertFile      string `env:"GRPC_NETWORK_CERT"           envDefault:""`
	EnableReflect bool   `env:"GRPC_NETWORK_ENABLE_REFLECT" envDefault:"false"`
	EnableMetrics bool   `env:"GRPC_NETWORK_ENABLE_METRICS" envDefault:"false"`
	EnableTraces  bool   `env:"GRPC_NETWORK_ENABLE_TRACER"  envDefault:"false"`
	GRPCAPIKey    string `env:"GRPC_API_KEY" envDefault:"test"`

	// gRPC gateway
	GatewayIP     string `env:"GRPC_GATEWAY_IP"      envDefault:"127.0.0.1"`
	GatewayPort   uint16 `env:"GRPC_GATEWAY_PORT"    envDefault:"8090"`
	GatewayAPIKey string `env:"GRPC_GATEWAY_API_KEY" envDefault:"test"`
}
