package grpc

type Gateway interface {
	ListenAndServe() error
}

type Server interface {
	Serve() error
}
